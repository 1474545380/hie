package models

import (
	"time"
)

type Costsettledetails struct {
	ID           uint `gorm:"primarykey"`
	CreatedAt    time.Time
	Memberid     uint `gorm:"column:memberid;type:varchar(100);" json:"memberid"`
	Roid         uint `gorm:"column:roid;type:varchar(32);" json:"roid"`
	Settleamount int  `gorm:"column:role;type:number(10);" json:"settleamount"`
	Userid       uint `gorm:"column:userid;type:varchar(100);" json:"userid"`
}

func (table *Costsettledetails) TableName() string {
	return "costsettledetails"
}
