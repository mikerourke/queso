package qemu

import "github.com/mikerourke/queso"

// WatchdogModel represents the model of hardware watchdog to emulate.
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
func Watchdog(model WatchdogModel) *queso.Option {
	return queso.NewOption("watchdog", string(model))
}

// WatchdogAction represents the action to perform when the watchdog timer
// expires.
type WatchdogAction string

const (
	// WatchdogActionReset forcefully resets the guest.
	WatchdogActionReset WatchdogAction = "reset"

	// WatchdogActionShutdown attempts to gracefully shut down the guest.
	// Note that this action requires that the guest responds to ACPI signals,
	// which it may not be able to do in the sort of situations where the
	// watchdog would have expired, and thus this action is not recommended
	// for production use.
	WatchdogActionShutdown WatchdogAction = "shutdown"

	// WatchdogActionPowerOff forcefully powers off the guest.
	WatchdogActionPowerOff WatchdogAction = "poweroff"

	// WatchdogActionInjectNMI injects a NMI into the guest.
	WatchdogActionInjectNMI WatchdogAction = "inject-nmi"

	// WatchdogActionPause pauses the guest.
	WatchdogActionPause WatchdogAction = "pause"

	// WatchdogActionDebug prints a debug message and continues.
	WatchdogActionDebug WatchdogAction = "debug"

	// WatchdogActionNone does nothing.
	WatchdogActionNone WatchdogAction = "none"
)

// WatchdogActionOnExpiration specifies the action to perform when the watchdog
// timer expires.
func WatchdogActionOnExpiration(action WatchdogAction) *queso.Option {
	return queso.NewOption("watchdog-action", string(action))
}
