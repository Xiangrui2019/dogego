package tasks

import "dogego/modules"

func RegisterCronTasks() {
	modules.ClearTimedJob()
	modules.AddTimedJob("@every 1m", TimeTask)
	modules.AddAyncJob(TimeTask1)
}
