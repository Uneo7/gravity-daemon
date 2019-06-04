package controllers

import "github.com/gin-gonic/gin"

func Create (c *gin.Context) {
	c.JSON(200, gin.H{
		"success": true,
		"message": "Daemon running",
	})
}
