package modules

import (
	"dogego/cache"
	"dogego/global"
	"log"
)

type RedisMQ struct {
}

func (mq *RedisMQ) Publish(queuename string, message string) error {
	err := cache.CacheClient.LPush(global.QueueNameKey(queuename), message).Err()

	if err != nil {
		return err
	}

	return nil
}

func (mq *RedisMQ) Custome(queuename string, cb func(message string) error) error {
	go func() {
		for {
			message, err := cache.CacheClient.RPop(
				global.QueueNameKey(queuename)).Result()

			if err != nil {
				log.Println(err)
			}

			err = cb(message)

			log.Printf("Execute Callback func Error: %s", err)
		}
	}()

	return nil
}
