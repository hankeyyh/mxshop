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

func InitConsulRegister() {
	conf := config.DefaultConfig().Consul
	registry = ConsulClient{
		Host: conf.Host,
		Port: conf.Port,
	}
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

func (r ConsulClient) GetServiceAddr(serviceName string) (string, error) {
	consulConfig := api.DefaultConfig()
	consulConfig.Address = fmt.Sprintf("%s:%d", r.Host, r.Port)

	client, err := api.NewClient(consulConfig)
	if err != nil {
		log.Error(context.Background(), "consul NewClient fail", log.Any("err", err))
		return "", err
	}
	servMap, err := client.Agent().ServicesWithFilter(fmt.Sprintf("Service == \"%s\"", serviceName))
	if err != nil {
		log.Error(context.Background(), "consul ServicesWithFilter fail", log.Any("err", err))
		return "", err
	}
	var addr string
	for _, v := range servMap {
		addr = fmt.Sprintf("%s:%d", v.Address, v.Port)
	}
	return addr, nil
}
