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
	"github.com/hankeyyh/mxshop/mxshop-srv/common/util"
	"github.com/hankeyyh/mxshop/mxshop-srv/user-srv/config"
	"github.com/hankeyyh/mxshop/mxshop-srv/user-srv/handler"
	"github.com/hankeyyh/mxshop/mxshop-srv/user-srv/log"
	"github.com/hankeyyh/mxshop/mxshop-srv/user-srv/proto"
	"github.com/hankeyyh/mxshop/mxshop-srv/user-srv/register/consul"
	"github.com/hashicorp/go-uuid"
	"google.golang.org/grpc"
	"net"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
)

func main() {
	if !config.IsDebug() {
		// 正式环境随机端口号，支持启动多个实例
		port, err := util.GetFreePort()
		if err != nil {
			panic(err)
		}
		config.DefaultConfig.Service.Port = port
	}
	serviceConf := config.DefaultConfig.Service
	host := *flag.String("host", serviceConf.Host, "Host address")
	port := *flag.Int("port", serviceConf.Port, "Port")
	flag.Parse()
	addr := host + ":" + strconv.Itoa(port)

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

	// 健康检查
	health_pb.RegisterHealthServer(server, &health_handler.HealthCheckService{})

	// 服务注册
	consulConf := config.DefaultConfig.Consul
	serviceId, err := uuid.GenerateUUID()
	if err != nil {
		panic(err)
	}
	consulClient := consul.NewRegistryClient(consulConf.Host, consulConf.Port)
	err = consulClient.Register(serviceConf.ServiceName,
		serviceId,
		serviceConf.ServiceTags,
		host,
		port)
	if err != nil {
		panic(err)
	}

	// 启动服务
	fmt.Printf("Server Running at %s\n", addr)
	err = server.Serve(listener)
	if err != nil {
		panic(err)
	}

	wg.Wait()
	log.Info(context.Background(), "Server Closed")
}
