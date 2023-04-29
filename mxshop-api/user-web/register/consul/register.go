package consul

import (
	"fmt"
	"github.com/hashicorp/consul/api"
)

type RegistryClient interface {
	Register(address string, port int, name string, tags []string, id string) error
	DeRegister(serviceId string) error
}

type Registry struct {
	Host string
	Port int
}

func NewRegistryClient(host string, port int) RegistryClient {
	return &Registry{
		Host: host,
		Port: port,
	}
}

func (r *Registry) Register(address string, port int, name string, tags []string, id string) error {
	conf := api.DefaultConfig()
	conf.Address = fmt.Sprintf("%s:%d", address, port)

	client, err := api.NewClient(conf)
	if err != nil {
		return err
	}
	check := &api.AgentServiceCheck{
		HTTP:                           fmt.Sprintf("http://%s:%s/health", address, port),
		Timeout:                        "5s",
		Interval:                       "5s",
		DeregisterCriticalServiceAfter: "10s",
	}

	// 生成注册对象
	registration := new(api.AgentServiceRegistration)
	registration.Name = name
	registration.Port = port
	registration.Tags = tags
	registration.ID = id
	registration.Check = check
	registration.Address = address

	if err = client.Agent().ServiceRegister(registration); err != nil {
		return err
	}
	return nil
}

func (r *Registry) DeRegister(serviceId string) error {
	cfg := api.DefaultConfig()
	cfg.Address = fmt.Sprintf("%s:%d", r.Host, r.Port)

	client, err := api.NewClient(cfg)
	if err != nil {
		return err
	}
	err = client.Agent().ServiceDeregister(serviceId)
	return err
}
