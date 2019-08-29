package conf

import (
	"dogego/cache"
	"dogego/executers"
	"dogego/models"
	"dogego/modules"
	"dogego/tasks"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()

	models.ConnectDatabase(os.Getenv("DATABASE_DSN"))
	cache.ConnectRedisCache()
	modules.InitAllModules()
	tasks.StartCronJobs(false)
	executers.TimeExecuter()
	executers.AsyncExecuter()
}
