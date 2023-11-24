// Package device is used to specify device drivers for use with QEMU.
package device

import (
	"fmt"

	"github.com/mikerourke/queso"
)

// Device represents a device used with QEMU.
type Device struct {
	Type       string
	properties []*queso.Property
}

// New returns a new instance of a [Device].
//
//	qemu-system-* -device <deviceType>
func New(deviceType string) *Device {
	return &Device{
		Type:       deviceType,
		properties: make([]*queso.Property, 0),
	}
}

// Option returns the option representation of the device.
func (d *Device) option() *queso.Option {
	return queso.NewOption("device", d.Type, d.properties...)
}

// Use adds a device driver with the specified name and properties.
// There are a _lot_ of options for deviceType. See the QEMU documentation for
// additional details.
func Use(deviceType string, properties ...*Property) *queso.Option {
	props := make([]*queso.Property, 0)

	for _, property := range properties {
		props = append(props, property.Property)
	}

	return queso.NewOption("device", deviceType, props...)
}

// IPMIBMC adds an IPMI BMC. This is a simulation of a hardware management
// interface processor that normally sits on a system. It provides a watchdog
// and the ability to reset and power control the system. You need to connect
// this to an IPMI interface to make it useful.
func IPMIBMC(id string, properties ...*Property) *queso.Option {
	props := []*Property{NewProperty("id", id)}

	if properties != nil {
		props = append(props, properties...)
	}

	return Use("ipmi-bmc-sim", props...)
}

// IPMIBMCExternal adds a connection to an external IPMI BMC simulator.
// Instead of locally emulating the BMC like the above item, instead connect to
// an external entity that provides the IPMI services.
//
// A connection is made to an external BMC simulator. If you do this, it is
// strongly recommended that you use the chardev.WithReconnect option to reconnect
// to the simulator if the connection is lost. Note that if this is not used
// carefully, it can be a security issue, as the interface has the ability to
// send resets, NMIs, and power off the VM. It's best if QEMU makes a connection
// to an external simulator running on a secure port on localhost, so neither the
// simulator nor QEMU is exposed to any outside network.
//
// See the "lanserv/README.vm" file in the OpenIPMI library for more details on the
// external interface.
func IPMIBMCExternal(
	id string,
	chardev string,
	properties ...*Property,
) *queso.Option {
	props := []*Property{
		NewProperty("id", id),
		NewProperty("chardev", chardev),
	}

	if properties != nil {
		props = append(props, properties...)
	}

	return Use("ipmi-bmc-extern", props...)
}

// KCSIPMIOnISABus adds a KCS IPMI interface on the ISA bus. This also adds a
// corresponding ACPI and SMBIOS entries, if appropriate. The bmc parameter
// represents the ID of a IPMIBMC or IPMIBMCExternal instance.
func KCSIPMIOnISABus(bmc string, properties ...*Property) *queso.Option {
	props := []*Property{NewProperty("bmc", bmc)}

	if properties != nil {
		props = append(props, properties...)
	}

	return Use("isa-ipmi-kcs", props...)
}

// BTIPMIOnISABus adds a BT IPMI interface on the ISA bus. The bmc parameter
// represents the ID of a IPMIBMC or IPMIBMCExternal instance.
func BTIPMIOnISABus(bmc string, properties ...*Property) *queso.Option {
	props := []*Property{NewProperty("bmc", bmc)}

	if properties != nil {
		props = append(props, properties...)
	}

	return Use("isa-ipmi-bt", props...)
}

// KCSIPMIOnPCIBus adds a KCS IPMI interface on the PCI bus. The bmc parameter
// represents the ID of a IPMIBMC or IPMIBMCExternal instance.
func KCSIPMIOnPCIBus(bmc string, properties ...*Property) *queso.Option {
	props := []*Property{NewProperty("bmc", bmc)}

	if properties != nil {
		props = append(props, properties...)
	}

	return Use("pci-ipmi-kcs", props...)
}

// BTIPMIOnPCIBus adds a BT IPMI interface on the PCI bus. The bmc parameter
// represents the ID of a IPMIBMC or IPMIBMCExternal instance.
func BTIPMIOnPCIBus(bmc string, properties ...*Property) *queso.Option {
	props := []*Property{NewProperty("bmc", bmc)}

	if properties != nil {
		props = append(props, properties...)
	}

	return Use("pci-ipmi-bc", props...)
}

