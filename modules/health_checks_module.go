package modules

type HealthCheckFunc func() error

var HealthChecksModule []*HealthCheckFunc

func InitHealthChecksModule() {
	HealthChecksModule = *new([]*HealthCheckFunc)
}
