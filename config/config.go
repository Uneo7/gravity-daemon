package config

import (
	"crypto/rsa"
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"io/ioutil"
	"log"
	"path"
)

type Config struct {
	Path   string
	Daemon Daemon `json:"daemon"`
	Panel  Panel  `json:"panel"`
}

type Daemon struct {
	Listen string `json:"listen"`
	Root   string `json:"root"`
	Tls    struct {
		Active bool   `json:"active"`
		Cert   string `json:"cert"`
		Key    string `json:"key"`
	} `json:"tls"`
}

type Panel struct {
	Address string `json:"address"`
	Key     string `json:"key"`
	PKey    *rsa.PublicKey
}

var config Config

func LoadConfig(config *Config) {

	cPath := path.Join(config.Path, "config.json")
	data, err := ioutil.ReadFile(cPath)

	if err != nil {
		log.Panic("An error as occurred while reading config : ", err.Error())
	}

	err = json.Unmarshal([]byte(data), &config)

	if err != nil {
		log.Panic("An error as occurred while parsing config : ", err.Error())
	}

	kPath := path.Join(config.Path, config.Panel.Key)
	pem, err := ioutil.ReadFile(kPath)

	if err != nil {
		log.Panic("An error as occurred while opening JWT key file : ", err.Error())
	}

	key, err := jwt.ParseRSAPublicKeyFromPEM(pem)

	if err != nil {
		log.Panic("An error as occurred while checking JWT key file : ", err.Error())
	}

	config.Panel.PKey = key
}
