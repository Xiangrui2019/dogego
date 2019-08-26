package utils

import (
	"log"
	"reflect"
	"runtime"
	"time"

	"github.com/gin-gonic/gin"
)

func RunTask(context *gin.Context, job func() error) {
	from := time.Now().UnixNano()
	err := job()
	to := time.Now().UnixNano()
	jobName := runtime.FuncForPC(reflect.ValueOf(job).Pointer()).Name()
	if err != nil {
		log.Printf("%s 执行失败: %dms\n", jobName, (to-from)/int64(time.Millisecond))
	} else {
		log.Printf("%s 执行成功: %dms\n", jobName, (to-from)/int64(time.Millisecond))
	}
}
