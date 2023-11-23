package fsdev

import (
	"fmt"

	"github.com/mikerourke/queso/qemu/cli"
)

// VirtIO9PType specifies the variant to be used for the virtio-9p machine type.
type VirtIO9PType string

const (
	VirtIO9PTypePCI    VirtIO9PType = "pci"
	VirtIO9PTypeCCW    VirtIO9PType = "ccw"
	VirtIO9PTypeDevice VirtIO9PType = "device"
)

// VirtIO9PDevice represents a device that uses virtio-9p.
type VirtIO9PDevice struct {
	// Type specifies the variant to be used.
	Type VirtIO9PType

	// DeviceID is the ID value of the file system device.
	DeviceID string

	// MountTag is the tag name to be used by the guest to mount this export point.
	MountTag string
}

// NewVirtIO9PDevice returns a new instance of [VirtIO9PDevice].
// typeOf specifies the variant to be used. deviceID is the ID value of the
// file system device. mountTag specifies the tag name to be used by the guest
// to mount this export point.
func NewVirtIO9PDevice(typeOf VirtIO9PType, deviceID string, mountTag string) *VirtIO9PDevice {
	return &VirtIO9PDevice{
		Type:     typeOf,
		DeviceID: deviceID,
		MountTag: mountTag,
	}
}

func (d *VirtIO9PDevice) option() *cli.Option {
	properties := []*cli.Property{
		cli.NewProperty("fsdev", d.DeviceID),
		cli.NewProperty("mount_tag", d.MountTag),
	}

	return cli.NewOption("device",
		fmt.Sprintf("virtio-9p-%s", d.Type), properties...)
}
