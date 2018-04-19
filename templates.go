package gota

const versionTemplateString = `{
  "latestVersion": "{{.FullVersion}}",
  "updateUrl": "itms-services://?action=download-manifest&url={{.PlistURL}}"
}`

const plistTemplateString = `<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
  <key>items</key>
  <array>
    <dict>
      <key>assets</key>
        <dict>
          <key>kind</key>
          <string>software-package</string>
          <key>url</key>
	  <string>{{.AppFile.DownloadURL}}</string>
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
	<string>{{.AppFile.Title}}</string>
      </dict>
    </dict>
  </array>
</dict>
</plist>`

var indexHTMLTemplateString = `<html>
  <head>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0, maximum-scale=1.0, user-scalable=no">
	<title>Install {{.AppFile.Title}}</title>
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

    <h1 style="text-align: center;">{{.AppFile.Title}}</h1>
    <div class="oneRow">
      <span class="download" id="ios">
        <a href="itms-services://?action=download-manifest&url={{.PlistURL}}" id="text" class="btn btn-lg btn-default" onclick="document.getElementById('finished').id = '';">
			Install {{.AppFile.Title}} {{.BundleVersion}} ({{.BuildNumber}})
        </a>
        <br>
		<p>Built on {{.AppFile.BuildDate}}</p>
      </span>
    </div>

    <h3 id="desktop">Please open this page on your iPhone!</h3>

    <p id="finished">
      App is being installed. Close Safari using the home button.
    </p>

    <p id="footnote">
	Contribute on Github!
    </p>
	<a href="https://github.com/bzon/gota">
		<img src="https://github.com/bzon/gota/blob/master/assets/Octocat.png?raw=true" id="githubLogo" style="width:42px;height:42px;border:0;"/>
	</a>
  </body>

  <script type='text/javascript'>
    document.getElementById("desktop").remove()
  </script>
</html>`
