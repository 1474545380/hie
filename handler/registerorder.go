package handler

import (
	"hie/main/help"
	"hie/main/models"
	"time"
)

func RegisterOrderCreate(memberid, departid, docid string, rotime int64) error {
	r := time.Unix(rotime, 0)
	data := models.Registerorder{
		Roid:      help.GetUUID(),
		CreatedAt: time.Time{},
		Memberid:  memberid,
		Departid:  departid,
		Docid:     docid,
		Rotime:    r,
		Status:    "01",
	}
	err := models.DB.Create(&data).Error
	return err
}

func RegisterOrderDelete(id string) error {
	r := new(models.Registerorder)
	err := models.DB.Where("roid = ?", id).Find(&r).Error
	err = models.DB.Delete(&r).Error
	return err
}

func RegisterOrderGet() ([]models.Registerorder, error) {
	r := make([]models.Registerorder, 0, 50)
	err := models.DB.Find(&r).Error
	return r, err
}

func RegisterOrderSelect(roid, mid, did string, rotime int64) ([]models.Registerorder, error) {
	r := time.Unix(rotime, 0)
	// 创建一个空字符串来存储查询条件
	query := ""
	// 创建一个空切片来存储查询参数
	args := []interface{}{}
	// 如果memberid不为空，添加到查询条件和查询参数中
	if roid != "" {
		query += "roid LIKE ?"
		args = append(args, "%"+roid+"%")
	}
	// 如果membername不为空，添加到查询条件和查询参数中
	if mid != "" {
		// 如果query不为空，添加OR连接符
		if query != "" {
			query += " OR "
		}
		query += "memberid LIKE ?"
		args = append(args, "%"+mid+"%")
	}
	// 如果credit不为空，添加到查询条件和查询参数中
	if did != "" {
		// 如果query不为空，添加OR连接符
		if query != "" {
			query += " OR "
		}
		query += "departid LIKE ?"
		args = append(args, "%"+did+"%")
	}
	// 如果tel不为空，添加到查询条件和查询参数中
	if r.IsZero() {
		// 如果query不为空，添加OR连接符
		if query != "" {
			query += " OR "
		}
		query += "role LIKE ?"
		args = append(args, "%"+r.Format("2006-01-02")+"%")
	}
	// 创建一个空切片来存储返回的会员
	rg := make([]models.Registerorder, 0, 50)
	// 使用gorm的Where方法和Limit方法来执行查询，并返回错误
	err := models.DB.Where(query, args...).Limit(100).Find(&rg).Error
	return rg, err
}
