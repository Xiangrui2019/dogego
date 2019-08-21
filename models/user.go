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
