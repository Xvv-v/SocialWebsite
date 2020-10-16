package routers

import (
	"SocialWebsite/api/controller"
	"SocialWebsite/api/middleware"

	"github.com/gin-gonic/gin"
)

//User 用户
func User() *gin.Engine {

	//不带中间件
	r := gin.New()
	userGroup := r.Group("/user")
	userGroup.Use(middleware.JWTAuthMiddleware())
	{
		userGroup.POST("/register", controller.NewUser)
		userGroup.POST("/login", controller.LogIn)
		userGroup.POST("/logout", controller.LogOut)
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
