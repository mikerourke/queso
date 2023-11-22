package qemu

import (
	"fmt"

	"github.com/mikerourke/queso"
)

type AcceleratorType string

const (
	// HAXM is the Intel Hardware Accelerated Execution Manager.
	// See https://github.com/intel/haxm for more details.
	// Deprecated: No longer supported, but kept for older versions of QEMU.
	HAXM AcceleratorType = "hax"

	// HVF is the Hypervisor.framework accelerator on macOS.
	// See https://developer.apple.com/documentation/hypervisor for more details.
	HVF AcceleratorType = "hvf"

	// KVM is the Kernel Virtual Machine, which is a Linux kernel module.
	// See https://wiki.qemu.org/Features/KVM for more details.
	KVM AcceleratorType = "kvm"

	// NVMM is a Type-2 hypervisor and hypervisor platform for NetBSD.
	// See https://m00nbsd.net/4e0798b7f2620c965d0dd9d6a7a2f296.html for more details.
	NVMM AcceleratorType = "nvmm"

	// TCG is the Tiny Code Generator, which is the core binary translation
	// engine for QEMU. See https://wiki.qemu.org/Features/TCG for more details.
	TCG AcceleratorType = "tcg"

	// WHPX is the Windows Hypervisor Platform accelerator (Hyper-V).
	// See https://docs.microsoft.com/en-us/xamarin/android/get-started/installation/android-emulator/hardware-acceleration
	// for more details.
	WHPX AcceleratorType = "whpx"

	// Xen is a Type-1 (bare metal) hypervisor.
	// See https://wiki.xenproject.org/wiki/Xen_Project_Software_Overview for
	// more details.
	Xen AcceleratorType = "xen"
)

// WithAccelerator can be used with in conjunction with the qemu.With method
// to define an accelerator with no additional properties.
func WithAccelerator(acceleratorType string) *queso.Option {
	return queso.NewOption("accel", acceleratorType)
}

type Accelerator struct {
	Type       string
	properties []*queso.Property
}

func NewAccelerator(acceleratorType string) *Accelerator {
	return &Accelerator{
		Type:       acceleratorType,
		properties: make([]*queso.Property, 0),
	}
}

func (a *Accelerator) option() *queso.Option {
	return queso.NewOption("accel", a.Type, a.properties...)
}

type NotifyOnVMExitOption string

const (
	// NotifyOnVMExitRun enables VM exit support on x86 host. "window" specifies the
	// corresponding notification window of time to trigger the VM exit if enabled.
	//
	// This feature can mitigate the CPU stuck issue due to event windows that don’t
	// open up for a specified amount of time.
	NotifyOnVMExitRun NotifyOnVMExitOption = "run"

	// NotifyOnVMExitInternalError does not notify and continues if the exit happens,
	// but it raises an internal error.
	NotifyOnVMExitInternalError NotifyOnVMExitOption = "internal-error"

	// NotifyOnVMExitDisable does not notify and does nothing when the exit happens.
	NotifyOnVMExitDisable NotifyOnVMExitOption = "disable"
)

func (a *Accelerator) SetNotifyOnVMExit(option NotifyOnVMExitOption, window int) *Accelerator {
	value := string(option)
	if option == NotifyOnVMExitRun {
		value = fmt.Sprintf("run,notify-window=%d", window)
	}

	a.properties = append(a.properties, queso.NewProperty("notify-vmexit", value))
	return a
}

type KVMAccelerator struct {
	*Accelerator
}

func NewKVMAccelerator() *KVMAccelerator {
	return &KVMAccelerator{
		Accelerator: &Accelerator{
			Type:       string(KVM),
			properties: make([]*queso.Property, 0),
		},
	}
}

func (a *KVMAccelerator) option() *queso.Option {
	return a.option()
}

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

// SetKernelIRQChip controls KVM in-kernel irqchip support. The default is full
// acceleration of the interrupt controllers. On x86, split irqchip reduces the
// kernel attack surface, at a performance cost for non-MSI interrupts. Disabling
// the in-kernel irqchip completely is not recommended except for debugging purposes.
func (a *KVMAccelerator) SetKernelIRQChip(mode KernelIRQChipMode) *KVMAccelerator {
	a.properties = append(a.properties, queso.NewProperty("kernel-irqchip", mode))
	return a
}

// SetKVMShadowMemory defines the size of the KVM shadow MMU.
func (a *KVMAccelerator) SetKVMShadowMemory(size int) *KVMAccelerator {
	a.properties = append(a.properties, queso.NewProperty("kvm-shadow-mem", size))
	return a
}

