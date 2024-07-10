package main

// import gin

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// Start the server
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}