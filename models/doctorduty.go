package models

import (
	"time"
)

type Doctorduty struct {
	Ddid    string    `gorm:"primarykey"`
	Docid   string    `gorm:"column:docid;type:varchar(100);" json:"docid"`
	Dutyday time.Time `gorm:"column:dutyday;type:date;" json:"dutyday"`
}

func (table *Doctorduty) TableName() string {
	return "doctorduty"
}
