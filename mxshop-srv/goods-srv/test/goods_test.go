package test

import (
	"context"
	"github.com/hankeyyh/mxshop/mxshop-srv/goods-srv/proto"
	"google.golang.org/grpc"
	"testing"
)

var (
	conn   *grpc.ClientConn
	client proto.GoodsClient
)

func init() {
	var err error
	conn, err = grpc.Dial("192.168.3.56:8084", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	client = proto.NewGoodsClient(conn)
}

func TestGoodsList(t *testing.T) {
	req := new(proto.GoodsFilterRequest)
	req.Pages = 1
	req.PagePerNums = 10
	req.TopCategory = 130358
	req.PriceMin = 0
	req.PriceMax = 100
	rsp, err := client.GoodsList(context.Background(), req)
	if err != nil {
		panic(err)
	}
	t.Log(rsp)
}

func TestBatchGetGoods(t *testing.T) {
	req := new(proto.BatchGoodsIdInfo)
	req.Id = []int32{421, 422, 423, 424}
	rsp, err := client.BatchGetGoods(context.Background(), req)
	if err != nil {
		panic(err)
	}
	t.Log(rsp)
}
