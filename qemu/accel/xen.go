package accel

import "github.com/mikerourke/queso/internal/cli"

// XenAccelerator represents an accelerator using Xen.
type XenAccelerator struct {
	*Accelerator
}

// NewXenAccelerator returns a new instace of XenAccelerator.
//
//	qemu-system-* -accel xen
func NewXenAccelerator() *XenAccelerator {
	return &XenAccelerator{
		NewAccelerator(Xen),
	}
}

// ToggleIGDPassThru controls whether Intel integrated graphics devices can be passed
// through to the guest.
//
//	qemu-system-* -accel xen igd-passthru=on|off
func (a *XenAccelerator) ToggleIGDPassThru(enabled bool) *XenAccelerator {
	a.properties = append(a.properties, cli.NewProperty("igd-passthru", enabled))
	return a
}
