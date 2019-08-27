package utils

import (
	"dogego/modules"

	"github.com/streadway/amqp"
)

func BuildChannel() (*amqp.Channel, error) {
	ch, err := modules.AMQPModule.Channel()

	if err != nil {
		return nil, err
	}

	defer ch.Close()

	return ch, nil
}

func BuildQueue(queuename string, ch *amqp.Channel) (*amqp.Queue, error) {
	queue, err := ch.QueueDeclare(
		queuename,
		false,
		false,
		false,
		false,
		amqp.Table{},
	)

	if err != nil {
		return nil, err
	}

	return &queue, nil
}

func BuildQueueChannel(queue_name string) (*amqp.Channel, *amqp.Queue, error) {
	ch, err := BuildChannel()

	if err != nil {
		return nil, nil, err
	}

	queue, err := BuildQueue(queue_name, ch)

	if err != nil {
		return ch, nil, err
	}

	return ch, queue, nil
}
