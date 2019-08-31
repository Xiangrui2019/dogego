package tasks

import (
	"dogego/models"
	"log"
	"time"
)

func TimeTask() error {
	time.Sleep(time.Second)
	log.Println(time.Now().Unix())
	return nil
}

func TimeTask1(data models.TaskData) error {
	time.Sleep(time.Second)
	log.Println(data)
	log.Println(time.Now().Unix())
	return nil
}
