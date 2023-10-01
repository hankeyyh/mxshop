package client

import (
	"context"
	"fmt"
	"github.com/hankeyyh/mxshop/mxshop-api/user-web/config"
	"github.com/hankeyyh/mxshop/mxshop-api/user-web/log"
	"github.com/hankeyyh/mxshop/mxshop-api/user-web/proto"
	_ "github.com/mbobakov/grpc-consul-resolver"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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
	conf := config.DefaultConfig.Client.UserSrv
	consulConf := config.DefaultConfig.Consul

	/* 我们可以直接用consul-agent拉取全部的服务ip，选择一个访问
	但为了使用负载均衡算法，这里使用基于consul的第三方的name resolver(scheme=consul)
	*/
	conn, err := grpc.Dial(fmt.Sprintf("consul://%s:%d/%s?wait=14s", consulConf.Host, consulConf.Port, conf.Name),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`))
	if err != nil {
		log.Error(context.Background(), "init grpc conn fail", log.Any("err", err))
		return err
	}

	UserSvrClient = proto.NewUserClient(conn)
	return nil
}
