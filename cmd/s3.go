// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"

	"github.com/bzon/gota/s3"
)

var upload s3.Upload

// s3Cmd represents the s3 command
var s3Cmd = &cobra.Command{
	Use:   "s3",
	Short: "Upload your apk or ipa file and create an over-the-air static site in an S3 Bucket directory",
	Long:  `Ensure that you have AWS_ACCESS_KEY and AWS_SECRET_ACCESS_KEY set in your environment variable.`,
	Run: func(cmd *cobra.Command, args []string) {
		if os.Getenv("AWS_ACCESS_KEY") == "" {
			log.Fatal("AWS_ACCESS_KEY env variable is not set")
		}
		if os.Getenv("AWS_SECRET_ACCESS_KEY") == "" {
			log.Fatal("AWS_SECRET_ACCESS_KEY env variable is not set")
		}
		app := NewMobileAppParser()
		if err := app.GenerateAssets(); err != nil {
			log.Fatal(err)
		}
		assets, err := s3.UploadAssets(app, upload.Bucket, destDir)
		if err != nil {
			log.Fatal(err)
		}
		for _, v := range assets {
			log.Println("file uploaded:", v)
			// write the index.html file (the ota link) to a file to gotalink.txt
			if strings.Contains(v, "index.html") {
				if err := ioutil.WriteFile("gotalink.txt", []byte(v), 0644); err != nil {
					log.Fatal(err)
				}
			}
			// write the ipa download link to a file ipalink.txt
			if strings.Contains(v, ".ipa") {
				if err := ioutil.WriteFile("ipalink.txt", []byte(v), 0644); err != nil {
					log.Fatal(err)
				}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(s3Cmd)
	s3Cmd.Flags().StringVar(&upload.Bucket, "bucket", "", "the amazon s3 bucket name")
	s3Cmd.MarkFlagRequired("bucket")
}
