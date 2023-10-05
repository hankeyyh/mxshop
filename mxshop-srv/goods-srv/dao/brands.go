package dao

import (
	"context"
	"errors"
	"gorm.io/gorm"
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

// GetAllBrands is a function to get a slice of record(s) from brands table in the mxshop_goods_srv database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - ErrNotFound, db Find error
func GetAllBrands(ctx context.Context, page, pagesize int, order string) (results []*model.Brands, totalRows int64, err error) {

	resultOrm := DB.Model(&model.Brands{})
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

// GetBrands is a function to get a single record from the brands table in the mxshop_goods_srv database
// error - ErrNotFound, db Find error
func GetBrands(ctx context.Context, argID int32) (record *model.Brands, err error) {
	record = &model.Brands{}
	if err = DB.First(record, argID).Error; err != nil {
		err = ErrNotFound
		return record, err
	}

	return record, nil
}

func GetBrandsByName(ctx context.Context, name string) (record *model.Brands, err error) {
	if err = DB.Where("name = ? and is_deleted = ?", name, 0).First(&record).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = ErrNotFound
			return
		}
	}
	return record, err
}

func GetBrandsList(ctx context.Context, idList []int32) (recList []*model.Brands, err error) {
	err = DB.Where("id in ? and is_deleted = ?", idList, 0).Find(&recList).Error
	return
}

// AddBrands is a function to add a single record to brands table in the mxshop_goods_srv database
// error - ErrInsertFailed, db save call failed
func AddBrands(ctx context.Context, record *model.Brands) (result *model.Brands, RowsAffected int64, err error) {
	db := DB.Save(record)
	if err = db.Error; err != nil {
		return nil, -1, ErrInsertFailed
	}

	return record, db.RowsAffected, nil
}

// UpdateBrands is a function to update a single record from brands table in the mxshop_goods_srv database
// error - ErrNotFound, db record for id not found
// error - ErrUpdateFailed, db meta data copy failed or db.Save call failed
func UpdateBrands(ctx context.Context, argID int32, updated *model.Brands) (result *model.Brands, RowsAffected int64, err error) {

	result = &model.Brands{}
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

// DeleteBrands is a function to delete a single record from brands table in the mxshop_goods_srv database
// error - ErrNotFound, db Find error
// error - ErrDeleteFailed, db Delete failed error
func DeleteBrands(ctx context.Context, argID int32) (rowsAffected int64, err error) {

	record := &model.Brands{}
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
