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

type Monitor struct {
	Usable
	Name       string
	properties []*queso.Property
}

// NewMonitor returns a new instance of a monitor, which can be used to monitor a
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

func Monitor1(name string, properties ...*MonitorProperty) *queso.Option {
	props := []*queso.Property{queso.NewProperty("chardev", name)}

	for _, property := range properties {
		props = append(props, property.Property)
	}

	table := queso.PropertiesTable(props)

	mode := table["mode"]
	pretty := table["pretty"]
	if mode == string(MonitorModeHMP) {
		if pretty == "on" {
			panic("you can only enable pretty when mode is QMP")
		}
	}

	return queso.NewOption("mon", "", props...)
}
