package global

import (
	"encoding/json"
	"fmt"
)

func AsyncTaskKey(taskname string, data interface{}) string {
	v, _ := json.Marshal(data)
	return fmt.Sprintf("%s/%s", taskname, v)
}
