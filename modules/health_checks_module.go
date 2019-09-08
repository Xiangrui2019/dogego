package modules

type HealthCheckFunc func() error
type HealthList []HealthCheckFunc

type HealthChecks struct {
}

var HealthCheckJobList HealthList

var HealthChecksModule *HealthChecks

func (hc *HealthChecks) AddHealthCheck(checkjob HealthCheckFunc) {
	HealthCheckJobList = append(HealthCheckJobList, checkjob)
}

func (hc *HealthChecks) Check() []string {
	errlist := make([]string, 0)

	for _, job := range HealthCheckJobList {
		err := job()

		if err != nil {
			errlist = append(errlist, err.Error())
		}
	}

	return errlist
}

func InitHealthChecksModule() {
	HealthCheckJobList = *new(HealthList)
	HealthChecksModule = new(HealthChecks)
}
