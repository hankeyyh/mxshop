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


CREATE TABLE `goods` (
  `id` int NOT NULL AUTO_INCREMENT,
  `add_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `is_deleted` tinyint(1) DEFAULT NULL,
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `category_id` int NOT NULL,
  `brand_id` int NOT NULL,
  `on_sale` tinyint(1) NOT NULL,
  `goods_sn` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `click_num` int NOT NULL,
  `sold_num` int NOT NULL,
  `fav_num` int NOT NULL,
  `stocks` int NOT NULL,
  `market_price` float NOT NULL,
  `shop_price` float NOT NULL,
  `goods_brief` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `ship_free` tinyint(1) NOT NULL,
  `images` json NOT NULL,
  `desc_images` json NOT NULL,
  `goods_front_image` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `is_new` tinyint(1) NOT NULL,
  `is_hot` tinyint(1) NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  KEY `goods2_category_id` (`category_id`) USING BTREE,
  KEY `goods2_brand_id` (`brand_id`) USING BTREE,
  CONSTRAINT `goods_ibfk_1` FOREIGN KEY (`category_id`) REFERENCES `category` (`id`) ON DELETE CASCADE ON UPDATE RESTRICT,
  CONSTRAINT `goods_ibfk_2` FOREIGN KEY (`brand_id`) REFERENCES `brands` (`id`) ON DELETE CASCADE ON UPDATE RESTRICT
) ENGINE=InnoDB AUTO_INCREMENT=841 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC

JSON Sample
-------------------------------------
{    "id": 92,    "add_time": "2178-08-10T22:23:44.133045991+08:00",    "is_deleted": 15,    "update_time": "2091-06-06T23:20:51.787692634+08:00",    "category_id": 51,    "brand_id": 23,    "on_sale": 83,    "goods_sn": "uTBuXyRYYtKwpYTCZVZJOYiVW",    "name": "AlsLMiPNCwPTUWOiBieigaZhg",    "click_num": 82,    "sold_num": 30,    "fav_num": 29,    "stocks": 46,    "market_price": 0.18026653,    "shop_price": 0.8860353,    "goods_brief": "BaFwYxOqRJAIYhcmHMQjirREf",    "ship_free": 86,    "images": "cTjYyZTTrZTpIwjhyETeKXvAH",    "desc_images": "gENTSCvshNmMIOHVJQpSvfHOU",    "goods_front_image": "fdWxZuuiiamQfhpSUhEcIuQML",    "is_new": 57,    "is_hot": 24}



*/

// Goods struct is a row record of the goods table in the mxshop_goods_srv database
type Goods struct {
	//[ 0] id                                             int                  null: false  primary: true   isArray: false  auto: true   col: int             len: -1      default: []
	ID int32 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:int;"`
	//[ 1] add_time                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [CURRENT_TIMESTAMP]
	AddTime time.Time `gorm:"column:add_time;type:timestamp;default:CURRENT_TIMESTAMP;"`
	//[ 2] is_deleted                                     tinyint              null: true   primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: []
	IsDeleted sql.NullInt64 `gorm:"column:is_deleted;type:tinyint;"`
	//[ 3] update_time                                    timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [CURRENT_TIMESTAMP]
	UpdateTime time.Time `gorm:"column:update_time;type:timestamp;default:CURRENT_TIMESTAMP;"`
	//[ 4] category_id                                    int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	CategoryID int32 `gorm:"column:category_id;type:int;"`
	//[ 5] brand_id                                       int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	BrandID int32 `gorm:"column:brand_id;type:int;"`
	//[ 6] on_sale                                        tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: []
	OnSale int32 `gorm:"column:on_sale;type:tinyint;"`
	//[ 7] goods_sn                                       varchar(50)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 50      default: []
	GoodsSn string `gorm:"column:goods_sn;type:varchar;size:50;"`
	//[ 8] name                                           varchar(100)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 100     default: []
	Name string `gorm:"column:name;type:varchar;size:100;"`
	//[ 9] click_num                                      int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	ClickNum int32 `gorm:"column:click_num;type:int;"`
	//[10] sold_num                                       int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	SoldNum int32 `gorm:"column:sold_num;type:int;"`
	//[11] fav_num                                        int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	FavNum int32 `gorm:"column:fav_num;type:int;"`
	//[12] stocks                                         int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	Stocks int32 `gorm:"column:stocks;type:int;"`
	//[13] market_price                                   float                null: false  primary: false  isArray: false  auto: false  col: float           len: -1      default: []
	MarketPrice float32 `gorm:"column:market_price;type:float;"`
	//[14] shop_price                                     float                null: false  primary: false  isArray: false  auto: false  col: float           len: -1      default: []
	ShopPrice float32 `gorm:"column:shop_price;type:float;"`
	//[15] goods_brief                                    varchar(200)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 200     default: []
	GoodsBrief string `gorm:"column:goods_brief;type:varchar;size:200;"`
	//[16] ship_free                                      tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: []
	ShipFree int32 `gorm:"column:ship_free;type:tinyint;"`
	//[17] images                                         json                 null: false  primary: false  isArray: false  auto: false  col: json            len: -1      default: []
	Images string `gorm:"column:images;type:json;"`
	//[18] desc_images                                    json                 null: false  primary: false  isArray: false  auto: false  col: json            len: -1      default: []
	DescImages string `gorm:"column:desc_images;type:json;"`
	//[19] goods_front_image                              varchar(200)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 200     default: []
	GoodsFrontImage string `gorm:"column:goods_front_image;type:varchar;size:200;"`
	//[20] is_new                                         tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: []
	IsNew int32 `gorm:"column:is_new;type:tinyint;"`
	//[21] is_hot                                         tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: []
	IsHot int32 `gorm:"column:is_hot;type:tinyint;"`
}

