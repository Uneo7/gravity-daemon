package utils

import (
	"io/ioutil"
	"log"
	"path/filepath"
	"regexp"
	"strings"
)

func parseStartCommand(server Server) string {

	command := server.Game.Commands.Start
	matcher := regexp.MustCompile(`\{\{([A-Za-z0-9\_\-]+)\}\}`)

	params := matcher.FindAllString(command, -1)

	for _, param := range params {

		cmd := strings.TrimLeft(param, "{{")
		cmd = strings.TrimRight(cmd, "}}")

		for _, value := range server.Server.Params {
			if value.Name == cmd {
				replacer := regexp.MustCompile(param)
				command = replacer.ReplaceAllString(command, value.Value)

				break
			}
		}
	}

	return command
}

func ParseLog(server Server) string {

	path := filepath.Join(server.Path, server.Game.Logs.Location)
	data, err := ioutil.ReadFile(path)

	if err != nil {
		log.Println("An error as occurred while reading server console : ", err.Error())
		return ""
	}

	eulaMatcher := regexp.MustCompile(server.Game.Logs.Eula)
	if matches := eulaMatcher.Find(data); len(matches) > 0 {
		return "eula"
	}

	server.GetPID()
	if server.Pid == 0 {
		return "stopped"
	}

	startedMatcher := regexp.MustCompile(server.Game.Logs.Started)
	if matches := startedMatcher.Find(data); len(matches) > 0 {
		return "started"
	}

	return "starting"
}
