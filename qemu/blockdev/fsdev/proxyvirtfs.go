package fsdev

import "github.com/mikerourke/queso"

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
	properties          []*queso.Property
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
		properties:          make([]*queso.Property, 0),
	}
}

func (d *VirtualProxyFileSystemDevice) option() *queso.Option {
	properties := append(d.properties,
		queso.NewProperty("mount_tag", d.MountTag),
		queso.NewProperty(string(d.SocketInterfaceType), d.SocketTarget))
	return queso.NewOption("virtfs", "proxy", properties...)
}

// SetProperty is used to add arbitrary properties to the [VirtualProxyFileSystemDevice].
func (d *VirtualProxyFileSystemDevice) SetProperty(key string, value interface{}) *VirtualProxyFileSystemDevice {
	d.properties = append(d.properties, queso.NewProperty(key, value))
	return d
}

// EnableWriteOut means that host page cache will be used to read and write data but
// write notification will be sent to the guest only when the data has been reported
// as written by the storage subsystem.
//
//	qemu-system-* -virtfs proxy,writeout=writeout
func (d *VirtualProxyFileSystemDevice) EnableWriteOut() *VirtualProxyFileSystemDevice {
	d.properties = append(d.properties,
		// The only supported value is "immediate".
		queso.NewProperty("writeout", "immediate"))
	return d
}

// SetMountTag specifies the tag name to be used by the guest to mount this export point.
//
//	qemu-system-* -virtfs proxy,mount_tag=tag
func (d *VirtualProxyFileSystemDevice) SetMountTag(tag string) *VirtualProxyFileSystemDevice {
	d.MountTag = tag
	return d
}

// SetSocketTarget sets the socket path or socket descriptor path (based
// on the [SocketInterfaceType] specified when the device was created).
//
//	qemu-system-* -virtfs proxy,socket=target
//	qemu-system-* -virtfs proxy,sock_fd=target
func (d *VirtualProxyFileSystemDevice) SetSocketTarget(target string) *VirtualProxyFileSystemDevice {
	d.SocketTarget = target
	return d
}

// ToggleReadOnly enables exporting 9P share as a readonly mount for guests.
// By default, read-write access is given.
//
//	qemu-system-* -virtfs proxy,readonly=on|off
func (d *VirtualProxyFileSystemDevice) ToggleReadOnly(enabled bool) *VirtualProxyFileSystemDevice {
	d.properties = append(d.properties, queso.NewProperty("readonly", enabled))
	return d
}
