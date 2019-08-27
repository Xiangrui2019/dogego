package modules

import (
	"dogego/cache"
	"dogego/global"
)

type Locker struct {
}

func (locker *Locker) Lock(lockname string) bool {
	result, err := cache.CacheClient.SetNX(
		global.LockKey(lockname), "true", 0).Result()

	if err != nil {
		return false
	}

	return result
}

func (locker *Locker) Free(lockname string) error {
	return cache.CacheClient.Del(lockname).Err()
}
