package router

import (
	"github.com/albingeorge/commently-service/app/controllers"
	"github.com/gin-gonic/gin"
)

var db = make(map[string]string)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
		"albin": "temp123", // user:foo password:bar
	}))

	authorized.POST("comment", controllers.CommentCreate)
	authorized.GET("comment", controllers.CommentFetch)
	authorized.GET("urls", controllers.UrlsFetch)

	return r
}
