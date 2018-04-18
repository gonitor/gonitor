package cpu

import (
	"github.com/shirou/gopsutil/cpu"
)

// ServiceGetSummaryPercent gets the percentage summary of cpu usage.
func ServiceGetSummaryPercent() ([]float64, error) {
	return cpu.Percent(0, false)
}

// ServiceGetCount gets the CPU counts.
func ServiceGetCount() (int, error) {
	return cpu.Counts(true)
}

// ServiceGetSummaryTime gets the time summary.
func ServiceGetSummaryTime() ([]cpu.TimesStat, error) {
	return cpu.Times(false)
}

// ServiceGetInfo gets the CPU information.
func ServiceGetInfo() ([]cpu.InfoStat, error) {
	return cpu.Info()
}

// ServiceGetPercent gets CPU usage in details.
func ServiceGetPercent() ([]float64, error) {
	return cpu.Percent(0, true)
}

// ServiceGetTime gets CPU time information in details.
func ServiceGetTime() ([]cpu.TimesStat, error) {
	return cpu.Times(true)
}
