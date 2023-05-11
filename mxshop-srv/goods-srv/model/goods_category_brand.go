package model

import (
	"database/sql"
	"time"

	"github.com/guregu/null"
	"github.com/satori/go.uuid"
)

var (
	_ = time.Second
	_ = sql.LevelDefault
	_ = null.Bool{}
	_ = uuid.UUID{}
)

/*
DB Table Details
-------------------------------------


CREATE TABLE `goods_category_brand` (
  `id` int NOT NULL AUTO_INCREMENT,
  `category_id` int NOT NULL,
  `brand_id` int NOT NULL,
  `add_time` datetime NOT NULL,
  `is_deleted` tinyint(1) DEFAULT NULL,
  `update_time` datetime NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `goodscategorybrand_category_id_brand_id` (`category_id`,`brand_id`) USING BTREE,
  KEY `goodscategorybrand_category_id` (`category_id`) USING BTREE,
  KEY `goodscategorybrand_brand_id` (`brand_id`) USING BTREE,
  CONSTRAINT `goods_category_brand_ibfk_1` FOREIGN KEY (`category_id`) REFERENCES `category` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `goods_category_brand_ibfk_2` FOREIGN KEY (`brand_id`) REFERENCES `brands` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE=InnoDB AUTO_INCREMENT=25799 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC

JSON Sample
-------------------------------------
{    "id": 12,    "category_id": 52,    "brand_id": 13,    "add_time": "2252-03-11T20:39:00.128265171+08:00",    "is_deleted": 67,    "update_time": "2203-04-05T18:49:17.866700754+08:00"}



*/

// GoodsCategoryBrand struct is a row record of the goods_category_brand table in the mxshop_goods_srv database
type GoodsCategoryBrand struct {
	//[ 0] id                                             int                  null: false  primary: true   isArray: false  auto: true   col: int             len: -1      default: []
	ID int32 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:int;"`
	//[ 1] category_id                                    int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	CategoryID int32 `gorm:"column:category_id;type:int;"`
	//[ 2] brand_id                                       int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	BrandID int32 `gorm:"column:brand_id;type:int;"`
	//[ 3] add_time                                       datetime             null: false  primary: false  isArray: false  auto: false  col: datetime        len: -1      default: []
	AddTime time.Time `gorm:"column:add_time;type:datetime;"`
	//[ 4] is_deleted                                     tinyint              null: true   primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: []
	IsDeleted sql.NullInt64 `gorm:"column:is_deleted;type:tinyint;"`
	//[ 5] update_time                                    datetime             null: false  primary: false  isArray: false  auto: false  col: datetime        len: -1      default: []
	UpdateTime time.Time `gorm:"column:update_time;type:datetime;"`
}

var goods_category_brandTableInfo = &TableInfo{
	Name: "goods_category_brand",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "id",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       true,
			IsAutoIncrement:    true,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "ID",
			GoFieldType:        "int32",
			JSONFieldName:      "id",
			ProtobufFieldName:  "id",
			ProtobufType:       "int32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "category_id",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "CategoryID",
			GoFieldType:        "int32",
			JSONFieldName:      "category_id",
			ProtobufFieldName:  "category_id",
			ProtobufType:       "int32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "brand_id",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "BrandID",
			GoFieldType:        "int32",
			JSONFieldName:      "brand_id",
			ProtobufFieldName:  "brand_id",
			ProtobufType:       "int32",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "add_time",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "datetime",
			DatabaseTypePretty: "datetime",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "datetime",
			ColumnLength:       -1,
			GoFieldName:        "AddTime",
			GoFieldType:        "time.Time",
			JSONFieldName:      "add_time",
			ProtobufFieldName:  "add_time",
			ProtobufType:       "google.protobuf.Timestamp",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "is_deleted",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "tinyint",
			DatabaseTypePretty: "tinyint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "tinyint",
			ColumnLength:       -1,
			GoFieldName:        "IsDeleted",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "is_deleted",
			ProtobufFieldName:  "is_deleted",
			ProtobufType:       "int32",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "update_time",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "datetime",
			DatabaseTypePretty: "datetime",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "datetime",
			ColumnLength:       -1,
			GoFieldName:        "UpdateTime",
			GoFieldType:        "time.Time",
			JSONFieldName:      "update_time",
			ProtobufFieldName:  "update_time",
			ProtobufType:       "google.protobuf.Timestamp",
			ProtobufPos:        6,
		},
	},
}

// TableName sets the insert table name for this struct type
func (g *GoodsCategoryBrand) TableName() string {
	return "goods_category_brand"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (g *GoodsCategoryBrand) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (g *GoodsCategoryBrand) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (g *GoodsCategoryBrand) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (g *GoodsCategoryBrand) TableInfo() *TableInfo {
	return goods_category_brandTableInfo
}
