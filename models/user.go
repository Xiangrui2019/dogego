package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	PhoneNumber string
	Password    string
	NickName    string
	Bio         string
	Status      string
	Avatar      string
}

const (
	PasswordCost = 12
	// Active 激活用户
	Active string = "active"
	// Inactive 未激活用户
	Inactive string = "inactive"
	// Suspend 被封禁用户
	Suspend string = "suspend"
)
