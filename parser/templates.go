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

var indexHTMLTemplateString = `<html>
  <head>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0, maximum-scale=1.0, user-scalable=no">
	<title>Install {{.Name}}</title>
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

    <h1 style="text-align: center;">{{.Name}}</h1>

	{{if .IsIOS -}}
	<a href="{{.PlistURL}}"><img src="appicon.png" id="appIcon"></a>
    <div class="oneRow">
      <span class="download" id="ios">
        <a href="itms-services://?action=download-manifest&amp;url={{.PlistURL}}" id="text" class="btn btn-lg btn-default" onclick="document.getElementById('finished').id = '';">
			Install {{.Name}} {{.Version}} ({{.Build}})
        </a>
        <br>
		<p>Uploaded on {{.UploadDate}}</p>
      </span>
    </div>

    <p id="finished">
      App is being installed. Close your Browser using the home button.
    </p>
	{{- else}}	
	<a href="{{.DownloadURL}}"><img src="appicon.png" id="appIcon"></a>
    <div class="oneRow">
      <span class="download" id="android">
        <a href="{{.DownloadURL}}" id="text" class="btn btn-lg btn-default" onclick="document.getElementById('finished').id = '';">
          Install {{.Name}} {{.Version}} ({{.Build}})
        </a>
        <br>
        <p>Uploaded on {{.UploadDate}}</p>
      </span>
    </div>

    <p id="finished">
      App is being installed. You might have to close the browser.
    </p>

	{{- end}}
    <p id="footnote">
	This is a beta version and is not meant for the public.
    </p>
  </body>

  <script type='text/javascript'>
    document.getElementById("desktop").remove()
  </script>
</html>`
