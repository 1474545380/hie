package handler

import (
	"hie/main/help"
	"hie/main/models"
	"time"
)

func DepartmentDetailChange(id, name, userid string) error {
	d := models.Department{}
	err := models.DB.Where("departid = ?", id).Take(&d).Error
	d.Name = name
	d.Userid = userid
	err = models.DB.Save(&d).Error
	return err
}

func DepartmentSelect(name, id string) ([]models.Department, error) {
	// 创建一个空字符串来存储查询条件
	query := ""
	// 创建一个空切片来存储查询参数
	args := []interface{}{}
	// 如果name不为空，添加到查询条件和查询参数中
	if name != "" {
		query += "name LIKE ?"
		args = append(args, "%"+name+"%")
	}
	// 如果id不为空，添加到查询条件和查询参数中
	if id != "" {
		// 如果query不为空，添加AND连接符
		if query != "" {
			query += " AND "
		}
		query += "departid LIKE ?"
		args = append(args, "%"+id+"%")
	}
	// 创建一个空切片来存储返回的部门
	d := []models.Department{}
	// 使用gorm的Model方法和Preload方法来指定模型和关联
	// 使用gorm的Where方法和Limit方法来执行查询，并返回错误
	err := models.DB.Model(new(models.Department)).Preload("User").Where(query, args...).Limit(100).Find(&d).Error
	return d, err
}

func IsDepartmentCreatd(name string) (int64, error) {
	var cnt int64
	err := models.DB.Where("name = ? ", name).Model(new(models.Department)).Count(&cnt).Error
	return cnt, err
}

func DepartmentCreat(name, userid string) error {
	data := models.Department{
		Departid:  help.GetUUID(),
		CreatedAt: time.Now(),
		Name:      name,
		Userid:    userid,
	}
	err := models.DB.Create(&data).Error
	return err
}

func DepartmentDelete(id string) error {
	d := new(models.Department)
	err := models.DB.Where("departid = ?", id).Find(&d).Error
	err = models.DB.Delete(&d).Error
	doc := new(models.Doctors)
	err = models.DB.Where("departid = ?", id).Find(&doc).Error
	err = models.DB.Delete(&doc).Error
	return err
}

func DepartmentGet() ([]models.Department, error) {
	d := make([]models.Department, 0, 50)
	err := models.DB.Model(new(models.Department)).Preload("User").Find(&d).Error
	return d, err
}
