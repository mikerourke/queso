// Package tpmdev is used to create TPM devices for use with QEMU. See
// https://qemu.readthedocs.io/en/latest/system/invocation.html#hxtool-7 for
// more details.
package tpmdev

import "github.com/mikerourke/queso"

// Backend represents a generic TPM backend device.
type Backend struct {
	// Type is the type of backend.
	Type       string
	properties []*queso.Property
}

// New returns a new TPM [Backend] instance with type specified.
func New(backendType string) *Backend {
	return &Backend{
		Type:       backendType,
		properties: make([]*queso.Property, 0),
	}
}

func (b *Backend) option() *queso.Option {
	return queso.NewOption("tpmdev", b.Type, b.properties...)
}

// SetProperty is used to add arbitrary properties to the [Backend].
func (b *Backend) SetProperty(key string, value interface{}) *Backend {
	b.properties = append(b.properties, queso.NewProperty(key, value))
	return b
}
