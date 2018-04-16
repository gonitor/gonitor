package load

import (
	"github.com/shirou/gopsutil/load"
)

// GetAverage .
func GetAverage() (*load.AvgStat, error) {
	return load.Avg()
}

// GetMisc .
func GetMisc() (*load.MiscStat, error) {
	return load.Misc()
}
