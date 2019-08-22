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
	router.Use(middlewares.Session(os.Getenv("SESSION_SECERT")))
	router.Use(middlewares.Cors(os.Getenv("CORS_DOMAIN")))
	router.Use(middlewares.CurrentUser())

	v1 := router.Group("/api/v1")
	{
		v1.POST("/ping", api.Ping)
		v1.POST("/user/register", api.UserRegister)
		v1.POST("/user/login", api.UserLogin)

		auth := v1.Use(middlewares.AuthRequired())
		{
			auth.POST("/user/change_password", api.UserChangePassword)
			auth.GET("/user/me", api.UserMe)
		}
	}

	return router
}
