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
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `category_id` bigint unsigned NOT NULL COMMENT '分类id',
  `brand_id` bigint unsigned NOT NULL COMMENT '品牌id',
  `on_sale` tinyint NOT NULL DEFAULT '1' COMMENT '是否上架',
  `goods_sn` varchar(100) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '商品唯一货号',
  `name` varchar(100) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '商品名称',
  `click_num` int NOT NULL DEFAULT '0' COMMENT '点击数',
  `sold_num` int NOT NULL DEFAULT '0' COMMENT '商品销售量',
  `fav_num` int NOT NULL DEFAULT '0' COMMENT '收藏数',
  `market_price` float NOT NULL DEFAULT '0' COMMENT '市场价格',
  `shop_price` float NOT NULL DEFAULT '0' COMMENT '本店价格',
  `goods_brief` varchar(200) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '商品简短描述',
  `ship_free` tinyint NOT NULL DEFAULT '1' COMMENT '是否承担运费,0:否，1:是',
  `images` json DEFAULT NULL COMMENT '商品轮播图',
  `desc_images` json DEFAULT NULL COMMENT '详情页图片',
  `goods_front_image` varchar(200) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '封面图',
  `is_new` tinyint NOT NULL DEFAULT '0' COMMENT '是否新品',
  `is_hot` tinyint NOT NULL DEFAULT '0' COMMENT '是否热销,0：否，1:是',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_goods_sn` (`goods_sn`),
  KEY `goods_FK` (`category_id`),
  KEY `goods_FK_1` (`brand_id`),
  CONSTRAINT `goods_FK` FOREIGN KEY (`category_id`) REFERENCES `category` (`id`) ON DELETE CASCADE,
  CONSTRAINT `goods_FK_1` FOREIGN KEY (`brand_id`) REFERENCES `brand` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='商品'

JSON Sample
-------------------------------------
{    "id": 79,    "category_id": 86,    "brand_id": 8,    "on_sale": 19,    "goods_sn": "DqcJOjPEvAdOtLgEboLtCeeyD",    "name": "KoWRrRGXaPhYGYJalgnbYKxvr",    "click_num": 73,    "sold_num": 16,    "fav_num": 50,    "market_price": 0.3482366,    "shop_price": 0.8294621,    "goods_brief": "TpfFuJkLixHENOiuSFhOWUARY",    "ship_free": 2,    "images": "seTAFhwWGfyONcqoRRrwxLCLf",    "desc_images": "gNOpHFEFpCIlSAciPbZgZSkJL",    "goods_front_image": "jNLwGkCyiwmRIqNEfsxcOqdQy",    "is_new": 93,    "is_hot": 67,    "created_at": "2070-05-31T04:30:42.142186995+08:00",    "updated_at": "2058-02-13T16:08:54.183859689+08:00"}


Comments
-------------------------------------
[ 0] column is set for unsigned
[ 1] column is set for unsigned
[ 2] column is set for unsigned



*/

// Goods struct is a row record of the goods table in the mxshop_goods_srv database
type Goods struct {
	//[ 0] id                                             ubigint              null: false  primary: true   isArray: false  auto: true   col: ubigint         len: -1      default: []
	ID uint64 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:ubigint;"` // 主键
	//[ 1] category_id                                    ubigint              null: false  primary: false  isArray: false  auto: false  col: ubigint         len: -1      default: []
	CategoryID uint64 `gorm:"column:category_id;type:ubigint;"` // 分类id
	//[ 2] brand_id                                       ubigint              null: false  primary: false  isArray: false  auto: false  col: ubigint         len: -1      default: []
	BrandID uint64 `gorm:"column:brand_id;type:ubigint;"` // 品牌id
	//[ 3] on_sale                                        tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [1]
	OnSale int32 `gorm:"column:on_sale;type:tinyint;default:1;"` // 是否上架
	//[ 4] goods_sn                                       varchar(100)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 100     default: []
	GoodsSn string `gorm:"column:goods_sn;type:varchar;size:100;"` // 商品唯一货号
	//[ 5] name                                           varchar(100)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 100     default: []
	Name string `gorm:"column:name;type:varchar;size:100;"` // 商品名称
	//[ 6] click_num                                      int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	ClickNum int32 `gorm:"column:click_num;type:int;default:0;"` // 点击数
	//[ 7] sold_num                                       int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	SoldNum int32 `gorm:"column:sold_num;type:int;default:0;"` // 商品销售量
	//[ 8] fav_num                                        int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	FavNum int32 `gorm:"column:fav_num;type:int;default:0;"` // 收藏数
	//[ 9] market_price                                   float                null: false  primary: false  isArray: false  auto: false  col: float           len: -1      default: [0]
	MarketPrice float32 `gorm:"column:market_price;type:float;default:0;"` // 市场价格
	//[10] shop_price                                     float                null: false  primary: false  isArray: false  auto: false  col: float           len: -1      default: [0]
	ShopPrice float32 `gorm:"column:shop_price;type:float;default:0;"` // 本店价格
	//[11] goods_brief                                    varchar(200)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 200     default: []
	GoodsBrief string `gorm:"column:goods_brief;type:varchar;size:200;"` // 商品简短描述
	//[12] ship_free                                      tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [1]
	ShipFree int32 `gorm:"column:ship_free;type:tinyint;default:1;"` // 是否承担运费,0:否，1:是
	//[13] images                                         json                 null: true   primary: false  isArray: false  auto: false  col: json            len: -1      default: []
	Images sql.NullString `gorm:"column:images;type:json;"` // 商品轮播图
	//[14] desc_images                                    json                 null: true   primary: false  isArray: false  auto: false  col: json            len: -1      default: []
	DescImages sql.NullString `gorm:"column:desc_images;type:json;"` // 详情页图片
	//[15] goods_front_image                              varchar(200)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 200     default: []
	GoodsFrontImage string `gorm:"column:goods_front_image;type:varchar;size:200;"` // 封面图
	//[16] is_new                                         tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [0]
	IsNew int32 `gorm:"column:is_new;type:tinyint;default:0;"` // 是否新品
	//[17] is_hot                                         tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [0]
	IsHot int32 `gorm:"column:is_hot;type:tinyint;default:0;"` // 是否热销,0：否，1:是
	//[18] created_at                                     timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [CURRENT_TIMESTAMP]
	CreatedAt time.Time `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP;"` // 创建时间
	//[19] updated_at                                     timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [CURRENT_TIMESTAMP]
	UpdatedAt time.Time `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP;"` // 更新时间

}

