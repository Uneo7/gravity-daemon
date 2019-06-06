package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type ServerConfig struct {
	Sid string `json:"sid"`
	Uid string `json:"uid"`

	Game struct {
		Id    string `json:"id"`
		Port  int    `json:"port"`
		Ports []int  `json:"ports"`
	} `json:"game"`

	Startup struct {
		Start       bool `json:"start"`
		AutoRestart bool `json:"auto_restart"`
	} `json:"startup"`

	Params []struct {
		Name  string `json:"name"`
		Value string `json:"value"`
	} `json:"params"`
}

func LoadServerConfig(server string) (config ServerConfig) {
	data, err := ioutil.ReadFile("./config/servers/" + server + ".json")

	if err != nil {
		log.Println("An error as occurred while reading server config : ", err.Error())
	}

	err = json.Unmarshal([]byte(data), &config)

	if err != nil {
		log.Println("An error as occurred while parsing server config : ", err.Error())
	}

	return
}

func DeleteServerConfig(server string) error {
	return os.RemoveAll("./config/servers/" + server + ".json")
}
