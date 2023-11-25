package accel

// HVFAccelerator represents an accelerator using the Hypervisor.framework
// accelerator on macOS. See https://developer.apple.com/documentation/hypervisor
// for more details.
type HVFAccelerator struct {
	*Accelerator
}

// NewHVFAccelerator returns a new instace of [HVFAccelerator].
//
//	qemu-system-* -accel hvf
func NewHVFAccelerator() *HVFAccelerator {
	return &HVFAccelerator{
		New(string(TypeHVF)),
	}
}
