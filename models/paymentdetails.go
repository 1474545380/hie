package models

import (
	"time"
)

type Paymentdetails struct {
	Id             string `gorm:"primarykey"`
	CreatedAt      time.Time
	Memberid       string  `gorm:"column:memberid;type:varchar(100);" json:"memberid"`
	Rechargeamount int     `gorm:"column:rechargeamount;type:decimal(10);" json:"rechargeamount"`
	Balance        int     `gorm:"column:balance;type:number(20);" json:"balance"`
	Rechargemethod string  `gorm:"column:rechargemethod;type:varchar(100);" json:"rechargemethod"`
	Userid         string  `gorm:"column:userid;type:varchar(20);" json:"userid"`
	Members        Members `gorm:"foreignKey:Memberid;references:Memberid"`
	User           User    `gorm:"foreignKey:Id;references:Userid"`
}

func (table *Paymentdetails) TableName() string {
	return "paymentdetails"
}
