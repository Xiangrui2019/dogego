package middlewares

import (
	"dogego/models"
	"dogego/serializer"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func CurrentUser() gin.HandlerFunc {
	return func(context *gin.Context) {
		session := sessions.Default(context)
		uid := session.Get("user_id")
		if uid != nil {
			user, err := models.GetUserById(uid)
			if err == nil {
				context.Set("user", user)
			}
		}
		context.Next()
	}
}

func AuthRequired(role []string) gin.HandlerFunc {
	return func(context *gin.Context) {
		if user, _ := context.Get("user"); user != nil {
			if v, ok := user.(*models.User); ok {
				for _, r := range role {
					if v.Role == r {
						context.Next()
						return
					}
				}
			}
		}

		context.JSON(http.StatusUnauthorized, serializer.Response{
			Code:    http.StatusUnauthorized,
			Message: "您还没有登录或没有权限.",
		})
		context.Abort()
	}
}
