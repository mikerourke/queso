package device

import "github.com/mikerourke/queso"

type IntelIOMMU struct {
	*Device
}

func NewIntelIOMMU() *IntelIOMMU {
	return &IntelIOMMU{New("intel-iommu")}
}

// AddressWidthBits represents the options available that can be set with the
// [IntelIOMMU.SetAddressWidthBits] method.
type AddressWidthBits int

const (
	// AddressWidthBits39 is used for 3-level IOMMU page tables.
	AddressWidthBits39 AddressWidthBits = 39

	// AddressWidthBits48 is used for 4-level IOMMU page tables.
	AddressWidthBits48 AddressWidthBits = 48
)

// SetAddressWidthBits specifies the address width of IOVA address space.
// The address space has 39 bits width for 3-level IOMMU page tables, and 48 bits
// for 4-level IOMMU page tables. Default is [AddressWidthBits39].
//
//	qemu-system-* -device intel-iommu,aw-bits=39|48
func (i *IntelIOMMU) SetAddressWidthBits(bits AddressWidthBits) *IntelIOMMU {
	i.properties = append(i.properties, queso.NewProperty("aw-bits", bits))
	return i
}

// TODO: Find out what interrupt remapping is.

// InterruptRemapping represents the options available that can be set with the
// [IntelIOMMU.SetInterruptRemapping] method.
type InterruptRemapping string

const (
	InterruptRemappingOn   InterruptRemapping = "on"
	InterruptRemappingOff  InterruptRemapping = "off"
	InterruptRemappingAuto InterruptRemapping = "auto"
)

// SetInterruptRemapping sets the interrupt remapping feature. It’s required to
// enable complete x2apic. Currently, it only supports kvm kernel-irqchip modes
// off or split, while full kernel-irqchip is not yet supported.
//
// The default value is [InterruptRemappingAuto], which will be decided by the
// mode of kernel-irqchip.
//
//	qemu-system-* -device intel-iommu,intremap=on|off|auto
func (i *IntelIOMMU) SetInterruptRemapping(remapping InterruptRemapping) *IntelIOMMU {
	i.properties = append(i.properties, queso.NewProperty("intremap", remapping))
	return i
}

// ToggleCachingMode enables or disables caching mode for the VT-d emulated device.
// When enabled, each guest DMA buffer mapping will generate an IOTLB invalidation
// from the guest IOMMU driver to the vIOMMU device in a synchronous way.
// It is required for a vfio-pci [Device] to work with the VT-d device, because
// host assigned devices requires to set up the DMA mapping on the host before
// guest DMA starts.
//
//	qemu-system-* -device intel-iommu,caching-mode=on|off
func (i *IntelIOMMU) ToggleCachingMode(enabled bool) *IntelIOMMU {
	i.properties = append(i.properties, queso.NewProperty("caching-mode", enabled))
	return i
}

// ToggleDeviceIOTLB enables or disables device-iotlb capability for the emulated
// VT-d device. So far virtio/vhost should be the only real user for this
// parameter, paired with ats=on configured for the device.
//
//	qemu-system-* -device intel-iommu,device-iotlb=on|off
func (i *IntelIOMMU) ToggleDeviceIOTLB(enabled bool) *IntelIOMMU {
	i.properties = append(i.properties, queso.NewProperty("device-iotlb", enabled))
	return i
}
