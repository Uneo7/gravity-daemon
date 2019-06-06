package controllers

import "github.com/gin-gonic/gin"

func ServerCreate(c *gin.Context) {
	c.JSON(200, gin.H{
		"success": true,
		"message": "Daemon running",
	})
}

func UserCreate(c *gin.Context) {
	c.JSON(200, gin.H{
		"success": true,
		"message": "Daemon running",
	})
}
