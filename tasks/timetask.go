package tasks

import (
	"log"
	"time"
)

func TimeTask() error {
	log.Println(time.Now().Unix())
	return nil
}
