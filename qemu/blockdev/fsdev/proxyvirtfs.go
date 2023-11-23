package fsdev

import "github.com/mikerourke/queso/qemu/cli"

// VirtualProxyFileSystemDevice represents a virtual file system device in which
// accesses to the filesystem are done by virtfs-proxy-helper(1).
//
// Deprecated: Use [VirtualLocalFileSystemDevice] instead.
type VirtualProxyFileSystemDevice struct {
	// MountTag is the tag name to be used by the guest to mount this export point.
	MountTag string

	// SocketTarget is the path or file descriptor of the socket.
	SocketTarget string

	// SocketInterfaceType is the type of socket interface to use (either path
	// or file descriptor).
	SocketInterfaceType SocketInterfaceType
	properties          []*cli.Property
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
		MountTag:            mountTag,
		SocketTarget:        socketTarget,
		SocketInterfaceType: socketInterfaceType,
		properties:          make([]*cli.Property, 0),
	}
}

// EnableWriteOut means that host page cache will be used to read and write data but
// write notification will be sent to the guest only when the data has been reported
// as written by the storage subsystem.
//
//	qemu-system-* -virtfs proxy,writeout=writeout
func (d *VirtualProxyFileSystemDevice) EnableWriteOut() *VirtualProxyFileSystemDevice {
	d.properties = append(d.properties,
		// The only supported value is "immediate".
		cli.NewProperty("writeout", "immediate"))
	return d
}

// SetMountTag specifies the tag name to be used by the guest to mount this export point.
//
//	qemu-system-* -virtfs proxy,mount_tag=tag
func (d *VirtualProxyFileSystemDevice) SetMountTag(tag string) *VirtualProxyFileSystemDevice {
	d.properties = append(d.properties, cli.NewProperty("mount_tag", tag))
	return d
}

// ToggleReadOnly enables exporting 9P share as a readonly mount for guests.
// By default, read-write access is given.
//
//	qemu-system-* -virtfs proxy,readonly=on|off
func (d *VirtualProxyFileSystemDevice) ToggleReadOnly(enabled bool) *VirtualProxyFileSystemDevice {
	d.properties = append(d.properties, cli.NewProperty("readonly", enabled))
	return d
}
