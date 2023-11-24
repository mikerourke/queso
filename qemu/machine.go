package qemu

import (
	"strings"

	"github.com/mikerourke/queso"
)

// VMWareIOPortFlag represents the flag to pass to the SetVMWareIOPort method
// for a [Machine].
type VMWareIOPortFlag string

const (
	// VMWareIOPortOn indicates that VMWare IO port emulation is enabled.
	VMWareIOPortOn VMWareIOPortFlag = "on"

	// VMWareIOPortOff indicates that VMWare IO port emulation is disabled.
	VMWareIOPortOff VMWareIOPortFlag = "off"

	// VMWareIOPortAuto indicates that VMWare IO port emulation is either enabled
	// or disabled based on the accelerator.
	VMWareIOPortAuto VMWareIOPortFlag = "auto"
)

// Machine selects the emulated machine by name.
type Machine struct {
	Name       string
	properties []*queso.Property
}

// NewMachine returns a new Machine instance.
//
// Example
//
//	machine := qemu.NewMachine().SetMemoryBackend("pc.ram")
//	qemu.New("qemu-system-x86_64").
//		With(
//			object.MemoryBackendFile("pc.ram",
//				object.WithMemorySize("512M"),
//				object.WithMemoryPath("/hugetlbfs"),
//				object.IsPrealloc(true),
//				object.IsShare(true)),
//			qemu.Memory("512M"),
//		).Use(machine)
//
//	qemu-system-x86_64 \
//		-object memory-backend-file,id=pc.ram,size=512M,mem-path=/hugetlbfs,prealloc=on,share=on \
//		-machine memory-backend=pc.ram \
//		-m 512M
func NewMachine() *Machine {
	return &Machine{
		Name:       "",
		properties: make([]*queso.Property, 0),
	}
}

func (m *Machine) option() *queso.Option {
	return queso.NewOption("machine", m.Name, m.properties...)
}

// SetProperty is used to add arbitrary properties to the [Machine].
func (m *Machine) SetProperty(key string, value interface{}) *Machine {
	m.properties = append(m.properties, queso.NewProperty(key, value))
	return m
}

// SetName sets the name of the machine.
func (m *Machine) SetName(name string) *Machine {
	m.Name = name
	return m
}

// SetAccelerators specifies one or more accelerators to use for a Machine. If there
// is more than one accelerator specified, the next one is used if the previous
// one fails to initialize.
func (m *Machine) SetAccelerators(types ...string) *Machine {
	nameStrings := make([]string, 0)
	for _, at := range types {
		nameStrings = append(nameStrings, string(at))
	}

	m.properties = append(m.properties,
		queso.NewProperty("accel", strings.Join(nameStrings, ":")))

	return m
}

// SetVMWareIOPort sets emulation of VMWare IO port, for vmmouse etc. for a Machine.
// VMWareIOPortAuto says to select the value based on accel. For AcceleratorXen accelerator,
// the default is VMWareIOPortOff otherwise the default is VMWareIOPortOn.
func (m *Machine) SetVMWareIOPort(port VMWareIOPortFlag) *Machine {
	m.properties = append(m.properties, queso.NewProperty("vmport", port))
	return m
}

// ToggleDumpGuestCore specifies whether to include guest memory in a core dump for
// a Machine.
func (m *Machine) ToggleDumpGuestCore(enabled bool) *Machine {
	m.properties = append(m.properties, queso.NewProperty("dump-guest-core", enabled))
	return m
}

// ToggleMemoryMerge enables or disables memory merge support for a Machine. This
// feature, when supported by the host, de-duplicates identical memory pages
// among VMs instances. This property is enabled by default.
func (m *Machine) ToggleMemoryMerge(enabled bool) *Machine {
	m.properties = append(m.properties, queso.NewProperty("mem-merge", enabled))
	return m
}

// ToggleAESKeyWrap enables or disables AES key wrapping support on s390-ccw hosts for
// a Machine. This feature controls whether AES wrapping keys will be created to
// allow execution of AES cryptographic functions. This property is enabled by
// default.
func (m *Machine) ToggleAESKeyWrap(enabled bool) *Machine {
	m.properties = append(m.properties, queso.NewProperty("aes-key-wrap", enabled))
	return m
}

// ToggleDEAKeyWrap enables or disables DEA key wrapping support on s390-ccw hosts
// for a Machine. This feature controls whether DEA wrapping keys will be created
// to allow execution of DEA cryptographic functions. This property is enabled
// by default.
func (m *Machine) ToggleDEAKeyWrap(enabled bool) *Machine {
	m.properties = append(m.properties, queso.NewProperty("dea-key-wrap", enabled))
	return m
}

// ToggleNVDIMM enables or disables NVDIMM support for a Machine. This property is
// disabled by default.
func (m *Machine) ToggleNVDIMM(enabled bool) *Machine {
	m.properties = append(m.properties, queso.NewProperty("nvdimm", enabled))
	return m
}

// SetMemoryEncryption specifies the memory encryption object to use for a
// Machine.
func (m *Machine) SetMemoryEncryption(id string) *Machine {
	m.properties = append(m.properties, queso.NewProperty("memory-encryption", id))
	return m
}

// ToggleHMAT enables or disables ACPI Heterogeneous Memory Attribute Table (HMAT)
// support for a Machine. This property is disabled by default.
func (m *Machine) ToggleHMAT(enabled bool) *Machine {
	m.properties = append(m.properties, queso.NewProperty("hmat", enabled))
	return m
}

// SetMemoryBackend allows use of a memory backend as main RAM for a Machine.
func (m *Machine) SetMemoryBackend(id string) *Machine {
	m.properties = append(m.properties, queso.NewProperty("memory-backend", id))
	return m
}
