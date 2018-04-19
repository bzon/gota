# GOTA

Over the Air installation for Android APK and iOS Ipa files!

![](./assets/gota_html.png)

Upload to Android APK and iOS IPA files to Nexus, AWS S3, and more.

## Command guide 

Upload to Nexus.

```bash
gota nexus -f MyApp.ipa -r http://localhost:8081/repository/maven-releases -g com.example -a MyApp 
```

Upload to AWS S3 bucket.

```bash
gota s3 -f MyApp.ipa -b s3://mysite.s3bucket.aws.amazon.com -d ios 
```
