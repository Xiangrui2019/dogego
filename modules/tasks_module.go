package modules

import (
	"reflect"
	"runtime"
)

const (
	TimeJob = iota
	AsyncJob
)

type Task struct {
	Taskname string
	Type     int
	Time     string
	Job      interface{}
}

type TimeTask func() error
type AsyncTask func(data interface{}) error

var TasksModule []*Task

func AddTimedJob(time string, job TimeTask) {
	jobName := runtime.FuncForPC(reflect.ValueOf(job).Pointer()).Name()
	TasksModule = append(TasksModule, &Task{
		Taskname: jobName,
		Job:      job,
		Time:     time,
		Type:     TimeJob,
	})
}

func AddAyncJob(job AsyncTask) {
	jobName := runtime.FuncForPC(reflect.ValueOf(job).Pointer()).Name()
	TasksModule = append(TasksModule, &Task{
		Taskname: jobName,
		Job:      job,
		Type:     AsyncJob,
	})
}

func ClearTimedJob() {
	TasksModule = *new([]*Task)
}
