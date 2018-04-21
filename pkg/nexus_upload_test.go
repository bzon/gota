package gota

import (
	"fmt"
	"testing"
	"time"
)

var nexus = Nexus{
	HostURL:        "http://localhost:8081",
	Username:       "admin",
	Password:       "admin123",
	SiteRepository: "site",
}

func TestNexusUpload(t *testing.T) {
	var testComponent = NexusComponent{
		SourceFile: "resources/index.html",
		Filename:   "index.html",
		Directory:  "go_upload_test",
	}
	uri, err := nexus.NexusUpload(testComponent)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("nexus url:", uri)
}

func TestGeneratedAssetsNexusUpload(t *testing.T) {
	var ipa = IOSIPA{
		AppFile: AppFile{
			Title:       "DarkSouls",
			BuildDate:   time.Now().Format(time.RFC822),
			BuildNumber: "99",
			SourceFile:  "resources/DarkSouls.ipa", // dummy file
		},
		BundleID:      "com.example.dark.souls",
		BundleVersion: "1.0.0",
	}
	var apk = AndroidAPK{
		AppFile: AppFile{
			Title:       "DarkSouls Android",
			BuildDate:   time.Now().Format(time.RFC822),
			BuildNumber: "22",
			SourceFile:  "resources/DarkSouls.apk", // dummy file
		},
		VersionName: "1.0.0",
		VersionCode: "100222333",
	}
	t.Run("upload ipa assets", func(t *testing.T) {
		assets, err := nexus.NexusUploadIOSAssets(&ipa, "nexus_ios_upload_test")
		if err != nil {
			t.Fatal(err)
		}
		for _, v := range assets {
			fmt.Println("nexus url:", v)
		}
	})
	t.Run("upload android assets", func(t *testing.T) {
		assets, err := nexus.NexusUploadAndroidAssets(&apk, "nexus_android_upload_test")
		if err != nil {
			t.Fatal(err)
		}
		for _, v := range assets {
			fmt.Println("nexus url:", v)
		}
	})
}
