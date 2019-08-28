package executers

import (
	"dogego/global"
	"dogego/modules"
	"dogego/utils"
	"log"
	"time"

	"github.com/streadway/amqp"
)

func TaskListenExecuter(ch *amqp.Channel, queue *amqp.Queue) error {
	msgs, err := ch.Consume(
		queue.Name,
		"",
		true,
		false,
		false,
		false,
		amqp.Table{},
	)

	if err != nil {
		return err
	}

	go func() {
		for d := range msgs {
			executeTask(&d)
		}
	}()

	return nil
}

func executeTask(d *amqp.Delivery) {
	for _, item := range modules.TasksModule {
		if item.Taskname == string(d.Body) {
			if !modules.LockerModule.Lock(item.Taskname, 0) {
				return
			}

			if item.Type != modules.TimeJob {
				return
			}

			job := item.Job.(modules.TimeTask)

			from := time.Now().UnixNano()
			err := job()
			to := time.Now().UnixNano()

			if err != nil {
				log.Printf("%s Execute Error: %dms\n", item.Taskname, (to-from)/int64(time.Millisecond))
			} else {
				log.Printf("%s Execute Success: %dms\n", item.Taskname, (to-from)/int64(time.Millisecond))
			}

			modules.LockerModule.Free(item.Taskname)
		}
	}
}

func TaskExcuter() error {
	ch, queue, err := utils.BuildQueueChannel(global.TimeTaskQueueKey())

	if err != nil {
		return err
	}

	return TaskListenExecuter(ch, queue)
}
