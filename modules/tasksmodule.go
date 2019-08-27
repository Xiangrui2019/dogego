package modules

import (
	"reflect"
	"runtime"
)

type Task struct {
	Taskname string
	Time     string
	Job      func() error
}

var TasksModule []*Task

func AddJob(time string, job func() error) {
	jobName := runtime.FuncForPC(reflect.ValueOf(job).Pointer()).Name()
	TasksModule = append(TasksModule, &Task{
		Taskname: jobName,
		Job:      job,
		Time:     time,
	})
}

func InitTaskModule() {

}
