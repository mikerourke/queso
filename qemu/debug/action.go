package debug

import "github.com/mikerourke/queso"

// Action modifies QEMU's default behavior for certain events. It provides a generic
// method for specifying the same behaviors that are modified by the NoReboot and
// NoShutdown options.
type Action struct {
	Event  string
	Action string
}

// NewAction returns a new instance of Action. These are passed into the qemu.Use
// method.
//
// # Examples
//
//	qemu.Use(
//		qemu.NewAction("panic", "none"),
//		qemu.NewAction("reboot", "shutdown"),
//		qemu.NewAction("shutdown", "pause"),
//		qemu.NewAction("watchdog", "pause"))
func NewAction(event string, action string) *Action {
	return &Action{
		Event:  event,
		Action: action,
	}
}

// Option returns the invoked option that gets converted to an argument when
// passed to QEMU.
func (a *Action) Option() *queso.Option {
	return queso.NewOption("action", "", queso.NewProperty(a.Event, a.Action))
}
