package utils

import (
	"dogego/global"
	"dogego/modules"
	"reflect"
	"runtime"

	"github.com/streadway/amqp"
)

func RunAsyncTask(job modules.AsyncTask, data interface{}) error {
	ch, queue, err := BuildQueueChannel(global.AsyncTaskQueueKey())

	if err != nil {
		return err
	}

	jobName := runtime.FuncForPC(reflect.ValueOf(job).Pointer()).Name()

	err = Publish(ch, queue, amqp.Publishing{
		ContentType:  "text/plain",
		Body:         []byte(global.AsyncTaskData(jobName, data)),
		DeliveryMode: amqp.Transient,
	})

	if err != nil {
		return err
	}

	return nil
}
