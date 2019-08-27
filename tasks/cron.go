package tasks

import (
	"dogego/utils"
	"fmt"

	"github.com/robfig/cron"
)

var Cron *cron.Cron

func StartCronJobs() {
	if Cron == nil {
		Cron = cron.New()
	}

	Cron.AddFunc("@every 10s", func() { utils.RunTask(TimeTask) })
	Cron.Start()

	fmt.Println("Cron Jobs started success.")
}
