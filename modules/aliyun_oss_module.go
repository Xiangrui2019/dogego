package modules

import (
	"log"
	"os"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

var AliyunOSSModule *oss.Client
var AliyunOSSBucketModule *oss.Bucket

func InitAliyunOSSModule() {
	var err error

	AliyunOSSModule, err = oss.New(os.Getenv("OSS_END_POINT"), os.Getenv("OSS_ACCESS_KEY_ID"), os.Getenv("OSS_ACCESS_KEY_SECRET"))

	if err != nil {
		log.Fatal(err)
	}

	AliyunOSSBucketModule, err = AliyunOSSModule.Bucket(os.Getenv("OSS_BUCKET"))

	if err != nil {
		log.Fatal(err)
	}
}
