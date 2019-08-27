package utils

import (
	"log"
	"time"
)

func RunTask(jobName string, job func() error) {
	from := time.Now().UnixNano()
	err := job()
	to := time.Now().UnixNano()

	if err != nil {
		log.Printf("%s Execute Error: %dms\n", jobName, (to-from)/int64(time.Millisecond))
	} else {
		log.Printf("%s Execute Success: %dms\n", jobName, (to-from)/int64(time.Millisecond))
	}
}
