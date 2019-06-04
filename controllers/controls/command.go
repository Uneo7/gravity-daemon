package controls

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"gravity-daemon/utils"
)

func Command(server utils.Server, c *gin.Context) {

	server = server.GetPID()

	if server.Pid == 0 {
		c.JSON(200, gin.H{
			"success": false,
			"message": "Server not running",
		})

		return
	}

	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(c.Request.Body)

	if err != nil {
		c.JSON(200, gin.H{
			"success": false,
			"message": "Error while parsing command",
		})

		return
	}

	command := buf.String()

	status := server.Command(command)

	if !status {
		c.JSON(200, gin.H{
			"success": false,
			"message": "Send failed (check console for more details)",
		})

		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"message": "Command sentpo",
	})
}
