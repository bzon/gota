package parser

import (
	"fmt"
	htmltemp "html/template"
	"os"
	"path/filepath"
	txttemp "text/template"
)

// AppFile contains common fields of APK and IPA file
type AppFile struct {
	Title, BuildDate, BuildNumber, DownloadURL, SourceFile string
}

// AndroidAPK fields for APK file upload
type AndroidAPK struct {
	AppFile
	VersionName, VersionCode string
}

// IOSIPA fields for IPA file upload
type IOSIPA struct {
	AppFile
	BundleID, BundleVersion string
	PlistURL                htmltemp.URL
}

// MobileApp allows IOSIPA and AndroidAPK structs to be able use generateCommonFiles function
type MobileApp interface {
	GenerateAssets() error
	FullVersion() string
	IsAndroid() bool
}

// GenerateAssets creates files for iOS
func (ipa IOSIPA) GenerateAssets() error {
	os.Remove("ios_assets")
	os.Mkdir("ios_assets", 0700)
	//	ipa.PlistURL = url.QueryEscape(ipa.PlistURL)
	if err := executeTemplate(ipa, "ios_assets/app.plist", plistTemplateString); err != nil {
		return err
	}
	if err := generateCommonFiles(ipa, "ios_assets"); err != nil {
		return err
	}
	return nil
}

// GenerateAssets creates files for Android
func (apk AndroidAPK) GenerateAssets() error {
	os.Remove("android_assets")
	os.Mkdir("android_assets", 0700)
	if err := generateCommonFiles(apk, "android_assets"); err != nil {
		return err
	}
	return nil
}

// FullVersion returns the full version for Android
func (apk AndroidAPK) FullVersion() string {
	return apk.VersionName + "." + apk.VersionCode
}

// FullVersion returns the full version for iOS
func (ipa IOSIPA) FullVersion() string {
	return ipa.BundleVersion + "." + ipa.BuildNumber
}

// IsAndroid is used for templating conditions. Returns true for android
func (apk AndroidAPK) IsAndroid() bool {
	return true
}

// IsAndroid is used for templating conditions. Returns true for iOS
func (ipa IOSIPA) IsAndroid() bool {
	return false
}

// Create index.html and version.json
func generateCommonFiles(app MobileApp, dir string) error {
	if err := executeTemplate(app, dir+"/index.html", indexHTMLTemplateString); err != nil {
		return err
	}
	if err := executeTemplate(app, dir+"/version.json", versionTemplateString); err != nil {
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
