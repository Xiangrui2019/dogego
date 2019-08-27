package global

import "fmt"

func LockKey(lockname string) string {
	return fmt.Sprintf("lock:%s", lockname)
}