// IntelIOMMU is used to enable Intel VT-d emulation within the guest.
// This can only be used with Q35 Machine instances.
func IntelIOMMU(properties ...*Property) *queso.Option {
	return Use("intel-iommu", properties...)
}

// Virtio9PVariant represents the variant of Virtio9P to use for a new
// Virtio9P.
type Virtio9PVariant string

const (
	Virtio9PPCI    Virtio9PVariant = "pci"
	Virtio9PCCW    Virtio9PVariant = "ccw"
	Virtio9PDevice Virtio9PVariant = "device"
)

// Virtio9P adds a Virtio 9P file system. The fsdev parameter corresponds
// to a filesystem device (see blockdev/fsdev.go). The mountTag parameter
// specifies the tag name to be used by the guest to mount this export point.
func Virtio9P(variant Virtio9PVariant, fsdev string, mountTag string) *queso.Option {
	name := fmt.Sprintf("virtio-9p-%s", variant)

	props := []*Property{
		NewProperty("fsdev", fsdev),
		NewProperty("mount_tag", mountTag),
	}

	return Use(name, props...)
}

// Property represents a property that can be used with the device option.
type Property struct {
	*queso.Property
}

// NewProperty returns a new instance of Property.
func NewProperty(key string, value interface{}) *Property {
	return &Property{
		Property: queso.NewProperty(key, value),
	}
}

// WithID defines the ID to use for the device.
func WithID(id string) *Property {
	return NewProperty("id", id)
}

// WithBus defines the bus name to use for the device.
func WithBus(name string) *Property {
	return NewProperty("bus", name)
}

// WithSlaveAddress defines a slave address to use for the BMC. The default is
// 0x20.
func WithSlaveAddress(addr int) *Property {
	return NewProperty("slave_addr", addr)
}

// WithSDRFile specifies a file containing raw Sensor Data Records (SDR) data.
func WithSDRFile(file string) *Property {
	return NewProperty("sdrfile", file)
}

// WithFRUAreaSize specifies the size of a Field Replaceable Unit (FRU) area.
// Default is 1024.
func WithFRUAreaSize(bytes int) *Property {
	return NewProperty("fruareasize", bytes)
}

// WithFRUFile specifies a file containing raw Field Replaceable Unit (FRU)
// inventory data.
func WithFRUFile(file string) *Property {
	return NewProperty("frudatafile", file)
}

// WithGUID specifies the value for the GUID for the BMC, in standard UUID format.
// If this is set, get "Get GUID" command to the BMC will return it. Otherwise,
// "Get GUID" will return an error.
func WithGUID(guid string) *Property {
	return NewProperty("guid", guid)
}

// WithIOPort defines the I/O address of the interface.
func WithIOPort(addr int) *Property {
	return NewProperty("ioport", addr)
}

// WithIRQ defines the interrupt to use.
func WithIRQ(interrupt int) *Property {
	return NewProperty("irq", interrupt)
}

// IsInterruptRemapping enables/disables interrupt remapping feature. It's
// required to enable complete x2apic. Currently, it only supports KVM
// qemu.WithKernelIRQChip modes qemu.KernelIRQChipOff or qemu.KernelIRQChipSplit,
// while qemu.KernelIRQChipOn is not yet supported. Default value is automatically
// decided by the mode of qemu.WithKernelIRQChip.
func IsInterruptRemapping(enabled bool) *Property {
	return NewProperty("intremap", enabled)
}

// IsCachingMode enables/disables caching mode for the VT-d emulated device.
// When IsCachingMode is true, each guest DMA buffer mapping will generate an IOTLB
// invalidation from the guest IOMMU driver to the vIOMMU device in a synchronous way.
// It is required for a VFIO PCI device to work with the VT-d device, because
// host assigned devices are required to set up the DMA mapping on the host
// before guest DMA starts. This property is disabled by default.
func IsCachingMode(enabled bool) *Property {
	return NewProperty("caching-mode", enabled)
}

// IsIOTLB enables/disabled device-iotlb capability for the emulated VT-d device.
// So far virtio/vhost should be the only real user for this parameter, paired
// with IsATS = enabled configured for the device. This property is disabled by
// default.
func IsIOTLB(enabled bool) *Property {
	return NewProperty("iotlb", enabled)
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
func WithAddressWidthBits(bits AddressWidthBits) *Property {
	return NewProperty("aw-bits", bits)
}
