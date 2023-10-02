package dao

import (
	"fmt"
	"github.com/hankeyyh/mxshop/mxshop-srv/goods-srv/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Init() error {
	var err error
	dbConf := config.DefaultConfig.Db.Mysql
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbConf.UserName, dbConf.Password, dbConf.Host, dbConf.Port, dbConf.DBName)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return err
}

func init() {
	Init()
}
