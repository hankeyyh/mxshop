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

// GetAllBanner is a function to get a slice of record(s) from banner table in the mxshop_goods_srv database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - ErrNotFound, db Find error
func GetAllBanner(ctx context.Context, page, pagesize int, order string) (results []*model.Banner, totalRows int64, err error) {

	resultOrm := DB.Model(&model.Banner{})
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

// GetBanner is a function to get a single record from the banner table in the mxshop_goods_srv database
// error - ErrNotFound, db Find error
func GetBanner(ctx context.Context, argID int32) (record *model.Banner, err error) {
	record = &model.Banner{}
	if err = DB.First(record, argID).Error; err != nil {
		err = ErrNotFound
		return record, err
	}

	return record, nil
}

// AddBanner is a function to add a single record to banner table in the mxshop_goods_srv database
// error - ErrInsertFailed, db save call failed
func AddBanner(ctx context.Context, record *model.Banner) (result *model.Banner, RowsAffected int64, err error) {
	db := DB.Save(record)
	if err = db.Error; err != nil {
		return nil, -1, ErrInsertFailed
	}

	return record, db.RowsAffected, nil
}

// UpdateBanner is a function to update a single record from banner table in the mxshop_goods_srv database
// error - ErrNotFound, db record for id not found
// error - ErrUpdateFailed, db meta data copy failed or db.Save call failed
func UpdateBanner(ctx context.Context, argID int32, updated *model.Banner) (result *model.Banner, RowsAffected int64, err error) {

	result = &model.Banner{}
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

// DeleteBanner is a function to delete a single record from banner table in the mxshop_goods_srv database
// error - ErrNotFound, db Find error
// error - ErrDeleteFailed, db Delete failed error
func DeleteBanner(ctx context.Context, argID int32) (rowsAffected int64, err error) {

	record := &model.Banner{}
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
