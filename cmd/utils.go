package cmd

import (
	"log"
	"time"

	"github.com/bzon/ipapk"

	"github.com/bzon/gota/parser"
)

func NewMobileAppParser() *parser.MobileApp {
	appInfo, err := ipapk.NewAppParser(srcFile)
	if err != nil {
		log.Fatal(err)
	}
	var app parser.MobileApp
	app.UploadDate = time.Now().Format(time.RFC1123)
	app.AppInfo = appInfo
	app.File = srcFile
	return &app
}
