package db

import (
	"gorm.io/gorm"
	"sync"
)

var (
	dbFactoryMap sync.Map
)

// 对应一个tag下可以有多个db, dbMap[addr]=DB
type dbFactory struct {
	dbMap sync.Map
}

type DB struct {
	config Config
	db     *gorm.DB
}
