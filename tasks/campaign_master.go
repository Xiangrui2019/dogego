package tasks

import "dogego/modules"

func CampaignMaster() {
	if modules.LockerModule.Lock("master") {
		StartCronJobs(true)
	}
}
