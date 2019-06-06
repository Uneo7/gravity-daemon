package main

import (
	"github.com/gin-gonic/gin"
	"gravity-daemon/controllers"
)

func LoadRouter() (router *gin.Engine) {
	router = gin.Default()

	router.GET("/", controllers.Index)

	api := router.Group("/api")

	//api.Use(middleware.GetAuth(config.Panel.PKey))
	{
		api.POST("/server/create", controllers.ServerCreate)
		api.POST("/user/create", controllers.UserCreate)

		api.GET("/:user/server/:id/*action", controllers.ControlsGet)
		api.POST("/:user/server/:id/*action", controllers.ControlsPost)

	}

	return
}
