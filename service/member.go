package service

import (
	"errors"
	"github.com/gin-gonic/gin"
	"hie/main/handler"
	"net/http"
	"strconv"
)

func MemberDataReceive(c *gin.Context) {
	name := c.PostForm("name")
	password := c.PostForm("password")
	realname := c.PostForm("realname")
	sex := c.PostForm("sex")
	age, _ := strconv.Atoi(c.PostForm("age"))
	tel := c.PostForm("tel")
	credit := c.PostForm("credit")
	balance, _ := strconv.Atoi(c.PostForm("balance"))
	anaphylaxis := c.PostForm("anaphylaxis")
	err := MemberCreate(password, name, realname, sex, tel, credit, anaphylaxis, balance, age)
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

func MemberCreate(password, name, realname, sex, tel, credit, anaphylaxis string, balance, age int) error {
	if password == "" || name == "" || realname == "" || sex == "" || age == 0 || tel == "" || credit == "" || anaphylaxis == "" {
		return errors.New("必填项不足")
	}
	//判断用户是否已存在
	cnt, err := handler.IsMemberCreated(name)
	if err != nil {
		return err
	}
	if cnt > 0 {
		return errors.New("已存在用户")
	}
	err = handler.MemberCreate(name, password, realname, sex, tel, credit, anaphylaxis, balance, age)
	if err != nil {
		return err
	}
	return nil
}

func MemberDelete(c *gin.Context) {
	memberId := c.PostForm("member_id")
	if memberId == "" {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "必填信息为空",
		})
		return
	}
	err := handler.MemberDelete(memberId)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "会员删除错误",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "OK",
	})
}

func MemberGet(c *gin.Context) {
	m, err := handler.MemberGet()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "会员查询错误",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg": map[string]interface{}{
			"data": m,
		},
	})
}

func MemberDetailChange(c *gin.Context) {
	id := c.PostForm("id")
	name := c.PostForm("name")
	password := c.PostForm("password")
	realname := c.PostForm("realname")
	sex := c.PostForm("sex")
	age, _ := strconv.Atoi(c.PostForm("age"))
	tel := c.PostForm("tel")
	credit := c.PostForm("credit")
	balance, _ := strconv.Atoi(c.PostForm("balance"))
	anaphylaxis := c.PostForm("anaphylaxis")
	if id == "" || password == "" || name == "" || realname == "" || sex == "" || age == 0 || tel == "" || credit == "" || anaphylaxis == "" {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "必填信息为空",
		})
		return
	}
	err := handler.MemberDetailChange(id, name, password, realname, sex, tel, credit, anaphylaxis, balance, age)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "change member err:" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "OK",
	})
}

// MemberSelect 模糊查询
func MemberSelect(c *gin.Context) {
	MemberId := c.PostForm("member_id")
	MemberName := c.PostForm("member_name")
	credit := c.PostForm("credit")
	tel := c.PostForm("tel")
	if MemberId == "" && MemberName == "" && credit == "" && tel == "" {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "必填信息为空",
		})
		return
	}
	m, err := handler.MemberSelect(MemberId, MemberName, credit, tel)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "select member err:" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"message": map[string]interface{}{
			"data": m,
		},
	})
}

// MemberRecharge 会员充值
func MemberRecharge(c *gin.Context) {
	memberid := c.PostForm("id")
	money := c.PostForm("money")
	rechargemethod := c.PostForm("rechargemethod")
	userid := c.PostForm("userid")
	err := handler.MemberRecharge(memberid, money, rechargemethod, userid)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "Recharge member err:",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "ok",
	})
}
