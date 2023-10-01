package client

import (
	"context"
	"fmt"
	"github.com/hankeyyh/mxshop/mxshop-api/user-web/config"
	"github.com/hankeyyh/mxshop/mxshop-api/user-web/log"
	"github.com/hankeyyh/mxshop/mxshop-api/user-web/proto"
	"github.com/hashicorp/consul/api"
	"github.com/hashicorp/consul/api/watch"
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
	// todo 1. 用户服务下线了，改ip端口了怎么办
	go func() {
		err := watchUserSrv(fmt.Sprintf("%s:%d", consulConf.Host, consulConf.Port))
		if err != nil {
			log.Error(context.Background(), "watchUserSrv fail", log.Any("err", err))
		}
	}()
	// todo 2. 多个goroutine用一个client，性能问题 - 连接池
	UserSvrClient = proto.NewUserClient(conn)
	return nil
}

func watchUserSrv(consulAddr string) error {
	plan, err := watch.Parse(map[string]interface{}{
		"type":    "service",
		"service": "user-srv",
	})
	if err != nil {
		return err
	}
	plan.Handler = func(u uint64, data interface{}) {
		switch d := data.(type) {
		case []*api.ServiceEntry:
			for _, i := range d {
				fmt.Printf("service %s change, host: %s, port: %d\n",
					i.Service.Service, i.Service.Address, i.Service.Port)
			}
		}
	}
	if err = plan.Run(consulAddr); err != nil {
		return err
	}
	return nil
}