var goodsTableInfo = &TableInfo{
	Name: "goods",
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
			Name:               "on_sale",
			Comment:            `是否上架`,
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
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "goods_sn",
			Comment:            `商品唯一货号`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(100)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       100,
			GoFieldName:        "GoodsSn",
			GoFieldType:        "string",
			JSONFieldName:      "goods_sn",
			ProtobufFieldName:  "goods_sn",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "name",
			Comment:            `商品名称`,
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
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "click_num",
			Comment:            `点击数`,
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
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "sold_num",
			Comment:            `商品销售量`,
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
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "fav_num",
			Comment:            `收藏数`,
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
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "market_price",
			Comment:            `市场价格`,
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
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
			Name:               "shop_price",
			Comment:            `本店价格`,
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
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
			Name:               "goods_brief",
			Comment:            `商品简短描述`,
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
			ProtobufPos:        12,
		},

		&ColumnInfo{
			Index:              12,
			Name:               "ship_free",
			Comment:            `是否承担运费,0:否，1:是`,
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
			ProtobufPos:        13,
		},

		&ColumnInfo{
			Index:              13,
			Name:               "images",
			Comment:            `商品轮播图`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "json",
			DatabaseTypePretty: "json",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "json",
			ColumnLength:       -1,
			GoFieldName:        "Images",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "images",
			ProtobufFieldName:  "images",
			ProtobufType:       "string",
			ProtobufPos:        14,
		},

		&ColumnInfo{
			Index:              14,
			Name:               "desc_images",
			Comment:            `详情页图片`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "json",
			DatabaseTypePretty: "json",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "json",
			ColumnLength:       -1,
			GoFieldName:        "DescImages",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "desc_images",
			ProtobufFieldName:  "desc_images",
			ProtobufType:       "string",
			ProtobufPos:        15,
		},

		&ColumnInfo{
			Index:              15,
			Name:               "goods_front_image",
			Comment:            `封面图`,
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
			ProtobufPos:        16,
		},

		&ColumnInfo{
			Index:              16,
			Name:               "is_new",
			Comment:            `是否新品`,
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
			ProtobufPos:        17,
		},

		&ColumnInfo{
			Index:              17,
			Name:               "is_hot",
			Comment:            `是否热销,0：否，1:是`,
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
			ProtobufPos:        18,
		},

		&ColumnInfo{
			Index:              18,
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
			ProtobufPos:        19,
		},

		&ColumnInfo{
			Index:              19,
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
			ProtobufPos:        20,
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
