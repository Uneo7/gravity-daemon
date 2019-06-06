package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"path"
)

type GameConfig struct {
	Name     string `json:"name"`
	Commands struct {
		Start string `json:"start"`
		Stop  string `json:"stop"`
	} `json:"commands"`

	Logs struct {
		Eula     string `json:"eula"`
		Started  string `json:"started"`
		Location string `json:"Location"`
	} `json:"logs"`

	Config struct {
		Location string `json:"location"`
		Replace  []struct {
			Source      string `json:"source"`
			Replacement string `json:"replacement"`
		}
	} `json:"config"`
}

func LoadGameConfig(game string) (gameConfig GameConfig) {
	cPath := path.Join(config.Path, game+".json")
	data, err := ioutil.ReadFile(cPath)

	if err != nil {
		log.Println("An error as occurred while reading game config : ", err.Error())
	}

	err = json.Unmarshal([]byte(data), &gameConfig)

	if err != nil {
		log.Println("An error as occurred while parsing game config : ", err.Error())
	}

	return
}
