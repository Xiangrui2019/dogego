package routers

import (
	"dogego/api"
	"dogego/middlewares"
	"os"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.Default()

	// 中间件, 顺序必须是这样
	router.Use(middlewares.Session(os.Getenv("SESSION_SECRET")))
	router.Use(middlewares.Cors(os.Getenv("CORS_DOMAIN")))
	router.Use(middlewares.CurrentUser())

	v1 := router.Group("/api/v1")
	{
		v1.POST("/ping", api.Ping)
		v1.POST("/user/register", api.UserRegister)
		v1.POST("/user/login", api.UserLogin)
		v1.POST("/task/test", api.TestAsyncTask)
		
		// 需要进行登录验证
		authed := v1.Group("")
		{
			// 使用登录验证中间件
			authed.Use(middlewares.AuthRequired())

			authed.PUT("/user/change_password", api.UserChangePassword)
			authed.PUT("/user/update_profile", api.UserUpdateProfile)
			authed.GET("/user/me", api.UserMe)
			authed.POST("/usr/logout", api.UserLogout)
		}
	}

	return router
}
