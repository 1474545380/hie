package service

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"hie/main/handler"
	"net/http"
)

func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	if password == "" || username == "" {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "必填信息为空",
		})
		return
	}
	err := handler.UserLogin(username, password)
	if err != nil {
		err = handler.DoctorLogin(username, password)
		if err != nil {
			err = handler.MemberLogin(username, password)
			if err != nil {
				if err == gorm.ErrRecordNotFound {
					c.JSON(http.StatusOK, gin.H{
						"code":    -1,
						"message": "用户名或密码错误",
					})
					return
				}
				c.JSON(http.StatusOK, gin.H{
					"code":    -1,
					"message": "get user err:" + err.Error(),
				})
				return
			}
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "OK",
	})
}
