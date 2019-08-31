package services

import (
	"dogego/models"
	"dogego/serializer"
	"dogego/tasks"
	"dogego/utils"
	"net/http"
)

type TestAsyncTaskService struct {
}

func (service *TestAsyncTaskService) Test() *serializer.Response {
	err := utils.RunAsyncTask(tasks.TimeTask1, models.TaskData{
		Data: "tg",
	})

	if err != nil {
		return utils.BuildErrorResponse(err)
	}

	return &serializer.Response{
		Code:    http.StatusOK,
		Message: "Successfly send task async task.",
	}
}
