package models

import (
	"time"
)

type Registerorder struct {
	Roid      string `gorm:"primarykey"`
	CreatedAt time.Time
	Memberid  string    `gorm:"column:memberid;type:varchar(100);" json:"memberid"`
	Departid  string    `gorm:"column:departid;type:varchar(100);" json:"departid"`
	Docid     string    `gorm:"column:docid;type:varchar(100);" json:"docid"`
	Rotime    time.Time `gorm:"column:rotime;type:date;" json:"rotime"`
	Status    string    `gorm:"column:status;type:varchar(10);" json:"status"`
}

func (table *Registerorder) TableName() string {
	return "registerorder"
}
