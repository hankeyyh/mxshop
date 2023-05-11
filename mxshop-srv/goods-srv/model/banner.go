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
  `add_time` datetime NOT NULL,
  `is_deleted` tinyint(1) DEFAULT NULL,
  `update_time` datetime NOT NULL,
  `image` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `url` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `index` int NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC

JSON Sample
-------------------------------------
{    "id": 84,    "add_time": "2174-07-22T12:45:09.682579992+08:00",    "is_deleted": 29,    "update_time": "2037-12-09T21:16:43.543966252+08:00",    "image": "RVEMrQVjKUrthliNThLQSkhju",    "url": "KfOqqsIVRprCqZtJujHlsWbdS",    "index": 43}



*/

// Banner struct is a row record of the banner table in the mxshop_goods_srv database
type Banner struct {
	//[ 0] id                                             int                  null: false  primary: true   isArray: false  auto: true   col: int             len: -1      default: []
	ID int32 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:int;"`
	//[ 1] add_time                                       datetime             null: false  primary: false  isArray: false  auto: false  col: datetime        len: -1      default: []
	AddTime time.Time `gorm:"column:add_time;type:datetime;"`
	//[ 2] is_deleted                                     tinyint              null: true   primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: []
	IsDeleted sql.NullInt64 `gorm:"column:is_deleted;type:tinyint;"`
	//[ 3] update_time                                    datetime             null: false  primary: false  isArray: false  auto: false  col: datetime        len: -1      default: []
	UpdateTime time.Time `gorm:"column:update_time;type:datetime;"`
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
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
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
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
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
