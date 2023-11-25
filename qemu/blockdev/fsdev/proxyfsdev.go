package fsdev

import "github.com/mikerourke/queso"

// ProxyFileSystemDevice represents a file system device in which accesses to the
// filesystem are done by virtfs-proxy-helper(1).
//
// Deprecated: Use [LocalFileSystemDevice] instead.
type ProxyFileSystemDevice struct {
	// ID is the unique identifier for the device.
	ID string

	// SocketTarget is the path or file descriptor of the socket.
	SocketTarget string

	// SocketInterfaceType is the type of socket interface to use (either path
	// or file descriptor).
	SocketInterfaceType SocketInterfaceType
	properties          []*queso.Property
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
		ID:                  id,
		SocketTarget:        socketTarget,
		SocketInterfaceType: socketInterfaceType,
		properties:          make([]*queso.Property, 0),
	}
}

func (d *ProxyFileSystemDevice) option() *queso.Option {
	properties := append(d.properties,
		queso.NewProperty("id", d.ID),
		queso.NewProperty(string(d.SocketInterfaceType), d.SocketTarget))
	return queso.NewOption("fsdev", "proxy", properties...)
}

// SetProperty is used to add arbitrary properties to the [ProxyFileSystemDevice].
func (d *ProxyFileSystemDevice) SetProperty(key string, value interface{}) *ProxyFileSystemDevice {
	d.properties = append(d.properties, queso.NewProperty(key, value))
	return d
}

// EnableWriteOut means that host page cache will be used to read and write data but
// write notification will be sent to the guest only when the data has been reported
// as written by the storage subsystem.
//
//	qemu-system-* -fsdev proxy,writeout=writeout
func (d *ProxyFileSystemDevice) EnableWriteOut() *ProxyFileSystemDevice {
	d.properties = append(d.properties,
		// The only supported value is "immediate".
		queso.NewProperty("writeout", "immediate"))
	return d
}

// SetID specifies the identifier for this device.
//
//	qemu-system-* -fsdev proxy,id=id
func (d *ProxyFileSystemDevice) SetID(id string) *ProxyFileSystemDevice {
	d.ID = id
	return d
}

// SetSocketTarget sets the socket path or socket descriptor path (based
// on the [SocketInterfaceType] specified when the device was created).
//
//	qemu-system-* -fsdev proxy,socket=target
//	qemu-system-* -fsdev proxy,sock_fd=target
func (d *ProxyFileSystemDevice) SetSocketTarget(target string) *ProxyFileSystemDevice {
	d.SocketTarget = target
	return d
}

// ToggleReadOnly enables exporting 9P share as a readonly mount for guests.
// By default, read-write access is given.
//
//	qemu-system-* -fsdev proxy,readonly=on|off
func (d *ProxyFileSystemDevice) ToggleReadOnly(enabled bool) *ProxyFileSystemDevice {
	d.properties = append(d.properties, queso.NewProperty("readonly", enabled))
	return d
}
