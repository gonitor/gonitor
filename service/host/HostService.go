package host

import (
	"github.com/shirou/gopsutil/host"
)

// ServiceGetInfo gets the host information.
func ServiceGetInfo() (*host.InfoStat, error) {
	return host.Info()
}

// ServiceGetTemperature gets the host temperature.
func ServiceGetTemperature() ([]host.TemperatureStat, error) {
	return host.SensorsTemperatures()
}
