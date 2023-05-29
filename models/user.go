package models

import (
	"time"
)

type User struct {
	Id        string `gorm:"primarykey"`
	CreatedAt time.Time
	Username  string `gorm:"column:username;type:varchar(100);" json:"username"`
	Password  string `gorm:"column:password;type:varchar(32);" json:"password"`
	Role      string `gorm:"column:role;type:varchar(20);" json:"role"`
	RealName  string `gorm:"column:realname;type:varchar(100);" json:"realname"`
	Sex       string `gorm:"column:sex;type:varchar(20);" json:"sex"`
	Age       int    `gorm:"column:age;type:int(11);" json:"age"`
	Tel       string `gorm:"column:tel;type:varchar(20);" json:"tel"`
	Address   string `gorm:"column:address;type:varchar(50);" json:"address"`
}

func (table *User) TableName() string {
	return "user"
}
