package service

import (
	"github.com/gin-gonic/gin"
	"hie/main/handler"
	"net/http"
	"time"
)

// 创建医嘱
func DoctorAdviceCreate(c *gin.Context) {
	// 获取请求参数
	memberid := c.PostForm("memberid")
	roid := c.PostForm("roid")
	results := c.PostForm("results")
	docid := c.PostForm("docid")
	prescription := c.PostForm("prescription")
	if memberid == "" || docid == "" || prescription == "" || roid == "" || results == "" {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "必填信息为空",
		})
		return
	}
	err := handler.DoctorAdviceCreate(memberid, roid, results, prescription, docid)
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

// 医嘱删除
func DoctorAdviceDelete(c *gin.Context) {
	// 获取请求参数
	daid := c.PostForm("daid")
	if daid == "" {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "必填信息为空",
		})
		return
	}
	// 删除数据
	if err := handler.DoctorAdviceDelete(daid); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// 返回成功响应
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "OK",
	})
}

// 获取全部医嘱
func DoctorAdviceGet(c *gin.Context) {
	// 查询数据
	das, err := handler.DoctorAdviceGet()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// 返回成功响应
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"message": map[string]interface{}{
			"data": das,
		},
	})
}

// 根据Memberid、Roid、CreatedAt进行模糊查询
func DoctorAdviceSelect(c *gin.Context) {
	// 获取请求参数
	memberid := c.PostForm("memberid")
	roid := c.PostForm("roid")
	start := c.PostForm("start")
	end := c.PostForm("end")
	if memberid == "" && roid == "" && start == "" && end == "" {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "必填信息为空",
		})
		return
	}
	var startTime time.Time
	var endTime time.Time
	var err error

	if start != "" {
		startTime, err = time.Parse(time.RFC3339, start)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid start time format"})
			return
		}
	} else {
		startTime = time.Time{}
	}

	if end != "" {
		endTime, err = time.Parse(time.RFC3339, end)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid end time format"})
			return
		}
	} else {
		endTime = time.Now()
	}
	das, err := handler.DoctorAdviceSelect(memberid, roid, startTime, endTime)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// 返回成功响应
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"message": map[string]interface{}{
			"data": das,
		},
	})
}

func DoctorAdviceByNotGrantGetAll(c *gin.Context) {
	p, err := handler.PrescribeGetAllByNotGrant()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "用户查询错误",
		})
		return
	}
	np := handler.DoctorAdviceGetByDaidList(p)
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg": map[string]interface{}{
			"data": np,
		},
	})
}
