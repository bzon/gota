package parser

import (
	"testing"
	"time"

	"github.com/bzon/ipapk"
)

func TestGenerateAssets(t *testing.T) {
	type tc struct {
		name string
		file string
	}
	tt := []tc{
		{"generate ios asset", "testdata/sample.ipa"},
		{"generate android asset", "testdata/sample.apk"},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			appInfo, err := ipapk.NewAppParser(tc.file)
			if err != nil {
				t.Fatal(err)
			}
			var app MobileApp
			app.UploadDate = time.Now().Format(time.RFC1123)
			app.AppInfo = appInfo
			app.File = tc.file
			if err := app.GenerateAssets(); err != nil {
				t.Fatal(err)
			}
		})
	}

}
