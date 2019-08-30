package http

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HttpProtocol struct {
	Router *gin.Engine
}

func NewHttpProtocol(router *gin.Engine) *HttpProtocol {
	return &HttpProtocol{
		Router: router,
	}
}

func (protocol *HttpProtocol) Start(addr string) error {
	httpServer := http.Server{
		Addr:    addr,
		Handler: protocol.Router,
	}

	log.Printf("Server started on %s", addr)
	err := httpServer.ListenAndServe()

	if err != nil {
		return err
	}

	return nil
}
