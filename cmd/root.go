// Copyright © 2018 John Bryan Sazon <bryansazon@hotmail.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gota",
	Short: "Go Over the Air installation for Android APK and iOS Ipa files! Source: https://github.com/bzon/gota",
	Run: func(cmd *cobra.Command, args []string) {
		validateAndParseArgs(cmd)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&title, "title", "", "application name to be displayed in the site")
	rootCmd.PersistentFlags().StringVar(&srcFile, "srcFile", "", "the apk or ipa file.")
	rootCmd.PersistentFlags().StringVar(&destDir, "destDir", "", "root directory of the site to create.")
	rootCmd.PersistentFlags().StringVar(&buildNumber, "buildNumber", "", "the apk or ipa build number.")
	rootCmd.PersistentFlags().StringVar(&bundleVersion, "bundleVersion", "", "if srcFile type is '.ipa', this is required.")
	rootCmd.PersistentFlags().StringVar(&bundleID, "bundleID", "", "if srcFile type is '.ipa', this is required. (example: com.example.bundleid)")
	rootCmd.PersistentFlags().StringVar(&versionName, "versionName", "", "if srcFile is '.apk', this is required.")
	rootCmd.PersistentFlags().StringVar(&versionCode, "versionCode", "", "if srcfile is '.apk', this is required.")
	rootCmd.MarkPersistentFlagRequired("title")
	rootCmd.MarkPersistentFlagRequired("srcFile")
	rootCmd.MarkPersistentFlagRequired("destDir")
	rootCmd.MarkPersistentFlagRequired("buildNumber")
}
