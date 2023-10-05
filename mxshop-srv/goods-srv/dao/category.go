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

// GetAllCategory is a function to get a slice of record(s) from category table in the mxshop_goods_srv database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - ErrNotFound, db Find error
func GetAllCategory(ctx context.Context, page, pagesize int, order string) (results []*model.Category, totalRows int64, err error) {

	resultOrm := DB.Model(&model.Category{})
	resultOrm.Count(&totalRows)

	if page > 0 {
		offset := (page - 1) * pagesize
		resultOrm = resultOrm.Offset(offset).Limit(pagesize)
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

// GetCategory is a function to get a single record from the category table in the mxshop_goods_srv database
// error - ErrNotFound, db Find error
func GetCategory(ctx context.Context, argID int32) (record *model.Category, err error) {
	record = &model.Category{}
	if err = DB.First(record, argID).Error; err != nil {
		err = ErrNotFound
		return record, err
	}

	return record, nil
}

func GetCategoryList(ctx context.Context, idList []int32) (recList []*model.Category, err error) {
	err = DB.Where("id in ? and is_deleted = ?", idList, 0).Find(&recList).Error
	return
}

// 查询最下级分类
func GetLowestCategoryList(ctx context.Context, topCategoryId int32) (recList []*model.Category, err error) {
	var topCategory model.Category
	if err = DB.Where("id = ? and is_deleted = ?", topCategoryId, 0).First(&topCategory).Error; err != nil {
		return
	}

	var parentIdList = []int32{topCategoryId}
	for {
		var childIdList []int32
		err = DB.Model(&model.Category{}).Where("parent_category_id in ? and is_deleted = ?", parentIdList, 0).
			Pluck("id", &childIdList).Error
		if err != nil {
			return
		}
		if len(childIdList) > 0 {
			parentIdList = childIdList
		} else {
			break
		}
	}
	//if topCategory.Level == 1 {
	//	// 二级分类
	//	subQuery := DB.Select("id").Where("parent_category_id = ? and is_deleted = ?", topCategoryId, 0).Table("category")
	//	// 三级分类
	//	err = DB.Where("parent_category_id in ? and is_deleted = ?", subQuery, 0).Pluck("id", &parentIdList).Error
	//	if err != nil {
	//		return
	//	}
	//} else if topCategory.Level == 2 {
	//	err = DB.Where("parent_category_id = ? and is_deleted = ?", topCategoryId, 0).Pluck("id", &parentIdList).Error
	//	if err != nil {
	//		return
	//	}
	//} else if topCategory.Level == 3 {
	//	// 三级分类已是最下层分类，直接返回
	//	return []*model.Category{&topCategory}, nil
	//} else {
	//	err = ErrWrongLevel
	//	return
	//}
	// 查询最下级分类
	err = DB.Where("id in ? and is_deleted = ?", parentIdList, 0).Find(&recList).Error
	return
}

// AddCategory is a function to add a single record to category table in the mxshop_goods_srv database
// error - ErrInsertFailed, db save call failed
func AddCategory(ctx context.Context, record *model.Category) (result *model.Category, RowsAffected int64, err error) {
	db := DB.Save(record)
	if err = db.Error; err != nil {
		return nil, -1, ErrInsertFailed
	}

	return record, db.RowsAffected, nil
}

// UpdateCategory is a function to update a single record from category table in the mxshop_goods_srv database
// error - ErrNotFound, db record for id not found
// error - ErrUpdateFailed, db meta data copy failed or db.Save call failed
func UpdateCategory(ctx context.Context, argID int32, updated *model.Category) (result *model.Category, RowsAffected int64, err error) {

	result = &model.Category{}
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

// DeleteCategory is a function to delete a single record from category table in the mxshop_goods_srv database
// error - ErrNotFound, db Find error
// error - ErrDeleteFailed, db Delete failed error
func DeleteCategory(ctx context.Context, argID int32) (rowsAffected int64, err error) {

	record := &model.Category{}
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
