package handler

import (
	"hie/main/help"
	"hie/main/models"
	"time"
)

func DoctorLogin(username, password string) error {
	docdata := new(models.Doctors)
	err := models.DB.Where("name = ? AND password = ?", username, password).First(&docdata).Error
	return err
}

func IsDoctorCreated(username string) (int64, error) {
	var cnt int64
	err := models.DB.Where("name = ?", username).Model(new(models.Doctors)).Count(&cnt).Error
	return cnt, err
}

func DoctorCreate(name, password, realname, departId, sex, tel, positional string, age int) error {
	data := models.Doctors{
		Docid:      help.GetUUID(),
		CreatedAt:  time.Time{},
		Name:       name,
		Password:   password,
		Role:       "04",
		Realname:   realname,
		Departid:   departId,
		Positional: positional,
		Sex:        sex,
		Age:        age,
		Tel:        tel,
	}
	err := models.DB.Create(&data).Error
	return err
}

func DoctorDelete(doctorId string) error {
	d := new(models.Doctors)
	err := models.DB.Where("docid = ?", doctorId).Find(&d).Error
	err = models.DB.Delete(&d).Error
	return err
}

func DoctorGet() ([]models.Doctors, error) {
	d := make([]models.Doctors, 0, 50)
	err := models.DB.Model(new(models.Doctors)).Preload("Department").Find(&d).Error
	return d, err
}

func DoctorDetailChange(id, name, password, tel, departId, realname, sex, positional string, age int) error {
	d := models.Doctors{}
	err := models.DB.Where("docid = ?", id).Take(&d).Error
	d.Name = name
	d.Password = password
	d.Tel = tel
	d.Age = age
	d.Departid = departId
	d.Realname = realname
	d.Sex = sex
	d.Positional = positional
	err = models.DB.Save(&d).Error
	return err
}

func DoctorSelect(doctorName, doctorId, department string) ([]models.Doctors, error) {
	// 创建一个空字符串来存储查询条件
	query := ""
	// 创建一个空切片来存储查询参数
	args := []interface{}{}
	// 如果doctorName不为空，添加到查询条件和查询参数中
	if doctorName != "" {
		query += "name LIKE ?"
		args = append(args, "%"+doctorName+"%")
	}
	// 如果doctorId不为空，添加到查询条件和查询参数中
	if doctorId != "" {
		// 如果query不为空，添加AND连接符
		if query != "" {
			query += " AND "
		}
		query += "id LIKE ?"
		args = append(args, "%"+doctorId+"%")
	}
	// 如果department不为空，添加到查询条件和查询参数中
	if department != "" {
		// 如果query不为空，添加AND连接符
		if query != "" {
			query += " AND "
		}
		query += "role LIKE ?"
		args = append(args, "%"+department+"%")
	}
	// 创建一个空切片来存储返回的医生
	d := []models.Doctors{}
	// 使用gorm的Where方法和Limit方法来执行查询，并返回错误
	err := models.DB.Model(new(models.Doctors)).Preload("Department").Where(query, args...).Limit(100).Find(&d).Error
	// 如果department不为空，使用gorm的Or方法来实现或逻辑，并返回错误
	if department != "" {
		err = models.DB.Model(new(models.Doctors)).Preload("Department").Or("username LIKE ?", "%"+department+"%").Find(&d).Error
	}
	return d, err
}

func GetDoctorByDepartId(id string) ([]models.Doctors, error) {
	d := make([]models.Doctors, 0, 50)
	err := models.DB.Model(new(models.Doctors)).Preload("Department").Where("departid = ?", id).Find(&d).Error
	return d, err
}
