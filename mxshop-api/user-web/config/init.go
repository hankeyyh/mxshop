package config

import (
	"github.com/BurntSushi/toml"
)

type logConfig struct {
	Level    string `toml:"level"`
	FilePath string `toml:"file_path"`
}

type config struct {
	Log logConfig `toml:"log"`
}

var (
	Conf config
)

func initConfig() {
	_, err := toml.DecodeFile("config/conf.toml", &Conf)
	if err != nil {
		panic(err)
	}
}

func init() {
	initConfig()
}
