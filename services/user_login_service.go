package services

import (
	"dogego/models"
	"dogego/serializer"
)

type UserLoginService struct {
}

func (service *UserLoginService) Login() (models.User, *serializer.Response) {}
