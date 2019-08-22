package services

import (
	"dogego/models"
	"dogego/serializer"
	"net/http"
)

type UserLoginService struct {
	PhoneNumber string `form:"phone_number" json:"phone_number" binding:"required,min=11,max=30"`
	Password    string `form:"password" json:"password" binding:"required,min=2,max=40"`
}

func (service *UserLoginService) Login() (models.User, *serializer.Response) {
	var user models.User

	if err := models.DB.Where("phone_number = ?", service.PhoneNumber).First(&user).Error; err != nil {
		return user, &serializer.Response{
			Code:    http.StatusBadRequest,
			Message: "账号或密码错误.",
		}
	}

	if user.CheckPassword(service.Password) == false {
		return user, &serializer.Response{
			Code:    http.StatusBadRequest,
			Message: "账号或密码错误.",
		}
	}

	return user, nil
}
