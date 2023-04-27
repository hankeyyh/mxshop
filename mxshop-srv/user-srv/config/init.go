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

type consulConfig struct {
	Host string `toml:"host"`
	Port int    `toml:"port"`
}

type serviceConfig struct {
	ServiceName string   `toml:"service_name"`
	ServiceTags []string `toml:"service_tags"`
	Port        int      `toml:"port"`
}

type Config struct {
	Db      map[string]dbConfig `toml:"db"`
	Log     logConfig           `toml:"log"`
	Consul  consulConfig        `toml:"consul"`
	Service serviceConfig       `toml:"service"`
}

var (
	conf Config
)

func initConfig() {
	_, err := toml.DecodeFile("config/conf.toml", &conf)
	if err != nil {
		panic(err)
	}
}

func DefaultConfig() Config {
	return conf
}

func init() {
	initConfig()
}
