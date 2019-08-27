package modules

type Task struct {
	Taskname string
	Job      func() error
}

var TasksModule []*Task
