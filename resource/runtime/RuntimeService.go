package runtime

import (
	"runtime"
)

// GetGOOS .
func GetGOOS() string {
	return runtime.GOOS
}
