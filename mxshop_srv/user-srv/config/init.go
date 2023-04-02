package config

import (
	"github.com/BurntSushi/toml"
)

type dbConfig struct {
	DBName   string `toml:"db_name"`
	UserName string `toml:"user_name"`
	Password string `toml:"password"`
	Host     string `toml:"host"`
	Port     int    `toml:"port"`
}

type logConfig struct {
	Level    string `toml:"level"`
	FilePath string `toml:"file_path"`
}

type config struct {
	Db  map[string]dbConfig `toml:"db"`
	Log logConfig           `toml:"log"`
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
