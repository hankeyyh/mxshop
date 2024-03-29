package dao

import (
	"context"
	"time"

	"github.com/hankeyyh/mxshop/mxshop-srv/goods-srv/model"

	"github.com/guregu/null"
	"github.com/satori/go.uuid"
)

var (
	_ = time.Second
	_ = null.Bool{}
	_ = uuid.UUID{}
)

// GetAllGoods is a function to get a slice of record(s) from goods table in the mxshop_goods_srv database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - ErrNotFound, db Find error
func GetAllGoods(ctx context.Context, page, pagesize int, order string) (results []*model.Goods, totalRows int64, err error) {

	resultOrm := DB.Model(&model.Goods{})
	resultOrm.Count(&totalRows)

	if page > 0 {
		offset := (page - 1) * pagesize
		resultOrm = resultOrm.Offset(offset).Limit(pagesize)
	} else {
		resultOrm = resultOrm.Limit(pagesize)
	}

	if order != "" {
		resultOrm = resultOrm.Order(order)
	}

	if err = resultOrm.Find(&results).Error; err != nil {
		err = ErrNotFound
		return nil, -1, err
	}

	return results, totalRows, nil
}

// GetGoods is a function to get a single record from the goods table in the mxshop_goods_srv database
// error - ErrNotFound, db Find error
func GetGoods(ctx context.Context, argID int32) (record *model.Goods, err error) {
	record = &model.Goods{}
	if err = DB.First(record, argID).Error; err != nil {
		err = ErrNotFound
		return record, err
	}

	return record, nil
}

type GoodsWhere struct {
	PriceMin       int32
	PriceMax       int32
	IsHot          bool
	IsNew          bool
	CategoryIdList []int32
	Pages          int32
	PagePerNums    int32
	KeyWords       string
	BrandId        int32
}

// 商品列表
func GetGoodsList(ctx context.Context, where GoodsWhere) (recList []*model.Goods, cnt int64, err error) {
	query := DB.Model(&model.Goods{})
	if where.PriceMin > 0 {
		query.Where("shop_price >= ?", where.PriceMin)
	}
	if where.PriceMax > 0 {
		query.Where("shop_price <= ?", where.PriceMax)
	}
	if where.IsHot {
		query.Where("is_hot = ?", 1)
	}
	if where.IsNew {
		query.Where("is_new = ?", 1)
	}
	if where.KeyWords != "" {
		query.Where("name like ?", "%"+where.KeyWords+"%")
	}
	if where.BrandId != 0 {
		query.Where("brand_id = ?", where.BrandId)
	}
	if len(where.CategoryIdList) > 0 {
		query.Where("category_id in ?", where.CategoryIdList)
	}
	err = query.Count(&cnt).Error
	if err != nil {
		return
	}
	offset := (where.Pages - 1) * where.PagePerNums
	err = query.Offset(int(offset)).Limit(int(where.PagePerNums)).Find(&recList).Error
	return
}

// 批量获取goods
func BatchGetGoods(ctx context.Context, idList []int32) (recList []*model.Goods, err error) {
	if err = DB.Where("id in ?", idList).Find(&recList).Error; err != nil {
		return nil, err
	}
	return recList, nil
}

// AddGoods is a function to add a single record to goods table in the mxshop_goods_srv database
// error - ErrInsertFailed, db save call failed
func AddGoods(ctx context.Context, record *model.Goods) (result *model.Goods, RowsAffected int64, err error) {
	db := DB.Save(record)
	if err = db.Error; err != nil {
		return nil, -1, ErrInsertFailed
	}

	return record, db.RowsAffected, nil
}

// UpdateGoods is a function to update a single record from goods table in the mxshop_goods_srv database
// error - ErrNotFound, db record for id not found
// error - ErrUpdateFailed, db meta data copy failed or db.Save call failed
func UpdateGoods(ctx context.Context, argID int32, updated *model.Goods) (result *model.Goods, RowsAffected int64, err error) {

	result = &model.Goods{}
	db := DB.First(result, argID)
	if err = db.Error; err != nil {
		return nil, -1, ErrNotFound
	}

	if err = Copy(result, updated); err != nil {
		return nil, -1, ErrUpdateFailed
	}

	db = db.Save(result)
	if err = db.Error; err != nil {
		return nil, -1, ErrUpdateFailed
	}

	return result, db.RowsAffected, nil
}

// DeleteGoods is a function to delete a single record from goods table in the mxshop_goods_srv database
// error - ErrNotFound, db Find error
// error - ErrDeleteFailed, db Delete failed error
func DeleteGoods(ctx context.Context, argID int32) (rowsAffected int64, err error) {

	record := &model.Goods{}
	db := DB.First(record, argID)
	if db.Error != nil {
		return -1, ErrNotFound
	}

	db = db.Delete(record)
	if err = db.Error; err != nil {
		return -1, ErrDeleteFailed
	}

	return db.RowsAffected, nil
}
