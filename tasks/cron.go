package tasks

import (
	"dogego/modules"
	"dogego/utils"
	"log"
	"time"

	"github.com/robfig/cron"
)

var Cron *cron.Cron

func StartCronJobs(locked bool) {
	Cron = cron.New()

	RegisterCronTasks()

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
				utils.PublishTask(item)
			})
		}
	}

	Cron.Start()

	log.Println("Cron Jobs started success.")
}
