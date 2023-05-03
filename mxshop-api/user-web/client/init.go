package client

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"mxshop-api/user-web/config"
	"mxshop-api/user-web/log"
	"mxshop-api/user-web/proto"
)

var (
	UserSvrClient proto.UserClient
)

func Init() error {
	if err := InitUserSrvClient(); err != nil {
		return err
	}
	return nil
}

func InitUserSrvClient() error {
	conf := config.DefaultConfig().UserSrvInfo
	addr := fmt.Sprintf("%s:%d", conf.Host, conf.Port)
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Error(context.Background(), "init grpc conn fail", log.Any("err", err))
		return err
	}
	UserSvrClient = proto.NewUserClient(conn)
	return nil
}
