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
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `category_id` bigint unsigned NOT NULL COMMENT '分类id',
  `brand_id` bigint unsigned NOT NULL COMMENT '品牌id',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `goods_category_brand_FK` (`category_id`),
  KEY `goods_category_brand_FK_1` (`brand_id`),
  CONSTRAINT `goods_category_brand_FK` FOREIGN KEY (`category_id`) REFERENCES `category` (`id`),
  CONSTRAINT `goods_category_brand_FK_1` FOREIGN KEY (`brand_id`) REFERENCES `brand` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='品牌分类'

JSON Sample
-------------------------------------
{    "id": 96,    "category_id": 99,    "brand_id": 27,    "created_at": "2244-02-10T18:04:18.202881266+08:00",    "updated_at": "2111-05-14T13:03:19.006701382+08:00"}


Comments
-------------------------------------
[ 0] column is set for unsigned
[ 1] column is set for unsigned
[ 2] column is set for unsigned



*/

// GoodsCategoryBrand struct is a row record of the goods_category_brand table in the mxshop_goods_srv database
type GoodsCategoryBrand struct {
	//[ 0] id                                             ubigint              null: false  primary: true   isArray: false  auto: true   col: ubigint         len: -1      default: []
	ID uint64 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:ubigint;"` // 主键
	//[ 1] category_id                                    ubigint              null: false  primary: false  isArray: false  auto: false  col: ubigint         len: -1      default: []
	CategoryID uint64 `gorm:"column:category_id;type:ubigint;"` // 分类id
	//[ 2] brand_id                                       ubigint              null: false  primary: false  isArray: false  auto: false  col: ubigint         len: -1      default: []
	BrandID uint64 `gorm:"column:brand_id;type:ubigint;"` // 品牌id
	//[ 3] created_at                                     timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [CURRENT_TIMESTAMP]
	CreatedAt time.Time `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP;"` // 创建时间
	//[ 4] updated_at                                     timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [CURRENT_TIMESTAMP]
	UpdatedAt time.Time `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP;"` // 更新时间

}

var goods_category_brandTableInfo = &TableInfo{
	Name: "goods_category_brand",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "id",
			Comment:            `主键`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "ubigint",
			DatabaseTypePretty: "ubigint",
			IsPrimaryKey:       true,
			IsAutoIncrement:    true,
			IsArray:            false,
			ColumnType:         "ubigint",
			ColumnLength:       -1,
			GoFieldName:        "ID",
			GoFieldType:        "uint64",
			JSONFieldName:      "id",
			ProtobufFieldName:  "id",
			ProtobufType:       "uint64",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "category_id",
			Comment:            `分类id`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "ubigint",
			DatabaseTypePretty: "ubigint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "ubigint",
			ColumnLength:       -1,
			GoFieldName:        "CategoryID",
			GoFieldType:        "uint64",
			JSONFieldName:      "category_id",
			ProtobufFieldName:  "category_id",
			ProtobufType:       "uint64",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "brand_id",
			Comment:            `品牌id`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "ubigint",
			DatabaseTypePretty: "ubigint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "ubigint",
			ColumnLength:       -1,
			GoFieldName:        "BrandID",
			GoFieldType:        "uint64",
			JSONFieldName:      "brand_id",
			ProtobufFieldName:  "brand_id",
			ProtobufType:       "uint64",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "created_at",
			Comment:            `创建时间`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "timestamp",
			DatabaseTypePretty: "timestamp",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "timestamp",
			ColumnLength:       -1,
			GoFieldName:        "CreatedAt",
			GoFieldType:        "time.Time",
			JSONFieldName:      "created_at",
			ProtobufFieldName:  "created_at",
			ProtobufType:       "uint64",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "updated_at",
			Comment:            `更新时间`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "timestamp",
			DatabaseTypePretty: "timestamp",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "timestamp",
			ColumnLength:       -1,
			GoFieldName:        "UpdatedAt",
			GoFieldType:        "time.Time",
			JSONFieldName:      "updated_at",
			ProtobufFieldName:  "updated_at",
			ProtobufType:       "uint64",
			ProtobufPos:        5,
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
