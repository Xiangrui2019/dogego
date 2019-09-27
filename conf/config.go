package conf

import (
	"dogego/cache"
	"dogego/executers"
	"dogego/healthchecks"
	"dogego/models"
	"dogego/modules"
	"dogego/protocol"
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
	healthchecks.RegisterHealthChecks()
	protocol.RegisterServers()
}
