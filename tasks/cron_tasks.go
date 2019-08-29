package tasks

import (
	"dogego/modules"
)

func RegisterCronTasks() {
	modules.ClearTimedJob()
	modules.AddTimedJob("@every 1m", TimeTask)
	modules.AddTimedJob("@every 2m", DeltaTask)
	modules.AddAyncJob(TimeTask1)
}
