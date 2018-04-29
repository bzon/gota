# GOTA

[![Go Report Card](https://goreportcard.com/badge/github.com/bzon/gota)](https://goreportcard.com/report/github.com/bzon/gota)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/bzon/gota)
[![GitHub tag](https://img.shields.io/github/tag/bzon/gota.svg)](https://github.com/bzon/gota/releases/)

Automate the beta testing distribution of your Android and iOS application files with gota.

Gota is a [Golang](http://golang.org/) powered Over the Air Installation site creation tool.

![](./docs/gota_workflow.png)

## Feature Checklist

* [x] Upload and generate site to a Nexus 3 Site Repository
* [x] Upload and generate site to an Amazon S3 bucket
* [ ] Upload and generate site to a Nexus 2 Site Repository

## Installation

Get the executable binary for your platform from the [Release Page](https://github.com/bzon/gota/releases/)

If you have Go installed, just run `go get github.com/bzon/gota`.

## User Guide

To see the required flags, use the --help flag.

```bash
gota --help
gota nexus --help
gota s3 --help
```

Gota creates a `gotalink.txt` and `ipalink.txt` (if uploading an ipa) that contains the url or direct download link.

If you are using a CI server, you can have it read these files for quickly getting the url that you can send to your team.

### Upload to S3 Bucket


```bash
# set the aws credentials
export AWS_ACCESS_KEY=xxxxx
export AWS_SECRET_ACCESS_KEY=xxxxx

./gota s3 --bucket example-s3-bucket --srcFile sample.ipa --destDir ios_bucket

2018/04/30 01:12:37 file uploaded: https://example-s3-bucket.s3.amazonaws.com/ios_bucket/1.0.0/4/appicon.png
2018/04/30 01:12:37 file uploaded: https://example-s3-bucket.s3.amazonaws.com/ios_bucket/1.0.0/version.json
2018/04/30 01:12:37 file uploaded: https://example-s3-bucket.s3.amazonaws.com/ios_bucket/1.0.0/4/index.html
2018/04/30 01:12:37 file uploaded: https://example-s3-bucket.s3.amazonaws.com/ios_bucket/1.0.0/4/sample.ipa
2018/04/30 01:12:37 file uploaded: https://example-s3-bucket.s3.amazonaws.com/ios_bucket/1.0.0/4/app.plist
```

__NOTE__: Currently, gota assigns a AES256 encrpytion and a public-read ACL to all files that are uploaded.
This may change to be configurable in the future.

### Upload to Nexus

The repository must be a [Raw Site Repository](https://help.sonatype.com/repomanager3/raw-repositories-and-maven-sites).

```bash
# set the nexus credentials
# this can also be set via command flags
export NEXUS_USER=admin
export NEXUS_PASSWORD=admin123

./gota nexus --nexusHost http://localhost:8081 \
            --nexusRepo site \
            --destDir nexus_android_repo \
            --srcFile build/outpus/apk/sample.apk \

file uploaded: http://localhost:8081/repository/site/nexus_android_repo/1.0.0/10222333/appicon.png
file uploaded: http://localhost:8081/repository/site/nexus_android_repo/1.0.0/version.json
file uploaded: http://localhost:8081/repository/site/nexus_android_repo/1.0.0/10222333/index.html
file uploaded: http://localhost:8081/repository/site/nexus_android_repo/1.0.0/10222333/sample.apk
```

__NOTE__: Currently supports only Nexus 3.

### Site Directory Layout

```bash
destDir
\__(ipa CFBundleShortVersion or apk versionName)
   \__version.json
   \__(ipa CFBundleVersion or apk versionCode)
	 \__appicon.png
	 \__(ipa or apk file)
	 \__app.plist (if ipa file)
	 \__index.html
```

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

### S3 Feature Test

Set these environment variables before running `go test` in s3 package.

```bash
AWS_ACCESS_KEY=xxxxx
AWS_SECRET_ACCESS_KEY=xxxxx
GOTEST_AWS_BUCKET=example-bucket
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
