// Package monitor corresponds to the options passable to the `-mon` flag.
package monitor

import "github.com/mikerourke/queso"

// Mode represents the monitor type. QEMU supports two monitors: the
// Human Monitor Protocol (HMP; for human interaction), and the QEMU Monitor
// Protocol (QMP; a JSON RPC-style protocol).
type Mode string

const (
	// ModeHMP indicates that the Human Monitor Protocol should be used.
	ModeHMP Mode = "readline"

	// ModeQMP indicates that the QEMU Monitor Protocol should be used.
	ModeQMP Mode = "control"
)

// Use returns a new instance of a monitor, which can be used to monitor a
// character device.
func Use(name string, properties ...*Property) *queso.Option {
	props := []*queso.Property{queso.NewProperty("chardev", name)}

	for _, property := range properties {
		props = append(props, property.Property)
	}

	table := queso.ToPropertiesTable(props)

	mode := table["mode"]
	pretty := table["pretty"]
	if mode == string(ModeHMP) {
		if pretty == "on" {
			panic("you can only enable pretty when mode is QMP")
		}
	}

	return queso.NewOption("mon", "", props...)
}

// Property represents a property to use with the Monitor option.
type Property struct {
	*queso.Property
}

// NewProperty returns a new instance of Property.
func NewProperty(key string, value interface{}) *Property {
	return &Property{
		Property: queso.NewProperty(key, value),
	}
}

// WithCharDevName sets the character device with the specified name to which
// the monitor is connected.
func WithCharDevName(name string) *Property {
	return NewProperty("chardev", name)
}

// WithMode sets the monitor mode to use. QEMU supports two monitors: the
// Human Monitor Protocol (ModeHMP), and the QEMU Monitor Protocol (ModeQMP).
func WithMode(mode Mode) *Property {
	return NewProperty("mode", mode)
}

// IsPretty sets the pretty option to "on". This option is only valid when Mode = ModeQMP,
// turning on JSON pretty printing to ease human reading and debugging.
func IsPretty(enabled bool) *Property {
	return NewProperty("pretty", enabled)
}
