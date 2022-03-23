## Develop GoLang Samples
간단한 예제들 구현한 모음집


### 개발환경
```shell
Windows + GoLand
Go 1.8
Gin Framework
(https://github.com/gin-gonic/gin)
GORM
(https://gorm.io/docs/)
```

### ENV
```shell
YOUTUBE_KEY
YOUTUBE_CHANNEL_ID

DB_USER
DB_PASSWORD
DB_HOST
DB_PORT
DB_DATABASE_NAME
```

### Samples
<details>
<summary>Youtube API</summary>

#### route / youtube
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
</details>



<details>
<summary>gin - swagger 연동</summary>

http://localhost:9091/swagger/index.html 접속하면 됨
```shell
https://github.com/swaggo/gin-swagger

# Windows 환경
export PATH=$(go env GOPATH)/bin:$PATH
```
</details>



<details>
<summary>Benchmark Test</summary>

#### route / benchmark
</details>



<details>
<summary>Scrap (goQuery)</summary>

#### route / scrap
```shell
https://github.com/PuerkitoBio/goquery
```
</details>
