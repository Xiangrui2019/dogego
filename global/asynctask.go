package global

import (
	"dogego/modules"
	"encoding/json"
	"fmt"
)

func AsyncTaskData(taskname string, data modules.TaskData) string {
	v, _ := json.Marshal(data)
	return fmt.Sprintf("%s#$#%s", taskname, v)
}
