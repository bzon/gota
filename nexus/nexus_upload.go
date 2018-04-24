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
	SourceFile, Filename, Directory string
}

// NexusUpload uploads a file to Nexus returns the uploaded file url
func (n *Nexus) NexusUpload(c NexusComponent) (string, error) {
	file, err := os.Open(c.SourceFile)
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

// NexusUploadIOSAssets wraps NexusUpload to upload all files for iOS to nexus
func (n *Nexus) NexusUploadIOSAssets(ipa *parser.IOSIPA, dir string) ([]string, error) {
	// Upload in directory with FullVersion() as name
	ipaSitePath := ipa.FullVersion() + "/" + filepath.Base(ipa.SourceFile)
	ipaPlistSitePath := ipa.FullVersion() + "/" + ipa.Title + ".plist"
	ipaIndexHTMLSitePath := ipa.FullVersion() + "/index.html"
	// assume the url before uploaded for templating
	ipa.DownloadURL = n.getRepoURL() + "/" + dir + "/" + ipaSitePath
	ipa.PlistURL = htmltemp.URL(n.getRepoURL() + "/" + dir + "/" + ipaPlistSitePath)
	// create the assets
	assets := []string{}
	if err := ipa.GenerateAssets(); err != nil {
		return assets, err
	}
	// upload assets
	uri, err := n.NexusUpload(NexusComponent{"ios_assets/version.json", "version.json", dir})
	if err != nil {
		return assets, err
	}
	assets = append(assets, uri)
	uri, err = n.NexusUpload(NexusComponent{"ios_assets/app.plist", ipaPlistSitePath, dir})
	if err != nil {
		return assets, err
	}
	assets = append(assets, uri)
	uri, err = n.NexusUpload(NexusComponent{"ios_assets/index.html", ipaIndexHTMLSitePath, dir})
	if err != nil {
		return assets, err
	}
	assets = append(assets, uri)
	uri, err = n.NexusUpload(NexusComponent{ipa.SourceFile, ipaSitePath, dir})
	if err != nil {
		return assets, err
	}
	assets = append(assets, uri)
	return assets, nil
}

// NexusUploadAndroidAssets wraps NexusUpload to upload all files for android to nexus
func (n *Nexus) NexusUploadAndroidAssets(apk *parser.AndroidAPK, dir string) ([]string, error) {
	// Upload in directory with FullVersion() as name
	apkSitePath := apk.FullVersion() + "/" + filepath.Base(apk.SourceFile)
	apkIndexHTMLSitePath := apk.FullVersion() + "/index.html"
	// assume the url before uploaded for templating
	apk.DownloadURL = n.getRepoURL() + "/" + dir + "/" + apkSitePath
	// create the assets
	assets := []string{}
	if err := apk.GenerateAssets(); err != nil {
		return assets, err
	}
	// upload assets
	uri, err := n.NexusUpload(NexusComponent{"android_assets/version.json", "version.json", dir})
	if err != nil {
		return assets, err
	}
	assets = append(assets, uri)
	uri, err = n.NexusUpload(NexusComponent{"android_assets/index.html", apkIndexHTMLSitePath, dir})
	if err != nil {
		return assets, err
	}
	assets = append(assets, uri)
	uri, err = n.NexusUpload(NexusComponent{apk.SourceFile, apkSitePath, dir})
	if err != nil {
		return assets, err
	}
	assets = append(assets, uri)
	return assets, nil
}
