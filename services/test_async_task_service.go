package services

import (
	"dogego/serializer"
	"dogego/tasks"
	"dogego/utils"
	"net/http"
)

type TestAsyncTaskService struct {
}

func (service *TestAsyncTaskService) Test() *serializer.Response {
	err := utils.RunAsyncTask(tasks.TimeTask1, "TH")

	if err != nil {
		return utils.BuildErrorResponse(err).Result()
	}

	return serializer.Response{
		Code:    http.StatusOK,
		Message: "Successfly send task async task.",
	}.Result()
}
