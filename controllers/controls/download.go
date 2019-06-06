package controls

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"gravity-daemon/utils"
)

func Download(server utils.Server, c *gin.Context) {

	server = server.GetPID()

	if server.Pid != 0 {
		c.JSON(200, gin.H{
			"success": false,
			"message": "Can't download the game, server is running",
		})

		return
	}

	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(c.Request.Body)

	if err != nil {
		c.JSON(200, gin.H{
			"success": false,
			"message": "Error while parsing the URL",
		})

		return
	}

	var data map[string]interface{}
	err = json.Unmarshal([]byte(buf.String()), &data)
	go utils.Download(server, data["name"].(string), data["url"].(string))

	c.JSON(200, gin.H{
		"success": true,
		"message": "Download in progress",
	})
}
