package qemu

import "github.com/mikerourke/queso"

const (
	// AccelHAXM is the Intel Hardware Accelerated Execution Manager.
	// See https://github.com/intel/haxm for more details.
	AccelHAXM = "hax"

	// AccelHVF is the Hypervisor.framework accelerator on macOS.
	// See https://developer.apple.com/documentation/hypervisor for more details.
	AccelHVF = "hvf"

	// AccelKVM is the Kernel Virtual Machine, which is a Linux kernel module.
	// See https://wiki.qemu.org/Features/KVM for more details.
	AccelKVM = "kvm"

	// AccelNVMM is a Type-2 hypervisor and hypervisor platform for NetBSD.
	// See https://m00nbsd.net/4e0798b7f2620c965d0dd9d6a7a2f296.html for more details.
	AccelNVMM = "nvmm"

	// AccelTCG is the Tiny Code Generator, which is the core binary translation
	// engine for QEMU. See https://wiki.qemu.org/Features/TCG for more details.
	AccelTCG = "tcg"

	// AccelWHPX is the Windows Hypervisor Platform accelerator (Hyper-V).
	// See https://docs.microsoft.com/en-us/xamarin/android/get-started/installation/android-emulator/hardware-acceleration
	// for more details.
	AccelWHPX = "whpx"

	// AccelXen is a Type-1 (bare metal) hypervisor.
	// See https://wiki.xenproject.org/wiki/Xen_Project_Software_Overview for
	// more details.
	AccelXen = "xen"
)

// KernelIRQChipMode represents the mode to use for the SetKernelIRQChip
// property for Accel.
type KernelIRQChipMode string

const (
	// KernelIRQChipOn indicates that full acceleration of the interrupt
	// controllers should be used.
	KernelIRQChipOn KernelIRQChipMode = "on"

	// KernelIRQChipOff disables full acceleration and should only be used for
	// debugging purposes.
	KernelIRQChipOff KernelIRQChipMode = "off"

	// KernelIRQChipSplit reduces the kernel attack surface, at a performance
	// cost for non-MSI interrupts.
	KernelIRQChipSplit KernelIRQChipMode = "split"
)

// ThreadFlag represents the type of TCG threads to use for Accel.
type ThreadFlag string

const (
	// ThreadSingle indicates that a single thread should be used with TCG.
	ThreadSingle ThreadFlag = "single"

	// ThreadMulti indicates that multiple threads should be used with TCG.
	ThreadMulti ThreadFlag = "multi"
)

func AccelType(accelType string) *queso.Option {
	return queso.NewOption("accel", accelType)
}

type Accel struct {
	Usable
	Type       string
	properties []*queso.Property
}

func NewAccel(accelType string) *Accel {
	return &Accel{
		Type: accelType,
	}
}

func (a *Accel) option() *queso.Option {
	return queso.NewOption("accel", "", a.properties...)
}

// ToggleIGDPassThru controls whether Intel integrated graphics devices can be passed
// through to the guest for Accel.
//
// This property can only be used with the Xen accelerator.
func (a *Accel) ToggleIGDPassThru(enabled bool) *Accel {
	a.properties = append(a.properties, queso.NewProperty("igd-passthru", enabled))
	return a
}

// SetKernelIRQChip controls KVM in-kernel IRQ chip support for Accel. The
// default is KernelIRQChipOn. See KernelIRQChipMode for more details.
func (a *Accel) SetKernelIRQChip(mode KernelIRQChipMode) *Accel {
	a.properties = append(a.properties, queso.NewProperty("kernel-irqchip", mode))
	return a
}

// SetKVMShadowMemory defines the size of the KVM shadow MMU.
func (a *Accel) SetKVMShadowMemory(size int) *Accel {
	a.properties = append(a.properties, queso.NewProperty("kvm-shadow-mem", size))
	return a
}

// ToggleSplitWX controls the use of split w^x mapping for the TCG code generation
// buffer for Accel. Some operating systems require this to be enabled, and in
// such a case this will default on. On other operating systems, this will default
// off, but one may enable this for testing or debugging.
func (a *Accel) ToggleSplitWX(enabled bool) *Accel {
	a.properties = append(a.properties, queso.NewProperty("split-wx", enabled))
	return a
}

// SetTranslationBlockCacheSize controls the size (in MiB) of the TCG
// translation block cache for Accel.
func (a *Accel) SetTranslationBlockCacheSize(megabytes int) *Accel {
	a.properties = append(a.properties, queso.NewProperty("tb-size", megabytes))
	return a
}

// SetThreadFlag controls number of TCG threads for Accel. When the TCG is
// multithreaded, there will be one thread per vCPU therefore taking advantage
// of additional host cores. The default is to enable multi-threading where both
// the back-end and front-ends support it and no incompatible TCG features have
// been enabled (e.g. icount/replay).
func (a *Accel) SetThreadFlag(flag ThreadFlag) *Accel {
	a.properties = append(a.properties, queso.NewProperty("thread", flag))
	return a
}

// SetDirtyRingSize controls the size of the per-vCPU dirty page ring buffer
// (number of entries for each vCPU) for Accel. It should be a value that is power
// of two, and it should be 1024 or bigger (but still less than the maximum value
// that the kernel supports). 4096 could be a good initial value if you have no
// idea which is the best. Set this value to 0 to disable the feature. By default,
// this feature is disabled (value = 0). When enabled, KVM will instead record
// dirty pages in a bitmap.
func (a *Accel) SetDirtyRingSize(bytes int) *Accel {
	a.properties = append(a.properties, queso.NewProperty("dirty-ring-size", bytes))
	return a
}
