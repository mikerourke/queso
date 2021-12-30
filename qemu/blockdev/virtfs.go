package blockdev

import "github.com/mikerourke/queso"

// VirtualFileSystem defines a new virtual filesystem. This is actually just a
// convenience shortcut for creating a file system device with the
// Virtio9P(Virtio9PPCI, ...) option (see qemu/device.go).
func VirtualFileSystem(
	deviceType string,
	mountTag string,
	properties ...*FileSystemDeviceProperty,
) *queso.Option {
	props := []*queso.Property{{"mount_tag", mountTag}}

	for _, property := range properties {
		props = append(props, property.Property)
	}

	return queso.NewOption("virtfs", deviceType, props...)
}

// LocalVirtualFileSystem defines a virtual filesystem for which accesses to the
// filesystem are done by QEMU. The exportPath parameter represents the export path
// for the filesystem. Files under this path will be available to the 9p client on
// the guest. The securityModel parameter specifies the security model to be used
// for the export path.
func LocalVirtualFileSystem(
	mountTag string,
	exportPath string,
	securityModel SecurityModel,
	properties ...*FileSystemDeviceProperty,
) *queso.Option {
	props := []*FileSystemDeviceProperty{
		NewFileSystemDeviceProperty("path", exportPath),
		NewFileSystemDeviceProperty("security_model", securityModel),
	}

	if properties != nil {
		props = append(props, properties...)
	}

	return VirtualFileSystem("local", mountTag, props...)
}

// ProxyVirtualFileSystem defines a virtual filesystem for which accesses to the
// filesystem are done by virtfs-proxy-helper. The socketType parameter specifies
// whether a file path or file descriptor should be used. The socketFileOrDescriptor
// parameter is the file path or fd to use (based on the socketType).
func ProxyVirtualFileSystem(
	mountTag string,
	socketType SocketType,
	socketFileOrDescriptor string,
	properties ...*FileSystemDeviceProperty,
) *queso.Option {
	props := []*FileSystemDeviceProperty{
		NewFileSystemDeviceProperty(string(socketType), socketFileOrDescriptor),
	}

	if properties != nil {
		props = append(props, properties...)
	}

	return VirtualFileSystem("proxy", mountTag, props...)
}

// SyntheticVirtualFileSystem defines a synthetic filesystem, which is only used by
// QTests.
func SyntheticVirtualFileSystem(mountTag string) *queso.Option {
	return VirtualFileSystem("synth", mountTag)
}

// MultiDevicesAction is used to specify how to deal with multiple devices. See
// the WithMultiDevicesAction for more details.
type MultiDevicesAction string

const (
	MultiDevicesActionForbid MultiDevicesAction = "forbid"
	MultiDevicesActionRemap  MultiDevicesAction = "remap"
	MultiDevicesActionWarn   MultiDevicesAction = "warn"
)

// WithMultiDevicesAction specifies how to deal with multiple devices being shared
// with a 9p export. Supported behaviours are either MultiDevicesActionRemap,
// MultiDevicesActionForbid, or MultiDevicesActionWarn. The latter is the default
// behaviour on which virtfs 9p expects only one device to be shared with the same
// export, and if more than one device is shared and accessed via the same 9p export
// then only a warning message is logged (once) by QEMU on host side.
//
// In order to avoid file ID collisions on guest you should either create a separate
// virtfs export for each device to be shared with guests (recommended way) or you
// might use MultiDevicesActionRemap instead which allows you to share multiple devices
// with only one export instead, which is achieved by remapping the original inode
// numbers from host to guest in a way that would prevent such collisions. Remapping
// inodes in such use cases is required because the original device IDs from host
// are never passed and exposed on guest. Instead, all files of an export shared
// with virtfs always share the same device id on guest. So two files with identical
// inode numbers but from actually different devices on host would otherwise cause
// a file ID collision and hence potential misbehaviour on guest.
//
// MultiDevicesActionForbid on the other hand assumes like MultiDevicesActionWarn
// that only one device is shared by the same export, however it will not only log
// a warning message but also deny access to additional devices on guest. Note
// though that MultiDevicesActionForbid does currently not block all possible file
// access operations (e.g. readdir() would still return entries from other devices).
func WithMultiDevicesAction(action MultiDevicesAction) *FileSystemDeviceProperty {
	return NewFileSystemDeviceProperty("multidevs", action)
}
