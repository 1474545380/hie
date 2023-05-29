package models

import (
	"gorm.io/gorm"
	"time"
)

type Doctoradvice struct {
	Daid         string  `gorm:"primarykey"`
	Roid         string  `gorm:"column:roid;type:varchar(100);" json:"roid"`
	Memberid     string  `gorm:"column:memberid;type:varchar(32);" json:"memberid"`
	Docid        string  `gorm:"column:docid;type:varchar(20);" json:"docid"`
	Results      string  `gorm:"column:results;type:varchar(128);" json:"results"`
	Prescription string  `gorm:"column:prescription;type:varchar(256);" json:"prescription"`
	Members      Members `gorm:"foreignKey:Memberid;references:Memberid"`
	Status       string  `gorm:"column:status;type:varchar(11);" json:"status"`
	CreatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}

func (table *Doctoradvice) TableName() string {
	return "doctoradvice"
}
