package qemu

import "github.com/mikerourke/queso"

// MonitorMode represents the monitor type. QEMU supports two monitors: the
// Human Monitor Protocol (HMP; for human interaction), and the QEMU Monitor
// Protocol (QMP; a JSON RPC-style protocol).
type MonitorMode string

const (
	// MonitorModeHMP indicates that the Human Monitor Protocol should be used.
	MonitorModeHMP MonitorMode = "readline"

	// MonitorModeQMP indicates that the QEMU Monitor Protocol should be used.
	MonitorModeQMP MonitorMode = "control"
)

// Monitor is used to set up a monitor connected to a character device.
type Monitor struct {
	Usable
	Name       string
	properties []*queso.Property
}

// NewMonitor returns a new instance of a Monitor, which can be used to monitor a
// character device.
func NewMonitor(name string) *Monitor {
	return &Monitor{
		Name:       name,
		properties: make([]*queso.Property, 0),
	}
}

func (m *Monitor) option() *queso.Option {
	table := queso.PropertiesTable(m.properties)

	mode := table["mode"]
	pretty := table["pretty"]
	if mode == string(MonitorModeHMP) {
		if pretty == "on" {
			panic("you can only enable pretty when mode is QMP")
		}
	}

	return queso.NewOption("mon", "", m.properties...)
}

// SetCharDevName sets the character device with the specified name to which
// the monitor is connected.
func (m *Monitor) SetCharDevName(name string) *Monitor {
	m.properties = append(m.properties, queso.NewProperty("chardev", name))
	return m
}

// SetMode sets the monitor mode to use. QEMU supports two monitors: the
// Human Monitor Protocol (HMP), and the QEMU Monitor Protocol (QMP).
func (m *Monitor) SetMode(mode MonitorMode) *Monitor {
	m.properties = append(m.properties, queso.NewProperty("mode", mode))
	return m
}

// TogglePretty sets the pretty option to "on". This option is only valid when
// MonitorMode = MonitorModeQMP, turning on JSON pretty printing to ease human
// reading and debugging.
func (m *Monitor) TogglePretty(enabled bool) *Monitor {
	m.properties = append(m.properties, queso.NewProperty("pretty", enabled))
	return m
}

// MonitorRedirect redirects the monitor to host device "device" (same devices as
// the serial port). The default device is vc in graphical mode and stdio in
// non-graphical mode. Use "none" for "device" to disable the default monitor.
func MonitorRedirect(device string) *queso.Option {
	return queso.NewOption("monitor", device)
}