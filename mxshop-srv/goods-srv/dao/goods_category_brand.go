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

// GetAllGoodsCategoryBrand is a function to get a slice of record(s) from goods_category_brand table in the mxshop_goods_srv database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - ErrNotFound, db Find error
func GetAllGoodsCategoryBrand(ctx context.Context, page, pagesize int, order string) (results []*model.GoodsCategoryBrand, totalRows int64, err error) {

	resultOrm := DB.Model(&model.GoodsCategoryBrand{})
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

// GetGoodsCategoryBrand is a function to get a single record from the goods_category_brand table in the mxshop_goods_srv database
// error - ErrNotFound, db Find error
func GetGoodsCategoryBrand(ctx context.Context, argID int32) (record *model.GoodsCategoryBrand, err error) {
	record = &model.GoodsCategoryBrand{}
	if err = DB.First(record, argID).Error; err != nil {
		err = ErrNotFound
		return record, err
	}

	return record, nil
}

func GetGoodsCategoryBrandListByCategoryId(ctx context.Context, categoryId int32) (results []*model.GoodsCategoryBrand, err error) {
	err = DB.Where("category_id = ? and is_deleted = ?", categoryId, 0).Find(&results).Error
	return
}

// AddGoodsCategoryBrand is a function to add a single record to goods_category_brand table in the mxshop_goods_srv database
// error - ErrInsertFailed, db save call failed
func AddGoodsCategoryBrand(ctx context.Context, record *model.GoodsCategoryBrand) (result *model.GoodsCategoryBrand, RowsAffected int64, err error) {
	db := DB.Save(record)
	if err = db.Error; err != nil {
		return nil, -1, ErrInsertFailed
	}

	return record, db.RowsAffected, nil
}

// UpdateGoodsCategoryBrand is a function to update a single record from goods_category_brand table in the mxshop_goods_srv database
// error - ErrNotFound, db record for id not found
// error - ErrUpdateFailed, db meta data copy failed or db.Save call failed
func UpdateGoodsCategoryBrand(ctx context.Context, argID int32, updated *model.GoodsCategoryBrand) (result *model.GoodsCategoryBrand, RowsAffected int64, err error) {

	result = &model.GoodsCategoryBrand{}
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

// DeleteGoodsCategoryBrand is a function to delete a single record from goods_category_brand table in the mxshop_goods_srv database
// error - ErrNotFound, db Find error
// error - ErrDeleteFailed, db Delete failed error
func DeleteGoodsCategoryBrand(ctx context.Context, argID int32) (rowsAffected int64, err error) {

	record := &model.GoodsCategoryBrand{}
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
