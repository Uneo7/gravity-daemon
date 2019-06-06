package utils

import (
	"log"
	"os"
	"path"
)

func LogToConsole(server Server, message string) {

	logfile := path.Join(server.Path, "output.log")

	f, err := os.OpenFile(logfile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		log.Println("Failed to open : ", logfile)
	}

	defer f.Close()

	if _, err = f.WriteString(message + "\n"); err != nil {
		log.Println("Failed to log data into : ", logfile)
	}
}
