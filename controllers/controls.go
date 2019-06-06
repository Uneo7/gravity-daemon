package controllers

import (
	"github.com/gin-gonic/gin"
	"gravity-daemon/controllers/controls"
	"gravity-daemon/utils"
	"net/http"
	"strings"
)

func getServer(c *gin.Context) (server utils.Server) {

	sid := c.Param("id")
	uid := c.Param("user")

	server.Load(uid, sid)

	if server.Sid == "" {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "Server not found",
		})

		return
	}

	return
}

func ControlsGet(c *gin.Context) {

	action := c.Param("action")

	server := getServer(c)
	if server.Sid == "" {
		return
	}

	switch strings.TrimLeft(action, "/") {
	case "start":
		controls.Start(server, c)

	case "stop":
		controls.Stop(server, c)

	case "pid":
		controls.Pid(&server, c)

	case "kill":
		controls.Kill(server, c)

	case "status":
		controls.Status(server, c)

	case "console":
		controls.Console(server, c)

	default:
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Invalid action",
		})

		return
	}
}

func ControlsPost(c *gin.Context) {

	action := c.Param("action")

	server := getServer(c)
	if server.Sid == "" {
		return
	}

	switch strings.TrimLeft(action, "/") {
	case "command":
		controls.Command(server, c)

	case "download":
		controls.Download(server, c)

	case "delete":
		controls.Destroy(server, c)

	default:
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Invalid action",
		})

		return
	}
}
