package services

import (
	"dogego/models"
	"dogego/serializer"
	"net/http"
)

type UserRegisterService struct {
	PhoneNumber     string `form:"phone_number" json:"phone_number" binding:"required,min=11,max=30"`
	NickName        string `form:"nick_name" json:"nick_name" binding:"required,min=2,max=100"`
	Password        string `form:"password" json:"password" binding:"required,min=2,max=40"`
	PasswordConfirm string `form:"password_confirm" json:"password_confirm" binding:"required,min=2,max=40"`
}

func (service *UserRegisterService) Valid() *serializer.Response {
	if service.PasswordConfirm != service.Password {
		return &serializer.Response{
			Code:    http.StatusBadRequest,
			Message: "两次输入的密码不相同",
		}
	}

	count := 0
	models.DB.Model(&models.User{}).Where("nick_name = ?", service.NickName).Count(&count)
	if count > 0 {
		return &serializer.Response{
			Code:    http.StatusBadRequest,
			Message: "昵称被占用",
		}
	}

	count = 0
	models.DB.Model(&models.User{}).Where("phone_number = ?", service.PhoneNumber).Count(&count)
	if count > 0 {
		return &serializer.Response{
			Code:    http.StatusBadRequest,
			Message: "电话号码已经注册",
		}
	}

	return nil
}

func (service *UserRegisterService) Register() (models.User, *serializer.Response) {
	user := models.User{
		PhoneNumber: service.PhoneNumber,
		NickName:    service.NickName,
		Bio:         "这个人很懒, 什么都没写....",
		Status:      models.Active,
		Avatar:      "",
	}

	if err := service.Valid(); err != nil {
		return user, err
	}

	if err := user.SetPassword(service.Password); err != nil {
		return user, &serializer.Response{
			Code:    http.StatusInternalServerError,
			Message: "设置密码时出错.",
			Error:   err.Error(),
		}
	}

	if err := models.RegisterUser(&user); err != nil {
		return user, &serializer.Response{
			Code:    http.StatusInternalServerError,
			Message: "注册用户时出错.",
			Error:   err.Error(),
		}
	}

	return user, nil
}
