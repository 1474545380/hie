package models

type Prescribe struct {
	Prescribeid string `gorm:"primarykey"`
	Daid        string `gorm:"column:daid;type:varchar(16);" json:"daid"`
	Roid        string `gorm:"column:roid;type:varchar(32);" json:"roid"`
	Drugsid     string `gorm:"column:drugsid;type:varchar(20);" json:"drugsid"`
	Num         int    `gorm:"column:num;type:number(100);" json:"num"`
	Descs       string `gorm:"column:descs;type:varchar(20);" json:"descs"`
	Status      string `gorm:"column:status;type:varchar(11);" json:"status"`
	Drugs       Drugs  `gorm:"foreignKey:Drugsid;references:Drugsid"`
}

func (table *Prescribe) TableName() string {
	return "prescribe"
}
