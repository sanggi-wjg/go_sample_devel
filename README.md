## Develop GoLang Samples
간단한 예제들 구현한 모음집

### 개발환경
```shell
Windows + GoLand
Go 1.7
Gin Framework
(https://github.com/gin-gonic/gin)
```

### Samples
* route / youtube
```shell
API Docs : https://developers.google.com/youtube/v3/docs?hl=ko
API KEY : https://console.cloud.google.com/apis

# Go Test
# Linux 환경
export YOUTUBE_KEY="xxxx"
export YOUTUBE_CHANNEL_ID="xxxx"

# Windows 환경
Run > Edit Configurations
Environment variables 추가
```

* gin - swagger 연동  
http://localhost:9091/swagger/index.html 접속하면 됨
```shell
https://github.com/swaggo/gin-swagger

# Windows 환경
export PATH=$(go env GOPATH)/bin:$PATH
```
* route / benchmark

* route / scrap
```shell
https://github.com/PuerkitoBio/goquery
```