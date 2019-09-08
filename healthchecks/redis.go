package healthchecks

import "dogego/cache"

func RedisHealthCheck() error {
	return cache.CacheClient.Ping().Err()
}
