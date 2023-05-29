package models

import (
	"time"
)

type Doctors struct {
	Docid      string `gorm:"primarykey"`
	CreatedAt  time.Time
	Name       string     `gorm:"column:name;type:varchar(100);" json:"name"`
	Password   string     `gorm:"column:password;type:varchar(32);" json:"password"`
	Role       string     `gorm:"column:role;type:varchar(20);" json:"role"`
	Realname   string     `gorm:"column:realname;type:varchar(100);" json:"realname"`
	Departid   string     `gorm:"column:departid;type:int(11);" json:"departid"`
	Positional string     `gorm:"column:positional;type:int(11);" json:"positional"`
	Sex        string     `gorm:"column:sex;type:varchar(20);" json:"sex"`
	Age        int        `gorm:"column:age;type:int(11);" json:"age"`
	Tel        string     `gorm:"column:tel;type:varchar(20);" json:"tel"`
	Department Department `gorm:"foreignKey:Departid;references:Departid"`
}

func (table *Doctors) TableName() string {
	return "doctors"
}
