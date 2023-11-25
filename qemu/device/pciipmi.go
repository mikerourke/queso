package device

import "github.com/mikerourke/queso"

// BTIPMIOnPCIBus represents a BT IPMI interface on the PCI bus.
type BTIPMIOnPCIBus struct {
	*Device
}

// NewBTIPMIOnPCIBus returns a new instance of [BTIPMIOnPCIBus]. bmc
// represents the ID of a [IPMIBMCSimulated] or [IPMIBMCExternal] instance.
//
//	qemu-system-* -device pci-ipmi-bt,bmc=id
func NewBTIPMIOnPCIBus(bmc string) *BTIPMIOnPCIBus {
	device := New("pci-ipmi-bt")
	device.properties = append(device.properties, queso.NewProperty("bmc", bmc))

	return &BTIPMIOnPCIBus{device}
}

// KCSIPMIOnPCIBus represents a KCS IPMI interface on the PCI bus.
type KCSIPMIOnPCIBus struct {
	*Device
}

// NewKCSIPMIOnPCIBus returns a new instance of [KCSIPMIOnPCIBus]. bmc
// represents the ID of a [IPMIBMCSimulated] or [IPMIBMCExternal] instance.
//
//	qemu-system-* -device pci-ipmi-kcs,bmc=id
func NewKCSIPMIOnPCIBus(bmc string) *KCSIPMIOnPCIBus {
	device := New("pci-ipmi-kcs")
	device.properties = append(device.properties, queso.NewProperty("bmc", bmc))

	return &KCSIPMIOnPCIBus{device}
}
