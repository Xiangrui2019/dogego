package tasks

import "dogego/modules"

func RegisterTasks() {
	modules.ClearTimedJob()
	modules.AddTimedJob("@every 3m", TimeTask)
	modules.AddAyncJob(TimeTask1)
}
