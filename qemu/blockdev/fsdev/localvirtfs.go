package fsdev

import "github.com/mikerourke/queso/qemu/cli"

// VirtualLocalFileSystemDevice represents a virtual file system device in which
// accesses to the filesystem are done by QEMU.
type VirtualLocalFileSystemDevice struct {
	// MountTag is the tag name to be used by the guest to mount this export point.
	MountTag string

	// Path is the export path for the file system device. Files under this path
	// will be available to the 9P client on the guest.
	Path string

	// SecurityModel is used to specify the security model to be used for the export
	// path in file system devices.
	SecurityModel SecurityModel
	properties    []*cli.Property
}

// NewVirtualLocalFileSystemDevice returns a new instance of [VirtualLocalFileSystemDevice].
// mountTag is the tag name to be used by the guest to mount this export point.
// path is the export path for the file system device. Files under this path
// will be available to the 9P client on the guest. model is the [SecurityModel]
// for the file system device.
//
//	qemu-system-* -virtfs local,mount_tag=tag,path=path,security_model=model
func NewVirtualLocalFileSystemDevice(
	mountTag string,
	path string,
	model SecurityModel,
) *VirtualLocalFileSystemDevice {
	return &VirtualLocalFileSystemDevice{
		MountTag:      mountTag,
		Path:          path,
		SecurityModel: model,
		properties:    make([]*cli.Property, 0),
	}
}

func (d *VirtualLocalFileSystemDevice) option() *cli.Option {
	properties := append(d.properties,
		cli.NewProperty("mount_tag", d.MountTag),
		cli.NewProperty("path", d.Path),
		cli.NewProperty("security_model", d.SecurityModel))
	return cli.NewOption("virtfs", "local", properties...)
}

// EnableWriteOut means that host page cache will be used to read and write data but
// write notification will be sent to the guest only when the data has been reported
// as written by the storage subsystem.
//
//	qemu-system-* -virtfs local,writeout=writeout
func (d *VirtualLocalFileSystemDevice) EnableWriteOut() *VirtualLocalFileSystemDevice {
	d.properties = append(d.properties,
		// The only supported value is "immediate".
		cli.NewProperty("writeout", "immediate"))
	return d
}

// SetDirectoryMode specifies the default mode for newly created directories on the
// host. Works only with [SecurityModel] set to [SecurityModelMappedXAttr] and
// [SecurityModelMappedFile].
//
//	qemu-system-* -virtfs local,dmode=mode
func (d *VirtualLocalFileSystemDevice) SetDirectoryMode(mode string) *VirtualLocalFileSystemDevice {
	d.properties = append(d.properties, cli.NewProperty("dmode", mode))
	return d
}

// SetFileMode specifies the default mode for newly created files on the
// host. Works only with [SecurityModel] set to [SecurityModelMappedXAttr] and
// [SecurityModelMappedFile].
//
//	qemu-system-* -virtfs local,fmode=mode
func (d *VirtualLocalFileSystemDevice) SetFileMode(mode string) *VirtualLocalFileSystemDevice {
	d.properties = append(d.properties, cli.NewProperty("fmode", mode))
	return d
}

// SetMountTag specifies the tag name to be used by the guest to mount this export point.
//
//	qemu-system-* -virtfs local,mount_tag=tag
func (d *VirtualLocalFileSystemDevice) SetMountTag(tag string) *VirtualLocalFileSystemDevice {
	d.MountTag = tag
	return d
}

// MultiDeviceSharing is used to specify how to deal with multiple devices being
// shared with a 9P export.
type MultiDeviceSharing string

const (
	// MultiDeviceSharingRemap allows you to share multiple devices with only one
	// export instead, which is achieved by remapping the original inode numbers from
	// host to guest in a way that would prevent such collisions.
	//
	// Remapping inodes in such use cases is required because the original device
	// IDs from host are never passed and exposed on guest. Instead, all files of
	// an export shared with virtfs always share the same device id on guest.
	// So two files with identical inode numbers but from actually different devices
	// on host would otherwise cause a file ID collision and hence potential
	// misbehaviours on guest.
	MultiDeviceSharingRemap MultiDeviceSharing = "remap"

	// MultiDeviceSharingForbid assumes like MultiDeviceSharingWarn that only one device
	// is shared by the same export, however it will not only log a warning message
	// but also deny access to additional devices on guest. Note though that this
	// option does currently not block all possible file access operations
	// (e.g. readdir() would still return entries from other devices).
	MultiDeviceSharingForbid MultiDeviceSharing = "forbid"

	// MultiDeviceSharingWarn should be used when virtfs 9P expects only one device
	// to be shared with the same export, and if more than one device is shared
	// and accessed via the same 9P export then only a warning message is logged
	// (once) by QEMU on host side.
	MultiDeviceSharingWarn MultiDeviceSharing = "warn"
)

// SetMultiDeviceSharing specifies how to deal with multiple devices being shared with a
// 9P export.
//
//	qemu-system-* -virtfs local,multidevs=sharing
func (d *VirtualLocalFileSystemDevice) SetMultiDeviceSharing(sharing MultiDeviceSharing) *VirtualLocalFileSystemDevice {
	d.properties = append(d.properties, cli.NewProperty("multidevs", string(sharing)))
	return d
}

// SetPath specifies the export path for the file system device. Files under
// this path will be available to the 9P client on the guest.
//
//	qemu-system-* -virtfs local,path=path
func (d *VirtualLocalFileSystemDevice) SetPath(path string) *VirtualLocalFileSystemDevice {
	d.Path = path
	return d
}

// SetSecurityModel specifies the security model to be used for this export path.
// See [SecurityModel] for additional details.
//
//	qemu-system-* -virtfs local,security_model=model
func (d *VirtualLocalFileSystemDevice) SetSecurityModel(model SecurityModel) *VirtualLocalFileSystemDevice {
	d.SecurityModel = model
	return d
}

// ToggleReadOnly enables exporting 9P share as a readonly mount for guests.
// By default, read-write access is given.
//
//	qemu-system-* -virtfs local,readonly=on|off
func (d *VirtualLocalFileSystemDevice) ToggleReadOnly(enabled bool) *VirtualLocalFileSystemDevice {
	d.properties = append(d.properties, cli.NewProperty("readonly", enabled))
	return d
}
