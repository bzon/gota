// https://support.magplus.com/hc/en-us/articles/203808598-iOS-Creating-an-Installation-Link-for-Your-Enterprise-App
package gota

import (
	"fmt"
	"html/template"
	"net/url"
	"os"
)

// AppFile contains common fields of APK and IPA file
type AppFile struct {
	Title, BuildDate, BuildNumber, DownloadURL string
}

// AndroidAppData fields for APK file upload
type AndroidAPK struct {
	AppFile
	VersionName, VersionCode string
}

// IOSIPA fields for IPA file upload
type IOSIPA struct {
	AppFile
	PlistURL, BundleID, BundleVersion string
}

type MobileApp interface {
	GenerateAssets() error
	FullVersion() string
}

// GenerateAssets creates files for iOS
func (ipa IOSIPA) GenerateAssets() error {
	os.Mkdir("ios_assets", 0755)
	ipa.PlistURL = url.QueryEscape(ipa.PlistURL)
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
	os.Mkdir("android_assets", 0755)
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
	templ := template.Must(template.New("templ").Parse(templateVar))
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
