package handler

import (
	"context"
	"github.com/hankeyyh/mxshop/mxshop-srv/goods-srv/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

type GoodsService struct {
	proto.UnimplementedGoodsServer
}

func (g GoodsService) GoodsList(ctx context.Context, request *proto.GoodsFilterRequest) (*proto.GoodsListResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (g GoodsService) BatchGetGoods(ctx context.Context, info *proto.BatchGoodsIdInfo) (*proto.GoodsListResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (g GoodsService) CreateGoods(ctx context.Context, info *proto.CreateGoodsInfo) (*proto.GoodsInfoResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (g GoodsService) DeleteGoods(ctx context.Context, info *proto.DeleteGoodsInfo) (*emptypb.Empty, error) {
	//TODO implement me
	panic("implement me")
}

func (g GoodsService) UpdateGoods(ctx context.Context, info *proto.CreateGoodsInfo) (*emptypb.Empty, error) {
	//TODO implement me
	panic("implement me")
}

func (g GoodsService) GetGoodsDetail(ctx context.Context, request *proto.GoodInfoRequest) (*proto.GoodsInfoResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (g GoodsService) GetAllCategorysList(ctx context.Context, empty *emptypb.Empty) (*proto.CategoryListResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (g GoodsService) GetSubCategory(ctx context.Context, request *proto.CategoryListRequest) (*proto.SubCategoryListResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (g GoodsService) CreateCategory(ctx context.Context, request *proto.CategoryInfoRequest) (*proto.CategoryInfoResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (g GoodsService) DeleteCategory(ctx context.Context, request *proto.DeleteCategoryRequest) (*emptypb.Empty, error) {
	//TODO implement me
	panic("implement me")
}

func (g GoodsService) UpdateCategory(ctx context.Context, request *proto.CategoryInfoRequest) (*emptypb.Empty, error) {
	//TODO implement me
	panic("implement me")
}

func (g GoodsService) BrandList(ctx context.Context, request *proto.BrandFilterRequest) (*proto.BrandListResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (g GoodsService) CreateBrand(ctx context.Context, request *proto.BrandRequest) (*proto.BrandInfoResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (g GoodsService) DeleteBrand(ctx context.Context, request *proto.BrandRequest) (*emptypb.Empty, error) {
	//TODO implement me
	panic("implement me")
}

func (g GoodsService) UpdateBrand(ctx context.Context, request *proto.BrandRequest) (*emptypb.Empty, error) {
	//TODO implement me
	panic("implement me")
}

func (g GoodsService) BannerList(ctx context.Context, empty *emptypb.Empty) (*proto.BannerListResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (g GoodsService) CreateBanner(ctx context.Context, request *proto.BannerRequest) (*proto.BannerResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (g GoodsService) DeleteBanner(ctx context.Context, request *proto.BannerRequest) (*emptypb.Empty, error) {
	//TODO implement me
	panic("implement me")
}

func (g GoodsService) UpdateBanner(ctx context.Context, request *proto.BannerRequest) (*emptypb.Empty, error) {
	//TODO implement me
	panic("implement me")
}

func (g GoodsService) CategoryBrandList(ctx context.Context, request *proto.CategoryBrandFilterRequest) (*proto.CategoryBrandListResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (g GoodsService) GetCategoryBrandList(ctx context.Context, request *proto.CategoryInfoRequest) (*proto.BrandListResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (g GoodsService) CreateCategoryBrand(ctx context.Context, request *proto.CategoryBrandRequest) (*proto.CategoryBrandResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (g GoodsService) DeleteCategoryBrand(ctx context.Context, request *proto.CategoryBrandRequest) (*emptypb.Empty, error) {
	//TODO implement me
	panic("implement me")
}

func (g GoodsService) UpdateCategoryBrand(ctx context.Context, request *proto.CategoryBrandRequest) (*emptypb.Empty, error) {
	//TODO implement me
	panic("implement me")
}

func (g GoodsService) mustEmbedUnimplementedGoodsServer() {
	//TODO implement me
	panic("implement me")
}
