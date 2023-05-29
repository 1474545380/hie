package models

import "time"

type Department struct {
	Departid  string `gorm:"primaryKey"`
	Name      string `gorm:"column:name;type:varchar(100);" json:"name"`
	Userid    string
	User      User `gorm:"foreignKey:id;references:userid"`
	CreatedAt time.Time
}

func (table *Department) TableName() string {
	return "department"
}
