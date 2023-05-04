package main

import (
	"context"
	"flag"
	"mxshop-api/user-web/client"
	"mxshop-api/user-web/config"
	"mxshop-api/user-web/log"
	"mxshop-api/user-web/register"
	"mxshop-api/user-web/router"
	"mxshop-api/user-web/util"
	"mxshop-api/user-web/validators"
	"strconv"
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

	//初始化logger
	log.Init()

	//初始化routers
	engine := router.Init()

	// 初始化 validator
	if err := validators.Init("zh"); err != nil {
		panic(err)
	}

	//consul
	register.InitConsulRegister()

	// 初始化user-srv-client连接
	if err := client.Init(); err != nil {
		panic(err)
	}

	// 服务注册
	err := register.DefaultRegistry().Register(serviceConf.ServiceName,
		serviceConf.ServiceName,
		serviceConf.ServiceTags,
		host,
		port,
	)
	if err != nil {
		log.Error(context.Background(), "consul.registry fail", log.Any("err", err))
		panic(err)
	}

	//run!!
	if err = engine.Run(addr); err != nil {
		log.Panic(context.Background(), "服务启动失败", log.Any("err", err))
	}

	//接收终止信号
}
