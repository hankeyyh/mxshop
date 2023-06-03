package register

import (
	"fmt"
	"github.com/hashicorp/consul/api"
)

type ConsulAgent struct {
	host   string
	port   int
	client *api.Client
}

func NewConsulAgent(host string, port int) (agent *ConsulAgent, err error) {
	consulConf := api.DefaultConfig()
	consulConf.Address = fmt.Sprintf("%s:%d", host, port)

	var client *api.Client
	client, err = api.NewClient(consulConf)
	if err != nil {
		return nil, err
	}

	agent = new(ConsulAgent)
	agent.host = host
	agent.port = port
	agent.client = client
	return
}

func (c *ConsulAgent) Register(conf ServiceAgentConf) error {
	registration := api.AgentServiceRegistration{
		ID:      conf.ServiceId,
		Name:    conf.ServiceName,
		Tags:    conf.ServiceTags,
		Port:    conf.ServicePort,
		Address: conf.ServiceHost,
		Check: &api.AgentServiceCheck{
			Interval:                       conf.HealthCheckInterval,
			GRPC:                           fmt.Sprintf("%s:%d/%s", conf.ServiceHost, conf.ServicePort, conf.ServiceName),
			DeregisterCriticalServiceAfter: conf.DeregisterCritical,
			GRPCUseTLS:                     false,
		},
	}
	return c.client.Agent().ServiceRegister(&registration)
}

func (c *ConsulAgent) Deregister(serviceId string) error {
	return c.client.Agent().ServiceDeregister(serviceId)
}
