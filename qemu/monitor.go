package qemu

import "github.com/mikerourke/queso"

// MonitorMode represents the monitor type. QEMU supports two monitors: the
// Human Monitor Protocol (HMP; for human interaction), and the QEMU Monitor
// Protocol (QMP; a JSON RPC-style protocol).
type MonitorMode string

const (
	// MonitorHMP indicates that the Human Monitor Protocol should be used.
	MonitorHMP MonitorMode = "readline"

	// MonitorQMP indicates that the QEMU Monitor Protocol should be used.
	MonitorQMP MonitorMode = "control"
)

// Monitor is used to set up a monitor connected to a character device.
type Monitor struct {
	// Name is the name of the character device to monitor.
	Name       string
	properties []*queso.Property
}

// NewMonitor returns a new instance of a [Monitor], which can be used to monitor
// a character device. name is the name of the character device to monitor.
//
//	qemu-system-* -mon <name>
func NewMonitor(name string) *Monitor {
	return &Monitor{
		Name:       name,
		properties: make([]*queso.Property, 0),
	}
}

// Option returns the invoked option that gets converted to an argument when
// passed to QEMU.
func (m *Monitor) Option() *queso.Option {
	table := queso.PropertiesTable(m.properties)

	mode := table["mode"]
	pretty := table["pretty"]
	if mode == string(MonitorHMP) {
		if pretty == "on" {
			panic("you can only enable pretty when mode is QMP")
		}
	}

	return queso.NewOption("mon", m.Name, m.properties...)
}

// SetMonitorMode sets the monitor mode to use. QEMU supports two monitors: the
// Human Monitor Protocol ([MonitorHMP]), and the QEMU Monitor Protocol ([MonitorQMP]).
//
//	qemu-system-* -mon <name>,mode=mode
func (m *Monitor) SetMonitorMode(mode MonitorMode) *Monitor {
	m.properties = append(m.properties, queso.NewProperty("mode", mode))
	return m
}

// TogglePretty sets the pretty option to "on". This option is only valid when
// MonitorMode = MonitorQMP, turning on JSON pretty printing to ease human
// reading and debugging.
//
//	qemu-system-* -mon <name>,pretty=on|off
func (m *Monitor) TogglePretty(enabled bool) *Monitor {
	m.properties = append(m.properties, queso.NewProperty("pretty", enabled))
	return m
}

// WithMonitorRedirect redirects the monitor to host device "device" (same devices
// as the serial port). The default device is "vc" in graphical mode and "stdio" in
// non-graphical mode. Use "none" for "device" to disable the default monitor.
//
//	qemu-system-* -monitor <device>
func WithMonitorRedirect(device string) *queso.Option {
	return queso.NewOption("monitor", device)
}
