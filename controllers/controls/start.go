package controls

import (
	"github.com/gin-gonic/gin"
	"gravity-daemon/utils"
)

func Start(server utils.Server, c *gin.Context) {

	server.GetPID()

	if server.Pid != 0 {
		c.JSON(200, gin.H{
			"success": false,
			"message": "Server already running",
			"pid":     server.Pid,
		})

		return
	}

	started := server.Start()

	if !started {
		c.JSON(200, gin.H{
			"success": false,
			"message": "Start failed (check console for more details)",
		})

		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"message": "Server starting",
		"pid":     server.Pid,
	})
}
