package handler

import (
	"hie/main/help"
	"hie/main/models"
	"time"
)

func UserLogin(username, password string) error {
	data := new(models.User)
	err := models.DB.Where("username = ? AND password = ?", username, password).First(&data).Error
	return err
}

func IsUserCreatedById(userid string) (int64, error) {
	var cnt int64
	err := models.DB.Where("id = ? ", userid).Model(new(models.User)).Count(&cnt).Error
	return cnt, err
}

func IsUserCreatedByName(username string) (int64, error) {
	var cnt int64
	err := models.DB.Where("username = ? ", username).Model(new(models.User)).Count(&cnt).Error
	return cnt, err
}

func UserCreate(username, password, role, realname, sex, tel, address string, age int) error {
	data := models.User{
		Id:        help.GetUUID(),
		CreatedAt: time.Now(),
		Username:  username,
		Password:  password,
		Role:      role,
		RealName:  realname,
		Sex:       sex,
		Age:       age,
		Tel:       tel,
		Address:   address,
	}
	err := models.DB.Create(&data).Error
	return err
}

func UserDelete(userId string) error {
	u := new(models.User)
	err := models.DB.Where("id = ?", userId).Find(&u).Error
	err = models.DB.Delete(&u).Error
	return err
}

func UserGet() ([]models.User, error) {
	u := make([]models.User, 0, 50)
	err := models.DB.Find(&u).Error
	return u, err
}

func UserDetailChange(id, username, password, tel, address, role, realname, sex string, age int) error {
	u := models.User{}
	err := models.DB.Where("id = ?", id).Take(&u).Error
	u.Username = username
	u.Password = password
	u.Tel = tel
	u.Age = age
	u.Address = address
	u.Role = role
	u.RealName = realname
	u.Sex = sex
	err = models.DB.Save(&u).Error
	return err
}

func UserSelect(username, userId, userRole string) ([]models.User, error) {
	// 创建一个空字符串来存储查询条件
	query := ""
	// 创建一个空切片来存储查询参数
	args := []interface{}{}
	// 如果username不为空，添加到查询条件和查询参数中
	if username != "" {
		query += "username LIKE ?"
		args = append(args, "%"+username+"%")
	}
	// 如果userId不为空，添加到查询条件和查询参数中
	if userId != "" {
		// 如果query不为空，添加AND连接符
		if query != "" {
			query += " AND "
		}
		query += "id LIKE ?"
		args = append(args, "%"+userId+"%")
	}
	// 如果userRole不为空，添加到查询条件和查询参数中
	if userRole != "" {
		// 如果query不为空，添加AND连接符
		if query != "" {
			query += " AND "
		}
		query += "role LIKE ?"
		args = append(args, "%"+userRole+"%")
	}
	// 创建一个空切片来存储返回的用户
	u := []models.User{}
	// 使用gorm的Where方法和Limit方法来执行查询，并返回错误
	err := models.DB.Where(query, args...).Limit(100).Find(&u).Error
	return u, err
}
