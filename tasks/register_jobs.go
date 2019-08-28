package tasks

import "dogego/modules"

func RegisterJobs() {
	modules.ClearTimedJob()
	modules.AddTimedJob("@every 3m", TimeTask)
	modules.AddAyncJob(TimeTask1)
}
