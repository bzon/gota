// The MIT License (MIT)

// Copyright (c) John Bryan Sazon <bryansazon@hotmail.com>

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

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
	SrcFile, DestFilePath string
}

// NexusUpload uploads a file to Nexus returns the uploaded file url
func (n *Nexus) NexusUpload(c NexusComponent) (string, error) {
	file, err := os.Open(c.SrcFile)
	if err != nil {
		return "", err
	}
	defer file.Close()
	uri := n.getRepoURL() + "/" + c.DestFilePath
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
func (n *Nexus) NexusUploadAssets(app *parser.MobileApp, destBaseDir string) ([]string, error) {
	// create the site path names and assume the url before uploaded for templating
	buildDir := destBaseDir + "/" + app.Version + "/" + app.Build
	appIconPath := buildDir + "/" + parser.AppIconFile
	appSitePath := buildDir + "/" + filepath.Base(app.File)
	appIndexHTMLSitePath := buildDir + "/" + parser.IndexHTMLFile
	app.DownloadURL = n.getRepoURL() + "/" + appSitePath

	// default directory of assets
	assetsDir := parser.AndroidAssetsDir

	// specific for ios
	var appPlistSitePath string
	if app.IsIOS() {
		assetsDir = parser.IOSAssetsDir
		appPlistSitePath = buildDir + "/" + parser.IOSPlistFile
		app.PlistURL = htmltemp.URL(n.getRepoURL() + "/" + appPlistSitePath)
	}

	// create the assets
	assets := []string{}
	if err := app.GenerateAssets(); err != nil {
		return assets, err
	}

	components := []NexusComponent{
		{assetsDir + "/" + parser.AppIconFile, appIconPath},
		{assetsDir + "/" + parser.VersionJsonFile, destBaseDir + "/" + app.Version + "/" + parser.VersionJsonFile},
		{assetsDir + "/" + parser.IndexHTMLFile, appIndexHTMLSitePath},
		{app.File, appSitePath},
	}
	if app.IsIOS() {
		components = append(components, NexusComponent{assetsDir + "/" + parser.IOSPlistFile, appPlistSitePath})
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
