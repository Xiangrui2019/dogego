package api

import (
	"dogego/serializer"
	"dogego/services"
	"dogego/utils"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func UserRegister(context *gin.Context) {
	service := services.UserRegisterService{}

	if err := context.ShouldBind(&service); err == nil {
		if user, err := service.Register(); err != nil {
			context.JSON(http.StatusInternalServerError, err)
		} else {
			context.JSON(http.StatusOK, serializer.BuildUserResponse(&user))
		}
	} else {
		context.JSON(http.StatusBadRequest, utils.BuildErrorResponse(err))
	}
}

func UserLogin(context *gin.Context) {
	service := services.UserLoginService{}

	if err := context.ShouldBind(&service); err == nil {
		if user, err := service.Login(); err != nil {
			context.JSON(http.StatusBadRequest, err)
		} else {
			session := sessions.Default(context)
			session.Clear()
			session.Set("user_id", user.ID)
			session.Save()

			context.JSON(http.StatusOK, serializer.BuildUserResponse(&user))
		}
	} else {
		context.JSON(http.StatusBadRequest, utils.BuildErrorResponse(err))
	}
}

func UserMe(context *gin.Context) {
	user := utils.CurrentUser(context)

	context.JSON(http.StatusOK, serializer.BuildUserResponse(user))
}

func UserUpdateProfile(context *gin.Context) {

}

func UserChangePassword(context *gin.Context) {

}

func UserLogout(context *gin.Context) {
	session := sessions.Default(context)
	session.Clear()
	session.Save()
	context.JSON(http.StatusOK, &serializer.Response{
		Code: http.StatusOK,
		Message: "登出成功.",
	})
}
