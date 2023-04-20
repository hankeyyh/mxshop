package main

import (
	"context"
	"flag"
	"fmt"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	health_handler "github.com/hankeyyh/mxshop/mxshop-srv/common/grpc-health/v1/handler"
	health_pb "github.com/hankeyyh/mxshop/mxshop-srv/common/grpc-health/v1/proto"
	"github.com/hankeyyh/mxshop/mxshop-srv/user-srv/handler"
	"github.com/hankeyyh/mxshop/mxshop-srv/user-srv/log"
	"github.com/hankeyyh/mxshop/mxshop-srv/user-srv/proto"
	"google.golang.org/grpc"
	"net"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
)

func main() {
	host := flag.String("host", "localhost", "Host address")
	port := flag.Int("port", 8083, "Port")
	addr := *host + ":" + strconv.Itoa(*port)

	// 中间件注册
	opt := grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
		grpc_ctxtags.UnaryServerInterceptor(),
		grpc_zap.UnaryServerInterceptor(log.DefaultLogger()),
	))

	// 创建服务
	server := grpc.NewServer(opt)
	proto.RegisterUserServer(server, &handler.UserService{})
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}

	// 主进程退出信号监听
	sgCh := make(chan os.Signal)
	signal.Notify(sgCh, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		<-sgCh
		server.GracefulStop()
		log.Info(context.Background(), "Server Closing...")
		wg.Done()
	}()

	// todo 服务注册
	health_pb.RegisterHealthServer(server, &health_handler.HealthCheckService{})

	// 启动服务
	fmt.Printf("Server Running at %s\n", addr)
	err = server.Serve(listener)
	if err != nil {
		panic(err)
	}

	wg.Wait()
	log.Info(context.Background(), "Server Closed")
}
