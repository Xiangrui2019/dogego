package tasks

import (
	"log"
	"reflect"
	"runtime"
	"time"

	"github.com/robfig/cron"
)

var Cron *cron.Cron

func Run(job func() error) {
	from := time.Now().UnixNano()
	err := job()
	to := time.Now().UnixNano()
	jobName := runtime.FuncForPC(reflect.ValueOf(job).Pointer()).Name()
	if err != nil {
		log.Printf("%s error: %dms\n", jobName, (to-from)/int64(time.Millisecond))
	} else {
		log.Printf("%s success: %dms\n", jobName, (to-from)/int64(time.Millisecond))
	}
}

func CronJob() {
	if Cron == nil {
		Cron = cron.New()
	}

	Cron.Start()
	log.Println("Cron job started.")
}
