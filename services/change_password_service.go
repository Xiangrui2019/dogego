package services

import (
	"dogego/models"
	"dogego/serializer"
	"net/http"
)

type ChangePasswordService struct {
	OldPassword     string `form:"old_password" json:"old_password" binding:"required,min=2,max=40"`
	Password        string `form:"password" json:"password" binding:"required,min=2,max=40"`
	PasswordConfirm string `form:"password_confirm" json:"password_confirm" binding:"required,min=2,max=40"`
}

func (service *ChangePasswordService) Valid(user *models.User) *serializer.Response {
	if err := user.CheckPassword(service.OldPassword); err != true {
		return &serializer.Response{
			Code:    http.StatusBadRequest,
			Message: "旧密码输入有误",
		}
	}

	if service.PasswordConfirm != service.Password {
		return &serializer.Response{
			Code:    http.StatusBadRequest,
			Message: "两次输入的密码不相同",
		}
	}

	return nil
}

func (service *ChangePasswordService) Change(user *models.User) *serializer.Response {
	if err := service.Valid(user); err != nil {
		return err
	}

	if err := user.SetPassword(service.Password); err != nil {
		return &serializer.Response{
			Code:    http.StatusInternalServerError,
			Message: "加密密码失败",
			Error:   err.Error(),
		}
	}

	if err := models.DB.Save(user).Error; err != nil {
		return &serializer.Response{
			Code:    http.StatusInternalServerError,
			Message: "保存密码失败",
			Error:   err.Error(),
		}
	}

	return &serializer.Response{
		Code: http.StatusOK,
		Message: "密码更新成功, 请重新登录.",
	}
}
