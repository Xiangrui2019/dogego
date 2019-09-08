package modules

type HealthCheckFunc func() error

var HealthChecksModule []*HealthCheckFunc

func AddHealthCheck(checkjob HealthCheckFunc) {
	HealthChecksModule = append(HealthChecksModule, &checkjob)
}

func InitHealthChecksModule() {
	HealthChecksModule = *new([]*HealthCheckFunc)
}
