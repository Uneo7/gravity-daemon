package controllers

import "github.com/gin-gonic/gin"

func Start (c *gin.Context) {
	c.JSON(200, gin.H{
		"success": true,
		"message": "Daemon running",
	})
}
