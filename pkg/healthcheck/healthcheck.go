package healthcheck

import (
	"sync"
	"sync/atomic"
)

var (
	healthCheck *HealthCheck
	once        sync.Once
)

type HealthCheck struct {
	status *atomic.Bool
}

func Get() *HealthCheck {
	once.Do(func() {
		status := &atomic.Bool{}
		status.Store(false)

		healthCheck = &HealthCheck{
			status: status,
		}
	})

	return healthCheck
}

func (hc *HealthCheck) MarkAsUp() {
	hc.status.Store(true)
}

func (hc *HealthCheck) MarkAsDown() {
	hc.status.Store(false)
}

func (hc *HealthCheck) IsAlive() bool {
	return hc.status.Load()
}
