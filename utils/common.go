package utils

import (
	"dogego/models"
	"dogego/serializer"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v8"
)

func RandStringRunes(n int) string {
	var letterRunes = []rune("1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func BuildErrorResponse(err error) *serializer.Response {
	if ve, ok := err.(validator.ValidationErrors); ok {
		for _, e := range ve {
			field := fmt.Sprintf("Field.%s", e.Field)
			tag := fmt.Sprintf("Tag.Valid.%s", e.Tag)
			return &serializer.Response{
				Code:    http.StatusBadRequest,
				Message: fmt.Sprintf("%s%s", field, tag),
				Error:   fmt.Sprint(err),
			}
		}
	}
	if _, ok := err.(*json.UnmarshalTypeError); ok {
		return &serializer.Response{
			Code:    http.StatusBadRequest,
			Message: "JSON类型不匹配",
			Error:   fmt.Sprint(err),
		}
	}

	return &serializer.Response{
		Code:    http.StatusBadRequest,
		Message: "参数错误",
		Error:   fmt.Sprint(err),
	}
}

func CurrentUser(context *gin.Context) *models.User {
	if user, _ := context.Get("user"); user != nil {
		if u, ok := user.(*models.User); ok {
			return u
		}
	}
	return nil
}
