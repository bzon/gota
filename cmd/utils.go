package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/bzon/gota/parser"
	"github.com/spf13/cobra"
)

// common
var title, srcFile, buildNumber, destDir string

// ios specific
var bundleID, bundleVersion string

// android specific
var versionName, versionCode string

func missingFlagError(cmd *cobra.Command, f string) {
	fmt.Printf("Error: required flag(s) \"%s\" not set\n", f)
	cmd.Usage()
	os.Exit(1)
}

func newApp() parser.MobileApp {
	appFile := parser.AppFile{
		Title:       title,
		BuildNumber: buildNumber,
		SourceFile:  srcFile,
	}
	if fileExt() == ".ipa" {
		return parser.IOSIPA{
			AppFile:       appFile,
			BundleID:      bundleID,
			BundleVersion: bundleVersion,
		}
	} else {
		return parser.AndroidAPK{
			AppFile:     appFile,
			VersionCode: versionCode,
			VersionName: versionName,
		}
	}
}

func fileExt() string {
	return filepath.Ext(srcFile)
}

// This function must be called before executing any commands function!
func validateAndParseArgs(cmd *cobra.Command) {
	if fileExt() == ".ipa" {
		if bundleVersion == "" {
			missingFlagError(cmd, "bundleVersion")
		}
		if bundleID == "" {
			missingFlagError(cmd, "bundleID")
		}
	} else if fileExt() == ".apk" {
		if versionName == "" {
			missingFlagError(cmd, "versionName")
		}
		if versionCode == "" {
			missingFlagError(cmd, "versionCode")
		}
	} else {
		fmt.Printf("Error: srcFile %s does not have a file extension of .apk or .ipa\n", srcFile)
		cmd.Usage()
		os.Exit(1)
	}
	if _, err := os.Stat(srcFile); os.IsNotExist(err) {
		fmt.Println(err)
		os.Exit(1)
	}
}
