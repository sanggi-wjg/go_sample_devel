package home

import "github.com/gin-gonic/gin"

func Home(c *gin.Context) {
	c.String(200, "GoLang sample develop")
}

func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
