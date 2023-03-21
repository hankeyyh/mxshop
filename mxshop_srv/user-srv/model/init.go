package model

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/hankeyyh/mxshop_user_srv/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

type DBConfig struct {
	DBName   string `toml:"db_name"`
	UserName string `toml:"user_name"`
	Password string `toml:"password"`
	Host     string `toml:"host"`
	Port     int    `toml:"port"`
}

func initDB() {
	var dbConfig DBConfig
	// todo toml加上名字
	_, err := toml.Decode(config.Str, &dbConfig)
	if err != nil {
		panic(err)
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbConfig.UserName,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName)

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("success!")
	}
}

func init() {
	initDB()
}
