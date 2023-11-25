package accel

// HAXMAccelerator represents an accelerator using the Intel Hardware Accelerated
// Execution Manager. See https://github.com/intel/haxm for more details.
//
// Deprecated: No longer supported, but kept for older versions of QEMU.
type HAXMAccelerator struct {
	*Accelerator
}

// NewHAXMAccelerator returns a new instace of [HAXMAccelerator].
//
//	qemu-system-* -accel haxm
func NewHAXMAccelerator() *HAXMAccelerator {
	return &HAXMAccelerator{
		New(string(TypeHAXM)),
	}
}
