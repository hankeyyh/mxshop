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

type Client struct {
	UserSrv UserSrvConfig `json:"user-srv" mapstructure:"user-srv"`
}

type UserSrvConfig struct {
	Host string `json:"host" mapstructure:"host" json:"host"`
	Port int    `json:"port" mapstructure:"port" json:"port"`
	Name string `json:"name" mapstructure:"name"`
}

type LogConfig struct {
	Level    string `json:"level" mapstructure:"level"`
	FilePath string `json:"file_path" mapstructure:"file_path"`
}

type JWTConfig struct {
	SigningKey string `json:"key" mapstructure:"key"`
}

type AliSmsConfig struct {
	ApiKey     string `json:"key" mapstructure:"key"`
	ApiSecrect string `json:"secrect" mapstructure:"secrect"`
}

type RedisConfig struct {
	Host   string `json:"host" mapstructure:"host"`
	Port   int    `json:"port" mapstructure:"port"`
	Expire int    `json:"expire" mapstructure:"expire"`
}

type ServiceConfig struct {
	ServiceName string   `json:"service_name" mapstructure:"service_name"`
	ServiceTags []string `json:"service_tags" mapstructure:"service_tags"`
	Host        string   `json:"host" mapstructure:"host"`
	Port        int      `json:"port" mapstructure:"port"`
}

type ConsulConfig struct {
	Host string `json:"host" mapstructure:"host"`
	Port int    `json:"port" mapstructure:"port"`
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

type ServerConfig struct {
	Service    ServiceConfig `json:"service" mapstructure:"service"`
	Client     Client        `json:"client" mapstructure:"client"`
	Log        LogConfig     `json:"log" mapstructure:"log"`
	JWTInfo    JWTConfig     `json:"jwt" mapstructure:"jwt"`
	AliSmsInfo AliSmsConfig  `json:"sms" mapstructure:"sms"`
	RedisInfo  RedisConfig   `json:"redis" mapstructure:"redis"`
	Consul     ConsulConfig  `json:"consul" mapstructure:"consul"`
}

var (
	DefaultConfig ServerConfig
	nacosConfig   NacosConfig
)

func Init() error {
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
	if err := Init(); err != nil {
		panic(err)
	}
}
