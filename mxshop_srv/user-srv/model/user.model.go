package model

type User struct {
	Id        int32  `gorm:"column:id;primaryKey"`
	Mobile    string `gorm:"column:mobile"`
	Password  string `gorm:"column:password"`
	Nickname  string `gorm:"column:nick_name"`
	HeadUrl   string `gorm:"column:head_url"`
	Birthday  uint64 `gorm:"column:birthday"`
	Address   string `gorm:"column:address"`
	Desc      string `gorm:"column:desc"`
	Gender    string `gorm:"column:gender"`
	Role      int32  `gorm:"column:role"`
	CreatedAt string `gorm:"column:created_at"`
	UpdatedAt string `gorm:"column:updated_at"`
}

func (User) TableName() string {
	return "user"
}
