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

	Cron.AddFunc("* 0 0 0 0", func() { utils.RunTask(TimeTask) })
	Cron.Start()

	fmt.Println("Cron Jobs started success.")
}
