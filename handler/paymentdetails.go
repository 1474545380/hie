package handler

import (
	"hie/main/help"
	"hie/main/models"
	"strconv"
)

func PaymentRecord(id, money, rechargemethod, userid string, balance int) error {
	var p models.Paymentdetails
	p.Id = help.GetUUID()
	p.Userid = userid
	p.Memberid = id
	p.Rechargemethod = rechargemethod
	intmoney, _ := strconv.Atoi(money)
	p.Rechargeamount = intmoney
	p.Balance = balance
	err := models.DB.Save(&p).Error
	return err
}

func PaymentGet() ([]models.Paymentdetails, error) {
	p := make([]models.Paymentdetails, 0, 50)
	err := models.DB.Preload("User").Preload("Members").Find(&p).Error
	return p, err
}

func PaymentDetailsSelect(mid string) ([]models.Paymentdetails, error) {
	// 创建一个空切片来存储返回的药品
	var p []models.Paymentdetails
	// 使用gorm的Where方法和Limit方法来执行查询，并返回错误
	err := models.DB.Preload("User").Preload("Members").Where("memberid LIKE ?", "%"+mid+"%").Limit(100).Find(&p).Error
	return p, err
}
