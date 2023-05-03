package register

import (
	"context"
	"fmt"
	"github.com/hashicorp/consul/api"
	"mxshop-api/user-web/config"
	"mxshop-api/user-web/log"
)

type ConsulClient struct {
	Host string
	Port int
}

func InitConsulRegister() error {
	conf := config.DefaultConfig().Consul
	serviceConf := config.DefaultConfig().Service
	registry := ConsulClient{
		Host: conf.Host,
		Port: conf.Port,
	}
	err := registry.Register(serviceConf.ServiceName,
		serviceConf.ServiceName,
		serviceConf.ServiceTags,
		"host.docker.internal", // todo 如何放入配置
		serviceConf.Port,
	)
	if err != nil {
		log.Error(context.Background(), "consul.registry fail", log.Any("err", err))
		return err
	}
	return nil
}

func (r ConsulClient) Register(serviceName string, serviceId string, tags []string, address string, port int, options ...map[string]interface{}) error {
	consulConfig := api.DefaultConfig()
	consulConfig.Address = fmt.Sprintf("%s:%d", r.Host, r.Port)

	client, err := api.NewClient(consulConfig)
	if err != nil {
		return err
	}
	// 注册配置项
	registry := new(api.AgentServiceRegistration)
	registry.Name = serviceName
	registry.ID = serviceId
	registry.Address = address
	registry.Tags = tags
	registry.Port = port
	registry.Check = &api.AgentServiceCheck{
		HTTP:                           fmt.Sprintf("http://%s:%d/health", address, port),
		Interval:                       "5s",
		Timeout:                        "5s",
		DeregisterCriticalServiceAfter: "10s",
	}
	err = client.Agent().ServiceRegister(registry)
	return err
}

func (r ConsulClient) Deregister(serviceId string) error {
	consulConfig := api.DefaultConfig()
	consulConfig.Address = fmt.Sprintf("%s:%d", r.Host, r.Port)

	client, err := api.NewClient(consulConfig)
	if err != nil {
		return err
	}
	err = client.Agent().ServiceDeregister(serviceId)
	return err
}
