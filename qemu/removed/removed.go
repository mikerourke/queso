// Package removed contains options that have been removed in newer versions
// of QEMU.
package removed

import "github.com/mikerourke/queso"

// EnableFIPS enables FIPS 140-2 compliance mode.
// Removed in v7.0.
func EnableFIPS() *queso.Option {
	return queso.NewOption("enable-fips", "")
}
