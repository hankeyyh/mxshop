package main

import (
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"github.com/hankeyyh/mxshop_user_srv/handler"
	"github.com/hankeyyh/mxshop_user_srv/log"
	//"github.com/hankeyyh/mxshop_user_srv/interceptor"
	"github.com/hankeyyh/mxshop_user_srv/proto"
	"google.golang.org/grpc"
	"net"
)

/*
todo 监听ip，端口作为启动参数传入
*/
func main() {
	opt := grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
		grpc_ctxtags.UnaryServerInterceptor(),
		grpc_zap.UnaryServerInterceptor(log.DefaultLogger()),
	))

	server := grpc.NewServer(opt)
	proto.RegisterUserServer(server, &handler.UserService{})
	listener, err := net.Listen("tcp", ":8083")
	if err != nil {
		panic(err)
	}

	// todo 主进程退出信号监听

	// 启动服务
	err = server.Serve(listener)
	if err != nil {
		panic(err)
	}

	// todo 服务注册

}
