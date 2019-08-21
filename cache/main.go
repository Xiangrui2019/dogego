package cache

import (
	"log"
	"os"

	"github.com/go-redis/redis"
)

var CacheClient *redis.Client

func ConnectRedisCache() {
	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: os.Getenv("REDIS_PASSWORD"),
	})

	_, err := client.Ping().Result()

	if err != nil {
		log.Fatal(err)
	}

	CacheClient = client
}