var goodsTableInfo = &TableInfo{
	Name: "goods",
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
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
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
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "on_sale",
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
			GoFieldName:        "OnSale",
			GoFieldType:        "int32",
			JSONFieldName:      "on_sale",
			ProtobufFieldName:  "on_sale",
			ProtobufType:       "int32",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "goods_sn",
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
			GoFieldName:        "GoodsSn",
			GoFieldType:        "string",
			JSONFieldName:      "goods_sn",
			ProtobufFieldName:  "goods_sn",
			ProtobufType:       "string",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "name",
			Comment:            ``,
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
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "click_num",
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
			GoFieldName:        "ClickNum",
			GoFieldType:        "int32",
			JSONFieldName:      "click_num",
			ProtobufFieldName:  "click_num",
			ProtobufType:       "int32",
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
			Name:               "sold_num",
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
			GoFieldName:        "SoldNum",
			GoFieldType:        "int32",
			JSONFieldName:      "sold_num",
			ProtobufFieldName:  "sold_num",
			ProtobufType:       "int32",
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
			Name:               "fav_num",
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
			GoFieldName:        "FavNum",
			GoFieldType:        "int32",
			JSONFieldName:      "fav_num",
			ProtobufFieldName:  "fav_num",
			ProtobufType:       "int32",
			ProtobufPos:        12,
		},

		&ColumnInfo{
			Index:              12,
			Name:               "stocks",
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
			GoFieldName:        "Stocks",
			GoFieldType:        "int32",
			JSONFieldName:      "stocks",
			ProtobufFieldName:  "stocks",
			ProtobufType:       "int32",
			ProtobufPos:        13,
		},

		&ColumnInfo{
			Index:              13,
			Name:               "market_price",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "float",
			DatabaseTypePretty: "float",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "float",
			ColumnLength:       -1,
			GoFieldName:        "MarketPrice",
			GoFieldType:        "float32",
			JSONFieldName:      "market_price",
			ProtobufFieldName:  "market_price",
			ProtobufType:       "float",
			ProtobufPos:        14,
		},

		&ColumnInfo{
			Index:              14,
			Name:               "shop_price",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "float",
			DatabaseTypePretty: "float",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "float",
			ColumnLength:       -1,
			GoFieldName:        "ShopPrice",
			GoFieldType:        "float32",
			JSONFieldName:      "shop_price",
			ProtobufFieldName:  "shop_price",
			ProtobufType:       "float",
			ProtobufPos:        15,
		},

		&ColumnInfo{
			Index:              15,
			Name:               "goods_brief",
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
			GoFieldName:        "GoodsBrief",
			GoFieldType:        "string",
			JSONFieldName:      "goods_brief",
			ProtobufFieldName:  "goods_brief",
			ProtobufType:       "string",
			ProtobufPos:        16,
		},

		&ColumnInfo{
			Index:              16,
			Name:               "ship_free",
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
			GoFieldName:        "ShipFree",
			GoFieldType:        "int32",
			JSONFieldName:      "ship_free",
			ProtobufFieldName:  "ship_free",
			ProtobufType:       "int32",
			ProtobufPos:        17,
		},

		&ColumnInfo{
			Index:              17,
			Name:               "images",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "json",
			DatabaseTypePretty: "json",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "json",
			ColumnLength:       -1,
			GoFieldName:        "Images",
			GoFieldType:        "string",
			JSONFieldName:      "images",
			ProtobufFieldName:  "images",
			ProtobufType:       "string",
			ProtobufPos:        18,
		},

		&ColumnInfo{
			Index:              18,
			Name:               "desc_images",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "json",
			DatabaseTypePretty: "json",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "json",
			ColumnLength:       -1,
			GoFieldName:        "DescImages",
			GoFieldType:        "string",
			JSONFieldName:      "desc_images",
			ProtobufFieldName:  "desc_images",
			ProtobufType:       "string",
			ProtobufPos:        19,
		},

		&ColumnInfo{
			Index:              19,
			Name:               "goods_front_image",
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
			GoFieldName:        "GoodsFrontImage",
			GoFieldType:        "string",
			JSONFieldName:      "goods_front_image",
			ProtobufFieldName:  "goods_front_image",
			ProtobufType:       "string",
			ProtobufPos:        20,
		},

		&ColumnInfo{
			Index:              20,
			Name:               "is_new",
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
			GoFieldName:        "IsNew",
			GoFieldType:        "int32",
			JSONFieldName:      "is_new",
			ProtobufFieldName:  "is_new",
			ProtobufType:       "int32",
			ProtobufPos:        21,
		},

		&ColumnInfo{
			Index:              21,
			Name:               "is_hot",
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
			GoFieldName:        "IsHot",
			GoFieldType:        "int32",
			JSONFieldName:      "is_hot",
			ProtobufFieldName:  "is_hot",
			ProtobufType:       "int32",
			ProtobufPos:        22,
		},
	},
}

// TableName sets the insert table name for this struct type
func (g *Goods) TableName() string {
	return "goods"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (g *Goods) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (g *Goods) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (g *Goods) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (g *Goods) TableInfo() *TableInfo {
	return goodsTableInfo
}
