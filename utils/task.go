package utils

import (
	"log"
	"reflect"
	"runtime"
	"time"
)

func RunTask(job func() error) {
	from := time.Now().UnixNano()
	err := job()
	to := time.Now().UnixNano()
	jobName := runtime.FuncForPC(reflect.ValueOf(job).Pointer()).Name()
	if err != nil {
		log.Printf("%s Execute Error: %dms\n", jobName, (to-from)/int64(time.Millisecond))
	} else {
		log.Printf("%s Execute Success: %dms\n", jobName, (to-from)/int64(time.Millisecond))
	}
}
