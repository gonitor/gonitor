package cpu

import (
	"github.com/shirou/gopsutil/cpu"
)

// ServiceGetSummaryPercent .
func ServiceGetSummaryPercent() ([]float64, error) {
	return cpu.Percent(0, false)
}

// ServiceGetCount .
func ServiceGetCount() (int, error) {
	return cpu.Counts(true)
}

// ServiceGetSummaryTime .
func ServiceGetSummaryTime() ([]cpu.TimesStat, error) {
	return cpu.Times(false)
}

// ServiceGetInfo .
func ServiceGetInfo() ([]cpu.InfoStat, error) {
	return cpu.Info()
}

// ServiceGetPercent .
func ServiceGetPercent() ([]float64, error) {
	return cpu.Percent(0, true)
}

// ServiceGetTime .
func ServiceGetTime() ([]cpu.TimesStat, error) {
	return cpu.Times(true)
}
