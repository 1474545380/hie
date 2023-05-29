package handler

import "hie/main/models"

type Returncs struct {
	Memid        string `json:"memid"`
	Memname      string `json:"memname"`
	Roid         string `json:"roid"`
	Balance      int    `json:"balance"`
	Settleamount int    `json:"settleamount"`
}

func GetallNotSettlement() ([]Returncs, error) {
	var das []models.Doctoradvice
	ret := make([]Returncs, 0, 50)
	err := models.DB.Preload("Members").Where("status = ?", "0").Find(&das).Error
	if err != nil {
		return nil, err
	}
	for _, d := range das {
		var ds Returncs
		pr, err := PrescribeGetBydaid(d.Daid)
		if err != nil {
			return nil, err
		}
		for _, p := range pr {
			ds.Settleamount += p.Drugs.Price * p.Num
		}
		ds.Balance = d.Members.Balance
		ds.Memid = d.Memberid
		ds.Memname = d.Members.Name
		ds.Roid = d.Roid
		ret = append(ret, ds)
	}
	return ret, err
}
