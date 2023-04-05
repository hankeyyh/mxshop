package main

import (
	"flag"
	"go.uber.org/zap"
	"mxshop-api/user-web/initialize"
	"strconv"
)

func main() {
	host := flag.String("host", "localhost", "Host address")
	port := flag.Int("port", 8021, "Port")
	flag.Parse()
	addr := *host + ":" + strconv.Itoa(*port)

	//初始化logger
	initialize.InitLogger()

	//初始化配置文件

	//初始化routers

	engine := initialize.InitRouter()

	//初始化srv的连接
	zap.S().Infof("服务启动 %s\n", addr)
	if err := engine.Run(addr); err != nil {
		zap.S().Panic("服务启动失败 err: ", err.Error())
	}

	//注册验证器

	//服务注册

	//接收终止信号
}
