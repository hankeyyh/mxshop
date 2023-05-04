package register

var (
	registry Register
)

func DefaultRegistry() Register {
	return registry
}

type Register interface {
	Register(serviceName string, serviceId string, tags []string, address string, port int, options ...map[string]interface{}) error
	Deregister(serviceId string) error
	GetServiceAddr(serviceName string) (string, error)
}
