package main

import (
	"context"
	"flag"
	"mxshop-api/user-web/client"
	"mxshop-api/user-web/config"
	"mxshop-api/user-web/log"
	"mxshop-api/user-web/register"
	"mxshop-api/user-web/router"
	"mxshop-api/user-web/validators"
	"strconv"
)

func main() {
	host := flag.String("host", "localhost", "Host address")
	port := flag.Int("port", 8021, "Port")
	flag.Parse()
	addr := *host + ":" + strconv.Itoa(*port)

	//初始化配置文件
	if err := config.Init(); err != nil {
		panic(err)
	}

	//初始化logger
	log.Init()

	//初始化routers
	engine := router.Init()

	// 初始化 validator
	if err := validators.Init("zh"); err != nil {
		panic(err)
	}

	// 初始化user-srv-client连接
	if err := client.Init(); err != nil {
		panic(err)
	}

	//服务注册
	if err := register.InitConsulRegister(); err != nil {
		panic(err)
	}

	//run!!
	log.Info(context.Background(), "服务启动 %s", log.Any("addr", addr))
	if err := engine.Run(addr); err != nil {
		log.Panic(context.Background(), "服务启动失败", log.Any("err", err))
	}

	//接收终止信号
}
