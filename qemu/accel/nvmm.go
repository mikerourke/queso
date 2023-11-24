package accel

// NVMMAccelerator represents a Type-2 hypervisor and hypervisor platform for
// NetBSD. See https://m00nbsd.net/4e0798b7f2620c965d0dd9d6a7a2f296.html for more
// details.
type NVMMAccelerator struct {
	*Accelerator
}

// NewNVMMAccelerator returns a new instace of [NVMMAccelerator].
//
//	qemu-system-* -accel nvmm
func NewNVMMAccelerator() *NVMMAccelerator {
	return &NVMMAccelerator{
		New(TypeNVMM),
	}
}
