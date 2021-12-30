package qemu

import (
	"strings"

	"github.com/mikerourke/queso"
)

// Machine selects the emulated machine by name.
func Machine(name string, properties ...*MachineProperty) *queso.Option {
	props := make([]*queso.Property, 0)

	for _, property := range properties {
		props = append(props, property.Property)
	}

	return queso.NewOption("machine", name, props...)
}

// MachineProperty represents a property that can be used with the Machine option.
type MachineProperty struct {
	*queso.Property
}

// NewMachineProperty returns a new instance of MachineProperty.
func NewMachineProperty(key string, value interface{}) *MachineProperty {
	return &MachineProperty{
		Property: queso.NewProperty(key, value),
	}
}

// WithAccel specifies one or more accelerators to use. If there is more than one
// accelerator specified, the next one is used if the previous one fails to
// initialize.
func WithAccel(types ...AccelType) *MachineProperty {
	nameStrings := make([]string, 0)
	for _, at := range types {
		nameStrings = append(nameStrings, string(at))
	}

	return NewMachineProperty("accel", strings.Join(nameStrings, ":"))
}

// VMWareIOPortFlag represents the flag to pass to the WithVMWareIOPort property.
type VMWareIOPortFlag string

const (
	VMWareIOPortOn   VMWareIOPortFlag = "on"
	VMWareIOPortOff  VMWareIOPortFlag = "off"
	VMWareIOPortAuto VMWareIOPortFlag = "auto"
)

// WithVMWareIOPort sets emulation of VMWare IO port, for vmmouse etc.
// VMWareIOPortAuto says to select the value based on accel. For XEN accelerator,
// the default is VMWareIOPortOff otherwise the default is VMWareIOPortOn.
func WithVMWareIOPort(port VMWareIOPortFlag) *MachineProperty {
	return NewMachineProperty("vmport", port)
}

// IsDumpGuestCore specifies whether to include guest memory in a core dump.
func IsDumpGuestCore(enabled bool) *MachineProperty {
	return NewMachineProperty("dump-guest-core", enabled)
}

// IsMemoryMerge enables or disables memory merge support. This feature, when
// supported by the host, de-duplicates identical memory pages among VMs
// instances (enabled by default).
func IsMemoryMerge(enabled bool) *MachineProperty {
	return NewMachineProperty("mem-merge", enabled)
}

// IsAESKeyWrap enables or disables AES key wrapping support on s390-ccw hosts.
// This feature controls whether AES wrapping keys will be created to allow
// execution of AES cryptographic functions (enabled by default).
func IsAESKeyWrap(enabled bool) *MachineProperty {
	return NewMachineProperty("aes-key-wrap", enabled)
}

// IsDEAKeyWrap enables or disables DEA key wrapping support on s390-ccw hosts.
// This feature controls whether DEA wrapping keys will be created to allow
// execution of DEA cryptographic functions (enabled by default).
func IsDEAKeyWrap(enabled bool) *MachineProperty {
	return NewMachineProperty("dea-key-wrap", enabled)
}

// IsNVDIMM enables or disables NVDIMM support (disabled by default).
func IsNVDIMM(enabled bool) *MachineProperty {
	return NewMachineProperty("nvdimm", enabled)
}

// WithMemoryEncryption specifies the memory encryption object to use.
func WithMemoryEncryption(id string) *MachineProperty {
	return NewMachineProperty("memory-encryption", id)
}

// IsHMAT enables or disables ACPI Heterogeneous Memory Attribute Table (HMAT)
// support (disabled by default).
func IsHMAT(enabled bool) *MachineProperty {
	return NewMachineProperty("hmat", enabled)
}

// WithMemoryBackend allows use of a memory backend as main RAM.
func WithMemoryBackend(id string) *MachineProperty {
	return NewMachineProperty("memory-backend", id)
}
