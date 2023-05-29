package service

import (
	"errors"
	"github.com/gin-gonic/gin"
	"hie/main/handler"
	"net/http"
	"strconv"
)

func DoctorDataReceive(c *gin.Context) {
	name := c.PostForm("name")
	password := c.PostForm("password")
	realname := c.PostForm("realname")
	departId := c.PostForm("deparid")
	positional := c.PostForm("positional")
	age, _ := strconv.Atoi(c.PostForm("age"))
	tel := c.PostForm("tel")
	sex := c.PostForm("sex")
	err := DoctorCreate(password, name, realname, departId, positional, tel, sex, age)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  err,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "OK",
	})
}

func DoctorCreate(password, name, realname, departId, positional, tel, sex string, age int) error {
	if password == "" || name == "" || realname == "" || sex == "" || age == 0 || tel == "" || departId == "" || positional == "" {
		return errors.New("必填项不足")
	}
	//判断用户是否已存在
	cnt, err := handler.IsDoctorCreated(name)
	if err != nil {
		return err
	}
	if cnt > 0 {
		return errors.New("用户已存在")
	}
	err = handler.DoctorCreate(name, password, realname, departId, sex, tel, positional, age)
	if err != nil {
		return err
	}
	return nil
}

func DoctorDelete(c *gin.Context) {
	doctorId := c.PostForm("doctor_id")
	if doctorId == "" {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "必填信息为空",
		})
		return
	}
	err := handler.DoctorDelete(doctorId)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "医生删除错误",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "OK",
	})
}

func GetAllDoctor(c *gin.Context) {
	d, err := handler.DoctorGet()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "医生查询错误",
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

func DoctorDetailChange(c *gin.Context) {
	id := c.PostForm("id")
	name := c.PostForm("name")
	password := c.PostForm("password")
	realname := c.PostForm("realname")
	departId := c.PostForm("deparid")
	positional := c.PostForm("positional")
	age, _ := strconv.Atoi(c.PostForm("age"))
	tel := c.PostForm("tel")
	sex := c.PostForm("sex")
	if id == "" || password == "" || name == "" || realname == "" || sex == "" || age == 0 || tel == "" || departId == "" || positional == "" {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "必填信息为空",
		})
		return
	}
	err := handler.DoctorDetailChange(id, name, password, tel, departId, realname, sex, positional, age)
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

func DoctorSelect(c *gin.Context) {
	doctorId := c.PostForm("doctor_id")
	doctorName := c.PostForm("doctor_name")
	department := c.PostForm("department")
	if doctorId == "" && doctorName == "" && department == "" {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "必填信息为空",
		})
		return
	}
	d, err := handler.DoctorSelect(doctorName, doctorId, department)
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
			"data": d,
		},
	})
}

func GetDoctorByDepartId(c *gin.Context) {
	departId := c.PostForm("depart_id")
	if departId == "" {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "必填信息为空",
		})
		return
	}
	d, err := handler.GetDoctorByDepartId(departId)
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
			"data": d,
		},
	})

}
