package middlewares

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Cors 跨域配置
func Cors(cors_domain string) gin.HandlerFunc {
	config := cors.DefaultConfig()
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Cookie"}
	if gin.Mode() != gin.ReleaseMode {
		config.AllowAllOrigins = true
	} else {
		config.AllowOrigins = []string{cors_domain}
	}
	config.AllowCredentials = true
	return cors.New(config)
}
