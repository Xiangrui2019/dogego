package conf

import (
	"dogego/cache"
	"dogego/models"
	"dogego/tasks"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()

	models.ConnectDatabase(os.Getenv("DATABASE_DSN"))
	cache.ConnectRedisCache()
	tasks.StartCronJobs()
}
