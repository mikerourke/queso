// Package device is used to specify device drivers for use with QEMU.
package device

import (
	"github.com/mikerourke/queso"
)

// Device represents a device used with QEMU.
type Device struct {
	Type       string
	properties []*queso.Property
}

// New returns a new instance of a [Device].
//
//	qemu-system-* -device <deviceType>
func New(deviceType string) *Device {
	return &Device{
		Type:       deviceType,
		properties: make([]*queso.Property, 0),
	}
}

// Option returns the option representation of the device.
func (d *Device) option() *queso.Option {
	return queso.NewOption("device", d.Type, d.properties...)
}

// SetProperty is used to add arbitrary properties to the [Device].
func (d *Device) SetProperty(key string, value interface{}) *Device {
	d.properties = append(d.properties, queso.NewProperty(key, value))
	return d
}
