package home

import "github.com/gin-gonic/gin"

// @BasePath /api/v1

// Home PingExample godoc
// @Summary Home
// @Schemes
// @Description home index
// @Tags Basic
// @Accept json
// @Produce json
// @Success 200 {string} GoLang sample develop
// @Router / [get]
func Home(c *gin.Context) {
	c.String(200, "GoLang sample develop")
}

func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
