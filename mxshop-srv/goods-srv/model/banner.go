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
  `id` int NOT NULL AUTO_INCREMENT,
  `add_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `is_deleted` tinyint(1) NOT NULL DEFAULT '0',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `image` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `url` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `index` int NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC

JSON Sample
-------------------------------------
{    "id": 8,    "add_time": "2281-06-02T12:50:00.12387918+08:00",    "is_deleted": 52,    "update_time": "2197-06-10T23:55:19.026363566+08:00",    "image": "ixrvDByFHfZtVufYKoNtNyawU",    "url": "smKBUjDiHLwFERrtPkbTnQuAD",    "index": 12}



*/

// Banner struct is a row record of the banner table in the mxshop_goods_srv database
type Banner struct {
	//[ 0] id                                             int                  null: false  primary: true   isArray: false  auto: true   col: int             len: -1      default: []
	ID int32 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:int;"`
	//[ 1] add_time                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [CURRENT_TIMESTAMP]
	AddTime time.Time `gorm:"column:add_time;type:timestamp;default:CURRENT_TIMESTAMP;"`
	//[ 2] is_deleted                                     tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [0]
	IsDeleted int32 `gorm:"column:is_deleted;type:tinyint;default:0;"`
	//[ 3] update_time                                    timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [CURRENT_TIMESTAMP]
	UpdateTime time.Time `gorm:"column:update_time;type:timestamp;default:CURRENT_TIMESTAMP;"`
	//[ 4] image                                          varchar(200)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 200     default: []
	Image string `gorm:"column:image;type:varchar;size:200;"`
	//[ 5] url                                            varchar(200)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 200     default: []
	URL string `gorm:"column:url;type:varchar;size:200;"`
	//[ 6] index                                          int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	Index int32 `gorm:"column:index;type:int;"`
}

var bannerTableInfo = &TableInfo{
	Name: "banner",
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
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
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
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
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
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "image",
			Comment:            ``,
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
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "url",
			Comment:            ``,
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
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "index",
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
			GoFieldName:        "Index",
			GoFieldType:        "int32",
			JSONFieldName:      "index",
			ProtobufFieldName:  "index",
			ProtobufType:       "int32",
			ProtobufPos:        7,
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
