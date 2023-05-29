package models

import (
	"time"
)

type Members struct {
	Memberid    string `gorm:"primarykey"`
	CreatedAt   time.Time
	Name        string `gorm:"column:name;type:varchar(100);" json:"name"`
	Password    string `gorm:"column:password;type:varchar(32);" json:"password"`
	Role        string `gorm:"column:role;type:varchar(20);" json:"role"`
	Realname    string `gorm:"column:realname;type:varchar(100);" json:"realname"`
	Credit      string `gorm:"column:credit;type:varchar(100);" json:"credit"`
	Sex         string `gorm:"column:sex;type:varchar(20);" json:"sex"`
	Age         int    `gorm:"column:age;type:int(11);" json:"age"`
	Tel         string `gorm:"column:tel;type:varchar(20);" json:"tel"`
	Balance     int    `gorm:"column:balance;type:varchar(20);" json:"balance"`
	Anaphylaxis string `gorm:"column:anaphylaxis;type:varchar(50);" json:"anaphylaxis"`
}

func (table *Members) TableName() string {
	return "members"
}
