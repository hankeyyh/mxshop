package handler

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hankeyyh/mxshop/mxshop-srv/goods-srv/dao"
	"github.com/hankeyyh/mxshop/mxshop-srv/goods-srv/log"
	"github.com/hankeyyh/mxshop/mxshop-srv/goods-srv/model"
	"github.com/hankeyyh/mxshop/mxshop-srv/goods-srv/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"strings"
)

type GoodsService struct {
	proto.UnimplementedGoodsServer
}

type CategoryListItemForJson struct {
	Id          int32                      `json:"id"`
	Name        string                     `json:"name"`
	Level       int32                      `json:"level"`
	ParentId    int32                      `json:"parent_id"`
	IsTab       int32                      `json:"is_tab"`
	SubCategory []*CategoryListItemForJson `json:"sub_category"`
}

func convertToCategoryInfoResponse(category *model.Category) *proto.CategoryInfoResponse {
	rsp := &proto.CategoryInfoResponse{
		Id:             category.ID,
		Name:           category.Name,
		ParentCategory: int32(category.ParentCategoryID.Int64),
		Level:          category.Level,
		IsTab:          category.IsTab == 1,
	}
	return rsp
}

func convertToCategoryListItemForJson(category *model.Category) *CategoryListItemForJson {
	return &CategoryListItemForJson{
		Id:       category.ID,
		Name:     category.Name,
		Level:    category.Level,
		ParentId: int32(category.ParentCategoryID.Int64),
		IsTab:    category.IsTab,
	}
}

func convertToBrandInfoResponse(brand *model.Brands) *proto.BrandInfoResponse {
	return &proto.BrandInfoResponse{
		Id:   brand.ID,
		Name: brand.Name,
		Logo: brand.Logo.String,
	}
}

func convertToBannerResponse(banner *model.Banner) *proto.BannerResponse {
	return &proto.BannerResponse{
		Id:    banner.ID,
		Index: banner.Index,
		Image: banner.Image,
		Url:   banner.URL,
	}
}

func convertToGoodsInfoResponse(goods *model.Goods, category *model.Category, brand *model.Brands) *proto.GoodsInfoResponse {
	rsp := &proto.GoodsInfoResponse{
		Id:              goods.ID,
		CategoryId:      goods.CategoryID,
		Name:            goods.Name,
		GoodsSn:         goods.GoodsSn,
		ClickNum:        goods.ClickNum,
		SoldNum:         goods.SoldNum,
		FavNum:          goods.FavNum,
		MarketPrice:     goods.MarketPrice,
		ShopPrice:       goods.ShopPrice,
		GoodsBrief:      goods.GoodsBrief,
		GoodsDesc:       "",
		ShipFree:        goods.ShipFree != 0,
		Images:          []string{goods.Images},
		DescImages:      []string{goods.DescImages},
		GoodsFrontImage: goods.GoodsFrontImage,
		IsNew:           goods.IsNew == 1,
		IsHot:           goods.IsHot == 1,
		OnSale:          goods.OnSale == 1,
		AddTime:         goods.AddTime.Unix(),
		Category: &proto.CategoryBriefInfoResponse{
			Id:   category.ID,
			Name: category.Name,
		},
		Brand: &proto.BrandInfoResponse{
			Id:   brand.ID,
			Name: brand.Name,
			Logo: brand.Logo.String,
		},
	}
	return rsp
}

func (g GoodsService) getGoodsInfoResponse(ctx context.Context, goods *model.Goods) (*proto.GoodsInfoResponse, error) {
	category, err := dao.GetCategory(ctx, goods.CategoryID)
	if err != nil {
		return nil, err
	}
	brand, err := dao.GetBrands(ctx, goods.BrandID)
	if err != nil {
		return nil, err
	}
	return convertToGoodsInfoResponse(goods, category, brand), nil
}

func (g GoodsService) getGoodsListResponse(ctx context.Context, goodsList []*model.Goods) ([]*proto.GoodsInfoResponse, error) {
	categoryIdList := make([]int32, 0)
	brandIdList := make([]int32, 0)
	for _, goods := range goodsList {
		categoryIdList = append(categoryIdList, goods.CategoryID)
		brandIdList = append(brandIdList, goods.BrandID)
	}
	// 分类列表
	categoryList, err := dao.GetCategoryList(ctx, categoryIdList)
	if err != nil {
		return nil, err
	}
	// 品牌列表
	brandList, err := dao.GetBrandsList(ctx, brandIdList)
	if err != nil {
		return nil, err
	}

	var categoryMap = make(map[int32]*model.Category)
	var brandsMap = make(map[int32]*model.Brands)
	for _, category := range categoryList {
		categoryMap[category.ID] = category
	}
	for _, brands := range brandList {
		brandsMap[brands.ID] = brands
	}

	// 结果
	data := make([]*proto.GoodsInfoResponse, 0, len(goodsList))
	for _, goods := range goodsList {
		data = append(data, convertToGoodsInfoResponse(goods, categoryMap[goods.CategoryID], brandsMap[goods.BrandID]))
	}
	return data, nil
}

