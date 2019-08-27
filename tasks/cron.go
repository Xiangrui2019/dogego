package tasks

import (
	"dogego/modules"
	"fmt"

	"github.com/robfig/cron"
)

var Cron *cron.Cron

func StartCronJobs(locked bool) {
	Cron = cron.New()

	if !locked {
		if !modules.LockerModule.Lock("master") {
			Cron.AddFunc("@every 2m", CampaignMaster)
			Cron.Start()
			return
		}
	}

	for _, item := range modules.TasksModule {
		Cron.AddFunc(item.Time, func() {})
	}

	fmt.Println("Cron Jobs started success.")
}
