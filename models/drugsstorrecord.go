package models

import (
	"time"
)

type Drugsstorrecord struct {
	ID          string `gorm:"primarykey"`
	CreatedAt   time.Time
	Drugsid     string    `gorm:"column:drugsid;type:varchar(100);" json:"drugsid"`
	Num         int       `gorm:"column:num;type:number(32);" json:"num"`
	Storagedate time.Time `gorm:"column:storagedate;type:date;" json:"storagedate"`
}

func (table *Drugsstorrecord) TableName() string {
	return "drugsstorrecord"
}
