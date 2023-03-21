package model

import (
	"context"
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
func (t TableUser) BatchUser(ctx context.Context) (recList []User, err error) {
	err = t.db.Find(&recList).Error
	return recList, err
}
