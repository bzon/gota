# GOTA

[![Go Report Card](https://goreportcard.com/badge/github.com/bzon/gota)](https://goreportcard.com/report/github.com/bzon/gota)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/bzon/gota)
[![GitHub tag](https://img.shields.io/github/tag/bzon/gota.svg)](https://github.com/bzon/gota/releases/)

Automate the beta testing distribution of your Android and iOS application files with gota.

Gota is a [Golang](http://golang.org/) powered Over the Air Installation site creation tool.

![](./docs/gota_html.png)

## Feature Checklist

* [x] Upload and generate site to a Nexus 3 Site Repository
* [ ] Upload and generate site to a Nexus 2 Site Repository (untested)
* [ ] Upload and generate site to an Amazon S3 bucket

## Installation

Get the executable binary for your platform from the [Release Page](https://github.com/bzon/gota/releases/).

## Commands Guide

Gota command help `gota --help`

```bash
Go Over the Air installation for Android APK and iOS Ipa files! Source: https://github.com/bzon/gota

Usage:
  gota [flags]
  gota [command]

Available Commands:
  help        Help about any command
  nexus       Upload your apk or ipa file and create an over-the-air static site in a Nexus Site repository

Flags:
      --buildNumber string     the apk or ipa build number.
      --bundleID string        if srcFile type is '.ipa', this is required. (example: com.example.bundleid)
      --bundleVersion string   if srcFile type is '.ipa', this is required.
      --destDir string         root directory of the site to create.
  -h, --help                   help for gota
      --srcFile string         the apk or ipa file.
      --title string           application name to be displayed in the site
      --versionCode string     if srcfile is '.apk', this is required.
      --versionName string     if srcFile is '.apk', this is required.

Use "gota [command] --help" for more information about a command.
```

Nexus command help `gota nexus --help`

```bash
Upload your apk or ipa file and create an over-the-air static site in a Nexus Site repository

Usage:
  gota nexus [flags]

Flags:
  -h, --help                   help for nexus
      --nexusHost string       nexus host url (including http protocol)
      --nexusPassword string   nexus password (can be passed as env variable $NEXUS_PASSWORD)
      --nexusRepo string       nexus site repository id (nexus v3 raw repository not maven!)
      --nexusUser string       nexus username (can be passed as env variable $NEXUS_USER)

Global Flags:
      --buildNumber string     the apk or ipa build number.
      --bundleID string        if srcFile type is '.ipa', this is required. (example: com.example.bundleid)
      --bundleVersion string   if srcFile type is '.ipa', this is required.
      --destDir string         root directory of the site to create.
      --srcFile string         the apk or ipa file.
      --title string           application name to be displayed in the site
      --versionCode string     if srcfile is '.apk', this is required.
      --versionName string     if srcFile is '.apk', this is required.
```

### Nexus APK Upload

Upload an APK file to a Nexus Site Repository

```bash
./gota nexus --nexusHost http://localhost:8081 \
            --nexusRepo site \
            --nexusUser admin \
            --nexusPassword admin123 \
            --destDir android \
            --buildNumber 1 \
            --srcFile pkg/resources/DarkSouls.apk \
            --title "DarkSouls" \
            --versionName "1.0.0" \
            --versionCode "10222333"

uploaded to nexus: http://localhost:8081/repository/site/android/version.json
uploaded to nexus: http://localhost:8081/repository/site/android/1.0.0.10222333/index.html
uploaded to nexus: http://localhost:8081/repository/site/android/1.0.0.10222333/DarkSouls.apk
```

You should now be able to install the APK file from your Android Phone by accessing the http://localhost:8081/repository/site/android/1.0.0.10222333/index.html URL.

Uploaded site structure

![](./docs/apk_nexus_uploaded.png)

### Nexus IPA Upload

Upload an IPA file to a Nexus Site Repository

```bash
./gota nexus --nexusHost http://localhost:8081 \
            --nexusRepo site \
            --nexusUser admin \
            --nexusPassword admin123 \
            --destDir ios \
            --buildNumber 1 \
            --srcFile pkg/resources/DarkSouls.ipa \
            --title DarkSouls \
            --bundleVersion 1.0.0 \
            --bundleID com.example.com

uploaded to nexus: http://localhost:8081/repository/site/ios/version.json
uploaded to nexus: http://localhost:8081/repository/site/ios/1.0.0.1/DarkSouls.plist
uploaded to nexus: http://localhost:8081/repository/site/ios/1.0.0.1/index.html
uploaded to nexus: http://localhost:8081/repository/site/ios/1.0.0.1/DarkSouls.ipa
```

You should now be able to install the IPA file from your iPhone by accessing the http://localhost:8081/repository/site/ios/1.0.0.1/index.html URL.

Uploaded site structure

![](./docs/ios_nexus_uploaded.png)

## Development Setup

### Build and Test

```bash
go get -v ./...
go test -v ./...
```

If you are on Windows, ensure to go get spf13/cobra's dependency for it. 

```bash
GOOS=windows go get -v -u github.com/spf13/cobra
go get -v ./...
go test -v ./...
```

### Nexus Feature Test

You must have a Nexus 3 server running in your machine. 

Get it easily via docker command: `docker run -d -p 8081:8081 --name nexus3`.

Then, create a Raw Repository with a repository id `site`.

### S3 Feature Test

WIP

## Inspirations and References

* [Fastlane Nexus Upload](https://docs.fastlane.tools/actions/nexus_upload/)
* [Fastlane S3 Plugin](https://github.com/joshdholtz/fastlane-plugin-s3/)
* [Creating an Installation Link for your enterprise App](https://support.magplus.com/hc/en-us/articles/203808598-iOS-Creating-an-Installation-Link-for-Your-Enterprise-App)
