package utils

import (
	"dogego/cache"
	"log"
	"reflect"
	"runtime"
	"time"
)

func RunTask(job func() error) {
	jobName := runtime.FuncForPC(reflect.ValueOf(job).Pointer()).Name()

	err := cache.CacheClient.SetNX(jobName, "true", 0).Err()
	if err != nil {
		return
	}
	from := time.Now().UnixNano()
	err = job()
	to := time.Now().UnixNano()
	if err != nil {
		log.Printf("%s Execute Error: %dms\n", jobName, (to-from)/int64(time.Millisecond))
	} else {
		log.Printf("%s Execute Success: %dms\n", jobName, (to-from)/int64(time.Millisecond))
	}
}
