package fsdev

import "github.com/mikerourke/queso/qemu/cli"

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
	properties          []*cli.Property
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
		properties:          make([]*cli.Property, 0),
	}
}

func (d *ProxyFileSystemDevice) option() *cli.Option {
	properties := append(d.properties,
		cli.NewProperty("id", d.ID),
		cli.NewProperty(string(d.SocketInterfaceType), d.SocketTarget))
	return cli.NewOption("fsdev", "proxy", properties...)
}

// EnableWriteOut means that host page cache will be used to read and write data but
// write notification will be sent to the guest only when the data has been reported
// as written by the storage subsystem.
//
//	qemu-system-* -fsdev proxy,writeout=writeout
func (d *ProxyFileSystemDevice) EnableWriteOut() *ProxyFileSystemDevice {
	d.properties = append(d.properties,
		// The only supported value is "immediate".
		cli.NewProperty("writeout", "immediate"))
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
	d.properties = append(d.properties, cli.NewProperty("readonly", enabled))
	return d
}
