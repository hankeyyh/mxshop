package config

import (
	"encoding/json"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"github.com/spf13/viper"
)

type dbConfig struct {
	Mysql db `json:"mysql" mapstructure:"mysql"`
}

type db struct {
	DBName   string `json:"db_name" mapstructure:"db_name"`
	UserName string `json:"user_name" mapstructure:"user_name"`
	Password string `json:"password" mapstructure:"password"`
	Host     string `json:"host" mapstructure:"host"`
	Port     int    `json:"port" mapstructure:"port"`
}

type logConfig struct {
	Level    string `json:"level" mapstructure:"level"`
	FilePath string `json:"file_path" mapstructure:"file_path"`
}

type consulConfig struct {
	Host string `json:"host" mapstructure:"host"`
	Port int    `json:"port" mapstructure:"port"`
}

type serviceConfig struct {
	ServiceName string   `json:"service_name" mapstructure:"service_name"`
	ServiceTags []string `json:"service_tags" mapstructure:"service_tags"`
	Host        string   `json:"host" mapstructure:"host"`
	Port        int      `json:"port" mapstructure:"port"`
}

type Config struct {
	Db      dbConfig      `json:"db" mapstructure:"db"`
	Log     logConfig     `json:"log" mapstructure:"log"`
	Consul  consulConfig  `json:"consul" mapstructure:"consul"`
	Service serviceConfig `json:"service" mapstructure:"service"`
}

type NacosConfig struct {
	NamespaceId string `mapstructure:"namespace_id"`
	Timeout     uint64 `mapstructure:"timeout"`
	LogDir      string `mapstructure:"log_dir"`
	CacheDir    string `mapstructure:"cache_dir"`
	LogLevel    string `mapstructure:"log_level"`
	IpAddr      string `mapstructure:"ip_addr"`
	Port        uint64 `mapstructure:"port"`
	DataId      string `mapstructure:"data_id"`
	Group       string `mapstructure:"group"`
}

var (
	DefaultConfig Config
	nacosConfig   NacosConfig
)

func Init() (err error) {
	configFilePath := "config-pro.toml"
	if IsDebug() {
		configFilePath = "config-debug.toml"
	}
	// viper read config
	viper.SetConfigFile(configFilePath)
	if err = viper.ReadInConfig(); err != nil {
		return
	}
	if err = viper.Unmarshal(&nacosConfig); err != nil {
		return
	}

	// nacos client
	clientConfig := constant.ClientConfig{
		NamespaceId: nacosConfig.NamespaceId,
		TimeoutMs:   nacosConfig.Timeout,
		LogDir:      nacosConfig.LogDir,
		CacheDir:    nacosConfig.CacheDir,
		LogLevel:    nacosConfig.LogLevel,
	}
	serverConfig := []constant.ServerConfig{
		{
			IpAddr: nacosConfig.IpAddr,
			Port:   nacosConfig.Port,
		},
	}
	nacosClient, err := clients.NewConfigClient(vo.NacosClientParam{
		ClientConfig:  &clientConfig,
		ServerConfigs: serverConfig,
	})
	content, err := nacosClient.GetConfig(vo.ConfigParam{
		DataId: nacosConfig.DataId,
		Group:  nacosConfig.Group,
	})
	err = json.Unmarshal([]byte(content), &DefaultConfig)

	return err
}

func GetEnv(key string) bool {
	viper.AutomaticEnv()
	return viper.GetBool(key)
}

func IsDebug() bool {
	return GetEnv("MXSHOP_DEBUG")
}

func init() {
	Init()
}
