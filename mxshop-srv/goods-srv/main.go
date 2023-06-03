package main

import (
	"flag"
	"fmt"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"github.com/hankeyyh/mxshop/mxshop-srv/goods-srv/config"
	"github.com/hankeyyh/mxshop/mxshop-srv/goods-srv/handler"
	"github.com/hankeyyh/mxshop/mxshop-srv/goods-srv/log"
	"github.com/hankeyyh/mxshop/mxshop-srv/goods-srv/proto"
	"github.com/hankeyyh/mxshop/mxshop-srv/goods-srv/register"
	"github.com/hankeyyh/mxshop/mxshop-srv/goods-srv/xhealth"
	"github.com/hankeyyh/mxshop/mxshop-srv/goods-srv/xsignal"
	"github.com/hashicorp/go-uuid"
	"google.golang.org/grpc"
	"net"
)

func main() {
	serviceConf := config.DefaultConfig.Service
	server := grpc.NewServer()

	// cmd参数
	host := flag.String("host", serviceConf.Host, "host")
	port := flag.Int("port", serviceConf.Port, "port")
	flag.Parse()

	proto.RegisterGoodsServer(server, &handler.GoodsService{})

	// 中间件注册
	grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
		grpc_zap.UnaryServerInterceptor(log.DefaultLogger())))

	// 健康检查
	xhealth.RegisterHealthSrv(serviceConf.ServiceName, server)

	// consul注册
	serviceId, err := uuid.GenerateUUID()
	if err != nil {
		panic(err)
	}
	serviceAgent := register.DefaultServiceAgent
	if err = serviceAgent.Register(register.ServiceAgentConf{
		ServiceId:           serviceId,
		ServiceName:         serviceConf.ServiceName,
		ServiceHost:         serviceConf.Host,
		ServicePort:         serviceConf.Port,
		ServiceTags:         serviceConf.ServiceTags,
		HealthCheckInterval: "5s",
		DeregisterCritical:  "10s",
	}); err != nil {
		panic(err)
	}

	// 退出信号
	xsignal.WaitShutdown(func(graceful bool) {
		if graceful {
			server.GracefulStop()
		} else {
			server.Stop()
		}
	})

	addr := fmt.Sprintf("%s:%d", *host, *port)
	listen, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}

	if err = server.Serve(listen); err != nil {
		panic(err)
	}
}
