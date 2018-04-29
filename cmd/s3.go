// The MIT License (MIT)

// Copyright (c) John Bryan Sazon <bryansazon@hotmail.com>

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

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
