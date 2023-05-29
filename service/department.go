package service

import (
	"github.com/gin-gonic/gin"
	"hie/main/handler"
	"net/http"
)

func DepartmentCreate(c *gin.Context) {
	name := c.PostForm("name")
	userid := c.PostForm("userid")
	if name == "" || userid == "" {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "必填信息为空",
		})
		return
	}
	//判断科室是否已存在
	cnt, err := handler.IsDepartmentCreatd(name)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "get user error:" + err.Error(),
		})
		return
	}
	if cnt > 0 {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "科室已存在",
		})
		return
	}
	//判断创建科室的用户是否存在
	cnt, err = handler.IsUserCreatedById(userid)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "get user error:" + err.Error(),
		})
		return
	}
	if cnt == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "用户不存在",
		})
		return
	}

	err = handler.DepartmentCreat(name, userid)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "creat department err:" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "OK",
	})
}

func DepartmentDelete(c *gin.Context) {
	departid := c.PostForm("departid")
	if departid == "" {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "必填信息为空",
		})
		return
	}
	err := handler.DepartmentDelete(departid)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "科室删除错误",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "OK",
	})
}

func DepartmentGet(c *gin.Context) {
	d, err := handler.DepartmentGet()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "科室查询错误",
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

func DepartmentDetail(c *gin.Context) {
	id := c.PostForm("departid")
	name := c.PostForm("name")
	userid := c.PostForm("userid")
	if id == "" || name == "" || userid == "" {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "必填信息为空",
		})
		return
	}
	err := handler.DepartmentDetailChange(id, name, userid)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "change department err:" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "OK",
	})
}

func DepartmentSelect(c *gin.Context) {
	name := c.PostForm("name")
	id := c.PostForm("id")
	if name == "" && id == "" {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "必填信息为空",
		})
		return
	}
	d, err := handler.DepartmentSelect(name, id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "select user err",
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
