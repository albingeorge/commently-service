package router

import (
	"net/http"

	"github.com/albingeorge/commently-service/controllers"
	"github.com/gin-gonic/gin"
)

var db = make(map[string]string)

func SetupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// Get user value
	r.GET("/user/:name", func(c *gin.Context) {
		user := c.Params.ByName("name")
		value, ok := db[user]
		if ok {
			c.JSON(http.StatusOK, gin.H{"user": user, "value": value})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
		}
	})

	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
		"albin": "temp123", // user:foo password:bar
	}))

	authorized.POST("comment", controllers.CommentCreate)
	authorized.GET("comment", controllers.CommentFetch)
	authorized.GET("urls", controllers.UrlsFetch)

	return r
}
