package debug

import "github.com/mikerourke/queso"

// Plugin represents a plugin loaded from a shared library file.
type Plugin struct {
	// File is the path to the shared library file.
	File       string
	properties []*queso.Property
}

// NewPlugin returns a new instance of [Plugin].
func NewPlugin(file string) *Plugin {
	return &Plugin{
		File:       file,
		properties: make([]*queso.Property, 0),
	}
}

// Option returns the invoked option that gets converted to an argument when
// passed to QEMU.
func (p *Plugin) Option() *queso.Option {
	properties := append(p.properties, queso.NewProperty("file", p.File))
	return queso.NewOption("plugin", "", properties...)
}

// SetProperty can be used to set arbitrary properties on the [Plugin].
func (p *Plugin) SetProperty(name string, value interface{}) *Plugin {
	p.properties = append(p.properties, queso.NewProperty(name, value))
	return p
}
