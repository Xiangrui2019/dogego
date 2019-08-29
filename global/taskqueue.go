package global

import "os"

func TimeTaskQueueKey() string {
	return os.Getenv("TIME_TASK_QUEUE")
}

func AsyncTaskQueueKey() string {
	return os.Getenv("ASYNC_TASK_QUEUE")
}
