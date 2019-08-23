package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RunTask(context *gin.Context, job func() error) {
	err := job()

	if err != nil {
		context.String(http.StatusInternalServerError, err.Error())
	} else {
		context.String(http.StatusOK, "")
	}
}
