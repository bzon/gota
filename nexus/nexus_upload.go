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
	appIconPath := app.Version + "/" + parser.AppIconFile
	appSitePath := app.Version + "/" + filepath.Base(app.File)
	appIndexHTMLSitePath := app.Version + "/" + parser.IndexHTMLFile
	app.DownloadURL = n.getRepoURL() + "/" + dir + "/" + appSitePath

	// default directory of assets
	assetsDir := parser.AndroidAssetsDir

	// specific for ios
	var appPlistSitePath string
	if app.IsIOS() {
		assetsDir = parser.IOSAssetsDir
		appPlistSitePath = app.Version + "/" + parser.IOSPlistFile
		app.PlistURL = htmltemp.URL(n.getRepoURL() + "/" + dir + "/" + appPlistSitePath)
	}

	// create the assets
	assets := []string{}
	if err := app.GenerateAssets(); err != nil {
		return assets, err
	}

	components := []NexusComponent{
		{assetsDir + "/" + parser.AppIconFile, appIconPath, dir},
		{assetsDir + "/" + parser.VersionJsonFile, parser.VersionJsonFile, dir},
		{assetsDir + "/" + parser.IndexHTMLFile, appIndexHTMLSitePath, dir},
		{app.File, appSitePath, dir},
	}
	if app.IsIOS() {
		components = append(components, NexusComponent{assetsDir + "/" + parser.IOSPlistFile, appPlistSitePath, dir})
	}

	for _, component := range components {
		uri, err := n.NexusUpload(component)
		if err != nil {
			return assets, err
		}
		assets = append(assets, uri)
	}

	return assets, nil
}
