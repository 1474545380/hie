package service

import (
	"github.com/gin-gonic/gin"
	"hie/main/handler"
	"net/http"
)

func PaymentGet(c *gin.Context) {
	p, err := handler.PaymentGet()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "充值记录查询错误",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg": map[string]interface{}{
			"data": p,
		},
	})
}

func PaymentDetailsSelect(c *gin.Context) {
	mid := c.PostForm("mid")
	if mid == "" {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "必填信息为空",
		})
		return
	}
	d, err := handler.PaymentDetailsSelect(mid)
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
