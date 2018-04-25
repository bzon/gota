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
	iosDir          = "ios_assets"
	androidDir      = "android_assets"
	iosPlistFile    = "app.plist"
	appIconFile     = "appicon.png"
	versionJsonFile = "version.json"
	indexHTMLFile   = "index.html"
	ipaExt          = ".ipa"
	apkExt          = ".apk"
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
	assetsDir := androidDir
	if app.IsIOS() {
		assetsDir = iosDir
	}
	os.Remove(assetsDir)
	os.Mkdir(assetsDir, 0700)
	if app.IsIOS() {
		err := executeTemplate(app, assetsDir+"/"+iosPlistFile, plistTemplateString)
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
	if filepath.Ext(app.File) == ipaExt {
		return true
	}
	return false
}

// Create index.html and version.json
func generateCommonFiles(app MobileApp, dir string) error {
	// Create the image png file
	if app.Icon != nil {
		img, err := os.Create(dir + "/" + appIconFile)
		if err != nil {
			return err
		}
		defer img.Close()
		if err := png.Encode(img, app.Icon); err != nil {
			return fmt.Errorf("failed creating an image: %v", err)
		}
	}
	if err := executeTemplate(app, dir+"/"+indexHTMLFile, indexHTMLTemplateString); err != nil {
		return err
	}
	if err := executeTemplate(app, dir+"/"+versionJsonFile, versionTemplateString); err != nil {
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
