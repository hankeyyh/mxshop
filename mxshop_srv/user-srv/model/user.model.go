package model

import "time"

type User struct {
	Id        int32     `gorm:"column:id;primaryKey"`
	Mobile    string    `gorm:"column:mobile"`
	Password  string    `gorm:"column:password"`
	Nickname  string    `gorm:"column:nick_name"`
	HeadUrl   string    `gorm:"column:head_url"`
	Birthday  time.Time `gorm:"column:birthday;type:date;default:1970-1-1"`
	Address   string    `gorm:"column:address"`
	Desc      string    `gorm:"column:desc"`
	Gender    string    `gorm:"column:gender;default:male"`
	Role      int32     `gorm:"column:role;default:1"`
	CreatedAt time.Time `gorm:"column:created_at;type:timestamp"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:timestamp"`
}

func (User) TableName() string {
	return "user"
}
