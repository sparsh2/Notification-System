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
	authG.PUT("/set-preference", setPreference)
	authG.POST("/set-user-details", setUserDetails)
	authG.POST("/register-service", registerService)
	authG.POST("/send-notification", sendNotification)

	return r
}

func registerService(c *gin.Context) {
	// TODO: Implement later. Not needed for now
	c.JSON(200, gin.H{
		"message": "Service registered successfully",
	})
}

func setUserDetails(c *gin.Context) {
	bytes, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Bad Request",
			"msg":   err.Error(),
		})
		return
	}
	userDetailsReq := &types.SetUserDetailsRequest{}
	err = json.Unmarshal(bytes, userDetailsReq)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Bad Request",
			"msg":   err.Error(),
		})
		return
	}
	if userDetailsReq.Email == "" || userDetailsReq.UserID == "" {
		c.JSON(400, gin.H{
			"error": "Bad Request",
			"msg":   "Email and User ID is required",
		})
		return
	}
	userDetailsReq.ServiceID = c.GetString("account")
	err = services.Service.SetUserDetails(userDetailsReq)
	if err != nil {
		c.JSON(500, gin.H{
			"error": "Internal Server Error",
			"msg":   err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "User details set successfully",
	})
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
	preferencesReq.ServiceID = c.GetString("account")
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

func sendNotification(c *gin.Context) {
	bytes, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Bad Request",
			"msg":   err.Error(),
		})
		return
	}

	notificationReq := &types.NotificationRequest{}
	err = json.Unmarshal(bytes, notificationReq)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Bad Request",
			"msg":   err.Error(),
		})
		return
	}

	// Validate required fields
	if notificationReq.Title == "" || notificationReq.Message == "" || notificationReq.UserID == "" {
		c.JSON(400, gin.H{
			"error": "Bad Request",
			"msg":   "Title, Message and User ID are required",
		})
		return
	}

	// Set service ID from auth context
	notificationReq.ServiceID = c.GetString("account")

	// Send to Kafka
	err = services.Kafka.SendNotification(notificationReq)
	if err != nil {
		c.JSON(500, gin.H{
			"error": "Internal Server Error",
			"msg":   err.Error(),
		})
		return
	}

	c.JSON(200, types.NotificationResponse{
		Success: true,
	})
}
