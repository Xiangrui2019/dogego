package services

import (
	"dogego/models"
	"dogego/serializer"
)

type UpdateProfileService struct {
	NickName string `form:"nick_name" json:"nick_name" binding:"required,min=2,max=100"`
	Bio      string `form:"bio" json:"bio" binding:"required,min=2,max=100"`
	Avatar   string `form:"avatar" json:"avatar" binding:"required,min=2,max=200"`
}

func (service *UpdateProfileService) Update(user *models.User) *serializer.Response {

}