// SetDirtyRingSize controls the size of the per-vCPU dirty page ring buffer
// (number of entries for each vCPU). It should be a value that is power
// of two, and it should be 1024 or bigger (but still less than the maximum value
// that the kernel supports). 4096 could be a good initial value if you have no
// idea which is the best. Set this value to 0 to disable the feature. By default,
// this feature is disabled (value = 0). When enabled, KVM will instead record
// dirty pages in a bitmap.
func (a *KVMAccelerator) SetDirtyRingSize(bytes int) *KVMAccelerator {
	a.properties = append(a.properties, queso.NewProperty("dirty-ring-size", bytes))
	return a
}

// SetEagerSplitSize specifies how many pages to break at a time. KVM implements dirty
// page logging at the PAGE_SIZE granularity and enabling dirty-logging on a huge-page
// requires breaking it into PAGE_SIZE pages in the first place. KVM on ARM does
// this splitting lazily by default.
//
// There are performance benefits in doing huge-page split eagerly, especially
// in situations where TLBI costs associated with break-before-make sequences
// are considerable and also if guest workloads are read intensive. The size
// needs to be a valid block size which is 1GB/2MB/4KB, 32MB/16KB and 512MB/64KB
// for 4KB/16KB/64KB PAGE_SIZE respectively.
//
// Be wary of specifying a higher size as it will have an impact on the memory.
// By default, this feature is disabled (i.e. size is set to 0).
func (a *KVMAccelerator) SetEagerSplitSize(size int) *KVMAccelerator {
	a.properties = append(a.properties, queso.NewProperty("eager-split-size", size))
	return a
}

type XENAccelerator struct {
	*Accelerator
}

func NewXenAccelerator() *XENAccelerator {
	return &XENAccelerator{
		Accelerator: &Accelerator{
			Type:       string(Xen),
			properties: make([]*queso.Property, 0),
		},
	}
}

func (a *XENAccelerator) option() *queso.Option {
	return a.option()
}

// ToggleIGDPassThru controls whether Intel integrated graphics devices can be passed
// through to the guest.
func (a *XENAccelerator) ToggleIGDPassThru(enabled bool) *XENAccelerator {
	a.properties = append(a.properties, queso.NewProperty("igd-passthru", enabled))
	return a
}

type TCGAccelerator struct {
	*Accelerator
}

func NewTCGAccelerator() *TCGAccelerator {
	return &TCGAccelerator{
		Accelerator: &Accelerator{
			Type:       string(TCG),
			properties: make([]*queso.Property, 0),
		},
	}
}

func (a *TCGAccelerator) option() *queso.Option {
	return a.option()
}

// ToggleOneInstructionPerTranslation makes the TCG accelerator put only one guest
// instruction into each translation block. This slows down emulation a lot, but
// can be useful in some situations, such as when trying to analyse the logs
// produced during debugging.
func (a *TCGAccelerator) ToggleOneInstructionPerTranslation(enabled bool) *TCGAccelerator {
	a.properties = append(a.properties, queso.NewProperty("one-insn-per-tb", enabled))
	return a
}

// ToggleSplitWX controls the use of split w^x mapping for the TCG code generation
// buffer. Some operating systems require this to be enabled, and in
// such a case this will default to true. On other operating systems, this will
// default to false, but one may enable this for testing or debugging.
func (a *TCGAccelerator) ToggleSplitWX(enabled bool) *TCGAccelerator {
	a.properties = append(a.properties, queso.NewProperty("split-wx", enabled))
	return a
}

// SetTranslationBlockCacheSize controls the size (in MiB) of the TCG
// translation block cache.
func (a *TCGAccelerator) SetTranslationBlockCacheSize(megabytes int) *TCGAccelerator {
	a.properties = append(a.properties, queso.NewProperty("tb-size", megabytes))
	return a
}

// ThreadingOption represents the type of TCG threads to use for Accel.
type ThreadingOption string

const (
	// SingleThreaded indicates that a single thread should be used with TCG.
	SingleThreaded ThreadingOption = "single"

	// MultiThreaded indicates that multiple threads should be used with TCG.
	MultiThreaded ThreadingOption = "multi"
)

// SetThreads controls number of TCG threads. When the TCG is multithreaded,
// there will be one thread per vCPU therefore taking advantage of additional
// host cores. The default is to enable multi-threading where both the back-end
// and front-ends support it and no incompatible TCG features have been enabled
// (e.g. icount/replay).
func (a *TCGAccelerator) SetThreads(option ThreadingOption) *TCGAccelerator {
	a.properties = append(a.properties, queso.NewProperty("thread", option))
	return a
}
