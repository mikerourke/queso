package accel

import "github.com/mikerourke/queso/qemu/cli"

// KVMAccelerator represents an accelerator using the Kernel Virtual Machine,
// which is a Linux kernel module. See https://wiki.qemu.org/Features/KVM
// for more details.
type KVMAccelerator struct {
	*Accelerator
}

// NewKVMAccelerator returns a new instace of [KVMAccelerator].
//
//	qemu-system-* -accel kvm
func NewKVMAccelerator() *KVMAccelerator {
	return &KVMAccelerator{
		New(TypeKVM),
	}
}

// SetDirtyRingSize controls the size of the per-vCPU dirty page ring buffer
// (number of entries for each vCPU). It should be a value that is power
// of two, and it should be 1024 or bigger (but still less than the maximum value
// that the kernel supports). 4096 could be a good initial value if you have no
// idea which is the best. Set this value to 0 to disable the feature. By default,
// this feature is disabled (value = 0). When enabled, KVM will instead record
// dirty pages in a bitmap.
//
//	qemu-system-* -accel kvm dirty-ring-size=bytes
func (a *KVMAccelerator) SetDirtyRingSize(bytes int) *KVMAccelerator {
	a.properties = append(a.properties, cli.NewProperty("dirty-ring-size", bytes))
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
//
//	qemu-system-* -accel kvm eager-split-size=size
func (a *KVMAccelerator) SetEagerSplitSize(size int) *KVMAccelerator {
	a.properties = append(a.properties, cli.NewProperty("eager-split-size", size))
	return a
}

// KernelIRQChipMode represents the mode to use for the SetKernelIRQChip
// property.
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
//
//	qemu-system-* -accel kvm kernel-irqchip=on|off|split
func (a *KVMAccelerator) SetKernelIRQChip(mode KernelIRQChipMode) *KVMAccelerator {
	a.properties = append(a.properties, cli.NewProperty("kernel-irqchip", mode))
	return a
}

// SetKVMShadowMemory defines the size of the KVM shadow MMU.
//
//	qemu-system-* -accel kvm kvm-shadow-mem=size
func (a *KVMAccelerator) SetKVMShadowMemory(size int) *KVMAccelerator {
	a.properties = append(a.properties, cli.NewProperty("kvm-shadow-mem", size))
	return a
}
