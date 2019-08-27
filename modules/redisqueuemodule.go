package modules

import (
	"log"
	"time"
)

type RedisQueue struct {
}

func (queue *RedisQueue) Publish(queuename string, message string) error {
	return nil
}

func (queue *RedisQueue) Receive(queuename string,
	callback func(message string) error) {
	go func() {
		for {
			err := queue.messageReceiver()
			log.Println(err)
		}
	}()
}

func (queue *RedisQueue) messageReceiver() error {
	time.Sleep(time.Duration)
	return nil
}
