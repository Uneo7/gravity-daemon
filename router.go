package main

import (
	"daemon/controllers"
	"daemon/middleware"
	"github.com/gin-gonic/gin"
)

func LoadRouter() (router *gin.Engine) {
	router = gin.Default()

	router.GET("/", controllers.Index)
	router.GET("/test", middleware.GetAuth(config.Panel.PKey), controllers.Index)

	return
}
