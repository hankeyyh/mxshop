package main

import (
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	_ "github.com/hankeyyh/mxshop/mxshop-srv/goods-srv/config"
	"github.com/hankeyyh/mxshop/mxshop-srv/goods-srv/handler"
	"github.com/hankeyyh/mxshop/mxshop-srv/goods-srv/log"
	"github.com/hankeyyh/mxshop/mxshop-srv/goods-srv/proto"
	"github.com/hankeyyh/mxshop/mxshop-srv/goods-srv/xhealth"
	"github.com/hankeyyh/mxshop/mxshop-srv/goods-srv/xsignal"
	"google.golang.org/grpc"
	"net"
)

func main() {
	server := grpc.NewServer()

	listen, err := net.Listen("tcp", ":1233")
	if err != nil {
		panic(err)
	}

	proto.RegisterGoodsServer(server, &handler.GoodsService{})

	// 中间件注册
	grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
		grpc_zap.UnaryServerInterceptor(log.DefaultLogger())))

	// 健康检查
	xhealth.RegisterHealthSrv("goods-srv", server)

	// todo consul注册

	// 退出信号
	xsignal.WaitShutdown(func(graceful bool) {
		if graceful {
			server.GracefulStop()
		} else {
			server.Stop()
		}
	})

	if err = server.Serve(listen); err != nil {
		panic(err)
	}
}
