package controls

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"gravity-daemon/config"
	"gravity-daemon/utils"
	"os"
)

func Destroy(server utils.Server, c *gin.Context) {

	server = server.GetPID()

	if server.Pid != 0 {
		status := server.Kill()

		if !status {
			c.JSON(200, gin.H{
				"success": false,
				"message": "Kill failed (check console for more details)",
			})

			return
		}
	}

	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(c.Request.Body)

	if err != nil {
		c.JSON(200, gin.H{
			"success": false,
			"message": "Error while parsing the params",
		})

		return
	}

	var data map[string]interface{}
	err = json.Unmarshal([]byte(buf.String()), &data)

	err = os.RemoveAll(server.Path)
	if err != nil {
		c.JSON(200, gin.H{
			"success": false,
			"message": "Failed to remove server",
		})

		return
	}

	err = config.DeleteServerConfig(server.Sid)

	if err != nil {
		c.JSON(200, gin.H{
			"success": false,
			"message": "Failed to remove config",
		})

		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"message": "Server removed",
	})
}
