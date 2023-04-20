package handler

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"grpc_health/v1/proto"
)

type HealthCheckService struct {
	proto.UnimplementedHealthServer
}

// todo 如何根据servicename注册多个健康状态
func (h HealthCheckService) Check(context.Context, *proto.HealthCheckRequest) (*proto.HealthCheckResponse, error) {
	rsp := new(proto.HealthCheckResponse)
	rsp.Status = proto.HealthCheckResponse_SERVING
	return rsp, nil
}
func (h HealthCheckService) Watch(*proto.HealthCheckRequest, proto.Health_WatchServer) error {
	return status.Errorf(codes.Unimplemented, "method Watch not implemented")
}
