package controls

import (
	"github.com/gin-gonic/gin"
	"gravity-daemon/utils"
)

func Stop(server utils.Server, c *gin.Context) {

	server = server.GetPID()

	if server.Pid == 0 {
		c.JSON(200, gin.H{
			"success": false,
			"message": "Server already stopped",
		})

		return
	}

	status := server.Stop()

	if !status {
		c.JSON(200, gin.H{
			"success": false,
			"message": "Stop failed (check console for more details)",
		})

		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"message": "Server stopping",
	})
}
