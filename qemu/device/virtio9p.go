package device

import (
	"fmt"

	"github.com/mikerourke/queso"
)

// Virtio9PVariant represents the variant of Virtio9P to use for a new [Virtio9P].
type Virtio9PVariant string

const (
	// Virtio9PCCW represents a virtual channel command word (CCW) device.
	Virtio9PCCW Virtio9PVariant = "ccw"

	// Virtio9PDevice represents a generic device.
	Virtio9PDevice Virtio9PVariant = "device"

	// Virtio9PPCI represents a virtual PCI device.
	Virtio9PPCI Virtio9PVariant = "pci"
)

// Virtio9P represents a device that uses virtio-9p.
type Virtio9P struct {
	// Variant specifies the variant to be used.
	Variant Virtio9PVariant

	// DeviceID is the ID value of the file system device.
	DeviceID string

	// MountTag is the tag name to be used by the guest to mount this export point.
	MountTag   string
	properties []*queso.Property
}

// NewVirtio9P returns a new instance of [Virtio9P]. variant specifies the variant
// to be used. deviceID is the ID value of the file system device. mountTag
// specifies the tag name to be used by the guest to mount this export point.
//
//	qemu-system-* -device virtio-9p-type,fsdev=id,mount_tag=mount_tag
func NewVirtio9P(variant Virtio9PVariant, deviceID string, mountTag string) *Virtio9P {
	return &Virtio9P{
		Variant:    variant,
		DeviceID:   deviceID,
		MountTag:   mountTag,
		properties: make([]*queso.Property, 0),
	}
}

func (d *Virtio9P) option() *queso.Option {
	d.properties = append(d.properties,
		queso.NewProperty("fsdev", d.DeviceID),
		queso.NewProperty("mount_tag", d.MountTag))

	return queso.NewOption("device",
		fmt.Sprintf("virtio-9p-%s", d.Variant), d.properties...)
}

// SetProperty is used to add arbitrary properties to the [Virtio9P].
func (d *Virtio9P) SetProperty(key string, value interface{}) *Virtio9P {
	d.properties = append(d.properties, queso.NewProperty(key, value))
	return d
}
