package service

import (
	"github.com/gin-gonic/gin"
	"hie/main/models"
	"net/http"
	"strconv"
)

func Create(c *gin.Context) {
	username := c.PostForm("name")
	password := c.PostForm("password")
	role := c.PostForm("role")
	realname := c.PostForm("realname")
	sex := c.PostForm("sex")
	age, _ := strconv.Atoi(c.PostForm("age"))
	tel := c.PostForm("tel")
	address := c.PostForm("address")
	credit := c.PostForm("credit")
	balance, _ := strconv.Atoi(c.PostForm("balance"))
	anaphylaxis := c.PostForm("anaphylaxis")
	departId := c.PostForm("deparid")
	positional := c.PostForm("positional")
	err := models.DB.Error
	switch role {
	case "01", "02", "03":
		err = UserCreate(password, username, role, realname, sex, tel, address, age)
		break
	case "04":
		err = DoctorCreate(password, username, realname, departId, positional, tel, sex, age)
		break
	case "05":
		err = MemberCreate(password, username, realname, sex, tel, credit, anaphylaxis, balance, age)
		break
	default:
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "请填写正确的角色",
		})
		return
	}
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "OK",
	})
}
