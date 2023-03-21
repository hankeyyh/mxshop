package model

import (
	"gorm.io/gorm"
	"sync"
)

var (
	userInstance TableUser
	once         sync.Once
)

type TableUser struct {
	db *gorm.DB
}

func UserInstance() TableUser {
	once.Do(func() {
		userInstance = TableUser{db: DB}
	})
	return userInstance
}

// BatchUser 批量查询用户
func (t TableUser) BatchUser(page int, pageSize int) (recList []User, err error) {
	err = t.db.Find(&recList).
		Offset((page - 1) * pageSize).
		Limit(pageSize).Error
	return recList, err
}

// GetUserCnt 查询用户数量
func (t TableUser) GetUserCnt() (cnt int64, err error) {
	err = t.db.Model(&User{}).Count(&cnt).Error
	return
}
