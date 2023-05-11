package main

import (
	"google.golang.org/grpc"
	"net"
)

func main() {
	server := grpc.NewServer()

	listen, err := net.Listen("tcp", ":1233")
	if err != nil {
		panic(err)
	}

	if err = server.Serve(listen); err != nil {
		panic(err)
	}
}
