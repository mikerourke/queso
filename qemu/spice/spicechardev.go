package spice

import (
	"github.com/mikerourke/queso"
	"github.com/mikerourke/queso/qemu/chardev"
)

// PortBackend connects to a spice port, allowing a Spice client to handle the
// traffic identified by a name (preferably a fqdn)., and is only available when
// spice support is built in. The debugLevel parameter is the debug level. The name
// parameter is the name of spice channel to connect to.
func PortBackend(id string, name string, properties ...*ChardevProperty) *queso.Option {
	props := []*queso.Property{
		queso.NewProperty("id", id),
		queso.NewProperty("name", name),
	}

	for _, property := range properties {
		props = append(props, property.Property)
	}

	return queso.NewOption("chardev", chardev.BackendTypeSpicePort, props...)
}

// VMCBackend connects to a spice virtual machine channel, such as `vdiport`,
// and is only available when spice support is built in. The name parameter is
// the name of spice channel to connect to.
func VMCBackend(id string, name string, properties ...*ChardevProperty) *queso.Option {
	props := []*queso.Property{
		queso.NewProperty("id", id),
		queso.NewProperty("name", name),
	}

	for _, property := range properties {
		props = append(props, property.Property)
	}

	return queso.NewOption("chardev", chardev.BackendTypeSpiceVMC, props...)
}

// ChardevProperty represents a property to associate with a character device Backend.
type ChardevProperty struct {
	*queso.Property
}

// NewChardevProperty returns a new instance of Property.
func NewChardevProperty(key string, value interface{}) *ChardevProperty {
	return &ChardevProperty{
		Property: queso.NewProperty(key, value),
	}
}

// WithDebugLevel returns the debug flag to set for Spice port and VMC.
// TODO: Find out if this is supposed to be a string or a number or something else.
func WithDebugLevel(value string) *ChardevProperty {
	return NewChardevProperty("debug", value)
}
