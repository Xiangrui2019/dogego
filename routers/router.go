package routers

import (
	"dogego/api"
	"dogego/auth"
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

		// 需要进行登录验证(auth.User)
		userauthed := v1.Group("")
		{
			// 使用登录验证中间件
			userauthed.Use(middlewares.AuthRequired(auth.User))

			userauthed.PUT("/user/change_password", api.UserChangePassword)
			userauthed.PUT("/user/update_profile", api.UserUpdateProfile)
			userauthed.GET("/user/me", api.UserMe)
			userauthed.POST("/usr/logout", api.UserLogout)
		}

		// 需要进行登录验证(auth.Admin)
		adminauthed := v1.Group("")
		{
			// 使用登录验证中间件
			adminauthed.Use(middlewares.AuthRequired(auth.Admin))

			adminauthed.GET("/auth/admintest", api.Test)
		}
	}

	return router
}
