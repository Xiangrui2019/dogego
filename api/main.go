package api

import (
	"dogego/serializer"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Ping(context *gin.Context) {
	context.JSON(http.StatusOK, &serializer.Response{
		Code:    http.StatusOK,
		Message: "Pong",
	})
}
