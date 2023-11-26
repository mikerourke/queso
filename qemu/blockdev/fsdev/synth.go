package fsdev

import "github.com/mikerourke/queso"

// SyntheticFileSystemDevice represents a synthetic file system, only used by
// QTests.
type SyntheticFileSystemDevice struct {
	*queso.Entity
}

// NewSyntheticFileSystemDevice returns a new instance of [SyntheticFileSystemDevice].
// id is a unique identifier for the device.
//
//	qemu-system-* -fsdev synth,id=id
func NewSyntheticFileSystemDevice(id string) *SyntheticFileSystemDevice {
	return &SyntheticFileSystemDevice{
		queso.NewEntity("fsdev", "synth"),
	}
}

// SetID specifies the identifier for this device.
//
//	qemu-system-* -fsdev synth,id=id
func (d *SyntheticFileSystemDevice) SetID(id string) *SyntheticFileSystemDevice {
	d.SetProperty("id", id)
	return d
}

// ToggleReadOnly enables exporting 9P share as a readonly mount for guests.
// By default, read-write access is given.
//
//	qemu-system-* -fsdev synth,readonly=on|off
func (d *SyntheticFileSystemDevice) ToggleReadOnly(enabled bool) *SyntheticFileSystemDevice {
	d.SetProperty("readonly", enabled)
	return d
}

// VirtualSyntheticFileSystemDevice represents a virtual synthetic file system,
// only used by QTests.
type VirtualSyntheticFileSystemDevice struct {
	*queso.Entity
}

// NewVirtualSyntheticFileSystemDevice returns a new instance of [VirtualSyntheticFileSystemDevice].
// mountTag is the tag name to be used by the guest to mount this export point.
//
//	qemu-system-* -virtfs synth,mount_tag=tag
func NewVirtualSyntheticFileSystemDevice(mountTag string) *VirtualSyntheticFileSystemDevice {
	return &VirtualSyntheticFileSystemDevice{
		queso.NewEntity("virtfs", "synth"),
	}
}

// SetMountTag specifies the tag name to be used by the guest to mount this export point.
//
//	qemu-system-* -virtfs synth,mount_tag=tag
func (d *VirtualSyntheticFileSystemDevice) SetMountTag(tag string) *VirtualSyntheticFileSystemDevice {
	d.SetProperty("mount_tag", tag)
	return d
}
