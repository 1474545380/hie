package service

import (
	"github.com/gin-gonic/gin"
	"hie/main/handler"
	"net/http"
	"strconv"
)

func DoctordutyCreate(c *gin.Context) {
	did := c.PostForm("did")
	dutydate, _ := strconv.Atoi(c.PostForm("dutydate"))
	if dutydate == 0 || did == "" {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "必填信息为空",
		})
		return
	}
	err := handler.DoctordutyCreate(did, int64(dutydate))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "creat Doctorduty err:" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "OK",
	})
}

func DoctordutyDelete(c *gin.Context) {
	did := c.PostForm("did")
	if did == "" {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "必填信息为空",
		})
		return
	}
	err := handler.DoctordutyDelete(did)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "医生值班删除错误",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "OK",
	})
}

func DoctordutyGet(c *gin.Context) {
	d, err := handler.DoctordutyGet()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "医生值班查询错误",
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

func DoctordutySelect(c *gin.Context) {
	did := c.PostForm("did")
	dutydate, _ := strconv.Atoi(c.PostForm("dutydate"))
	if did == "" && dutydate == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "必填信息为空",
		})
		return
	}
	d, err := handler.DoctordutySelect(did, int64(dutydate))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "select Doctorduty err",
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
