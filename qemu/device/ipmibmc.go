package device

import "github.com/mikerourke/queso"

// IPMIBMCExternal represents a connection to an external IPMI BMC simulator.
// Instead of locally emulating the BMC like in [IPMIBMCSimulated], connect to
// an external entity that provides the IPMI services.
//
// A connection is made to an external BMC simulator. If you do this, it is
// strongly recommended that you use the SetReconnectTimeout method on
// chardev.SocketBackend to reconnect to the simulator if the connection is
// lost. Note that if this is not used carefully, it can be a security issue, as
// the interface has the ability to send resets, NMIs, and power off the VM.
// It's best if QEMU makes a connection to an external simulator running on a
// secure port on localhost, so neither the simulator nor QEMU is exposed to any
// outside network.
//
// See the "lanserv/README.vm" file in the OpenIPMI library for more details on
// the external interface.
type IPMIBMCExternal struct {
	*Device
}

// NewIPMIBMCExternal returns a new instance of [IPMIBMCExternal].
// id is a unique identifier for the device. chardev is the ID of the character
// device representing the connection to the external BMC simulator.
//
//	qemu-system-* -device ipmi-bmc-extern,id=id,chardev=chardev
func NewIPMIBMCExternal(id string, chardev string) *IPMIBMCExternal {
	device := New("ipmi-bmc-sim").
		SetProperty("id", id).
		SetProperty("chardev", chardev)

	return &IPMIBMCExternal{device}
}

// SetSlaveAddress defines the slave address to use for the BMC.
//
//	qemu-system-* -device ipmi-bmc-extern,slave_addr=addr
func (i *IPMIBMCExternal) SetSlaveAddress(addr int) *IPMIBMCExternal {
	i.properties = append(i.properties, queso.NewProperty("slave_addr", addr))
	return i
}

// IPMIBMCSimulated represents an IPMI BMC. This is a simulation of a hardware
// management interface processor that normally sits on a system. It provides a
// watchdog and the ability to reset and power control the system. You need to
// connect this to an IPMI interface to make it useful.
type IPMIBMCSimulated struct {
	*Device
}

// NewIPMIBMCSimulated returns a new instance of [IPMIBMCSimulated]. id is a
// unique identifier for the device.
//
//	qemu-system-* -device ipmi-bmc-sim,id=id
func NewIPMIBMCSimulated(id string) *IPMIBMCSimulated {
	device := New("ipmi-bmc-sim")
	device.properties = append(device.properties, queso.NewProperty("id", id))

	return &IPMIBMCSimulated{device}
}

// SetFRUAreaSize sets the size of a Field Replaceable Unit (FRU) area. The
// default is 1024.
//
//	qemu-system-* -device ipmi-bmc-sim,fruareasize=size
func (i *IPMIBMCSimulated) SetFRUAreaSize(size int) *IPMIBMCSimulated {
	i.properties = append(i.properties, queso.NewProperty("fruareasize", size))
	return i
}

// SetFRUDataFile sets the file containing raw Field Replaceable Unit (FRU)
// inventory data. The default is none.
//
//	qemu-system-* -device ipmi-bmc-sim,frudatafile=file
func (i *IPMIBMCSimulated) SetFRUDataFile(file int) *IPMIBMCSimulated {
	i.properties = append(i.properties, queso.NewProperty("frudatafile", file))
	return i
}

// SetGUID sets the value for the GUID for the BMC, in standard UUID format. If
// this is set, get “Get GUID” command to the BMC will return it.
// Otherwise, “Get GUID” will return an error.
//
//	qemu-system-* -device ipmi-bmc-sim,guid=uuid
func (i *IPMIBMCSimulated) SetGUID(guid string) *IPMIBMCSimulated {
	i.properties = append(i.properties, queso.NewProperty("guid", guid))
	return i
}

// SetSDRFile sets the file containing raw Sensor Data Records (SDR) data.
// The default is none.
//
//	qemu-system-* -device ipmi-bmc-sim,sdrfile=file
func (i *IPMIBMCSimulated) SetSDRFile(file string) *IPMIBMCSimulated {
	i.properties = append(i.properties, queso.NewProperty("sdrfile", file))
	return i
}

// SetSlaveAddress defines the slave address to use for the BMC. The default is 0x20.
//
//	qemu-system-* -device ipmi-bmc-sim,slave_addr=addr
func (i *IPMIBMCSimulated) SetSlaveAddress(addr int) *IPMIBMCSimulated {
	i.properties = append(i.properties, queso.NewProperty("slave_addr", addr))
	return i
}
