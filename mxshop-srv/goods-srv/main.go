package main

import (
	"github.com/hankeyyh/mxshop/mxshop-srv/goods-srv/handler"
	"github.com/hankeyyh/mxshop/mxshop-srv/goods-srv/proto"
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

	if err = server.Serve(listen); err != nil {
		panic(err)
	}
}
