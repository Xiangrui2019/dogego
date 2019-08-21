package models

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

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

func RegisterUser(user *User) error {
	err := DB.Create(&user).Error

	if err != nil {
		return nil
	}

	return err
}

func GetUserById(ID interface{}) (*User, error) {
	var user User

	err := DB.First(&user, ID).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func UpdateUserProfile(user *User) error {
	err := DB.Save(&user).Error

	if err != nil {
		return err
	}

	return nil
}

func (user *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PasswordCost)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}

func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err == nil
}
