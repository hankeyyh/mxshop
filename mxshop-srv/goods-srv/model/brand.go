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


CREATE TABLE `brand` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `name` varchar(100) COLLATE utf8mb4_general_ci NOT NULL COMMENT '名称',
  `logo` varchar(100) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '图标',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='品牌'

JSON Sample
-------------------------------------
{    "id": 77,    "name": "YatcUBdLTThTLIqfnwRgfcXuJ",    "logo": "XpELHjeaQwMpnlaVKgSAjWvRk",    "created_at": "2119-10-19T07:37:08.16287268+08:00",    "updated_at": "2217-04-21T05:27:57.07868844+08:00"}


Comments
-------------------------------------
[ 0] column is set for unsigned



*/

// Brand struct is a row record of the brand table in the mxshop_goods_srv database
type Brand struct {
	//[ 0] id                                             ubigint              null: false  primary: true   isArray: false  auto: true   col: ubigint         len: -1      default: []
	ID uint64 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:ubigint;"` // 主键
	//[ 1] name                                           varchar(100)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 100     default: []
	Name string `gorm:"column:name;type:varchar;size:100;"` // 名称
	//[ 2] logo                                           varchar(100)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 100     default: []
	Logo string `gorm:"column:logo;type:varchar;size:100;"` // 图标
	//[ 3] created_at                                     timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [CURRENT_TIMESTAMP]
	CreatedAt time.Time `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP;"` // 创建时间
	//[ 4] updated_at                                     timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [CURRENT_TIMESTAMP]
	UpdatedAt time.Time `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP;"` // 更新时间

}

var brandTableInfo = &TableInfo{
	Name: "brand",
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
			Name:               "name",
			Comment:            `名称`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(100)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       100,
			GoFieldName:        "Name",
			GoFieldType:        "string",
			JSONFieldName:      "name",
			ProtobufFieldName:  "name",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "logo",
			Comment:            `图标`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(100)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       100,
			GoFieldName:        "Logo",
			GoFieldType:        "string",
			JSONFieldName:      "logo",
			ProtobufFieldName:  "logo",
			ProtobufType:       "string",
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
func (b *Brand) TableName() string {
	return "brand"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (b *Brand) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (b *Brand) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (b *Brand) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (b *Brand) TableInfo() *TableInfo {
	return brandTableInfo
}
