package controllers

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"strconv"
)

func ServerCreate(c *gin.Context) {
	c.JSON(200, gin.H{
		"success": true,
		"message": "Daemon running",
	})
}

func UserCreate(c *gin.Context) {

	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(c.Request.Body)

	if err != nil {
		log.Println(err)
		c.JSON(200, gin.H{
			"success": false,
			"message": "Error while parsing the params",
		})

		return
	}

	var data map[string]interface{}
	err = json.Unmarshal([]byte(buf.String()), &data)

	if _, ok := data["username"]; !ok {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Invalid parameters",
		})

		return
	}

	rootPath := filepath.Join(config.Daemon.Root, data["username"].(string))
	err = os.MkdirAll(rootPath, 0700)

	path := filepath.Join(rootPath, "servers")
	err = os.MkdirAll(path, 0700)

	path = filepath.Join(rootPath, "backups")
	err = os.MkdirAll(path, 0700)

	cmd := exec.Command("/usr/sbin/useradd", "-d", rootPath, "-s /sbin/nologin", data["username"].(string))
	_, err = cmd.CombinedOutput()
	log.Println(cmd.Args)
	log.Println(err)

	id, err := user.Lookup(data["username"].(string))

	uid, err := strconv.Atoi(id.Uid)
	gid, err := strconv.Atoi(id.Gid)

	err = os.Chown(path, uid, gid)

	if err != nil {
		c.JSON(200, gin.H{
			"success": false,
			"message": "Unable to create account",
		})

		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"message": "User created",
	})
}
