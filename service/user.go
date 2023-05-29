package service

import (
	"errors"
	"github.com/gin-gonic/gin"
	"hie/main/handler"
	"net/http"
	"strconv"
)

func UserDataReceive(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	role := c.PostForm("role")
	realname := c.PostForm("realname")
	sex := c.PostForm("sex")
	age, _ := strconv.Atoi(c.PostForm("age"))
	tel := c.PostForm("tel")
	address := c.PostForm("address")
	if role != "01" && role != "02" && role != "03" {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "请填写正确的角色",
		})
		return
	}
	err := UserCreate(password, username, role, realname, sex, tel, address, age)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": err,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "OK",
	})
}

func UserCreate(password, username, role, realname, sex, tel, address string, age int) error {
	if password == "" || username == "" || role == "" || realname == "" || sex == "" || age == 0 || tel == "" || address == "" {
		return errors.New("必填信息为空")
	}
	//判断用户是否已存在
	cnt, err := handler.IsUserCreatedByName(username)
	if err != nil {
		return err
	}
	if cnt > 0 {
		return errors.New("用户已存在")
	}
	err = handler.UserCreate(username, password, role, realname, sex, tel, address, age)
	if err != nil {
		return err
	}
	return nil
}

func UserDelete(c *gin.Context) {
	userId := c.PostForm("userid")
	if userId == "" {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "必填信息为空",
		})
		return
	}
	err := handler.UserDelete(userId)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "用户删除错误",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "OK",
	})
}

func GetAllUser(c *gin.Context) {
	u, err := handler.UserGet()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "用户查询错误",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg": map[string]interface{}{
			"data": u,
		},
	})
}

func UserDetailChange(c *gin.Context) {
	id := c.PostForm("id")
	username := c.PostForm("username")
	password := c.PostForm("password")
	role := c.PostForm("role")
	realname := c.PostForm("realname")
	sex := c.PostForm("sex")
	age, _ := strconv.Atoi(c.PostForm("age"))
	tel := c.PostForm("tel")
	address := c.PostForm("address")
	if id == "" || password == "" || username == "" || role == "" || realname == "" || sex == "" || age == 0 || tel == "" || address == "" {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "必填信息为空",
		})
		return
	} else if role != "01" && role != "02" && role != "03" {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "请填写正确的角色",
		})
		return
	}
	err := handler.UserDetailChange(id, username, password, tel, address, role, realname, sex, age)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "change user err:" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "OK",
	})
}

func UserSelect(c *gin.Context) {
	username := c.PostForm("username")
	userId := c.PostForm("id")
	userRole := c.PostForm("role")
	if username == "" && userId == "" && userRole == "" {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "必填信息为空",
		})
		return
	}
	u, err := handler.UserSelect(username, userId, userRole)
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
			"data": u,
		},
	})
}
