package main

import "C"
import (
	"flag"
	"github.com/gin-gonic/gin"
	Cfg "gravity-daemon/config"
	"gravity-daemon/controllers"
	"gravity-daemon/utils"
	"log"
	"os"
	"path"
)

var config Cfg.Config
var router *gin.Engine

func main() {

	configPath := flag.String("c", "", "Config path")
	flag.Parse()

	if *configPath == "" {
		log.Panic("Config path not found")
	}

	config.Path = *configPath
	cPath := path.Join(config.Path, "config.json")

	log.Println("Loading configuration from : " + cPath)

	if _, err := os.Stat(cPath); os.IsNotExist(err) {
		log.Panic("Config not found")
	}

	Cfg.LoadConfig(&config)

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
