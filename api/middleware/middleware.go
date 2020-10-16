package api

import (
	"SocialWebsite/infrastructure/general"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

//JWTAuthMiddleware jwt鉴权中间件
func JWTAuthMiddleware() func(c *gin.Context) {

	return func(c *gin.Context) {
		//token放在请求头中，以Bearer开头
		//取出header中的Authorization字段
		authHead := c.Request.Header.Get("Authorization")
		if authHead == "" {
			c.JSON(http.StatusOK, gin.H{
				"code": 200,
				"msg":  "请求头中auth为空",
				"data": "",
			})
			c.Abort()
			return
		}
		//将Bearer与token分割开
		//按照空格分隔
		parts := strings.SplitN(authHead, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			//如果parts长度不为2且第一个元素值不为Bearer时
			c.JSON(http.StatusOK, gin.H{
				"code": 200,
				"msg":  "auth格式错误",
				"data": "",
			})
			c.Abort()
			return
		}
		//开始检验token
		claim, err := general.ParaseAccessToken(parts[1])
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 200,
				"msg":  "无效token", //重新登陆
				"data": "",
			})
			return
		}
		//将当前请求的id写入请求的上下文
		c.Set("userid", claim.Userid)
		c.Next()
	}
}
