package global

import "fmt"

func QueueNameKey(queuename string) string {
	return fmt.Sprintf("queue:%s", queuename)
}
