package gota

import (
	"testing"
	"time"
)

func TestGenerateAssets(t *testing.T) {
	var ipa = IOSIPA{
		AppFile: AppFile{
			Title:       "DarkSouls",
			BuildDate:   time.Now().Format(time.RFC822),
			DownloadURL: "http://s3.amazon.com/bucket/DarkSouls.ipa",
			BuildNumber: "99",
			SourceFile:  "resources/DarkSouls.ipa", // dummy file
		},
		PlistURL:      "http://s3.amazon.com/bucket/DarkSouls.plist",
		BundleID:      "com.example.dark.souls",
		BundleVersion: "1.0.0",
	}
	var apk = AndroidAPK{
		AppFile: AppFile{
			Title:       "DarkSouls Android",
			BuildDate:   time.Now().Format(time.RFC822),
			DownloadURL: "http://s3.amazon.com/bucket/DarkSouls.apk",
			BuildNumber: "22",
			SourceFile:  "resources/DarkSouls.apk", // dummy file
		},
		VersionName: "1.0.0",
		VersionCode: "100222333",
	}
	t.Run("generate ipa assets", func(t *testing.T) {
		if err := ipa.GenerateAssets(); err != nil {
			t.Fatal(err)
		}
	})
	t.Run("generate apk assets", func(t *testing.T) {
		if err := apk.GenerateAssets(); err != nil {
			t.Fatal(err)
		}
	})
}
