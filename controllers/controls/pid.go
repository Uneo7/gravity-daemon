package controls

import (
	"github.com/gin-gonic/gin"
	"gravity-daemon/utils"
)

func Pid(server *utils.Server, c *gin.Context) {

	server.GetPID()

	c.JSON(200, gin.H{
		"success": true,
		"pid":     server.Pid,
	})
}
