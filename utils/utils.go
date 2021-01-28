package utils

import (
	"fmt"
)

// FormatBytes - Convert bytes to human readable string
func FormatBytes(i int64) (result string) {
	const (
		KiB = 1024
		MiB = 1048576
		GiB = 1073741824
		TiB = 1099511627776
	)
	switch {
	case i >= TiB:
		result = fmt.Sprintf("%.02fTiB", float64(i)/TiB)
	case i >= GiB:
		result = fmt.Sprintf("%.02fGiB", float64(i)/GiB)
	case i >= MiB:
		result = fmt.Sprintf("%.02fMiB", float64(i)/MiB)
	case i >= KiB:
		result = fmt.Sprintf("%.02fKiB", float64(i)/KiB)
	default:
		result = fmt.Sprintf("%dB", i)
	}
	return
}
