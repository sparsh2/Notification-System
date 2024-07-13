package web

import (
	"encoding/json"
	"io"

	"github.com/gin-gonic/gin"
	"github.com/sparsh2/notification-system/src/components/notification-server/services"
	"github.com/sparsh2/notification-system/src/components/notification-server/types"
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
	bytes, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Bad Request",
			"msg":   err.Error(),
		})
		return
	}
	preferencesReq := &types.SetPreferenceRequest{}
	err = json.Unmarshal(bytes, preferencesReq)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Bad Request",
			"msg":   err.Error(),
		})
		return
	}
	preferencesReq.ServiceID = c.GetString("service_id")
	err = services.Service.SetPreference(preferencesReq)
	if err != nil {
		c.JSON(500, gin.H{
			"error": "Internal Server Error",
			"msg":   err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "Preference set successfully",
	})
}
