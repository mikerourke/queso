package accel

// WHPXAccelerator represents an accelerator using the Windows Hypervisor Platform
// accelerator (Hyper-V). See https://docs.microsoft.com/en-us/xamarin/android/get-started/installation/android-emulator/hardware-acceleration
// for more details.
type WHPXAccelerator struct {
	*Accelerator
}

// NewWHPXAccelerator returns a new instace of [WHPXAccelerator].
//
//	qemu-system-* -accel whpx
func NewWHPXAccelerator() *WHPXAccelerator {
	return &WHPXAccelerator{
		New(string(TypeWHPX)),
	}
}
