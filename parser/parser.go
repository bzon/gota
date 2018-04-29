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

package parser

import (
	"fmt"
	htmltemp "html/template"
	"image/png"
	"os"
	"path/filepath"
	txttemp "text/template"

	"github.com/bzon/ipapk"
)

const (
	IOSAssetsDir     = "ios_assets"
	AndroidAssetsDir = "android_assets"
	IOSPlistFile     = "app.plist"
	AppIconFile      = "appicon.png"
	VersionJsonFile  = "version.json"
	IndexHTMLFile    = "index.html"
	IPAExt           = ".ipa"
	APKExt           = ".apk"
)

// AppFile contains common fields of APK and IPA file
type MobileApp struct {
	*ipapk.AppInfo
	UploadDate  string
	DownloadURL string
	PlistURL    htmltemp.URL
	File        string
}

// GenerateAssets creates the site assets that will be uploaded along with the ipa or apk file
func (app MobileApp) GenerateAssets() error {
	assetsDir := AndroidAssetsDir
	if app.IsIOS() {
		assetsDir = IOSAssetsDir
	}
	os.Remove(assetsDir)
	os.Mkdir(assetsDir, 0700)
	if app.IsIOS() {
		err := executeTemplate(app, assetsDir+"/"+IOSPlistFile, plistTemplateString)
		if err != nil {
			return err
		}
	}
	err := generateCommonFiles(app, assetsDir)
	if err != nil {
		return err
	}
	return nil
}

// IsIOS is used for templating conditions
func (app MobileApp) IsIOS() bool {
	if filepath.Ext(app.File) == IPAExt {
		return true
	}
	return false
}

// Create index.html and version.json
func generateCommonFiles(app MobileApp, dir string) error {
	// Create the image png file
	if app.Icon != nil {
		img, err := os.Create(dir + "/" + AppIconFile)
		if err != nil {
			return err
		}
		defer img.Close()
		if err := png.Encode(img, app.Icon); err != nil {
			return fmt.Errorf("failed creating an image: %v", err)
		}
	}
	if err := executeTemplate(app, dir+"/"+IndexHTMLFile, indexHTMLTemplateString); err != nil {
		return err
	}
	if err := executeTemplate(app, dir+"/"+VersionJsonFile, versionTemplateString); err != nil {
		return err
	}
	return nil
}

func executeTemplate(app MobileApp, filename, templateVar string) error {
	if filepath.Ext(filename) == ".html" {
		templ := htmltemp.Must(htmltemp.New("templ").Parse(templateVar))
		file, err := os.Create(filename)
		if err != nil {
			return fmt.Errorf("failed creating %s got error %v", filename, err)
		}
		defer file.Close()
		if err := templ.Execute(file, app); err != nil {
			return fmt.Errorf("failed templating %s got error %v", filename, err)
		}
		return nil
	}
	templ := txttemp.Must(txttemp.New("templ").Parse(templateVar))
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed creating %s got error %v", filename, err)
	}
	defer file.Close()
	if err := templ.Execute(file, app); err != nil {
		return fmt.Errorf("failed templating %s got error %v", filename, err)
	}
	return nil
}
