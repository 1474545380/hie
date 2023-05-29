package service

import (
	"github.com/gin-gonic/gin"
	"hie/main/handler"
	"net/http"
	"strconv"
)

func RegisterOrderCreate(c *gin.Context) {
	memberid := c.PostForm("memberid")
	departid := c.PostForm("departid")
	docid := c.PostForm("docid")
	rotime, _ := strconv.Atoi(c.PostForm("rotime"))
	if memberid == "" || departid == "" || docid == "" || rotime == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "必填信息为空",
		})
		return
	}
	err := handler.RegisterOrderCreate(memberid, departid, docid, int64(rotime))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "creat order err:" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "OK",
	})

}

func RegisterOrderDelete(c *gin.Context) {
	id := c.PostForm("id")
	if id == "" {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "必填信息为空",
		})
		return
	}
	err := handler.RegisterOrderDelete(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "订单删除错误",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "OK",
	})
}

func RegisterOrderGet(c *gin.Context) {
	r, err := handler.RegisterOrderGet()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "订单查询错误",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg": map[string]interface{}{
			"data": r,
		},
	})
}

func RegisterOrderSelect(c *gin.Context) {
	roid := c.PostForm("roid")
	mid := c.PostForm("mid")
	did := c.PostForm("did")
	rotime, _ := strconv.Atoi(c.PostForm("rotime"))
	if roid == "" && mid == "" && did == "" && rotime == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "必填信息为空",
		})
		return
	}
	r, err := handler.RegisterOrderSelect(roid, mid, did, int64(rotime))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "select user err:" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"message": map[string]interface{}{
			"data": r,
		},
	})
}
