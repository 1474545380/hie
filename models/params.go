package models

import (
	"time"
)

type Params struct {
	ID         uint `gorm:"primarykey"`
	CreatedAt  time.Time
	PCode      string `gorm:"column:p_code;type:varchar(64);" json:"p_code"`
	PValue     string `gorm:"column:p_value;type:varchar(8);" json:"p_value"`
	PName      string `gorm:"column:p_name;type:varchar(64);" json:"p_name"`
	ParentCode string `gorm:"column:parent_code;type:varchar(64);" json:"parent_code"`
	PFlag      string `gorm:"column:p_flag;type:varchar(1);" json:"p_flag"`
	DispSn     string `gorm:"column:disp_sn;type:varchar(8);" json:"disp_sn"`
}

func (table *Params) TableName() string {
	return "params"
}
