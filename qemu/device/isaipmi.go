package device

import "github.com/mikerourke/queso"

// BTIPMIOnISABus represents a BT IPMI interface on the ISA bus. This also
// adds a corresponding ACPI and SMBIOS entries, if appropriate.
type BTIPMIOnISABus struct {
	*Device
}

// NewBTIPMIOnISABus returns a new instance of [BTIPMIOnISABus]. bmc represents
// the ID of a [IPMIBMCSimulated] or [IPMIBMCExternal] instance.
//
//	qemu-system-* -device isa-ipmi-bt,bmc=id
func NewBTIPMIOnISABus(bmc string) *BTIPMIOnISABus {
	device := New("isa-ipmi-bt")
	device.properties = append(device.properties, queso.NewProperty("bmc", bmc))

	return &BTIPMIOnISABus{device}
}

// SetIOAddress defines the I/O address of the interface. The default is 0xe4
// for BT.
//
//	qemu-system-* -device isa-ipmi-kcs,ioport=addr
func (i *BTIPMIOnISABus) SetIOAddress(addr int) *BTIPMIOnISABus {
	i.properties = append(i.properties, queso.NewProperty("ioport", addr))
	return i
}

// SetIRQInterrupt defines the interrupt to use. The default is 5. To disable
// interrupts, set this to 0.
//
//	qemu-system-* -device isa-ipmi-kcs,irq=value
func (i *BTIPMIOnISABus) SetIRQInterrupt(value int) *BTIPMIOnISABus {
	i.properties = append(i.properties, queso.NewProperty("irq", value))
	return i
}

// KCSIPMIOnISABus represents a KCS IPMI interface on the ISA bus. This also
// adds a corresponding ACPI and SMBIOS entries, if appropriate.
type KCSIPMIOnISABus struct {
	*Device
}

// NewKCSIPMIOnISABus returns a new instance of [KCSIPMIOnISABus]. bmc represents
// the ID of a [IPMIBMCSimulated] or [IPMIBMCExternal] instance.
//
//	qemu-system-* -device isa-ipmi-kcs,bmc=id
func NewKCSIPMIOnISABus(bmc string) *KCSIPMIOnISABus {
	device := New("isa-ipmi-kcs")
	device.properties = append(device.properties, queso.NewProperty("bmc", bmc))

	return &KCSIPMIOnISABus{device}
}

// SetIOAddress defines the I/O address of the interface. The default is 0xca0
// for KCS.
//
//	qemu-system-* -device isa-ipmi-kcs,ioport=addr
func (i *KCSIPMIOnISABus) SetIOAddress(addr int) *KCSIPMIOnISABus {
	i.properties = append(i.properties, queso.NewProperty("ioport", addr))
	return i
}

// SetIRQInterrupt defines the interrupt to use. The default is 5. To disable
// interrupts, set this to 0.
//
//	qemu-system-* -device isa-ipmi-kcs,irq=value
func (i *KCSIPMIOnISABus) SetIRQInterrupt(value int) *KCSIPMIOnISABus {
	i.properties = append(i.properties, queso.NewProperty("irq", value))
	return i
}
