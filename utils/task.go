package utils

import (
	"dogego/global"
	"dogego/modules"
	"log"
	"reflect"
	"runtime"
)

func RunAsyncTask(job modules.AsyncTask, data interface{}) error {
	jobName := runtime.FuncForPC(reflect.ValueOf(job).Pointer()).Name()

	err := modules.RedisMQModule.Publish(
		global.AsyncTaskQueueKey(),
		global.AsyncTaskData(jobName, &modules.TaskData{
			Data: data,
		}),
	)

	log.Println("Async Task Commit Success: ", jobName)

	if err != nil {
		return err
	}

	return nil
}

func PublishTask(data *modules.Task) error {
	err := modules.RedisMQModule.Publish(
		global.TimeTaskQueueKey(),
		data.Taskname,
	)

	log.Println("Timed Task Commit Success: ", data.Taskname)

	if err != nil {
		return err
	}

	return nil
}
