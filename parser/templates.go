package parser

const versionTemplateString = `{
"latestVersion": "{{.FullVersion}}",
{{if .IsAndroid -}}
  "updateUrl": "{{.DownloadURL}}"
{{else -}}
  "updateUrl": "itms-services://?action=download-manifest&amp;url={{.PlistURL}}"
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
	<string>{{.BundleID}}</string>
        <key>bundle-version</key>
	<string>{{.BundleVersion}}</string>
        <key>kind</key>
        <string>software</string>
        <key>title</key>
	<string>{{.Title}}</string>
      </dict>
    </dict>
  </array>
</dict>
</plist>`

var indexHTMLTemplateString = `<html>
  <head>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0, maximum-scale=1.0, user-scalable=no">
	<title>Install {{.Title}}</title>
  </head>
  <body>
    <style type="text/css">
      * {
        font-family: "Helvetica Neue",Helvetica,Arial,sans-serif;
        text-align: center;
        background-color: #f5f5f5;
      }
      .oneRow {
        width: 100%;
        overflow: auto;
        overflow-y: hidden;
        white-space: nowrap;
        text-align: center;
      }
      .download {
        margin: 30px;
        font-size: 130%;
      }
      #appIcon {
        -webkit-border-radius: 22.544%;
        -moz-border-radius: 22.544%;
        -ms-border-radius: 22.544%;
          border-radius: 22.544%;
        margin-bottom: 30px;
      }
      a {
        text-decoration: none;
        color: blue;
      }
      a:hover {
        text-decoration: underline;
      }
      #footnote {
        color: #737373;
        font-size: 14px;
      }
      #finished { display: none; }
      #fastlaneLogo {
        text-align: center;
        max-width: 150px;
        margin-top: 10px;
      }
    </style>

    <h1 style="text-align: center;">{{.Title}}</h1>
	{{if .IsAndroid -}}
    <div class="oneRow">
      <span class="download" id="android">
        <a href="{{.DownloadURL}}" id="text" class="btn btn-lg btn-default" onclick="document.getElementById('finished').id = '';">
          Install {{.Title}} {{.VersionName}} ({{.VersionCode}})
        </a>
        <br>
        <p>Built on {{.BuildDate}}</p>
      </span>
    </div>

    <h3 id="desktop">Please open this page on your Android device!</h3>

    <p id="finished">
      App is being installed. You might have to close the browser.
    </p>
	{{- else}}	
    <div class="oneRow">
      <span class="download" id="ios">
        <a href="itms-services://?action=download-manifest&amp;url={{.PlistURL}}" id="text" class="btn btn-lg btn-default" onclick="document.getElementById('finished').id = '';">
			Install {{.Title}} {{.BundleVersion}} ({{.BuildNumber}})
        </a>
        <br>
		<p>Built on {{.BuildDate}}</p>
      </span>
    </div>

    <h3 id="desktop">Please open this page on your iPhone!</h3>

    <p id="finished">
      App is being installed. Close Safari using the home button.
    </p>
	{{- end}}
    <p id="footnote">
	Contribute on Github!
    </p>
	<a href="https://github.com/bzon/gota">
		<img src="https://github.com/bzon/gota/blob/master/resources/Octocat.png?raw=true" id="githubLogo" style="width:42px;height:42px;border:0;"/>
	</a>
  </body>

  <script type='text/javascript'>
    document.getElementById("desktop").remove()
  </script>
</html>`
