package tasks

import (
	"dogego/modules"
	"log"
	"time"

	"github.com/robfig/cron"
)

var Cron *cron.Cron

func StartCronJobs(locked bool) {
	Cron = cron.New()

	RegisterTasks()

	if !locked {
		if !modules.LockerModule.Lock("master", time.Minute*2) {
			Cron.AddFunc("@every 2m", CampaignMasterTask)
			Cron.Start()
			return
		}
	}

	Cron.AddFunc("@every 1m", ClifeMasterTask)

	for _, item := range modules.TasksModule {
		if item.Type == modules.TimeJob {
			Cron.AddFunc(item.Time, func() {
				PublishTask(item)
			})
		}
	}

	Cron.Start()

	log.Println("Cron Jobs started success.")
}
