package middlewares

import (
	"log"
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

// Session 初始化session
func Session(secret string) gin.HandlerFunc {
	store, err := redis.NewStore(10, "tcp", os.Getenv("REDIS_ADDR"), os.Getenv("REDIS_PASSWORD"), []byte(secret))
	if err != nil {
		log.Fatal(err)
	}
	store.Options(sessions.Options{HttpOnly: false, MaxAge: 30 * 86400, Path: "/"})
	return sessions.Sessions("gin-session", store)
}
