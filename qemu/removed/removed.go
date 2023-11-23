// Package removed contains options that have been removed in newer versions
// of QEMU.
package removed

import (
	"github.com/mikerourke/queso/internal/cli"
)

// EnableFIPS enables FIPS 140-2 compliance mode.
//
// This option restricted usage of certain cryptographic algorithms when the host is operating in
// FIPS mode. If FIPS compliance is required, QEMU should be built with the libgcrypt or
// gnutls library enabled as a cryptography provider. Neither the nettle library, nor the
// built-in cryptography provider are supported on FIPS enabled hosts.
//
// Removed in v7.1.
func EnableFIPS() *cli.Option {
	return cli.NewOption("enable-fips", "")
}

// WatchdogModel represents the model of hardware watchdog to emulate.
//
// Removed in v7.2. Use `-device` instead.
type WatchdogModel string

const (
	// WatchdogIBase700 represents the iBASE 700, which is a very simple ISA
	// watchdog with a single timer.
	WatchdogIBase700 WatchdogModel = "ib700"

	// WatchdogIntel6300ESB represents the Intel 6300ESB I/O controller hub is a
	// much more full-featured PCI-based dual-timer watchdog.
	WatchdogIntel6300ESB WatchdogModel = "i6300esb"

	// WatchdogDiagnose288 represents a virtual watchdog for s390x backed by the
	// Diagnose 288 Hypercall (currently KVM only).
	WatchdogDiagnose288 WatchdogModel = "diag288"
)

// Watchdog creates a virtual hardware watchdog device. Once enabled (by a guest action),
// the watchdog must be periodically polled by an agent inside the guest or else
// the guest will be restarted. Choose a model for which your guest has drivers.
// Only one watchdog can be enabled for a guest.
//
// Removed in v7.2. Use `-device` instead.
func Watchdog(model WatchdogModel) *cli.Option {
	return cli.NewOption("watchdog", string(model))
}
