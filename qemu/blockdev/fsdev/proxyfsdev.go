package fsdev

import "github.com/mikerourke/queso"

// ProxyFileSystemDevice represents a file system device in which accesses to the
// filesystem are done by virtfs-proxy-helper(1).
//
// Deprecated: Use [LocalFileSystemDevice] instead.
type ProxyFileSystemDevice struct {
	*queso.Entity
	SocketInterfaceType SocketInterfaceType
}

// NewProxyFileSystemDevice returns a new instance of [ProxyFileSystemDevice].
// id is a unique identifier for the device. socketTarget is the path
// or file descriptor of the socket. socketInterfaceType is the type of
// socket interface to use (either path or file descriptor).
//
//	qemu-system-* -fsdev proxy,id=id,socket=target
//	qemu-system-* -fsdev proxy,id=id,sock_fd=target
//
// Deprecated: Use [NewLocalFileSystemDevice] instead.
func NewProxyFileSystemDevice(
	id string,
	socketTarget string,
	socketInterfaceType SocketInterfaceType,
) *ProxyFileSystemDevice {
	return &ProxyFileSystemDevice{
		queso.NewEntity("fsdev", "proxy").
			SetProperty("id", id).
			SetProperty(string(socketInterfaceType), socketTarget),
		socketInterfaceType,
	}
}

// EnableWriteOut means that host page cache will be used to read and write data but
// write notification will be sent to the guest only when the data has been reported
// as written by the storage subsystem.
//
//	qemu-system-* -fsdev proxy,writeout=writeout
func (d *ProxyFileSystemDevice) EnableWriteOut() *ProxyFileSystemDevice {
	// The only supported value is "immediate".
	d.SetProperty("writeout", "immediate")
	return d
}

// SetID specifies the identifier for this device.
//
//	qemu-system-* -fsdev proxy,id=id
func (d *ProxyFileSystemDevice) SetID(id string) *ProxyFileSystemDevice {
	d.UpsertProperty("id", id)
	return d
}

// SetSocketTarget sets the socket path or socket descriptor path (based
// on the [SocketInterfaceType] specified when the device was created).
//
//	qemu-system-* -fsdev proxy,socket=target
//	qemu-system-* -fsdev proxy,sock_fd=target
func (d *ProxyFileSystemDevice) SetSocketTarget(target string) *ProxyFileSystemDevice {
	d.UpsertProperty(string(d.SocketInterfaceType), target)
	return d
}

// ToggleReadOnly enables exporting 9P share as a readonly mount for guests.
// By default, read-write access is given.
//
//	qemu-system-* -fsdev proxy,readonly=on|off
func (d *ProxyFileSystemDevice) ToggleReadOnly(enabled bool) *ProxyFileSystemDevice {
	d.SetProperty("readonly", enabled)
	return d
}
