package handler

import (
	"hie/main/help"
	"hie/main/models"
)

func PrescribeGetBydaid(daid string) ([]models.Prescribe, error) {
	var ps []models.Prescribe
	err := models.DB.Preload("Drugs").Where("daid = ?", daid).Find(&ps).Error
	return ps, err
}

func PrescribeDelete(prescribeid string) error {
	err := models.DB.Delete(&models.Prescribe{}, "prescribeid = ?", prescribeid).Error
	return err
}

func PrescribeCreate(roid, daid, descs, drugsid string, num int) error {
	var p models.Prescribe
	// 生成主键
	p.Prescribeid = help.GetUUID()
	p.Num = num
	p.Descs = descs
	p.Daid = daid
	p.Roid = roid
	p.Drugsid = drugsid
	p.Status = "0"
	// 插入数据
	err := models.DB.Create(&p).Error
	return err
}

func PrescribeGetAllByNotGrant() ([]models.Prescribe, error) {
	p := make([]models.Prescribe, 0, 50)
	err := models.DB.Where("status = 0").Find(&p).Error
	return p, err
}

func PrescribeStatusChange(daid string) error {
	var p []models.Prescribe
	err := models.DB.Where("daid = ?", daid).Find(&p).Error
	if err != nil {
		return err
	}
	for i, _ := range p {
		p[i].Status = "1"
	}
	err = models.DB.Save(&p).Error
	return err
}
