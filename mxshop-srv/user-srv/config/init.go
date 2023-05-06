package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type dbConfig struct {
	Mysql db `mapstructure:"mysql"`
}

type db struct {
	DBName   string `mapstructure:"db_name"`
	UserName string `mapstructure:"user_name"`
	Password string `mapstructure:"password"`
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
}

type logConfig struct {
	Level    string `mapstructure:"level"`
	FilePath string `mapstructure:"file_path"`
}

type consulConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

type serviceConfig struct {
	ServiceName string   `mapstructure:"service_name"`
	ServiceTags []string `mapstructure:"service_tags"`
	Host        string   `mapstructure:"host"`
	Port        int      `mapstructure:"port"`
}

type Config struct {
	Db      dbConfig      `mapstructure:"db"`
	Log     logConfig     `mapstructure:"log"`
	Consul  consulConfig  `mapstructure:"consul"`
	Service serviceConfig `mapstructure:"service"`
}

var (
	DefaultConfig Config
)

func initConfig() error {
	debug := GetEnvInfo("MXSHOP_DEBUG")
	configPrefix := "config"
	configFileName := fmt.Sprintf("%s-pro.toml", configPrefix)
	if debug {
		configFileName = fmt.Sprintf("%s-debug.toml", configPrefix)
	}

	v := viper.New()
	v.SetConfigFile(configFileName)
	if err := v.ReadInConfig(); err != nil {
		return err
	}
	if err := v.Unmarshal(&DefaultConfig); err != nil {
		return err
	}
	return nil
}

func GetEnvInfo(env string) bool {
	viper.AutomaticEnv()
	return viper.GetBool(env)
}

// 是否测试环境
func IsDebug() bool {
	return GetEnvInfo("MXSHOP_DEBUG")
}

func init() {
	if err := initConfig(); err != nil {
		panic(err)
	}
}
