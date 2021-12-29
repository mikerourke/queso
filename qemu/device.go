package qemu

import (
	"fmt"

	"github.com/mikerourke/queso"
)

// Device is used to create a generic device in which arbitrary options can
// be specified.
func Device(name string, properties ...*DeviceProperty) *queso.Option {
	props := make([]*queso.Property, 0)

	for _, property := range properties {
		props = append(props, property.Property)
	}

	return queso.NewOption("device", name, props...)
}

// newDeviceOption returns a new Option instance for a device. The idField
// parameter represents the field name for the identifier. This is required because
// some devices use `id=[ID]` while others use `bmc=[ID]`.
func newDeviceOption(name string, idField string, id string, properties ...*DeviceProperty) *queso.Option {
	props := make([]*queso.Property, 0)

	if idField != "" && id != "" {
		props = append(props, queso.NewProperty(idField, id))
	}

	for _, property := range properties {
		props = append(props, property.Property)
	}

	return queso.NewOption("device", name, props...)
}

// DeviceDriver adds a device driver with the specified name and properties.
func DeviceDriver(name string, properties ...*DeviceProperty) *queso.Option {
	props := make([]*queso.Property, 0)

	for _, property := range properties {
		props = append(props, property.Property)
	}

	return queso.NewOption("device", name, props...)
}

// DeviceIPMIBMC adds an IPMI BMC. This is a simulation of a hardware management interface
// processor that normally sits on a system. It provides a watchdog and the ability
// to reset and power control the system. You need to connect this to an IPMI
// interface to make it useful.
func DeviceIPMIBMC(id string, properties ...*DeviceProperty) *queso.Option {
	return newDeviceOption("ipmi-bmc-sim", "id", id, properties...)
}

// DeviceIPMIBMCExternal adds a connection to an external IPMI BMC simulator.
// Instead of locally emulating the BMC like the above item, instead connect to an
// external entity that provides the IPMI services.
//
// A connection is made to an external BMC simulator. If you do this, it is strongly
// recommended that you use the WithReconnect chardev option to reconnect to the
// simulator if the connection is lost. Note that if this is not used carefully, it
// can be a security issue, as the interface has the ability to send resets, NMIs,
// and power off the VM. It’s best if QEMU makes a connection to an external simulator
// running on a secure port on localhost, so neither the simulator nor QEMU is exposed
// to any outside network.
//
// See the "lanserv/README.vm" file in the OpenIPMI library for more details on the
// external interface.
func DeviceIPMIBMCExternal(id string, chardev string, properties ...*DeviceProperty) *queso.Option {
	props := []*DeviceProperty{NewDeviceProperty("chardev", chardev)}

	if properties != nil {
		props = append(props, properties...)
	}

	return newDeviceOption("ipmi-bmc-sim", "id", id, props...)
}

// DeviceISAIPMIKCS adds a KCS IPMI interface on the ISA bus. This also adds a
// corresponding ACPI and SMBIOS entries, if appropriate. The bmc parameter
// // represents the ID of a DeviceIPMIBMC or DeviceIPMIBMCExternal instance.
func DeviceISAIPMIKCS(bmc string, properties ...*DeviceProperty) *queso.Option {
	return newDeviceOption("isa-ipmi-kcs", "bmc", bmc, properties...)
}

// DeviceISAIPMIBT adds a BT IPMI interface on the ISA bus. The bmc parameter
// represents the ID of a DeviceIPMIBMC or DeviceIPMIBMCExternal instance.
func DeviceISAIPMIBT(bmc string, properties ...*DeviceProperty) *queso.Option {
	return newDeviceOption("isa-ipmi-bt", "bmc", bmc, properties...)
}

// DevicePCIIPMIKCS adds a KCS IPMI interface on the PCI bus. The bmc parameter
// // represents the ID of a DeviceIPMIBMC or DeviceIPMIBMCExternal instance.
func DevicePCIIPMIKCS(bmc string, properties ...*DeviceProperty) *queso.Option {
	return newDeviceOption("pci-ipmi-kcs", "bmc", bmc, properties...)
}

// DevicePCIIPMIBT adds a BT IPMI interface on the PCI bus. The bmc parameter
// // represents the ID of a DeviceIPMIBMC or DeviceIPMIBMCExternal instance.
func DevicePCIIPMIBT(bmc string, properties ...*DeviceProperty) *queso.Option {
	return newDeviceOption("pci-ipmi-bc", "bmc", bmc, properties...)
}

// DeviceIntelIOMMU is used to enable Intel VT-d emulation within the guest.
// This can only be used with Q35 Machine instances.
func DeviceIntelIOMMU(properties ...*DeviceProperty) *queso.Option {
	return newDeviceOption("intel-iommu", "", "", properties...)
}

