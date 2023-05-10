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


CREATE TABLE `banner` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `image` varchar(200) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '图片url',
  `url` varchar(200) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '访问url',
  `index` int NOT NULL DEFAULT '0' COMMENT '轮播顺序',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='轮播的商品'

JSON Sample
-------------------------------------
{    "id": 20,    "image": "hPVebcSsgonpyjEmNpKuwjbPx",    "url": "hIMmkGdlnDABtvoilOOcumdOB",    "index": 87,    "created_at": "2155-05-31T02:29:42.04095111+08:00",    "updated_at": "2122-04-06T16:45:16.469172059+08:00"}


Comments
-------------------------------------
[ 0] column is set for unsigned



*/

// Banner struct is a row record of the banner table in the mxshop_goods_srv database
type Banner struct {
	//[ 0] id                                             ubigint              null: false  primary: true   isArray: false  auto: true   col: ubigint         len: -1      default: []
	ID uint64 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:ubigint;"` // 主键
	//[ 1] image                                          varchar(200)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 200     default: []
	Image string `gorm:"column:image;type:varchar;size:200;"` // 图片url
	//[ 2] url                                            varchar(200)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 200     default: []
	URL string `gorm:"column:url;type:varchar;size:200;"` // 访问url
	//[ 3] index                                          int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	Index int32 `gorm:"column:index;type:int;default:0;"` // 轮播顺序
	//[ 4] created_at                                     timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [CURRENT_TIMESTAMP]
	CreatedAt time.Time `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP;"` // 创建时间
	//[ 5] updated_at                                     timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [CURRENT_TIMESTAMP]
	UpdatedAt time.Time `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP;"` // 更新时间

}

var bannerTableInfo = &TableInfo{
	Name: "banner",
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
			Name:               "image",
			Comment:            `图片url`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(200)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       200,
			GoFieldName:        "Image",
			GoFieldType:        "string",
			JSONFieldName:      "image",
			ProtobufFieldName:  "image",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "url",
			Comment:            `访问url`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(200)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       200,
			GoFieldName:        "URL",
			GoFieldType:        "string",
			JSONFieldName:      "url",
			ProtobufFieldName:  "url",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "index",
			Comment:            `轮播顺序`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "Index",
			GoFieldType:        "int32",
			JSONFieldName:      "index",
			ProtobufFieldName:  "index",
			ProtobufType:       "int32",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
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
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
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
			ProtobufPos:        6,
		},
	},
}

// TableName sets the insert table name for this struct type
func (b *Banner) TableName() string {
	return "banner"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (b *Banner) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (b *Banner) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (b *Banner) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (b *Banner) TableInfo() *TableInfo {
	return bannerTableInfo
}
