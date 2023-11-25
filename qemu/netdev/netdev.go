package netdev

import "github.com/mikerourke/queso"

// Backend represents a generic network backend.
type Backend struct {
	// Type is the type of the network backend.
	Type       string
	properties []*queso.Property
}

// New returns a new instance of [Backend]. backendType is the type of backend
// to create.
//
//	qemu-system-* -netdev <backendType>
func New(backendType string) *Backend {
	return &Backend{
		Type:       backendType,
		properties: make([]*queso.Property, 0),
	}
}

func (b *Backend) option() *queso.Option {
	return queso.NewOption("netdev", b.Type, b.properties...)
}

// SetProperty is used to add arbitrary properties to the [Backend].
func (b *Backend) SetProperty(key string, value interface{}) *Backend {
	b.properties = append(b.properties, queso.NewProperty(key, value))
	return b
}
