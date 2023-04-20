package model

import (
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
	err = t.db.Offset((page - 1) * pageSize).
		Limit(pageSize).
		Find(&recList).Error
	return recList, err
}

// GetUserCnt 查询用户数量
func (t TableUser) GetUserCnt() (cnt int64, err error) {
	err = t.db.Model(&User{}).Count(&cnt).Error
	return
}

// GetUserByMobile 根据手机号查询用户
func (t TableUser) GetUserByMobile(mobile string) (rec User, err error) {
	err = t.db.Where("mobile = ?", mobile).First(&rec).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = status.Error(codes.NotFound, err.Error())
	}
	return
}

// GetUser 查询用户
func (t TableUser) GetUser(id int32) (rec User, err error) {
	err = t.db.Where("id = ?", id).First(&rec).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = status.Error(codes.NotFound, err.Error())
	}
	return
}

func (t TableUser) CreateUser(rec *User) error {
	err := t.db.Create(rec).Error
	return err
}

func (t TableUser) UpdateUser(rec User) error {
	err := t.db.Save(&rec).Error
	return err
}
