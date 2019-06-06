package controls

import (
	"github.com/gin-gonic/gin"
	"gravity-daemon/utils"
)

func Status(server utils.Server, c *gin.Context) {

	server.GetPID()
	status := utils.ParseLog(server)
	cpu, ram := utils.Resources(server)

	c.JSON(200, gin.H{
		"success": true,
		"status":  status,
		"cpu":     cpu,
		"ram":     ram,
	})
}
