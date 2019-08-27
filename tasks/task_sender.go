package tasks

import (
	"dogego/global"
	"dogego/modules"
	"dogego/utils"

	"github.com/streadway/amqp"
)

func PublishTask(data *modules.Task) error {
	ch, queue, err := utils.BuildQueueChannel(global.TimeTaskQueueKey())

	err = ch.Publish(
		"",
		queue.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(data.Taskname),
		},
	)

	if err != nil {
		return err
	}

	return nil
}
