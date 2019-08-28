package tasks

import (
	"dogego/global"
	"dogego/modules"
	"dogego/utils"
	"log"

	"github.com/streadway/amqp"
)

func PublishTask(data *modules.Task) error {
	ch, queue, err := utils.BuildQueueChannel(global.TimeTaskQueueKey())

	if err != nil {
		return err
	}

	err = utils.Publish(ch, queue, amqp.Publishing{
		ContentType:  "text/plain",
		Body:         []byte(data.Taskname),
		DeliveryMode: amqp.Transient,
	})

	log.Println("Task Publish Success: ", data.Taskname)

	if err != nil {
		return err
	}

	return nil
}
