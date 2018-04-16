package cpu

import (
	"github.com/shirou/gopsutil/cpu"
)

// GetSummaryPercent .
func GetSummaryPercent() ([]float64, error) {
	return cpu.Percent(0, false)
}

// GetCount .
func GetCount() (int, error) {
	return cpu.Counts(true)
}

// GetSummaryTime .
func GetSummaryTime() ([]cpu.TimesStat, error) {
	return cpu.Times(false)
}

// GetInfo .
func GetInfo() ([]cpu.InfoStat, error) {
	return cpu.Info()
}

// GetPercent .
func GetPercent() ([]float64, error) {
	return cpu.Percent(0, true)
}

// GetTime .
func GetTime() ([]cpu.TimesStat, error) {
	return cpu.Times(true)
}
