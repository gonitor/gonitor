package host

import (
	"github.com/shirou/gopsutil/host"
)

// GetInfo .
func GetInfo() (*host.InfoStat, error) {
	return host.Info()
}

// GetTemperature .
func GetTemperature() ([]host.TemperatureStat, error) {
	return host.SensorsTemperatures()
}
