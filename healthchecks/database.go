package healthchecks

import "dogego/models"

func DatabaseHealthCheck() error {
	return models.DB.DB().Ping()
}
