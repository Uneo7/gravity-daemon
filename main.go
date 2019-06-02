package main

import (
	"github.com/gin-gonic/gin"
	"log"
)
var config Config
var router *gin.Engine

func main() {

	config = LoadConfig()
	router = LoadRouter()

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