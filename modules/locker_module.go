package modules

import (
	"dogego/cache"
	"dogego/global"
	"time"
)

type Locker struct {
}

var LockerModule *Locker

func (locker *Locker) Lock(lockname string, locktime time.Duration) bool {
	result, err := cache.CacheClient.SetNX(
		global.LockKey(lockname), "true", locktime).Result()

	if err != nil {
		return false
	}

	return result
}

func (locker *Locker) Free(lockname string) error {
	return cache.CacheClient.Del(global.LockKey(lockname)).Err()
}

func InitLockerModule() {
	LockerModule = new(Locker)
}