// Virtio9PVariant represents the variant of Virtio9P to use for a new
// DeviceVirtio9P.
type Virtio9PVariant string

const (
	Virtio9PPCI    Virtio9PVariant = "pci"
	Virtio9PCCW    Virtio9PVariant = "ccw"
	Virtio9PDevice Virtio9PVariant = "device"
)

// DeviceVirtio9P adds a Virtio 9P file system. The deviceID parameter corresponds
// to a filesystem device (see blockdev/fsdev.go). The mountTag specifies the
// tag name to be used by the guest to mount this export point.
func DeviceVirtio9P(variant Virtio9PVariant, deviceID string, mountTag string) *queso.Option {
	name := fmt.Sprintf("virtio-9p-%s", variant)

	return newDeviceOption(name, "fsdev", deviceID,
		NewDeviceProperty("mount_tag", mountTag))
}

// DeviceProperty represents a property that can be used with the device option.
type DeviceProperty struct {
	*queso.Property
}

// NewDeviceProperty returns a new instance of a DeviceProperty.
func NewDeviceProperty(key string, value interface{}) *DeviceProperty {
	return &DeviceProperty{
		Property: queso.NewProperty(key, value),
	}
}

// WithSlaveAddr defines a slave address to use for the BMC. The default is 0x20.
func WithSlaveAddr(addr int) *DeviceProperty {
	return NewDeviceProperty("slave_addr", addr)
}

// WithSDRFile specifies a file containing raw Sensor Data Records (SDR) data.
func WithSDRFile(filename string) *DeviceProperty {
	return NewDeviceProperty("sdrfile", filename)
}

// WithFRUAreaSize specifies the size of a Field Replaceable Unit (FRU) area.
// Default is 1024.
func WithFRUAreaSize(bytes int) *DeviceProperty {
	return NewDeviceProperty("fruareasize", bytes)
}

// WithFRUFile specifies a file containing raw Field Replaceable Unit (FRU)
// inventory data.
func WithFRUFile(filename string) *DeviceProperty {
	return NewDeviceProperty("frudatafile", filename)
}

// WithGUID specifies the value for the GUID for the BMC, in standard UUID format.
// If this is set, get "Get GUID" command to the BMC will return it. Otherwise,
// "Get GUID" will return an error.
func WithGUID(guid string) *DeviceProperty {
	return NewDeviceProperty("guid", guid)
}

// WithIOPort defines the I/O address of the interface.
func WithIOPort(addr int) *DeviceProperty {
	return NewDeviceProperty("ioport", addr)
}

// WithIRQ defines the interrupt to use.
func WithIRQ(interrupt int) *DeviceProperty {
	return NewDeviceProperty("irq", interrupt)
}

// IsInterruptRemapping enables/disables interrupt remapping feature. It's required to
// enable complete x2apic. Currently, it only supports KVM WithKernelIRQChip modes
// KernelIRQChipOff or KernelIRQChipSplit, while KernelIRQChipOn is not yet supported.
// Default value is automatically decided by the mode of WithKernelIRQChip.
func IsInterruptRemapping(enabled bool) *DeviceProperty {
	return NewDeviceProperty("intremap", enabled)
}

// IsCachingMode enables/disables caching mode for the VT-d emulated device.
// When IsCachingMode is true, each guest DMA buffer mapping will generate an IOTLB
// invalidation from the guest IOMMU driver to the vIOMMU device in a synchronous way.
// It is required for a VFIO PCI device to work with the VT-d device, because host assigned
// devices are required to set up the DMA mapping on the host before guest DMA starts.
// Default is disabled.
func IsCachingMode(enabled bool) *DeviceProperty {
	return NewDeviceProperty("caching-mode", enabled)
}

// IsIOTLB enables/disabled device-iotlb capability for the emulated VT-d device.
// So far virtio/vhost should be the only real user for this parameter, paired with
// IsATS = enabled configured for the device. Default is disabled.
func IsIOTLB(enabled bool) *DeviceProperty {
	return NewDeviceProperty("iotlb", enabled)
}

// AddressWidthBits represents the options available for the WithAddressWidthBits
// property.
type AddressWidthBits int

const (
	// AddressWidthBits39 is used for 3-level IOMMU page tables.
	AddressWidthBits39 AddressWidthBits = 39

	// AddressWidthBits48 is used for 4-level IOMMU page tables.
	AddressWidthBits48 AddressWidthBits = 48
)

// WithAddressWidthBits specifies the address width of IOVA address space.
// The address space has 39 bits width for 3-level IOMMU page tables, and 48 bits
// for 4-level IOMMU page tables. Default is AddressWidthBits39.
func WithAddressWidthBits(bits AddressWidthBits) *DeviceProperty {
	return NewDeviceProperty("aw-bits", bits)
}
