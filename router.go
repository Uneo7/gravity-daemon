package main

import (
	"daemon/controllers"
	"daemon/middleware"
	"github.com/gin-gonic/gin"
)

func LoadRouter() (router *gin.Engine) {
	router = gin.Default()

	router.GET("/", controllers.Index)

	api := router.Group("/api")

	api.Use(middleware.GetAuth(config.Panel.PKey))
	{
		router.GET("/test", controllers.Index)
	}
	return
}
