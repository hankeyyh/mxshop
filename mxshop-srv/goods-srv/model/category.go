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
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `parent_category_id` int DEFAULT NULL,
  `level` int NOT NULL,
  `is_tab` tinyint(1) NOT NULL,
  `url` varchar(300) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `add_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `is_deleted` tinyint NOT NULL DEFAULT '0',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE,
  KEY `category_name` (`name`) USING BTREE,
  KEY `category_parent_category_id` (`parent_category_id`) USING BTREE,
  KEY `category_url` (`url`) USING BTREE,
  CONSTRAINT `category_ibfk_1` FOREIGN KEY (`parent_category_id`) REFERENCES `category` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE=InnoDB AUTO_INCREMENT=238023 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC

JSON Sample
-------------------------------------
{    "id": 82,    "name": "jXiaZgRFnfeBTeSPLaQDvLyQJ",    "parent_category_id": 51,    "level": 50,    "is_tab": 44,    "url": "QADqRlqDeBhvDMqlFePYvkRkA",    "add_time": "2248-02-16T13:19:09.595912119+08:00",    "is_deleted": 35,    "update_time": "2024-05-31T23:33:04.421092795+08:00"}



*/

// Category struct is a row record of the category table in the mxshop_goods_srv database
type Category struct {
	//[ 0] id                                             int                  null: false  primary: true   isArray: false  auto: true   col: int             len: -1      default: []
	ID int32 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:int;"`
	//[ 1] name                                           varchar(50)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 50      default: []
	Name string `gorm:"column:name;type:varchar;size:50;"`
	//[ 2] parent_category_id                             int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	ParentCategoryID sql.NullInt64 `gorm:"column:parent_category_id;type:int;"`
	//[ 3] level                                          int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	Level int32 `gorm:"column:level;type:int;"`
	//[ 4] is_tab                                         tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: []
	IsTab int32 `gorm:"column:is_tab;type:tinyint;"`
	//[ 5] url                                            varchar(300)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 300     default: []
	URL string `gorm:"column:url;type:varchar;size:300;"`
	//[ 6] add_time                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [CURRENT_TIMESTAMP]
	AddTime time.Time `gorm:"column:add_time;type:timestamp;default:CURRENT_TIMESTAMP;"`
	//[ 7] is_deleted                                     tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [0]
	IsDeleted int32 `gorm:"column:is_deleted;type:tinyint;default:0;"`
	//[ 8] update_time                                    timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [CURRENT_TIMESTAMP]
	UpdateTime time.Time `gorm:"column:update_time;type:timestamp;default:CURRENT_TIMESTAMP;"`
}

var categoryTableInfo = &TableInfo{
	Name: "category",
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
			Name:               "name",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(50)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       50,
			GoFieldName:        "Name",
			GoFieldType:        "string",
			JSONFieldName:      "name",
			ProtobufFieldName:  "name",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "parent_category_id",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "ParentCategoryID",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "parent_category_id",
			ProtobufFieldName:  "parent_category_id",
			ProtobufType:       "int32",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "level",
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
			GoFieldName:        "Level",
			GoFieldType:        "int32",
			JSONFieldName:      "level",
			ProtobufFieldName:  "level",
			ProtobufType:       "int32",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "is_tab",
			Comment:            ``,
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
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "url",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(300)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       300,
			GoFieldName:        "URL",
			GoFieldType:        "string",
			JSONFieldName:      "url",
			ProtobufFieldName:  "url",
			ProtobufType:       "string",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "add_time",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "timestamp",
			DatabaseTypePretty: "timestamp",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "timestamp",
			ColumnLength:       -1,
			GoFieldName:        "AddTime",
			GoFieldType:        "time.Time",
			JSONFieldName:      "add_time",
			ProtobufFieldName:  "add_time",
			ProtobufType:       "uint64",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "is_deleted",
			Comment:            ``,
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
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "update_time",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "timestamp",
			DatabaseTypePretty: "timestamp",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "timestamp",
			ColumnLength:       -1,
			GoFieldName:        "UpdateTime",
			GoFieldType:        "time.Time",
			JSONFieldName:      "update_time",
			ProtobufFieldName:  "update_time",
			ProtobufType:       "uint64",
			ProtobufPos:        9,
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
