package utils

import Cfg "gravity-daemon/config"

var config Cfg.Config

func SetConfig(c Cfg.Config) {
	config = c
}
