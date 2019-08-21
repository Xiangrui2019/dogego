package middlewares

import (
	"dogego/models"
	"dogego/serializer"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func CurrentUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		uid := session.Get("user_id")
		if uid != nil {
			user, err := models.GetUserById(uid)
			if err == nil {
				c.Set("user", &user)
			}
		}
		c.Next()
	}
}

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		if user, _ := c.Get("user"); user != nil {
			if _, ok := user.(*models.User); ok {
				c.Next()
				return
			}
		}

		c.JSON(http.StatusUnauthorized, serializer.Response{
			Status: http.StatusUnauthorized,
			Msg:    "需要登录",
		})
		c.Abort()
	}
}
