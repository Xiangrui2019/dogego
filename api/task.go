package api

import (
	"dogego/services"

	"github.com/gin-gonic/gin"
)

func TestAsyncTask(context *gin.Context) {
	service := services.TestAsyncTaskService{}

	resp := service.Test()

	context.JSON(resp.Code, resp)
}
