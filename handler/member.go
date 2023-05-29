package handler

import (
	"hie/main/help"
	"hie/main/models"
	"strconv"
	"time"
)

func MemberLogin(username, password string) error {
	memberdata := new(models.Members)
	err := models.DB.Where("name = ? AND password = ?", username, password).First(&memberdata).Error
	return err
}

func IsMemberCreated(name string) (int64, error) {
	var cnt int64
	err := models.DB.Where("name = ?", name).Model(new(models.Members)).Count(&cnt).Error
	return cnt, err

}

func MemberCreate(name, password, realname, sex, tel, credit, anaphylaxis string, balance, age int) error {
	data := models.Members{
		Memberid:    help.GetUUID(),
		CreatedAt:   time.Time{},
		Name:        name,
		Password:    password,
		Role:        "05",
		Realname:    realname,
		Credit:      credit,
		Sex:         sex,
		Age:         age,
		Tel:         tel,
		Balance:     balance,
		Anaphylaxis: anaphylaxis,
	}
	err := models.DB.Create(&data).Error
	return err
}

func MemberDelete(id string) error {
	m := new(models.Members)
	err := models.DB.Where("memberid = ?", id).Find(&m).Error
	err = models.DB.Delete(&m).Error
	return err
}

func MemberGet() ([]models.Members, error) {
	m := make([]models.Members, 0, 50)
	err := models.DB.Find(&m).Error
	return m, err
}

func MemberDetailChange(id, name, password, realname, sex, tel, credit, anaphylaxis string, balance, age int) error {
	m := models.Members{}
	err := models.DB.Where("memberid = ?", id).Take(&m).Error
	m.Name = name
	m.Password = password
	m.Realname = realname
	m.Credit = credit
	m.Sex = sex
	m.Age = age
	m.Tel = tel
	m.Balance = balance
	m.Anaphylaxis = anaphylaxis
	err = models.DB.Save(&m).Error
	return err
}

func MemberSelect(memberid, membername, credit, tel string) ([]models.Members, error) {
	// 创建一个空字符串来存储查询条件
	query := ""
	// 创建一个空切片来存储查询参数
	args := []interface{}{}
	// 如果memberid不为空，添加到查询条件和查询参数中
	if memberid != "" {
		query += "memberid LIKE ?"
		args = append(args, "%"+memberid+"%")
	}
	// 如果membername不为空，添加到查询条件和查询参数中
	if membername != "" {
		// 如果query不为空，添加OR连接符
		if query != "" {
			query += " OR "
		}
		query += "name LIKE ?"
		args = append(args, "%"+membername+"%")
	}
	// 如果credit不为空，添加到查询条件和查询参数中
	if credit != "" {
		// 如果query不为空，添加OR连接符
		if query != "" {
			query += " OR "
		}
		query += "credit LIKE ?"
		args = append(args, "%"+credit+"%")
	}
	// 如果tel不为空，添加到查询条件和查询参数中
	if tel != "" {
		// 如果query不为空，添加OR连接符
		if query != "" {
			query += " OR "
		}
		query += "tel LIKE ?"
		args = append(args, "%"+tel+"%")
	}
	// 创建一个空切片来存储返回的会员
	m := []models.Members{}
	// 使用gorm的Where方法和Limit方法来执行查询，并返回错误
	err := models.DB.Where(query, args...).Limit(100).Find(&m).Error
	return m, err
}

func MemberRecharge(id, money, rechargemethod, userid string) error {
	m := models.Members{}
	err := models.DB.Where("memberid = ?", id).Take(&m).Error
	lastMoney, _ := strconv.Atoi(money)
	m.Balance += lastMoney
	err = models.DB.Save(&m).Error
	err = PaymentRecord(id, money, rechargemethod, userid, m.Balance)
	if err != nil {
		return err
	}

	return err
}
