package client

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"mxshop-api/user-web/config"
	"mxshop-api/user-web/log"
	"mxshop-api/user-web/proto"
	"mxshop-api/user-web/register"
)

var (
	UserSvrClient proto.UserClient
)

func Init() error {
	if err := InitUserSrvClient(); err != nil {
		return err
	}
	return nil
}

func InitUserSrvClient() error {
	conf := config.DefaultConfig().Client.UserSrv
	registry := register.DefaultRegistry()
	// 服务发现
	// todo consul记录的user-srv地址是host.docker.internal 在宿主机上访问不了
	addr, err := registry.GetServiceAddr(conf.Name)
	if err != nil {
		return err
	}
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Error(context.Background(), "init grpc conn fail", log.Any("err", err))
		return err
	}
	// todo 1. 用户服务下线了，改ip端口了怎么办
	// todo 2. 多个goroutine用一个client，性能问题 - 连接池
	UserSvrClient = proto.NewUserClient(conn)
	return nil
}
