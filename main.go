package main

import "C"
import (
	"github.com/gin-gonic/gin"
	Cfg "gravity-daemon/config"
	"gravity-daemon/controllers"
	"gravity-daemon/utils"
	"log"
	"os"
)

var config Cfg.Config
var router *gin.Engine

func main() {

	if _, err := os.Stat("./config/config.json"); os.IsNotExist(err) {
		err = utils.Copy("./config/config.dist.json", "./config/config.json")
		if err != nil {
			log.Panic("An error as occurred while setping config file : ", err.Error())
		}
	}
	config = Cfg.LoadConfig()

	utils.SetConfig(config)
	controllers.SetConfig(config)

	router = LoadRouter()

	if _, err := os.Stat(config.Daemon.Root); os.IsNotExist(err) {
		err = os.Mkdir(config.Daemon.Root, 0600)
		if err != nil {
			log.Panic("An error as occurred creating server folder : ", err.Error())
		}
	}

	if config.Daemon.Tls.Active {
		err := router.RunTLS(config.Daemon.Listen, config.Daemon.Tls.Cert, config.Daemon.Tls.Key)

		if err != nil {
			log.Panic("An error as occurred while starting up TLS server : ", err.Error())
		}
	} else {
		err := router.Run(config.Daemon.Listen)

		if err != nil {
			log.Panic("An error as occurred while starting up server : ", err.Error())
		}
	}
}
