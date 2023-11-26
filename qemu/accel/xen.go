package accel

// XenAccelerator represents an accelerator using Xen.
type XenAccelerator struct {
	*Accelerator
}

// NewXenAccelerator returns a new instance of [XenAccelerator].
//
//	qemu-system-* -accel xen
func NewXenAccelerator() *XenAccelerator {
	return &XenAccelerator{
		New(string(TypeXen)),
	}
}

// ToggleIGDPassThru controls whether Intel integrated graphics devices can be
// passed through to the guest.
//
//	qemu-system-* -accel xen,igd-passthru=on|off
func (a *XenAccelerator) ToggleIGDPassThru(enabled bool) *XenAccelerator {
	a.SetProperty("igd-passthru", enabled)
	return a
}
