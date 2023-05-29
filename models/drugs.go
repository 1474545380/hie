package models

import (
	"time"
)

type Drugs struct {
	Drugsid       string    `gorm:"primarykey"`
	Name          string    `gorm:"column:name;type:varchar(100);" json:"name"`
	Price         int       `gorm:"column:price;type:number(10);" json:"price"`
	Purchaseprice int       `gorm:"column:purchaseprice;type:number(10);" json:"purchaseprice"`
	Num           int       `gorm:"column:num;type:number(15);" json:"num"`
	Introducedate time.Time `gorm:"column:introducedate;type:date;" json:"introducedate"`
	Productdate   time.Time `gorm:"column:productdate;type:date;" json:"productdate"`
	Qualityperiod time.Time `gorm:"column:qualityperiod;type:date;" json:"qualityperiod"`
	Supplyunit    string    `gorm:"column:supplyunit;type:varchar(50);" json:"supplyunit"`
	Productunit   string    `gorm:"column:productunit;type:varchar(50);" json:"productunit"`
}

func (table *Drugs) TableName() string {
	return "drugs"
}
