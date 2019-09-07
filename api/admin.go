package api

import "github.com/gin-gonic/gin"

func Test(context *gin.Context) {
	context.JSON(200, "dffd")
}
