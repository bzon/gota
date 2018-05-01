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

const versionTemplateString = `{
"latestVersion": "{{.Version}}.{{.Build}}",
{{if .IsIOS -}}
  "updateUrl": "itms-services://?action=download-manifest&amp;url={{.PlistURL}}"
{{else -}}
  "updateUrl": "{{.DownloadURL}}"
{{end -}}
}`

const plistTemplateString = `<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
  <key>items</key>
  <array>
    <dict>
      <key>assets</key>
	  <array>
        <dict>
          <key>kind</key>
          <string>software-package</string>
          <key>url</key>
	  <string>{{.DownloadURL}}</string>
        </dict>
      </array>
      <key>metadata</key>
      <dict>
        <key>bundle-identifier</key>
	<string>{{.BundleId}}</string>
        <key>bundle-version</key>
	<string>{{.Version}}</string>
        <key>kind</key>
        <string>software</string>
        <key>title</key>
	<string>{{.Name}}</string>
      </dict>
    </dict>
  </array>
</dict>
</plist>`

var indexHTMLTemplateString = `<!DOCTYPE HTML5>
<html>
    <head>
        <meta name="viewport" content="width=device-width, initial-scale=1.0">
        <link rel="stylesheet" href="https://fonts.googleapis.com/icon?family=Material+Icons">
        <link rel="stylesheet" href="https://code.getmdl.io/1.3.0/material.indigo-pink.min.css">
        <script defer src="https://code.getmdl.io/1.3.0/material.min.js"></script>
        <style>
		  	* {
				font-family: "Helvetica Neue",Helvetica,Arial,sans-serif;
		  	}
            .main-container {
                height: inherit;
                width: inherit;
                margin-top: 2%;
            }
            .demo-card-wide.mdl-card {
                width: 512px;
                margin: 0 auto;
            }
            .demo-card-wide > .mdl-card__title {
                color: #fff;
                height: 116px;
                background: #3c8fc6 center / cover;
                text-indent: 0px;   
                vertical-align:bottom;             
            }
            .demo-card-wide > .mdl-card__details {
				margin: 0 auto;
				text-align: center;
            }
            .demo-card-wide > .mdl-card__menu {
                color: #fff;
            }
            .logo {
                background: url('appicon.png') center / cover;
                display: block; 
                float: left;
                width: 80px;
                height: 80px;
            }
            .logo-title {
                display: block; 
                float: left;
                text-indent: 2%;
                height: 80px;
                line-height: 80px;
            }
            .release-note {
                margin: 0 auto; 
                margin-top: 10px; 
                width: 512px;
            }
        </style>
    </head>
    <body>
        <div class="main-container">
            <!-- elease Details -->
            <div class="demo-card-wide mdl-card mdl-shadow--2dp">
                <div class="mdl-card__title">
                    <div class="logo"></div>
                    <div class="logo-title mdl-card__title-text">{{.Name}}</div>
                </div>
                <div class="mdl-card__supporting-text" style="padding: 0; width: 100%">
                    <table class="mdl-data-table mdl-js-data-table" style="width: 100%; border: 0px">
                        <tbody>
							{{if .IsIOS -}}
                            <tr>
                    	        <td class="mdl-data-table__cell--non-numeric">CFBundleShortVersion</td>
                                <td class="mdl-data-table__cell--non-numeric">{{.Version}}</td>
							</tr>
							<tr>
                    	        <td class="mdl-data-table__cell--non-numeric">CFBundleVersion</td>
                                <td class="mdl-data-table__cell--non-numeric">{{.Build}}</td>
							</tr>
							{{else -}}
							<tr>
                    	        <td class="mdl-data-table__cell--non-numeric">Version Name</td>
                                <td class="mdl-data-table__cell--non-numeric">{{.Version}}</td>
							</tr>
							<tr>
                    	        <td class="mdl-data-table__cell--non-numeric">Version Code</td>
                                <td class="mdl-data-table__cell--non-numeric">{{.Build}}</td>
							</tr>
							{{end -}}
							<tr>
                    	        <td class="mdl-data-table__cell--non-numeric">Bundle ID</td>
                                <td class="mdl-data-table__cell--non-numeric">{{.BundleId}}</td>
							</tr>
							<tr>
                    	        <td class="mdl-data-table__cell--non-numeric">Upload Date</td>
                                <td class="mdl-data-table__cell--non-numeric">{{.UploadDate}}</td>
                            </tr>
                        </tbody>
                    </table>
                </div>

                <div class="mdl-card__actions mdl-card--border">
					{{if .IsIOS -}}
                    <a href="itms-services://?action=download-manifest&amp;url={{.PlistURL}}" class="mdl-button mdl-button--raised mdl-button--colored mdl-js-button mdl-js-ripple-effect">
                    Install
                    </a>
					{{else -}}
					<a href="{{.DownloadURL}}" class="mdl-button mdl-button--raised mdl-button--colored mdl-js-button mdl-js-ripple-effect">
                    Install
                    </a>
					{{end -}}
                </div>
                <div class="mdl-card__menu">
                    <button class="mdl-button mdl-button--icon mdl-js-button mdl-js-ripple-effect">
                    <i class="material-icons">share</i>
                    </button>
                </div>
            </div>
            <!-- Release Notes -->
            <table class="mdl-data-table mdl-shadow--2dp mdl-js-data-table release-note">
                <thead>
					<tr>
                       <th class="mdl-data-table__cell--non-numeric">Date</th>
                       <th class="mdl-data-table__cell--non-numeric">Author</th>
                       <th class="mdl-data-table__cell--non-numeric">Subject</th>
					</tr>
                </thead>
                <tbody>
					{{range .Changelogs}}
                    <tr>
                        <td class="mdl-data-table__cell--non-numeric">{{.Date}}</td>
                        <td class="mdl-data-table__cell--non-numeric">{{.Author}}</td>
                        <td class="mdl-data-table__cell--non-numeric">{{.Subject}}</td>
                    </tr>
					{{end}}
                </tbody>
            </table>
        </div>
    </body>
</html>
`
