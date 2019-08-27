package tasks

import (
	"dogego/utils"

	"github.com/robfig/cron"
)

var Cron *cron.Cron

func StartCronJobs() {
	if Cron == nil {
		Cron = cron.New()
	}

	Cron.AddFunc("* 0 0 0 0", func() { utils.RunTask(TimeTask) })
}
