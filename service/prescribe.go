package service

import (
	"github.com/gin-gonic/gin"
	"hie/main/handler"
	"net/http"
	"strconv"
)

// 创建开药
func PrescribeCreate(c *gin.Context) {
	// 获取请求参数
	drugsid := c.PostForm("drugsid")
	daid := c.PostForm("daid")
	roid := c.PostForm("roid")
	num, _ := strconv.Atoi(c.PostForm("num"))
	descs := c.PostForm("descs")
	if roid == "" || descs == "" || daid == "" || num == 0 || drugsid == "" {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "必填信息为空",
		})
		return
	}
	err := handler.PrescribeCreate(roid, daid, descs, drugsid, num)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// 返回成功响应
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "OK",
	})
}

// 删除开药
func PrescribeDelete(c *gin.Context) {
	// 获取请求参数
	prescribeid := c.PostForm("prescribeid")
	if prescribeid == "" {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "必填信息为空",
		})
		return
	}
	// 删除数据
	err := handler.PrescribeDelete(prescribeid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// 返回成功响应
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "OK",
	})
}

// PrescribeGetBydaid 根据daid获取开药列表
func PrescribeGetBydaid(c *gin.Context) {
	// 获取请求参数
	daid := c.PostForm("daid")
	if daid == "" {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "必填信息为空",
		})
		return
	}
	// 查询数据
	ps, err := handler.PrescribeGetBydaid(daid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// 返回成功响应
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"message": map[string]interface{}{
			"data": ps,
		},
	})
}

func PrescribeGrant(c *gin.Context) {
	// 获取请求参数
	daid := c.PostForm("daid")
	if daid == "" {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "必填信息为空",
		})
		return
	}
	err := handler.DoctorAdviceGrant(daid)
	// 查询数据
	err = handler.PrescribeStatusChange(daid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// 返回成功响应
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "ok",
	})
}
