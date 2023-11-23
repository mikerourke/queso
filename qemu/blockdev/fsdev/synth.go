package fsdev

import "github.com/mikerourke/queso/qemu/cli"

// SyntheticFileSystemDevice represents a synthetic file system, only used by
// QTests.
type SyntheticFileSystemDevice struct {
	// ID is the unique identifier for the device.
	ID         string
	properties []*cli.Property
}

// NewSyntheticFileSystemDevice returns a new instance of [SyntheticFileSystemDevice].
// id is a unique identifier for the device.
//
//	qemu-system-* -fsdev synth,id=id
func NewSyntheticFileSystemDevice(id string) *SyntheticFileSystemDevice {
	return &SyntheticFileSystemDevice{
		ID:         id,
		properties: make([]*cli.Property, 0),
	}
}

// SetID specifies the identifier for this device.
//
//	qemu-system-* -fsdev synth,id=id
func (d *SyntheticFileSystemDevice) SetID(id string) *SyntheticFileSystemDevice {
	d.ID = id
	return d
}

// ToggleReadOnly enables exporting 9P share as a readonly mount for guests.
// By default, read-write access is given.
//
//	qemu-system-* -fsdev synth,readonly=on|off
func (d *SyntheticFileSystemDevice) ToggleReadOnly(enabled bool) *SyntheticFileSystemDevice {
	d.properties = append(d.properties, cli.NewProperty("readonly", enabled))
	return d
}

// VirtualSyntheticFileSystemDevice represents a virtual synthetic file system,
// only used by QTests.
type VirtualSyntheticFileSystemDevice struct {
	// MountTag is the tag name to be used by the guest to mount this export point.
	MountTag   string
	properties []*cli.Property
}

// NewVirtualSyntheticFileSystemDevice returns a new instance of [VirtualSyntheticFileSystemDevice].
// mountTag is the tag name to be used by the guest to mount this export point.
//
//	qemu-system-* -virtfs synth,mount_tag=tag
func NewVirtualSyntheticFileSystemDevice(mountTag string) *VirtualSyntheticFileSystemDevice {
	return &VirtualSyntheticFileSystemDevice{
		MountTag:   mountTag,
		properties: make([]*cli.Property, 0),
	}
}

// SetMountTag specifies the tag name to be used by the guest to mount this export point.
//
//	qemu-system-* -virtfs synth,mount_tag=tag
func (d *VirtualSyntheticFileSystemDevice) SetMountTag(tag string) *VirtualSyntheticFileSystemDevice {
	d.MountTag = tag
	return d
}
