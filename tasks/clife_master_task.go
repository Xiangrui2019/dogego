package tasks

import (
	"dogego/cache"
	"dogego/global"
	"log"
	"time"
)

func ClifeMasterTask() {
	cache.CacheClient.Set(global.LockKey("master"), "true", time.Minute*2)
	log.Println("Continued life Success.")
}
