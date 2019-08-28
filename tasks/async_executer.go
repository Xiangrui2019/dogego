package tasks

import (
	"dogego/global"
	"dogego/modules"
	"dogego/utils"
	"encoding/json"
	"log"
	"strings"
	"time"

	"github.com/streadway/amqp"
)

func AsyncListenExecuter(ch *amqp.Channel, queue *amqp.Queue) error {
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
			executeAsyncTask(&d)
		}
	}()

	return nil
}

func executeAsyncTask(d *amqp.Delivery) {
	for _, item := range modules.TasksModule {
		l := strings.Split(string(d.Body), "#$#")

		if item.Taskname == l[0] {
			var data interface{}

			if !modules.LockerModule.Lock(item.Taskname, 0) {
				return
			}

			job := item.Job.(modules.AsyncTask)
			err := json.Unmarshal([]byte(l[1]), &data)

			if err != nil {
				return
			}

			from := time.Now().UnixNano()
			err = job(data)
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

func AsyncExecuter() error {
	ch, queue, err := utils.BuildQueueChannel(global.AsyncTaskQueueKey())

	if err != nil {
		return err
	}

	return AsyncListenExecuter(ch, queue)
}
