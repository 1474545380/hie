package service

import (
	"github.com/gin-gonic/gin"
	"hie/main/handler"
	"net/http"
	"strconv"
)

func DrugCreate(c *gin.Context) {
	name := c.PostForm("name")
	price, _ := strconv.Atoi(c.PostForm("price"))
	purchaseprice, _ := strconv.Atoi(c.PostForm("purchaseprice"))
	num, _ := strconv.Atoi(c.PostForm("num"))
	introducedate, _ := strconv.Atoi(c.PostForm("introducedate"))
	productdate, _ := strconv.Atoi(c.PostForm("productdate"))
	qualityperiod, _ := strconv.Atoi(c.PostForm("qualityperiod"))
	supplyunit := c.PostForm("supplyunit")
	productunit := c.PostForm("productunit")
	if price == 0 || name == "" || purchaseprice == 0 || num == 0 || productdate == 0 || introducedate == 0 || qualityperiod == 0 || supplyunit == "" || productunit == "" {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "必填信息为空",
		})
		return
	}
	err := handler.DrugCreate(name, supplyunit, productunit, price, purchaseprice, num, int64(introducedate), int64(productdate), int64(qualityperiod))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "creat drug err:" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "OK",
	})
}

func DrugDelete(c *gin.Context) {
	drugId := c.PostForm("drug_id")
	if drugId == "" {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "必填信息为空",
		})
		return
	}
	err := handler.DrugDelete(drugId)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "药品删除错误",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "OK",
	})
}

func DrugGet(c *gin.Context) {
	d, err := handler.DrugGet()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "药品查询错误",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg": map[string]interface{}{
			"data": d,
		},
	})
}

func DrugDetailChange(c *gin.Context) {
	id := c.PostForm("id")
	name := c.PostForm("name")
	price, _ := strconv.Atoi(c.PostForm("price"))
	purchaseprice, _ := strconv.Atoi(c.PostForm("purchaseprice"))
	num, _ := strconv.Atoi(c.PostForm("num"))
	introducedate, _ := strconv.Atoi(c.PostForm("introducedate"))
	productdate, _ := strconv.Atoi(c.PostForm("productdate"))
	qualityperiod, _ := strconv.Atoi(c.PostForm("qualityperiod"))
	supplyunit := c.PostForm("supplyunit")
	productunit := c.PostForm("productunit")
	if id == "" || price == 0 || name == "" || purchaseprice == 0 || num == 0 || productdate == 0 || introducedate == 0 || qualityperiod == 0 || supplyunit == "" || productunit == "" {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "必填信息为空",
		})
		return
	}
	err := handler.DrugDetailChange(id, name, supplyunit, productunit, price, purchaseprice, num, int64(qualityperiod), int64(introducedate), int64(productdate))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "change doctor err:" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "OK",
	})

}

func DrugSelect(c *gin.Context) {
	name := c.PostForm("name")
	id := c.PostForm("id")
	if name == "" && id == "" {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "必填信息为空",
		})
		return
	}
	d, err := handler.DrugSelect(name, id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "select drugs err",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"message": map[string]interface{}{
			"data": d,
		},
	})

}

func IncreaseQuantity(c *gin.Context) {
	id := c.PostForm("id")
	quantity := c.PostForm("quantity")
	err := handler.IncreaseQuantity(id, quantity)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "IncreaseQuantity err:",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "ok",
	})
}
