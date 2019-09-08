package healthchecks

import "dogego/modules"

func RegisterHealthChecks() {
	modules.HealthChecksModule.AddHealthCheck(DatabaseHealthCheck)
	modules.HealthChecksModule.AddHealthCheck(RedisHealthCheck)
}
