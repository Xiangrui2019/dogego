package tasks

import (
	"dogego/modules"
	"log"
	"time"
)

func CampaignMaster() {
	if modules.LockerModule.Lock("master", time.Minute*2) {
		StartCronJobs(true)
		log.Println("Campaign master Success.")
	}
}
