package api

import (
	"dogego/services"
	"dogego/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserRegister(context *gin.Context) {
	service := services.UserRegisterService{}

	if err := context.ShouldBind(&service); err == nil {
		if user, err := service.Register(); err != nil {
			context.JSON(http.StatusInternalServerError, err)
		} else {
			context.JSON(http.StatusOK, user)
		}
	} else {
		context.JSON(http.StatusBadRequest, utils.BuildErrorResponse(err))
	}
}

func UserLogin(context *gin.Context) {

}

func UserMe(context *gin.Context) {

}

func UserLogout(context *gin.Context) {

}
