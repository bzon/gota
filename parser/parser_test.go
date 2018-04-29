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
