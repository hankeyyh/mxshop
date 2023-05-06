package config

import (
	"encoding/json"
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/config_client"
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
	if err := v.Unmarshal(&nacosConfig); err != nil {
		return err
	}

	// nacos 使用1.x版本，避免鉴权配置
	nacosClient, err := newNacosClient(nacosConfig)
	if err != nil {
		return err
	}
	data, err := nacosClient.GetConfig(vo.ConfigParam{
		DataId: nacosConfig.DataId,
		Group:  nacosConfig.Group,
	})
	if err != nil {
		return err
	}
	err = json.Unmarshal([]byte(data), &DefaultConfig)
	if err != nil {
		return err
	}
	return nil
}

func GetEnvInfo(env string) bool {
	viper.AutomaticEnv()
	return viper.GetBool(env)
}

// IsDebug 是否测试环境
func IsDebug() bool {
	return GetEnvInfo("MXSHOP_DEBUG")
}

func newNacosClient(conf NacosConfig) (client config_client.IConfigClient, err error) {
	clientConfig := *constant.NewClientConfig(
		constant.WithNamespaceId(conf.NamespaceId), //When namespace is public, fill in the blank string here.
		constant.WithTimeoutMs(conf.Timeout),
		constant.WithNotLoadCacheAtStart(true),
		constant.WithLogDir(conf.LogDir),
		constant.WithCacheDir(conf.CacheDir),
		constant.WithLogLevel(conf.LogLevel),
	)

	serverConfigList := []constant.ServerConfig{
		*constant.NewServerConfig(conf.IpAddr, conf.Port),
	}

	return clients.NewConfigClient(vo.NacosClientParam{
		ClientConfig:  &clientConfig,
		ServerConfigs: serverConfigList,
	})
}

func init() {
	if err := initConfig(); err != nil {
		panic(err)
	}
}
