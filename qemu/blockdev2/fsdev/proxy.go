package fsdev

import "github.com/mikerourke/queso/internal/cli"

type SocketTarget string

const (
	SocketTargetPath       SocketTarget = "socket"
	SocketTargetDescriptor SocketTarget = "sock_fd"
)

type anyProxyDevice struct {
	Identifier string
	Socket     string
	Target     SocketTarget
	properties []*cli.Property
}

// EnableWriteOut means that host page cache will be used to read and write data but
// write notification will be sent to the guest only when the data has been reported
// as written by the storage subsystem.
//
//	qemu-system-* -fsdev proxy,writeout=writeout
func (d *anyProxyDevice) EnableWriteOut() *anyProxyDevice {
	d.properties = append(d.properties,
		// The only supported value is "immediate".
		cli.NewProperty("writeout", "immediate"))
	return d
}

func (d *anyProxyDevice) SetSocket(path string) *anyProxyDevice {
	d.Socket = path
	return d
}

func (d *anyProxyDevice) ToggleReadOnly(enabled bool) *anyProxyDevice {
	d.properties = append(d.properties, cli.NewProperty("readonly", enabled))
	return d
}

type ProxyFileSystemDevice struct {
	*anyProxyDevice
}

func NewProxyFileSystemDevice(id string, socket string, target SocketTarget) *ProxyFileSystemDevice {
	return &ProxyFileSystemDevice{
		&anyProxyDevice{
			Identifier: id,
			Socket:     socket,
			Target:     target,
			properties: nil,
		},
	}
}

func (d *ProxyFileSystemDevice) option() *cli.Option {
	properties := append(d.properties,
		cli.NewProperty("id", d.Identifier),
		cli.NewProperty(string(d.Target), d.Socket))
	return cli.NewOption("fsdev", "proxy", properties...)
}

func (d *ProxyFileSystemDevice) SetID(id string) *ProxyFileSystemDevice {
	d.Identifier = id
	return d
}

type VirtualProxyFileSystemDevice struct {
	*anyProxyDevice
}

func NewVirtualProxyFileSystemDevice(id string, socket string, target SocketTarget) *VirtualProxyFileSystemDevice {
	return &VirtualProxyFileSystemDevice{
		&anyProxyDevice{
			Identifier: id,
			Socket:     socket,
			Target:     target,
			properties: nil,
		},
	}
}

func (d *VirtualProxyFileSystemDevice) SetMountTag(tag string) *VirtualProxyFileSystemDevice {
	d.properties = append(d.properties, cli.NewProperty("mount_tag", tag))
	return d
}
