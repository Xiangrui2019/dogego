package services

import (
	"dogego/models"
	"dogego/serializer"
	"net/http"
)

type UpdateProfileService struct {
	NickName string `form:"nick_name" json:"nick_name" binding:"required,min=2,max=100"`
	Bio      string `form:"bio" json:"bio" binding:"required,min=2,max=100"`
	Avatar   string `form:"avatar" json:"avatar" binding:"required,min=2,max=200"`
}

func (service *UpdateProfileService) Update(user *models.User) *serializer.Response {
	user.NickName = service.NickName
	user.Bio = service.Bio
	user.Avatar = service.Avatar

	if err := models.UpdateUserProfile(user); err != nil {
		return &serializer.Response{
			Code:    http.StatusInternalServerError,
			Message: "更新信息出错.",
			Error:   err.Error(),
		}
	}

	return &serializer.Response{
		Code:    http.StatusOK,
		Message: "更新用户信息成功.",
	}
}