func (g GoodsService) GoodsList(ctx context.Context, request *proto.GoodsFilterRequest) (*proto.GoodsListResponse, error) {
	whereCond := dao.GoodsWhere{
		PriceMin:    request.GetPriceMin(),
		PriceMax:    request.GetPriceMax(),
		IsHot:       request.GetIsHot(),
		IsNew:       request.GetIsNew(),
		Pages:       request.GetPages(),
		PagePerNums: request.GetPagePerNums(),
		KeyWords:    request.GetKeyWords(),
		BrandId:     request.GetBrand(),
	}
	if request.TopCategory != 0 {
		// 找到最下级分类
		categoryList, err := dao.GetLowestCategoryList(ctx, request.TopCategory)
		if err != nil {
			return nil, err
		}
		var categoryIdList = make([]int32, 0, len(categoryList))
		for _, category := range categoryList {
			categoryIdList = append(categoryIdList, category.ID)
		}
		// 商品归属于最下级分类
		whereCond.CategoryIdList = categoryIdList
	}
	// 货物列表
	goodsList, total, err := dao.GetGoodsList(ctx, whereCond)
	if err != nil {
		return nil, err
	}
	// goods model list -> goods pb list
	goodsInfoResponseList, err := g.getGoodsListResponse(ctx, goodsList)
	if err != nil {
		return nil, err
	}
	rsp := new(proto.GoodsListResponse)
	rsp.Total = total
	rsp.Data = goodsInfoResponseList
	return rsp, nil
}

func (g GoodsService) BatchGetGoods(ctx context.Context, request *proto.BatchGoodsIdInfo) (*proto.GoodsListResponse, error) {
	idList := request.GetId()
	goodsList, err := dao.BatchGetGoods(ctx, idList)
	if err != nil {
		return nil, err
	}
	// goods model list -> goods pb list
	goodsInfoResponseList, err := g.getGoodsListResponse(ctx, goodsList)
	if err != nil {
		return nil, err
	}

	rsp := new(proto.GoodsListResponse)
	rsp.Total = int64(len(goodsList))
	rsp.Data = goodsInfoResponseList
	return rsp, nil
}

func (g GoodsService) CreateGoods(ctx context.Context, request *proto.CreateGoodsInfo) (*proto.GoodsInfoResponse, error) {
	categoryId := request.GetCategoryId()
	brandId := request.GetBrandId()

	// 检查分类
	_, err := dao.GetCategory(ctx, categoryId)
	if err != nil {
		log.Info("category not exist")
		return nil, err
	}
	// 检查品牌
	_, err = dao.GetBrands(ctx, brandId)
	if err != nil {
		log.Info("brand not exist")
		return nil, err
	}

	// 创建goods
	var onSale int32 = 0
	if request.OnSale {
		onSale = 1
	}
	var shipFree int32 = 0
	if request.ShipFree {
		shipFree = 1
	}
	var isNew int32 = 0
	if request.IsNew {
		isNew = 1
	}
	var isHot int32 = 0
	if request.IsHot {
		isHot = 1
	}
	// 将每个字符串用双引号引起来 - images
	quotedStrs := make([]string, len(request.Images))
	for i, s := range request.Images {
		quotedStrs[i] = fmt.Sprintf("%q", s)
	}
	images := "[" + strings.Join(quotedStrs, ",") + "]"

	// 将每个字符串用双引号引起来 - descImages
	quotedStrs = make([]string, len(request.DescImages))
	for i, s := range request.Images {
		quotedStrs[i] = fmt.Sprintf("%q", s)
	}
	descImages := "[" + strings.Join(quotedStrs, ",") + "]"

	goods := &model.Goods{
		CategoryID:      categoryId,
		BrandID:         brandId,
		OnSale:          onSale,
		GoodsSn:         request.GoodsSn,
		Name:            request.Name,
		Stocks:          request.Stocks,
		MarketPrice:     request.MarketPrice,
		ShopPrice:       request.ShopPrice,
		GoodsBrief:      request.GoodsBrief,
		ShipFree:        shipFree,
		Images:          images,
		DescImages:      descImages,
		GoodsFrontImage: request.GoodsFrontImage,
		IsNew:           isNew,
		IsHot:           isHot,
	}
	createdGoods, _, err := dao.AddGoods(ctx, goods)
	if err != nil {
		log.Error("dao.AddGoods fail", log.Any("err", err))
		return nil, err
	}
	return g.getGoodsInfoResponse(ctx, createdGoods)
}

