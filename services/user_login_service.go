package services

import (
	"dogego/models"
	"dogego/serializer"
)

type UserLoginService struct {
	PhoneNumber string `form:"user_name" json:"user_name" binding:"required,min=11,max=30"`
	Password    string `form:"password" json:"password" binding:"required,min=2,max=40"`
}

func (service *UserLoginService) Login() (models.User, *serializer.Response) {}
