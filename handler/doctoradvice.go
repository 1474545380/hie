package handler

import (
	"hie/main/help"
	"hie/main/models"
	"strings"
	"time"
)

func DoctorAdviceCreate(memberid, roid, results, prescription, docid string) error {
	var da models.Doctoradvice
	// 生成主键
	da.Daid = help.GetUUID()
	da.Memberid = memberid
	da.Roid = roid
	da.Docid = docid
	da.Status = "0"
	da.Results = results
	da.Prescription = prescription
	// 插入数据
	err := models.DB.Create(&da).Error
	return err
}

func DoctorAdviceDelete(id string) error {
	err := models.DB.Delete(&models.Doctoradvice{}, "daid = ?", id).Error
	return err
}

func DoctorAdviceGet() ([]models.Doctoradvice, error) {
	var das []models.Doctoradvice
	err := models.DB.Preload("Members").Find(&das).Error
	return das, err
}

func DoctorAdviceSelect(memberid, roid string, startTime, endTime time.Time) ([]models.Doctoradvice, error) { // 构造查询条件
	var conditions []string
	var values []interface{}
	if memberid != "" {
		conditions = append(conditions, "memberid LIKE ?")
		values = append(values, "%"+memberid+"%")
	}
	if roid != "" {
		conditions = append(conditions, "roid LIKE ?")
		values = append(values, "%"+roid+"%")
	}
	conditions = append(conditions, "created_at BETWEEN ? AND ?")
	values = append(values, startTime)
	values = append(values, endTime)

	// 查询数据
	var das []models.Doctoradvice
	err := models.DB.Where(strings.Join(conditions, " AND "), values...).Find(&das).Error
	return das, err
}

func DoctorAdviceGetByDaid(id string) models.Doctoradvice {
	var das models.Doctoradvice
	models.DB.Where("daid = ?", id).Find(&das)
	return das
}

func DoctorAdviceGetByDaidList(plist []models.Prescribe) []models.Doctoradvice {
	np := make([]models.Doctoradvice, 0, 10)
	for _, prescribe := range plist {
		np = append(np, DoctorAdviceGetByDaid(prescribe.Daid))
	}
	return np
}

func DoctorAdviceGrant(daid string) error {
	var das models.Doctoradvice
	err := models.DB.Where("daid = ?", daid).Find(&das).Error
	das.Status = "1"
	err = models.DB.Save(&das).Error
	return err
}
