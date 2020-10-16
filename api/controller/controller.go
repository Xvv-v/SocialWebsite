package api

import (
	"SocialWebsite/api/model"
	"SocialWebsite/domain/user/models"
	"fmt"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
)

//NewUser 新用户注册
func NewUser(c *gin.Context) {

	fmt.Println("api.Newuser()")
	var r model.UserVO
	if err := c.ShouldBindJSON(&r); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "错误请求",
			"err":  err.Error(),
			"data": "",
		})
		return
	}
	fmt.Println("name:"+r.Username, "number:"+r.Telephone, "email:"+r.Email, "pwd:"+r.Password)
	//开始检验请求数据
	//检查手机号是否匹配 ^1[3|4|5|7|8]\d{9}$
	if r.Telephone != "" {
		//手机号不为空，即用手机号注册的，检查手机号
		if ok, _ := regexp.MatchString("^1[3|4|5|7|8]\\d{9}$", r.Telephone); !ok {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": 401,
				"msg":  "手机号码格式错误",
				"data": "",
			})
			return
		}
	} else {
		//使用邮箱注册，检查邮箱 /^([A-Za-z0-9_\-\.\u4e00-\u9fa5])+\@([A-Za-z0-9_\-\.])+\.([A-Za-z]{2,8})$/
		if ok, _ := regexp.MatchString("/^([A-Za-z0-9_\\-\\.\u4e00-\u9fa5])+\\@([A-Za-z0-9_\\-\\.])+\\.([A-Za-z]{2,8})$/",
			r.Email); !ok {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": 401,
				"msg":  "邮箱格式错误",
				"data": "",
			})
			return
		}
	}
	//检查密码  ^(?=.*?[A-Z])(?=.*?[a-z])(?=.*?[0-9])(?=.*?[#?!@$%^&*-]).{8,}$
	if len(r.Password) < 8 {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "密码不能小于八位",
			"data": "",
		})
	}
	//全都没有问题了，进行下步操作
	u := models.NewUserDO(r.Telephone, r.Email, r.Password)
	err := u.Register()
	if err == nil {
		//没有注册过
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "注册成功！",
			"data": "",
		})
		return
	}
	//账号已存在
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  err.Error(),
		"data": "",
	})
	return
}

//LogIn 登录
func LogIn(c *gin.Context) {

	var login model.LoginVO
	if err := c.ShouldBind(&login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "请求错误",
			"data": "",
		})
	}
	//将账号密码数据传给domain层
	token, err := models.NewUserDO(login.Telephone, login.Email, login.Password).Login(true)
	if token != nil {
		//登陆成功，将令牌返回
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "登陆成功",
			"data": token,
		})
		return
	} else if err.Error() == "" {
		//账号或密码错误
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  err.Error(),
			"data": "",
		})
	}
	//账号未注册
	c.JSON(http.StatusBadRequest, gin.H{
		"code": 400,
		"msg":  err.Error(),
		"data": "",
	})
}

//LogOut 登出
func LogOut() {

}
