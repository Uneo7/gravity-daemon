package controls

import (
	"github.com/gin-gonic/gin"
	"gravity-daemon/utils"
	"io/ioutil"
	"path/filepath"
)

func Console(server utils.Server, c *gin.Context) {

	path := filepath.Join(server.Path, server.Game.Logs.Location)
	data, err := ioutil.ReadFile(path)

	if err != nil {
		c.JSON(200, gin.H{
			"success": false,
			"message": "An error as occurred while reading server console",
		})
		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"console": string(data),
	})
}
