package web

import (
	"database/sql"
	"encoding/json"
	"io"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"github.com/sparsh2/notification-system/src/components/notification-server/config"
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
	// c.JSON(200, gin.H{
	// 	"message": "pong",
	// })
	cfg := mysql.Config{
		User:   config.Configs.DBConfig.User,
		Passwd: config.Configs.DBConfig.Password,
		Net:    "tcp",
		Addr:   config.Configs.DBConfig.Host,
		DBName: config.Configs.DBConfig.DBName,
	}
	// Get a database handle.
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		c.JSON(500, gin.H{
			"error": "Internal Server Error",
			"msg":   err.Error(),
		})
		return
	}
	defer db.Close()
	bytes, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Bad Request",
			"msg":   err.Error(),
		})
		return
	}
	var preference map[string]string
	err = json.Unmarshal(bytes, &preference)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Bad Request",
			"msg":   err.Error(),
		})
		return
	}
	db.Exec("INSERT INTO user_preferences (user_id, service_id, preferences) VALUES (?, ?, ?)", c.GetString("account"), c.PostForm("preference"))
}
