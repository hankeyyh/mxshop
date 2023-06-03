package register

import (
	"github.com/hankeyyh/mxshop/mxshop-srv/goods-srv/config"
)

var (
	DefaultServiceAgent ServiceAgent
)

type ServiceAgentConf struct {
	ServiceId   string
	ServiceName string
	ServiceHost string
	ServicePort int
	ServiceTags []string

	// 健康检查
	HealthCheckInterval string
	DeregisterCritical  string
}

type ServiceAgent interface {
	Register(conf ServiceAgentConf) error
	Deregister(serviceId string) error
}

func Init() error {
	consulConf := config.DefaultConfig.Consul
	var err error
	DefaultServiceAgent, err = NewConsulAgent(consulConf.Host, consulConf.Port)
	return err
}

func init() {
	Init()
}
