package qemu

import "github.com/mikerourke/queso"

// AccelType represents the hardware accelerator to use with QEMU.
type AccelType string

const (
	// AccelHAXM is the Intel Hardware Accelerated Execution Manager.
	// See https://github.com/intel/haxm for more details.
	AccelHAXM AccelType = "hax"

	// AccelHVF is the Hypervisor.framework accelerator on macOS.
	// See https://developer.apple.com/documentation/hypervisor for more details.
	AccelHVF AccelType = "hvf"

	// AccelKVM is the Kernel Virtual Machine, which is a Linux kernel module.
	// See https://wiki.qemu.org/Features/KVM for more details.
	AccelKVM AccelType = "kvm"

	// AccelNVMM is a Type-2 hypervisor and hypervisor platform for NetBSD.
	// See https://m00nbsd.net/4e0798b7f2620c965d0dd9d6a7a2f296.html for more details.
	AccelNVMM AccelType = "nvmm"

	// AccelTCG is the Tiny Code Generator, which is the core binary translation
	// engine for QEMU. See https://wiki.qemu.org/Features/TCG for more details.
	AccelTCG AccelType = "tcg"

	// AccelWHPX is the Windows Hypervisor Platform accelerator (Hyper-V).
	// See https://docs.microsoft.com/en-us/xamarin/android/get-started/installation/android-emulator/hardware-acceleration
	// for more details.
	AccelWHPX AccelType = "whpx"

	// AccelXen is a Type-1 (bare metal) hypervisor.
	// See https://wiki.xenproject.org/wiki/Xen_Project_Software_Overview for
	// more details.
	AccelXen AccelType = "xen"
)

// AccelProperty represents a property that can be used with the Accel option.
type AccelProperty struct {
	*queso.Property
}

func newAccelProperty(key string, value interface{}) *AccelProperty {
	return &AccelProperty{
		Property: &queso.Property{key, value},
	}
}

// Accel is used to enable an accelerator.
func Accel(accelType AccelType, properties ...*AccelProperty) *queso.Option {
	props := make([]*queso.Property, 0)

	for _, property := range properties {
		props = append(props, property.Property)
	}

	return queso.NewOption("accel", string(accelType), props...)
}

// IsIGDPassThru controls whether Intel integrated graphics devices can be passed
// through to the guest.
//
// This property can only be used with the XEN accelerator.
func IsIGDPassThru(enabled bool) *AccelProperty {
	return newAccelProperty("igd-passthru", enabled)
}

// KernelIRQChipFlag represents the option to use for the WithKernelIRQChip
// property.
type KernelIRQChipFlag string

const (
	// KernelIRQChipOn indicates that full acceleration of the interrupt controllers
	// should be used.
	KernelIRQChipOn KernelIRQChipFlag = "on"

	// KernelIRQChipOff disables full acceleration and should only be used for debugging
	// purposes.
	KernelIRQChipOff KernelIRQChipFlag = "off"

	// KernelIRQChipSplit reduces the kernel attack surface, at a performance cost
	// for non-MSI interrupts.
	KernelIRQChipSplit KernelIRQChipFlag = "split"
)

// WithKernelIRQChip controls KVM in-kernel IRQ chip support. The default is
// KernelIRQChipOn. See KernelIRQChipFlag for more details.
func WithKernelIRQChip(flag KernelIRQChipFlag) *AccelProperty {
	return newAccelProperty("kernel-irqchip", flag)
}

// WithKVMShadowMemory defines the size of the KVM shadow MMU.
func WithKVMShadowMemory(size int) *AccelProperty {
	return newAccelProperty("kvm-shadow-mem", size)
}

// IsSplitWX controls the use of split w^x mapping for the TCG code generation
// buffer. Some operating systems require this to be enabled, and in such a case
// this will default on. On other operating systems, this will default off, but
// one may enable this for testing or debugging.
func IsSplitWX(enabled bool) *AccelProperty {
	return newAccelProperty("split-wx", enabled)
}

// WithTranslationBlockCacheSize controls the size (in MiB) of the TCG translation
// block cache.
func WithTranslationBlockCacheSize(size int) *AccelProperty {
	return newAccelProperty("tb-size", size)
}

// ThreadFlag represents the type of TCG threads to use.
type ThreadFlag string

const (
	// ThreadSingle indicates that a single thread should be used with TCG.
	ThreadSingle ThreadFlag = "single"

	// ThreadMulti indicates that multiple threads should be used with TCG.
	ThreadMulti ThreadFlag = "multi"
)

// WithThread controls number of TCG threads. When the TCG is multi-threaded, there
// will be one thread per vCPU therefore taking advantage of additional host cores.
// The default is to enable multi-threading where both the back-end and front-ends
// support it and no incompatible TCG features have been enabled (e.g. icount/replay).
func WithThread(flag ThreadFlag) *AccelProperty {
	return newAccelProperty("thread", flag)
}

// WithDirtyRingSize controls the size of the per-vCPU dirty page ring buffer
// (number of entries for each vCPU). It should be a value that is power of
// two, and it should be 1024 or bigger (but still less than the maximum value that
// the kernel supports). 4096 could be a good initial value if you have no idea
// which is the best. Set this value to 0 to disable the feature. By default, this
// feature is disabled (value = 0). When enabled, KVM will instead record
// dirty pages in a bitmap.
func WithDirtyRingSize(size int) *AccelProperty {
	return newAccelProperty("dirty-ring-size", size)
}