func (g GoodsService) DeleteGoods(ctx context.Context, request *proto.DeleteGoodsInfo) (*emptypb.Empty, error) {
	id := request.GetId()
	_, err := dao.DeleteGoods(ctx, id)
	if err != nil {
		log.Error("dao.DeleteGoods fail", log.Any("err", err))
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (g GoodsService) UpdateGoods(ctx context.Context, request *proto.CreateGoodsInfo) (*emptypb.Empty, error) {
	categoryId := request.GetCategoryId()
	brandId := request.GetBrandId()
	goodsId := request.GetId()

	// 检查分类
	_, err := dao.GetCategory(ctx, categoryId)
	if err != nil {
		log.Info("category not exist")
		return nil, err
	}
	// 检查品牌
	_, err = dao.GetBrands(ctx, brandId)
	if err != nil {
		log.Info("brand not exist")
		return nil, err
	}
	// 检查商品
	goods, err := dao.GetGoods(ctx, goodsId)
	if err != nil {
		log.Info("goods not exist")
		return nil, err
	}

	// 创建goods
	var onSale int32 = 0
	if request.OnSale {
		onSale = 1
	}
	var shipFree int32 = 0
	if request.ShipFree {
		shipFree = 1
	}
	var isNew int32 = 0
	if request.IsNew {
		isNew = 1
	}
	var isHot int32 = 0
	if request.IsHot {
		isHot = 1
	}
	// 将每个字符串用双引号引起来 - images
	quotedStrs := make([]string, len(request.Images))
	for i, s := range request.Images {
		quotedStrs[i] = fmt.Sprintf("%q", s)
	}
	images := "[" + strings.Join(quotedStrs, ",") + "]"

	// 将每个字符串用双引号引起来 - descImages
	quotedStrs = make([]string, len(request.DescImages))
	for i, s := range request.Images {
		quotedStrs[i] = fmt.Sprintf("%q", s)
	}
	descImages := "[" + strings.Join(quotedStrs, ",") + "]"

	goods.BrandID = brandId
	goods.CategoryID = categoryId
	goods.Name = request.GetName()
	goods.GoodsSn = request.GetGoodsSn()
	goods.MarketPrice = request.GetMarketPrice()
	goods.ShopPrice = request.GetShopPrice()
	goods.GoodsBrief = request.GetGoodsBrief()
	goods.ShipFree = shipFree
	goods.Images = images
	goods.DescImages = descImages
	goods.GoodsFrontImage = request.GoodsFrontImage
	goods.IsNew = isNew
	goods.IsHot = isHot
	goods.OnSale = onSale

	_, _, err = dao.UpdateGoods(ctx, goodsId, goods)
	if err != nil {
		log.Error("dao.UpdateGoods fail", log.Any("err", err))
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (g GoodsService) GetGoodsDetail(ctx context.Context, request *proto.GoodInfoRequest) (*proto.GoodsInfoResponse, error) {
	id := request.GetId()

	goods, err := dao.GetGoods(ctx, id)
	if err != nil {
		log.Error("dao.GetGoods fail", log.Any("err", err))
		return nil, err
	}
	return g.getGoodsInfoResponse(ctx, goods)
}

func (g GoodsService) GetAllCategorysList(ctx context.Context, empty *emptypb.Empty) (*proto.CategoryListResponse, error) {
	// 找出所有分类
	categoryList, cnt, err := dao.GetAllCategory(ctx, 0, 0, "")
	if err != nil {
		return nil, err
	}

	// 将这些分类放入rsp.data
	data := make([]*proto.CategoryInfoResponse, 0, len(categoryList))
	level1Map := make(map[int32]*CategoryListItemForJson)
	level2Map := make(map[int32]*CategoryListItemForJson)
	level3List := make([]*CategoryListItemForJson, 0)
	for _, category := range categoryList {
		data = append(data, convertToCategoryInfoResponse(category))
		if category.Level == 1 {
			level1Map[category.ID] = convertToCategoryListItemForJson(category)
		} else if category.Level == 2 {
			level2Map[category.ID] = convertToCategoryListItemForJson(category)
		} else if category.Level == 3 {
			level3List = append(level3List, convertToCategoryListItemForJson(category))
		}
	}

	// 整理各个level，将一级，二级，三级分类依此嵌套起来
	for _, category3 := range level3List {
		if parentCategory, ok := level2Map[category3.ParentId]; ok {
			parentCategory.SubCategory = append(parentCategory.SubCategory, category3)
		}
	}
	for _, category2 := range level2Map {
		if parentCategory, ok := level1Map[category2.ParentId]; ok {
			parentCategory.SubCategory = append(parentCategory.SubCategory, category2)
		}
	}
	level1List := make([]*CategoryListItemForJson, 0, len(level1Map))
	for _, v := range level1Map {
		level1List = append(level1List, v)
	}
	jsonData, err := json.Marshal(level1List)
	if err != nil {
		return nil, err
	}

	// 组装数据
	rsp := &proto.CategoryListResponse{
		Total:    cnt,
		Data:     data,
		JsonData: string(jsonData),
	}

	return rsp, nil
}

func (g GoodsService) GetSubCategory(ctx context.Context, request *proto.CategoryListRequest) (*proto.SubCategoryListResponse, error) {
	categoryId := request.GetId()
	category, err := dao.GetCategory(ctx, categoryId)
	if errors.Is(err, dao.ErrNotFound) {
		return nil, err
	}
	if err != nil {
		log.Error("dao.GetCategory fail", log.Any("err", err))
		return nil, err
	}

	// 子分类
	subCategoryList, err := dao.GetCategoryListByParentId(ctx, categoryId)
	if err != nil {
		return nil, err
	}

	// 组装数据
	categoryPb := convertToCategoryInfoResponse(category)
	subCategoryPbList := make([]*proto.CategoryInfoResponse, 0, len(subCategoryList))
	for _, subCategory := range subCategoryList {
		subCategoryPbList = append(subCategoryPbList, convertToCategoryInfoResponse(subCategory))
	}

	rsp := &proto.SubCategoryListResponse{
		Total:        int64(len(subCategoryPbList)),
		Info:         categoryPb,
		SubCategorys: subCategoryPbList,
	}

	return rsp, nil
}

func (g GoodsService) CreateCategory(ctx context.Context, request *proto.CategoryInfoRequest) (*proto.CategoryInfoResponse, error) {
	var isTab int32 = 0
	if request.GetIsTab() {
		isTab = 1
	}
	var parentCategoryId int32 = 0
	valid := false
	if request.GetParentCategory() != 0 {
		parentCategoryId = request.GetParentCategory()
		valid = true
	}

	category := &model.Category{
		Name:             request.GetName(),
		ParentCategoryID: sql.NullInt64{Int64: int64(parentCategoryId), Valid: valid},
		Level:            request.GetLevel(),
		IsTab:            isTab,
		URL:              "",
	}
	result, _, err := dao.AddCategory(ctx, category)
	if err != nil {
		log.Error("dao.AddCategory fail", log.Any("err", err))
		return nil, err
	}
	rsp := &proto.CategoryInfoResponse{
		Id:             result.ID,
		Name:           result.Name,
		ParentCategory: int32(result.ParentCategoryID.Int64),
		Level:          result.Level,
		IsTab:          result.IsTab == 1,
	}
	return rsp, nil
}

func (g GoodsService) DeleteCategory(ctx context.Context, request *proto.DeleteCategoryRequest) (*emptypb.Empty, error) {
	id := request.GetId()
	_, err := dao.DeleteCategory(ctx, id)
	if err != nil {
		log.Error("dao.DeleteCategory fail", log.Any("err", err))
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (g GoodsService) UpdateCategory(ctx context.Context, request *proto.CategoryInfoRequest) (*emptypb.Empty, error) {
	id := request.GetId()
	category, err := dao.GetCategory(ctx, id)
	if errors.Is(err, dao.ErrNotFound) {
		return nil, err
	}
	if err != nil {
		log.Error("dao.GetCategory fail", log.Any("err", err))
		return nil, err
	}

	if request.Name != "" {
		category.Name = request.Name
	}
	if request.ParentCategory != 0 {
		category.ParentCategoryID = sql.NullInt64{Int64: int64(request.ParentCategory), Valid: true}
	}
	if request.Level != 0 {
		category.Level = request.Level
	}
	var isTab int32 = 0
	if request.IsTab {
		isTab = 1
	}
	category.IsTab = isTab
	_, _, err = dao.UpdateCategory(ctx, id, category)
	if err != nil {
		log.Error("dao.UpdateCategory fail", log.Any("err", err))
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (g GoodsService) BrandList(ctx context.Context, request *proto.BrandFilterRequest) (*proto.BrandListResponse, error) {
	page := request.GetPages()
	pageSize := request.GetPagePerNums()

	brandList, cnt, err := dao.GetAllBrands(ctx, int(page), int(pageSize), "")
	if err != nil {
		log.Error("dao.GetAllBrands fail", log.Any("err", err))
		return nil, err
	}

	brandPbList := make([]*proto.BrandInfoResponse, 0, len(brandList))
	for _, brand := range brandList {
		brandPbList = append(brandPbList, convertToBrandInfoResponse(brand))
	}
	rsp := &proto.BrandListResponse{
		Total: cnt,
		Data:  brandPbList,
	}
	return rsp, nil
}

func (g GoodsService) CreateBrand(ctx context.Context, request *proto.BrandRequest) (*proto.BrandInfoResponse, error) {
	name := request.GetName()
	logo := request.GetLogo()
	brand, err := dao.GetBrandsByName(ctx, name)
	if !errors.Is(err, dao.ErrNotFound) {
		if err != nil { // 意外错误
			return nil, err
		}
		return convertToBrandInfoResponse(brand), nil
	}

	brand = &model.Brands{
		Name: name,
		Logo: sql.NullString{String: logo, Valid: true},
	}

	result, _, err := dao.AddBrands(ctx, brand)
	if err != nil {
		log.Error("dao.Add.Brands fail", log.Any("err", err))
		return nil, err
	}
	return convertToBrandInfoResponse(result), nil
}

func (g GoodsService) DeleteBrand(ctx context.Context, request *proto.BrandRequest) (*emptypb.Empty, error) {
	id := request.GetId()
	_, err := dao.DeleteBrands(ctx, id)
	if err != nil {
		log.Error("dao.Add.DeleteBrands fail", log.Any("err", err))
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (g GoodsService) UpdateBrand(ctx context.Context, request *proto.BrandRequest) (*emptypb.Empty, error) {
	id := request.GetId()
	brand, err := dao.GetBrands(ctx, id)
	if err != nil {
		return nil, err
	}
	if request.GetName() != "" {
		brand.Name = request.GetName()
	}
	if request.GetLogo() != "" {
		brand.Logo = sql.NullString{String: request.GetLogo(), Valid: true}
	}

	_, _, err = dao.UpdateBrands(ctx, id, brand)
	if err != nil {
		log.Error("dao.UpdateBrands fail", log.Any("err", err))
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (g GoodsService) BannerList(ctx context.Context, empty *emptypb.Empty) (*proto.BannerListResponse, error) {
	bannerList, cnt, err := dao.GetAllBanner(ctx, 0, 0, "")
	if err != nil {
		log.Error("dao.GetAllBanner fail", log.Any("err", err))
		return nil, err
	}
	bannerPbList := make([]*proto.BannerResponse, 0, len(bannerList))
	for _, banner := range bannerList {
		bannerPbList = append(bannerPbList, convertToBannerResponse(banner))
	}

	rsp := &proto.BannerListResponse{
		Total: cnt,
		Data:  bannerPbList,
	}
	return rsp, nil
}

func (g GoodsService) CreateBanner(ctx context.Context, request *proto.BannerRequest) (*proto.BannerResponse, error) {
	banner := &model.Banner{
		Image: request.GetImage(),
		URL:   request.GetUrl(),
		Index: request.GetIndex(),
	}
	result, _, err := dao.AddBanner(ctx, banner)
	if err != nil {
		log.Error("dao.AddBanner fail", log.Any("err", err))
		return nil, err
	}
	return convertToBannerResponse(result), nil
}

func (g GoodsService) DeleteBanner(ctx context.Context, request *proto.BannerRequest) (*emptypb.Empty, error) {
	id := request.GetId()
	_, err := dao.DeleteBanner(ctx, id)
	if err != nil {
		log.Error("dao.DeleteBanner", log.Any("err", err))
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (g GoodsService) UpdateBanner(ctx context.Context, request *proto.BannerRequest) (*emptypb.Empty, error) {
	id := request.GetId()
	banner, err := dao.GetBanner(ctx, id)
	if err != nil {
		log.Error("dao.GetBanner fail", log.Any("err", err))
		return nil, err
	}
	if request.GetImage() != "" {
		banner.Image = request.GetImage()
	}
	if request.GetUrl() != "" {
		banner.URL = request.GetUrl()
	}
	if request.GetIndex() != 0 {
		banner.Index = request.GetIndex()
	}

	_, _, err = dao.UpdateBanner(ctx, id, banner)
	if err != nil {
		log.Error("dao.UpdateBanner fail", log.Any("err", err))
		return nil, err
	}
	return &emptypb.Empty{}, nil
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
