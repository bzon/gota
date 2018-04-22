[![forthebadge](https://forthebadge.com/images/badges/built-with-love.svg)](https://forthebadge.com)

# GOTA

Automate the beta testing distribution of your Android and iOS application files with gota. Gota is a [Golang](http://golang.org/) powered Over the Air Installation site creation tool.

![](./docs/gota_html.png)

## Feature Checklist

* [x] Upload and generate site to a Nexus 3 Site Repository
* [ ] Upload and generate site to a Nexus 2 Site Repository (untested)
* [ ] Upload and generate site to an Amazon S3 bucket

## User Guide

### Nexus APK Upload

Upload an APK file to a Nexus Site Repository

```bash
./gota nexus --nexusHost http://localhost:8081 --nexusRepo site --nexusUser admin --nexusPassword admin123 --destDir android --buildNumber 1 --srcFile pkg/resources/DarkSouls.apk --title "DarkSouls" --versionName "1.0.0" --versionCode "10222333"
uploaded to nexus: http://localhost:8081/repository/site/android/version.json
uploaded to nexus: http://localhost:8081/repository/site/android/1.0.0.10222333/index.html
uploaded to nexus: http://localhost:8081/repository/site/android/1.0.0.10222333/DarkSouls.apk
```

Uploaded site structure

![](./docs/apk_nexus_uploaded.png)

### Nexus IPA Upload

Upload an IPA file to a Nexus Site Repository

```bash
./gota nexus --nexusHost http://localhost:8081 --nexusRepo site --nexusUser admin --nexusPassword admin123 --destDir ios --buildNumber 1 --srcFile pkg/resources/DarkSouls.ipa --title DarkSouls --bundleVersion 1.0.0 --bundleID com.example.com
uploaded to nexus: http://localhost:8081/repository/site/ios/version.json
uploaded to nexus: http://localhost:8081/repository/site/ios/1.0.0.1/DarkSouls.plist
uploaded to nexus: http://localhost:8081/repository/site/ios/1.0.0.1/index.html
uploaded to nexus: http://localhost:8081/repository/site/ios/1.0.0.1/DarkSouls.ipa
```

Uploaded site structure

![](./docs/ios_nexus_uploaded.png)

## Development Setup

### Nexus

You must have a Nexus 3 server running in your machine. Get it via docker: `docker run -d -p 8081:8081 --name nexus3` then create a Raw Repository with a repository id `site`.

```bash
go get -v ./...
go test -v ./...
```

### S3

WIP

## Inspirations and References

* [Fastlane Nexus Upload](https://docs.fastlane.tools/actions/nexus_upload/)
* [Fastlane S3 Plugin](https://github.com/joshdholtz/fastlane-plugin-s3/)
* [Creating an Installation Link for your enterprise App](https://support.magplus.com/hc/en-us/articles/203808598-iOS-Creating-an-Installation-Link-for-Your-Enterprise-App)
