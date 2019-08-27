package tasks

import "dogego/modules"

func RegisterJobs() {
	modules.AddTimedJob("@every 10s", TimeTask)
}
