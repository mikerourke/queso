package debug

import "github.com/mikerourke/queso"

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
