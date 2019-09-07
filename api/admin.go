package api

import (
	"dogego/utils"

	"github.com/gin-gonic/gin"
)

func Test(context *gin.Context) {
	context.JSON(200, utils.CurrentUser(context))
}
