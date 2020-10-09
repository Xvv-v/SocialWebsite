package model

import (
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
)

//Register 定义注册数据模型
type Register struct {
	Username     string `json:"username,omitempty" binding:"username"`         //用户名
	Telephone    string `json:"telephone,omitempty" bindin:"telephone"`        //手机号
	Email        string `json:"email,omitempty" binding:"email"`               //邮箱
	Password     string `json:"password,omitempty" binding:"password"`         //密码
	Verification string `json:"verification,omitempty" binding:"verification"` //验证码
	Gender       string `json:"gender,omitempty" binding:"gender"`             //性别
	Region       string `json:"region,omitempty" binding:"region"`             //地区
	Education    string `json:"education,omitempty" binding:"education"`       //教育经历
}

//Login 定义登陆数据模型
type Login struct {
	Account  string `json:"account，omitempty" binding:"account"`   //账号
	Password string `json:"password，omitempty" binding:"password"` //密码
}

//NewUser 新用户注册
func NewUser(c *gin.Context) {

	var r Register
	if err := c.ShouldBind(&r); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "错误请求",
			"data": "",
		})
	}
	//检查手机号是否匹配 ^1[3|4|5|7|8]\d{9}$
	if ok, _ := regexp.MatchString("^1[3|4|5|7|8]d{9}$", r.Telephone); !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 401,
			"msg":  "手机号码格式错误",
			"data": "",
		})
	}
	//检查邮箱 ^([A-Za-z0-9_\-\.])+\@([A-Za-z0-9_\-\.])+\.([A-Za-z]{2,4})$
	if ok, _ := regexp.MatchString("^([A-Za-z0-9_-.])+@([A-Za-z0-9_-.])+.([A-Za-z]{2,4})$", r.Email); !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 401,
			"msg":  "邮箱格式错误",
			"data": "",
		})
	}
	//检查验证码

	//检查密码
}

//LogIn 登录
func LogIn(c *gin.Context) {

	var login Login
	if err := c.ShouldBind(&login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "请求错误",
			"data": "",
		})
	}
	//将账号密码数据传给
}
