package test

import (
	"context"
	"github.com/hankeyyh/mxshop/mxshop-srv/goods-srv/proto"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
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

func TestCreateGoods(t *testing.T) {
	req := &proto.CreateGoodsInfo{
		Name:            "黑色鲤鱼",
		GoodsSn:         "123123",
		Stocks:          123,
		MarketPrice:     12,
		ShopPrice:       38,
		GoodsBrief:      "黑色鲤鱼",
		ShipFree:        false,
		Images:          []string{"1.img", "2.img"},
		DescImages:      []string{"1.img", "2.img"},
		GoodsFrontImage: "1.img",
		IsNew:           true,
		IsHot:           true,
		OnSale:          true,
		CategoryId:      130361,
		BrandId:         616,
	}
	rsp, err := client.CreateGoods(context.Background(), req)
	if err != nil {
		panic(err)
	}
	t.Log(rsp)
}

func TestUpdateGoods(t *testing.T) {
	req := &proto.CreateGoodsInfo{
		Id:              841,
		Name:            "黑色鲤鱼2",
		GoodsSn:         "123123",
		Stocks:          123,
		MarketPrice:     12,
		ShopPrice:       38,
		GoodsBrief:      "黑色鲤鱼",
		ShipFree:        false,
		Images:          []string{"1.img", "2.img"},
		DescImages:      []string{"1.img", "2.img"},
		GoodsFrontImage: "1.img",
		IsNew:           true,
		IsHot:           true,
		OnSale:          true,
		CategoryId:      130361,
		BrandId:         616,
	}
	rsp, err := client.UpdateGoods(context.Background(), req)
	if err != nil {
		panic(err)
	}
	t.Log(rsp)
}

func TestGetGoodsDetail(t *testing.T) {
	req := &proto.GoodInfoRequest{
		Id: 788,
	}
	rsp, err := client.GetGoodsDetail(context.Background(), req)
	if err != nil {
		panic(err)
	}
	t.Log(rsp)
}

func TestGetAllCategorysList(t *testing.T) {
	rsp, err := client.GetAllCategorysList(context.Background(), &emptypb.Empty{})
	if err != nil {
		panic(err)
	}
	t.Log(rsp)
}

func TestGetSubCategory(t *testing.T) {
	req := &proto.CategoryListRequest{
		Id: 130358,
	}
	rsp, err := client.GetSubCategory(context.Background(), req)
	if err != nil {
		panic(err)
	}
	t.Log(rsp)
}

func TestCreateAndDeleteCategory(t *testing.T) {
	req := &proto.CategoryInfoRequest{
		Name:           "测试分类2级",
		ParentCategory: 130358,
		Level:          2,
		IsTab:          false,
	}
	rsp, err := client.CreateCategory(context.Background(), req)
	if err != nil {
		panic(err)
	}
	t.Log(rsp)

	upReq := &proto.CategoryInfoRequest{
		Id:             rsp.GetId(),
		Name:           "测试分类3级",
		ParentCategory: 130364,
		Level:          3,
		IsTab:          false,
	}
	upRsp, err := client.UpdateCategory(context.Background(), upReq)
	if err != nil {
		panic(err)
	}
	t.Log(upRsp)

	delReq := &proto.DeleteCategoryRequest{Id: rsp.GetId()}
	delRsp, err := client.DeleteCategory(context.Background(), delReq)
	if err != nil {
		panic(delRsp)
	}
	t.Log(delRsp)
}
