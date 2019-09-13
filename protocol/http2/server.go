package http2

import (
	"crypto/tls"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type Http2Protocol struct {
	Router *gin.Engine
}

func NewHttp2Protocol(router *gin.Engine) *Http2Protocol {
	return &Http2Protocol{
		Router: router,
	}
}

func (protocol *Http2Protocol) Start(addr string) error {
	http2Server := http.Server{
		Addr:    addr,
		Handler: protocol.Router,
	}

	if _, err := tls.LoadX509KeyPair(os.Getenv("TLS_PEM"), os.Getenv("TLS_KEY")); err != nil {
		log.Fatal(err)
	}

	log.Printf("HTTP2 Server started on %s", addr)
	err := http2Server.ListenAndServeTLS(
		os.Getenv("TLS_PEM"), os.Getenv("TLS_KEY"))

	if err != nil {
		return err
	}

	return nil
}
