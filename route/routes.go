package route

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"go_sample_devel/docs"
	"go_sample_devel/route/home"
	"go_sample_devel/route/scrap"
	"go_sample_devel/route/youtube"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	// swagger
	docs.SwaggerInfo.BasePath = "/api/v1"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	v1 := r.Group("/api/v1")
	{
		eg := v1.Group("")
		{
			eg.GET("/", home.Home)
		}
	}

	// home
	r.GET("/ping", home.Ping)

	// scrap
	r.GET("/scrap/naver/news", scrap.GetNaverNews)
	r.GET("/scrap/naver/news/save", scrap.SaveNaverNews)

	// youtube
	r.GET("/youtube/channel/stat", youtube.GetYoutubeChannelStat)

	return r
}
