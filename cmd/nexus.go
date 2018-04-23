// Copyright © 2018 John Bryan Sazon <bryansazon@hotmail.com>
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
	"fmt"
	"os"

	"github.com/bzon/gota/parser"

	nexuspkg "github.com/bzon/gota/nexus"
	"github.com/spf13/cobra"
)

var nexus nexuspkg.Nexus

// nexusCmd represents the nexus command
var nexusCmd = &cobra.Command{
	Use:   "nexus",
	Short: "Upload your apk or ipa file and create an over-the-air static site in a Nexus Site repository",
	Run: func(cmd *cobra.Command, args []string) {
		validateAndParseArgs(cmd)
		var assets []string
		var err error
		app := newApp()
		switch app.(type) {
		case parser.IOSIPA:
			ipa := app.(parser.IOSIPA)
			assets, err = nexus.NexusUploadIOSAssets(&ipa, destDir)
		case parser.AndroidAPK:
			apk := app.(parser.AndroidAPK)
			assets, err = nexus.NexusUploadAndroidAssets(&apk, destDir)
		}
		if err != nil {
			fmt.Printf("failed uploading assets: %+v", err)
			os.Exit(1)
		}
		for _, a := range assets {
			fmt.Printf("uploaded to nexus: %s\n", a)
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
