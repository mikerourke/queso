// Package chardev is used to add character device backends.
package chardev

import "github.com/mikerourke/queso"

// Backend represents a generic character device backend of any allowed type.
type Backend struct {
	// Type represents the type of the character device backend.
	Type string

	// ID is the unique ID, which can be any string up to 127 characters long.
	ID         string
	properties []*queso.Property
}

// New returns a new instance of [Backend]. backendType is the type of
// backend, this is a string rather than a constant to allow for future backend
// types. id is the unique ID, which can be any string up to 127 characters long.
//
//	qemu-system-* -chardev <backendType>,id=id
func New(backendType string, id string) *Backend {
	return &Backend{
		Type:       backendType,
		ID:         id,
		properties: make([]*queso.Property, 0),
	}
}

func (b *Backend) option() *queso.Option {
	properties := append(b.properties, queso.NewProperty("id", b.ID))

	return queso.NewOption("chardev", b.Type, properties...)
}

// SetProperty can be used to set arbitrary properties on the [Backend].
func (b *Backend) SetProperty(name string, value interface{}) *Backend {
	b.properties = append(b.properties, queso.NewProperty(name, value))
	return b
}

// ToggleMultiplexing enables or disables the character device to be used in multiplexing
// mode by multiple front-ends. Specify true to enable this mode. A multiplexer is a "1:N"
// device, and here the "1" end is your specified chardev backend, and the "N" end is
// the various parts of QEMU that can talk to a character device. If you create a
// character device with ID of "myid" and multiplexing enabled, QEMU will create a
// multiplexer with your specified ID, and you can then configure multiple front
// ends to use that character device ID for their input/output.
//
// Up to four different front ends can be connected to a single multiplexed character
// device. (Without multiplexing enabled, a character device can only be used by a
// single front end.)
//
// For instance, you could use this to allow a single stdio character device to be
// used by two serial ports and the QEMU monitor:
//
//	TODO: Add examples using Queso rather than QEMU.
func (b *Backend) ToggleMultiplexing(enabled bool) *Backend {
	b.properties = append(b.properties, queso.NewProperty("mux", enabled))
	return b
}
