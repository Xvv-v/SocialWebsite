package routers

import (
	"github.com/gin-gonic/gin"
)

//Account 账户
func Account() *gin.Engine {

	r := gin.Default()
	accGroup := r.Group("/account")
	{
		accGroup.POST("/register", func(c *gin.Context) {

		})
		accGroup.POST("/login", func(c *gin.Context) {

		})
		accGroup.POST("/logout", func(c *gin.Context) {

		})
	}
	return r
}

func user() *gin.Engine {
	r := gin.Default()
	userGroup := r.Group("/user")
	{
		userGroup.GET("/personal", func(c *gin.Context) {

		})
		userGroup.GET("/public", func(c *gin.Context) {

		})
	}
	return r
}

//SetupRouters 设置路由
func SetupRouters() *gin.Engine {

	r := gin.Default()

	accGroup := r.Group("/account")
	{
		accGroup.POST("/register", func(c *gin.Context) {

		})
		accGroup.POST("/login", func(c *gin.Context) {

		})
		accGroup.POST("/logout", func(c *gin.Context) {

		})
	}

	userGroup := r.Group("/user")
	{
		userGroup.GET("/personal", func(c *gin.Context) {

		})
		userGroup.GET("/public", func(c *gin.Context) {

		})
	}

	return r
}
