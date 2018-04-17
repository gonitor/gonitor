package load

import (
	"github.com/shirou/gopsutil/load"
)

// ServiceGetAverage .
func ServiceGetAverage() (*load.AvgStat, error) {
	return load.Avg()
}

// GetMisc .
func GetMisc() (*load.MiscStat, error) {
	return load.Misc()
}
