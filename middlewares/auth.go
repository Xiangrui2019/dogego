package middlewares

import (
	"dogego/models"

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
