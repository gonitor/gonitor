package host

import (
	"github.com/shirou/gopsutil/host"
)

// ServiceGetInfo .
func ServiceGetInfo() (*host.InfoStat, error) {
	return host.Info()
}

// ServiceGetTemperature .
func ServiceGetTemperature() ([]host.TemperatureStat, error) {
	return host.SensorsTemperatures()
}
