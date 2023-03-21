package config

import "os"

var (
	Str string
)

func initConfig() {
	content, err := os.ReadFile("config/conf.toml")
	if err != nil {
		panic(err)
	}
	Str = string(content)
}

func init() {
	initConfig()
}
