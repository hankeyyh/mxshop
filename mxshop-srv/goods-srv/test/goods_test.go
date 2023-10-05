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

func TestCreateUpdateDeleteCategory(t *testing.T) {
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

func TestBrandList(t *testing.T) {
	req := &proto.BrandFilterRequest{
		Pages:       1,
		PagePerNums: 20,
	}
	rsp, err := client.BrandList(context.Background(), req)
	if err != nil {
		panic(err)
	}
	t.Log(rsp)
}

func TestCreateUpdateDeleteBrand(t *testing.T) {
	req := &proto.BrandRequest{
		Name: "测试brand",
		Logo: "测试logo",
	}
	rsp, err := client.CreateBrand(context.Background(), req)
	if err != nil {
		panic(err)
	}
	t.Log(rsp)

	upReq := &proto.BrandRequest{
		Id:   rsp.GetId(),
		Name: "测试brand2",
		Logo: "测试logo2",
	}
	upRsp, err := client.UpdateBrand(context.Background(), upReq)
	if err != nil {
		panic(err)
	}
	t.Log(upRsp)

	delReq := &proto.BrandRequest{
		Id: rsp.GetId(),
	}
	delRsp, err := client.DeleteBrand(context.Background(), delReq)
	if err != nil {
		panic(err)
	}
	t.Log(delRsp)
}

func TestBannerList(t *testing.T) {
	rsp, err := client.BannerList(context.Background(), &emptypb.Empty{})
	if err != nil {
		panic(err)
	}
	t.Log(rsp)
}

func TestCreateUpdateDeleteBanner(t *testing.T) {
	req := &proto.BannerRequest{
		Index: 1,
		Image: "测试image",
		Url:   "测试url",
	}
	rsp, err := client.CreateBanner(context.Background(), req)
	if err != nil {
		panic(err)
	}
	t.Log(rsp)

	upReq := &proto.BannerRequest{
		Id:    rsp.GetId(),
		Index: 2,
		Image: "测试image2",
		Url:   "测试url2",
	}
	upRsp, err := client.UpdateBanner(context.Background(), upReq)
	if err != nil {
		panic(err)
	}
	t.Log(upRsp)

	delReq := &proto.BannerRequest{Id: rsp.GetId()}
	delRsp, err := client.DeleteBanner(context.Background(), delReq)
	if err != nil {
		panic(err)
	}
	t.Log(delRsp)
}

func TestCategoryBrandList(t *testing.T) {
	req := &proto.CategoryBrandFilterRequest{
		Pages:       1,
		PagePerNums: 20,
	}
	rsp, err := client.CategoryBrandList(context.Background(), req)
	if err != nil {
		panic(err)
	}
	t.Log(rsp)
}

func TestGetCategoryBrandList(t *testing.T) {
	req := &proto.CategoryInfoRequest{Id: 130366}
	rsp, err := client.GetCategoryBrandList(context.Background(), req)
	if err != nil {
		panic(err)
	}
	t.Log(rsp)
}

func TestCreateUpdateDeleteCategoryBrand(t *testing.T) {
	req := &proto.CategoryBrandRequest{
		CategoryId: 130358,
		BrandId:    856,
	}
	rsp, err := client.CreateCategoryBrand(context.Background(), req)
	if err != nil {
		panic(err)
	}
	t.Log(rsp)

	upReq := &proto.CategoryBrandRequest{
		Id:         rsp.GetId(),
		CategoryId: 135505,
		BrandId:    1072,
	}
	upRsp, err := client.UpdateCategoryBrand(context.Background(), upReq)
	if err != nil {
		panic(err)
	}
	t.Log(upRsp)

	delReq := &proto.CategoryBrandRequest{Id: rsp.GetId()}
	delRsp, err := client.DeleteCategoryBrand(context.Background(), delReq)
	if err != nil {
		panic(err)
	}
	t.Log(delRsp)
}
