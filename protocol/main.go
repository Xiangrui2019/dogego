package protocol

import (
	"log"

	"github.com/doge-soft/doge_server_http"
	"github.com/gin-gonic/gin"
)

var Servers []func(router *gin.Engine) error

// 注册Protocol的地方
func RegisterServers() {
	RegisterProtocol(doge_server_http.HTTPServerProtocol)
}

// 注册Protocol的工具函数
func RegisterProtocol(protocol func(router *gin.Engine) error) {
	Servers = append(Servers, protocol)
}

// 启动所有服务器的函数
func StartAllServers(router *gin.Engine) {
	for _, server := range Servers {
		go func(server func(router *gin.Engine) error) {
			err := server(router)

			if err != nil {
				log.Fatal(err)
			}
		}(server)
	}
}
