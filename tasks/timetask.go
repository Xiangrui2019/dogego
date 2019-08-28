package tasks

import (
	"log"
	"time"
)

func TimeTask() error {
	time.Sleep(time.Second)
	log.Println(time.Now().Unix())
	return nil
}

func TimeTask1(data interface{}) error {
	time.Sleep(time.Second)
	log.Println(time.Now().Unix())
	return nil
}
