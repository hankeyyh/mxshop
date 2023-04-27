package consul

import (
	"fmt"
	"github.com/hashicorp/consul/api"
)

type RegistryClient struct {
	Host string
	Port int
}

func NewRegistryClient(host string, port int) RegistryClient {
	return RegistryClient{
		Host: host,
		Port: port,
	}
}

func (r RegistryClient) Register(serviceName string, serviceId string, tags []string, address string, port int, options ...map[string]interface{}) error {
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
		GRPC:                           fmt.Sprintf("%s:%d", address, port),
		GRPCUseTLS:                     false,
		Interval:                       "5s",
		Timeout:                        "5s",
		DeregisterCriticalServiceAfter: "10s",
	}
	err = client.Agent().ServiceRegister(registry)
	return err
}

func (r RegistryClient) Deregister(serviceId string) error {
	consulConfig := api.DefaultConfig()
	consulConfig.Address = fmt.Sprintf("%s:%d", r.Host, r.Port)

	client, err := api.NewClient(consulConfig)
	if err != nil {
		return err
	}
	err = client.Agent().ServiceDeregister(serviceId)
	return err
}
