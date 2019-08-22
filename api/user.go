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
			context.JSON(err.Code, err)
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
			context.JSON(err.Code, err)
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
	service := services.UpdateProfileService{}
	user := utils.CurrentUser(context)

	if err := context.ShouldBind(&service); err == nil {
		res := service.Update(user)

		context.JSON(res.Code, res)
	} else {
		context.JSON(http.StatusBadRequest, utils.BuildErrorResponse(err))
	}
}

func UserChangePassword(context *gin.Context) {
	service := services.ChangePasswordService{}
	user := utils.CurrentUser(context)

	if err := context.ShouldBind(&service); err == nil {
		res := service.Change(user)
		session := sessions.Default(context)
		session.Clear()
		session.Save()
		context.JSON(res.Code, res)
	} else {
		context.JSON(http.StatusBadRequest, utils.BuildErrorResponse(err))
	}
}

func UserLogout(context *gin.Context) {
	session := sessions.Default(context)
	session.Clear()
	session.Save()
	context.JSON(http.StatusOK, &serializer.Response{
		Code:    http.StatusOK,
		Message: "登出成功.",
	})
}
