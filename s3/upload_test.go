package s3

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/bzon/gota/parser"
	"github.com/bzon/ipapk"
)

func TestUploadFile(t *testing.T) {
	if os.Getenv("GOTEST_AWS_BUCKET") == "" {
		t.Fatal("GOTEST_AWS_BUCKET env variable is not set.")
	}
	if os.Getenv("AWS_ACCESS_KEY") == "" {
		t.Fatal("AWS_ACCESS_KEY env variable is not set.")
	}
	if os.Getenv("AWS_SECRET_ACCESS_KEY") == "" {
		t.Fatal("AWS_SECRET_ACCESS_KEY env variable is not set.")
	}
	file, err := os.Create("samplefile.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()
	defer os.Remove("samplefile.txt")
	upload := Upload{
		Bucket:       os.Getenv("GOTEST_AWS_BUCKET"),
		SrcFile:      "samplefile.txt",
		DestFilePath: "dir1/dir2/samplefile.txt",
	}
	fileUrl, err := UploadFile(upload)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(fileUrl)
}

func TestUploadAssets(t *testing.T) {
	type tc struct {
		name    string
		destDir string
		file    string
	}
	tt := []tc{
		{"upload ios assets", "ios_test", "../parser/testdata/sample.ipa"},
		{"upload android assets", "android_test", "../parser/testdata/sample.apk"},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			appInfo, err := ipapk.NewAppParser(tc.file)
			if err != nil {
				t.Fatal(err)
			}
			var app parser.MobileApp
			app.UploadDate = time.Now().Format(time.RFC1123)
			app.AppInfo = appInfo
			app.File = tc.file
			if err := app.GenerateAssets(); err != nil {
				t.Fatal(err)
			}
			assets, err := UploadAssets(&app, os.Getenv("GOTEST_AWS_BUCKET"), tc.destDir)
			if err != nil {
				t.Fatal(err)
			}
			for _, v := range assets {
				fmt.Println("s3 url:", v)
			}
		})
	}
}
