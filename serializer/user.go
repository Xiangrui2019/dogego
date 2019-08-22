package serializer

import (
	"dogego/models"
	"net/http"
)

type User struct {
	ID          uint   `json:"id"`
	PhoneNumber string `json:"phone_number"`
	Nickname    string `json:"nickname"`
	Bio         string `json:"bio"`
	Status      string `json:"status"`
	Avatar      string `json:"avatar"`
	CreatedAt   int64  `json:"created_at"`
}

// BuildUser 序列化用户
func BuildUser(user *models.User) User {
	return User{
		ID:          user.ID,
		PhoneNumber: user.PhoneNumber,
		Nickname:    user.NickName,
		Bio:         user.Bio,
		Status:      user.Status,
		Avatar:      user.Avatar,
		CreatedAt:   user.CreatedAt.Unix(),
	}
}

// BuildUserResponse 序列化用户响应
func BuildUserResponse(user *models.User) *Response {
	return &Response{
		Code:    http.StatusOK,
		Message: "序列化用户成功.",
		Data:    BuildUser(user),
	}
}
