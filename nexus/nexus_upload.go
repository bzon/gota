package nexus

import (
	"fmt"
	htmltemp "html/template"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	"github.com/bzon/gota/parser"
)

// Nexus contains the fields requied for accessing a nexus server
type Nexus struct {
	SiteRepository, HostURL, Username, Password string
}

// NexusComponent contains the fields that will be passed as a parameter for NexusUpload
type NexusComponent struct {
	File, Filename, Directory string
}

// NexusUpload uploads a file to Nexus returns the uploaded file url
func (n *Nexus) NexusUpload(c NexusComponent) (string, error) {
	file, err := os.Open(c.File)
	if err != nil {
		return "", err
	}
	defer file.Close()
	uri := n.getRepoURL() + "/" + c.Directory + "/" + c.Filename
	req, err := http.NewRequest("PUT", uri, file)
	if err != nil {
		return "", err
	}
	req.SetBasicAuth(n.Username, n.Password)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusCreated {
		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return "", err
		}
		return "", fmt.Errorf("%s: %s", resp.Status, string(b))
	}
	return uri, nil
}

func (n *Nexus) getRepoURL() string {
	return n.HostURL + "/repository/" + n.SiteRepository
}

// NexusUploadAssets uploads the generated files by the parser package along with the ipa or apk file
func (n *Nexus) NexusUploadAssets(app *parser.MobileApp, dir string) ([]string, error) {
	// create the site path names and assume the url before uploaded for templating
	appSitePath := app.Version + "/" + filepath.Base(app.File)
	appIndexHTMLSitePath := app.Version + "/index.html"
	app.DownloadURL = n.getRepoURL() + "/" + dir + "/" + appSitePath

	// default directory of assets
	assetsDir := "android_assets"

	// specific for ios
	var appPlistSitePath string
	if app.IsIOS() {
		assetsDir = "ios_assets"
		appPlistSitePath = app.Version + "/" + app.Name + ".plist"
		app.PlistURL = htmltemp.URL(n.getRepoURL() + "/" + dir + "/" + appPlistSitePath)
	}

	// create the assets
	assets := []string{}
	if err := app.GenerateAssets(); err != nil {
		return assets, err
	}

	// upload assets
	uri, err := n.NexusUpload(NexusComponent{assetsDir + "/version.json", "version.json", dir})
	if err != nil {
		return assets, err
	}
	assets = append(assets, uri)
	if app.IsIOS() {
		uri, err = n.NexusUpload(NexusComponent{assetsDir + "/app.plist", appPlistSitePath, dir})
		if err != nil {
			return assets, err
		}
		assets = append(assets, uri)
	}
	uri, err = n.NexusUpload(NexusComponent{assetsDir + "/index.html", appIndexHTMLSitePath, dir})
	if err != nil {
		return assets, err
	}
	assets = append(assets, uri)
	uri, err = n.NexusUpload(NexusComponent{app.File, appSitePath, dir})
	if err != nil {
		return assets, err
	}
	assets = append(assets, uri)
	return assets, nil
}
