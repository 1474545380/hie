package handler

import (
	"hie/main/help"
	"hie/main/models"
	"time"
)

func DoctordutyCreate(did string, dutydate int64) error {
	d := time.Unix(dutydate, 0)
	data := models.Doctorduty{
		Ddid:    help.GetUUID(),
		Docid:   did,
		Dutyday: d,
	}
	err := models.DB.Create(&data).Error
	return err
}

func DoctordutyDelete(id string) error {
	d := new(models.Doctorduty)
	err := models.DB.Where("ddid = ?", id).Find(&d).Error
	err = models.DB.Delete(&d).Error
	return err
}

func DoctordutyGet() ([]models.Doctorduty, error) {
	d := make([]models.Doctorduty, 0, 50)
	err := models.DB.Find(&d).Error
	return d, err
}

func DoctordutySelect(did string, dutyDate int64) ([]models.Doctorduty, error) {
	r := time.Unix(dutyDate, 0)
	// 创建一个空字符串来存储查询条件
	query := ""
	// 创建一个空切片来存储查询参数
	args := []interface{}{}
	// 如果memberid不为空，添加到查询条件和查询参数中
	if did != "" {
		query += "docid LIKE ?"
		args = append(args, "%"+did+"%")
	}
	// 如果tel不为空，添加到查询条件和查询参数中
	if r.IsZero() {
		// 如果query不为空，添加OR连接符
		if query != "" {
			query += " OR "
		}
		query += "dutyday LIKE ?"
		args = append(args, "%"+r.Format("2006-01-02")+"%")
	}
	// 创建一个空切片来存储返回的会员
	rg := make([]models.Doctorduty, 0, 50)
	// 使用gorm的Where方法和Limit方法来执行查询，并返回错误
	err := models.DB.Where(query, args...).Limit(100).Find(&rg).Error
	return rg, err

}
