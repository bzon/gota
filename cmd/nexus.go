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

	nexuspkg "github.com/bzon/gota/nexus"
	"github.com/spf13/cobra"
)

var nexus nexuspkg.Nexus

// nexusCmd represents the nexus command
var nexusCmd = &cobra.Command{
	Use:   "nexus",
	Short: "Upload your apk or ipa file and create an over-the-air static site in a Nexus Site repository",
	Run: func(cmd *cobra.Command, args []string) {
		app := NewMobileAppParser()
		if err := app.GenerateAssets(); err != nil {
			log.Fatal(err)
		}
		assets, err := nexus.NexusUploadAssets(app, destDir)
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
	rootCmd.AddCommand(nexusCmd)
	nexusCmd.Flags().StringVar(&nexus.HostURL, "nexusHost", "", "nexus host url (including http protocol)")
	nexusCmd.Flags().StringVar(&nexus.Username, "nexusUser", "", "nexus username (can be passed as env variable $NEXUS_USER)")
	nexusCmd.Flags().StringVar(&nexus.Password, "nexusPassword", "", "nexus password (can be passed as env variable $NEXUS_PASSWORD)")
	nexusCmd.Flags().StringVar(&nexus.SiteRepository, "nexusRepo", "", "nexus site repository id (nexus v3 raw repository not maven!)")
	nexusCmd.MarkFlagRequired("nexusHost")
	nexusCmd.MarkFlagRequired("nexusRepo")
	if os.Getenv("NEXUS_USER") == "" {
		nexusCmd.MarkFlagRequired("nexusUser")
	} else {
		nexus.Username = os.Getenv("NEXUS_USER")
	}
	if os.Getenv("NEXUS_PASSWORD") == "" {
		nexusCmd.MarkFlagRequired("nexusPassword")
	} else {
		nexus.Password = os.Getenv("NEXUS_PASSWORD")
	}
}
