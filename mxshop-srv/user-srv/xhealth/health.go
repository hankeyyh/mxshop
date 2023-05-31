package xhealth

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
)

var healthSrv *health.Server

func init() {
	healthSrv = health.NewServer()
}

func RegisterHealthSrv(serviceName string, s grpc.ServiceRegistrar) {
	SetServingStatus(serviceName, grpc_health_v1.HealthCheckResponse_SERVING)
	grpc_health_v1.RegisterHealthServer(s, healthSrv)
}

func SetServingStatus(serviceName string, status grpc_health_v1.HealthCheckResponse_ServingStatus) {
	healthSrv.SetServingStatus(serviceName, status)
}

func Shutdown() {
	healthSrv.Shutdown()
}

func Resume() {
	healthSrv.Resume()
}
