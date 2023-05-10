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


CREATE TABLE `category` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `name` varchar(100) COLLATE utf8mb4_general_ci NOT NULL COMMENT '名称',
  `level` int NOT NULL DEFAULT '1' COMMENT '级别',
  `is_tab` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否显示在首页tab，0:否，1:是',
  `parent_category` bigint unsigned DEFAULT '0' COMMENT '父类别',
  `is_deleted` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否删除，0:否，1:是',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `category_FK` (`parent_category`),
  CONSTRAINT `category_FK` FOREIGN KEY (`parent_category`) REFERENCES `category` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='分类'

JSON Sample
-------------------------------------
{    "id": 13,    "name": "ahYZJUgEVhRMXIvtQHZAoRyYK",    "level": 88,    "is_tab": 94,    "parent_category": 17,    "is_deleted": 51,    "created_at": "2301-09-29T23:00:22.007087025+08:00",    "updated_at": "2196-01-25T05:53:34.816929147+08:00"}


Comments
-------------------------------------
[ 0] column is set for unsigned
[ 4] column is set for unsigned



*/

// Category struct is a row record of the category table in the mxshop_goods_srv database
type Category struct {
	//[ 0] id                                             ubigint              null: false  primary: true   isArray: false  auto: true   col: ubigint         len: -1      default: []
	ID uint64 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:ubigint;"` // 主键
	//[ 1] name                                           varchar(100)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 100     default: []
	Name string `gorm:"column:name;type:varchar;size:100;"` // 名称
	//[ 2] level                                          int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [1]
	Level int32 `gorm:"column:level;type:int;default:1;"` // 级别
	//[ 3] is_tab                                         tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [0]
	IsTab int32 `gorm:"column:is_tab;type:tinyint;default:0;"` // 是否显示在首页tab，0:否，1:是
	//[ 4] parent_category                                ubigint              null: true   primary: false  isArray: false  auto: false  col: ubigint         len: -1      default: [0]
	ParentCategory sql.NullInt64 `gorm:"column:parent_category;type:ubigint;default:0;"` // 父类别
	//[ 5] is_deleted                                     tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [0]
	IsDeleted int32 `gorm:"column:is_deleted;type:tinyint;default:0;"` // 是否删除，0:否，1:是
	//[ 6] created_at                                     timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [CURRENT_TIMESTAMP]
	CreatedAt time.Time `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP;"` // 创建时间
	//[ 7] updated_at                                     timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [CURRENT_TIMESTAMP]
	UpdatedAt time.Time `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP;"` // 更新时间

}

var categoryTableInfo = &TableInfo{
	Name: "category",
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
			Name:               "level",
			Comment:            `级别`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "Level",
			GoFieldType:        "int32",
			JSONFieldName:      "level",
			ProtobufFieldName:  "level",
			ProtobufType:       "int32",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "is_tab",
			Comment:            `是否显示在首页tab，0:否，1:是`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "tinyint",
			DatabaseTypePretty: "tinyint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "tinyint",
			ColumnLength:       -1,
			GoFieldName:        "IsTab",
			GoFieldType:        "int32",
			JSONFieldName:      "is_tab",
			ProtobufFieldName:  "is_tab",
			ProtobufType:       "int32",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "parent_category",
			Comment:            `父类别`,
			Notes:              `column is set for unsigned`,
			Nullable:           true,
			DatabaseTypeName:   "ubigint",
			DatabaseTypePretty: "ubigint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "ubigint",
			ColumnLength:       -1,
			GoFieldName:        "ParentCategory",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "parent_category",
			ProtobufFieldName:  "parent_category",
			ProtobufType:       "uint64",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "is_deleted",
			Comment:            `是否删除，0:否，1:是`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "tinyint",
			DatabaseTypePretty: "tinyint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "tinyint",
			ColumnLength:       -1,
			GoFieldName:        "IsDeleted",
			GoFieldType:        "int32",
			JSONFieldName:      "is_deleted",
			ProtobufFieldName:  "is_deleted",
			ProtobufType:       "int32",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
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
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
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
			ProtobufPos:        8,
		},
	},
}

// TableName sets the insert table name for this struct type
func (c *Category) TableName() string {
	return "category"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (c *Category) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (c *Category) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (c *Category) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (c *Category) TableInfo() *TableInfo {
	return categoryTableInfo
}
