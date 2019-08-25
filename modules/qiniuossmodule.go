package modules

import (
	"os"

	"github.com/qiniu/api.v7/auth/qbox"
	"github.com/qiniu/api.v7/storage"
)

var QiniuOSSModule *qbox.Mac
var QiniuOSSBucketModule *storage.PutPolicy

func InitQiniuOSSModule() {
	QiniuOSSModule = qbox.NewMac(os.Getenv("QINIU_AK"), os.Getenv("QINIU_SK"))
	QiniuOSSBucketModule = &storage.PutPolicy{
		Scope: os.Getenv("QINIU_BUCKET"),
	}
}
