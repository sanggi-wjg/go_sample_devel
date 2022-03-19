package route

import (
	"github.com/gin-gonic/gin"
	"go_sample_devel/route/home"
	"go_sample_devel/route/youtube"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	r.GET("/", home.Home)
	r.GET("/ping", home.Ping)

	r.GET("/youtube/subs/count", youtube.GetYoutubeChannelStat)

	return r
}
