package fsdev

import "github.com/mikerourke/queso"

// VirtualProxyFileSystemDevice represents a virtual file system device in which
// accesses to the filesystem are done by virtfs-proxy-helper(1).
//
// Deprecated: Use [VirtualLocalFileSystemDevice] instead.
type VirtualProxyFileSystemDevice struct {
	*queso.Entity
	socketInterfaceType SocketInterfaceType
}

// NewVirtualProxyFileSystemDevice returns a new instance of [VirtualProxyFileSystemDevice].
// mountTag is the tag name to be used by the guest to mount this export point.
// socketTarget is the path or file descriptor of the socket. socketInterfaceType
// is the type of socket interface to use (either path or file descriptor).
//
//	qemu-system-* -virtfs proxy,id=id,socket=target
//	qemu-system-* -virtfs proxy,id=id,sock_fd=target
//
// Deprecated: Use [NewVirtualLocalFileSystemDevice] instead.
func NewVirtualProxyFileSystemDevice(
	mountTag string,
	socketTarget string,
	socketInterfaceType SocketInterfaceType,
) *VirtualProxyFileSystemDevice {
	return &VirtualProxyFileSystemDevice{
		queso.NewEntity("virtfs", "proxy").
			SetProperty("mount_tag", mountTag).
			SetProperty(string(socketInterfaceType), socketTarget),
		socketInterfaceType,
	}
}

// EnableWriteOut means that host page cache will be used to read and write data but
// write notification will be sent to the guest only when the data has been reported
// as written by the storage subsystem.
//
//	qemu-system-* -virtfs proxy,writeout=writeout
func (d *VirtualProxyFileSystemDevice) EnableWriteOut() *VirtualProxyFileSystemDevice {
	// The only supported value is "immediate".
	d.SetProperty("writeout", "immediate")
	return d
}

// SetMountTag specifies the tag name to be used by the guest to mount this export point.
//
//	qemu-system-* -virtfs proxy,mount_tag=tag
func (d *VirtualProxyFileSystemDevice) SetMountTag(tag string) *VirtualProxyFileSystemDevice {
	d.UpsertProperty("mount_tag", tag)
	return d
}

// SetSocketTarget sets the socket path or socket descriptor path (based
// on the [SocketInterfaceType] specified when the device was created).
//
//	qemu-system-* -virtfs proxy,socket=target
//	qemu-system-* -virtfs proxy,sock_fd=target
func (d *VirtualProxyFileSystemDevice) SetSocketTarget(target string) *VirtualProxyFileSystemDevice {
	d.UpsertProperty(string(d.socketInterfaceType), target)
	return d
}

// ToggleReadOnly enables exporting 9P share as a readonly mount for guests.
// By default, read-write access is given.
//
//	qemu-system-* -virtfs proxy,readonly=on|off
func (d *VirtualProxyFileSystemDevice) ToggleReadOnly(enabled bool) *VirtualProxyFileSystemDevice {
	d.SetProperty("readonly", enabled)
	return d
}
