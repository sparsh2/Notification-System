package web

import (
	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	authG := r.Group("/api/v1/")
	authG.Use(AuthMiddleware)
	authG.POST("/set-preference", setPreference)

	return r
}

func setPreference(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}