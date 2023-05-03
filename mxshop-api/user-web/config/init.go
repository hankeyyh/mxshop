package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Client struct {
	UserSrv UserSrvConfig `mapstructure:"user-srv"`
}

type UserSrvConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port int    `mapstructure:"port" json:"port"`
}

type LogConfig struct {
	Level    string `mapstructure:"level"`
	FilePath string `mapstructure:"file_path"`
}

type JWTConfig struct {
	SigningKey string `mapstructure:"key"`
}

type AliSmsConfig struct {
	ApiKey     string `mapstructure:"key"`
	ApiSecrect string `mapstructure:"secrect"`
}

type RedisConfig struct {
	Host   string `mapstructure:"host"`
	Port   int    `mapstructure:"port"`
	Expire int    `mapstructure:"expire"`
}

type ServiceConfig struct {
	ServiceName string   `mapstructure:"service_name"`
	ServiceTags []string `mapstructure:"service_tags"`
	Port        int      `mapstructure:"port"`
}

type ConsulConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

type ServerConfig struct {
	Service    ServiceConfig `mapstructure:"service"`
	Client     Client        `mapstructure:"client"`
	Log        LogConfig     `mapstructure:"log"`
	JWTInfo    JWTConfig     `mapstructure:"jwt"`
	AliSmsInfo AliSmsConfig  `mapstructure:"sms"`
	RedisInfo  RedisConfig   `mapstructure:"redis"`
	Consul     ConsulConfig  `mapstructure:"consul"`
}

var defaultConfig ServerConfig

func DefaultConfig() ServerConfig {
	return defaultConfig
}

func Init() error {
	debug := GetEnvInfo("MXSHOP_DEBUG")
	configPrefix := "config"
	configFileName := fmt.Sprintf("user-web/%s-pro.toml", configPrefix)
	if debug {
		configFileName = fmt.Sprintf("user-web/%s-debug.toml", configPrefix)
	}

	v := viper.New()
	v.SetConfigFile(configFileName)
	if err := v.ReadInConfig(); err != nil {
		return err
	}

	if err := v.Unmarshal(&defaultConfig); err != nil {
		return err
	}
	return nil
}

func GetEnvInfo(env string) bool {
	viper.AutomaticEnv()
	return viper.GetBool(env)
}