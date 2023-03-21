package test

import (
	"context"
	"fmt"
	"github.com/hankeyyh/mxshop_user_srv/proto"
	"google.golang.org/grpc"
	"testing"
)

var (
	conn   *grpc.ClientConn
	client proto.UserClient
)

func init() {
	var err error
	conn, err = grpc.Dial("127.0.0.1:8083", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	client = proto.NewUserClient(conn)
}

func TestGetUserList(t *testing.T) {
	defer conn.Close()
	rsp, err := client.GetUserList(context.Background(), &proto.PageInfo{
		Pn:    1,
		PSize: 10,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(rsp)
}
