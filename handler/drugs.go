package handler

import (
	"hie/main/help"
	"hie/main/models"
	"strconv"
	"time"
)

func DrugCreate(name, supplyunit, productunit string, price, purchaseprice, num int, introducedate, productdate, qualityperiod int64) error {
	qualityperiod += productdate
	i := time.Unix(introducedate, 0)
	p := time.Unix(productdate, 0)
	q := time.Unix(qualityperiod, 0)
	data := models.Drugs{
		Drugsid:       help.GetUUID(),
		Name:          name,
		Price:         price,
		Purchaseprice: purchaseprice,
		Num:           num,
		Introducedate: i,
		Productdate:   p,
		Qualityperiod: q,
		Supplyunit:    supplyunit,
		Productunit:   productunit,
	}
	err := models.DB.Create(&data).Error
	return err
}

func DrugDelete(id string) error {
	d := new(models.Drugs)
	err := models.DB.Where("drugsid = ?", id).Find(&d).Error
	err = models.DB.Delete(&d).Error
	return err
}

func DrugGet() ([]models.Drugs, error) {
	d := make([]models.Drugs, 0, 50)
	err := models.DB.Find(&d).Error
	return d, err
}

func DrugDetailChange(id, name, supplyunit, productunit string, price, purchaseprice, num int, introducedate, productdate, qualityperiod int64) error {
	d := new(models.Drugs)
	qualityperiod += productdate
	i := time.Unix(introducedate, 0)
	p := time.Unix(productdate, 0)
	q := time.Unix(qualityperiod, 0)
	err := models.DB.Where("drugsid = ?", id).Find(&d).Error
	d.Name = name
	d.Supplyunit = supplyunit
	d.Purchaseprice = purchaseprice
	d.Productunit = productunit
	d.Price = price
	d.Num = num
	d.Introducedate = i
	d.Productdate = p
	d.Qualityperiod = q
	err = models.DB.Save(&d).Error
	return err
}

func DrugSelect(name, id string) ([]models.Drugs, error) {
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
		query += "drugsid LIKE ?"
		args = append(args, "%"+id+"%")
	}
	// 创建一个空切片来存储返回的药品
	d := []models.Drugs{}
	// 使用gorm的Where方法和Limit方法来执行查询，并返回错误
	err := models.DB.Where(query, args...).Limit(100).Find(&d).Error
	return d, err
}

func IncreaseQuantity(id, quantity string) error {
	d := models.Drugs{}
	err := models.DB.Where("drugsid = ?", id).Take(&d).Error
	lastnum, _ := strconv.Atoi(quantity)
	d.Num += lastnum
	err = models.DB.Save(&d).Error
	return err
}
