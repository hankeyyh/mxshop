package main

import (
	"github.com/hankeyyh/mxshop_user_srv/handler"
	"github.com/hankeyyh/mxshop_user_srv/proto"
	"google.golang.org/grpc"
	"net"
)

/*
todo log pkg
todo 中间件打印return err，打印耗时
todo 监听ip，端口作为启动参数传入
*/
func main() {
	server := grpc.NewServer()
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
