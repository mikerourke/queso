package chardev

import "github.com/mikerourke/queso/qemu/cli"

// UDPBackend sends all traffic from the guest to a remote host over UDP.
type UDPBackend struct {
	*Backend
	// Port is the port number to connect to.
	Port int
}

// NewUDPBackend returns a new instance of [UDPBackend]. id is the unique ID, which
// can be any string up to 127 characters long. port is the port number to connect to.
//
//	qemu-system-* -chardev udp,id=id,port=port
func NewUDPBackend(id string, port int) *UDPBackend {
	return &UDPBackend{
		Backend: NewBackend("udp", id),
		Port:    port,
	}
}

// SetLocalAddress specifies the local address to bind to. If not specified it
// defaults to 0.0.0.0.
//
//	qemu-system-* -chardev udp,localaddr=addr
func (b *UDPBackend) SetLocalAddress(addr string) *UDPBackend {
	b.properties = append(b.properties, cli.NewProperty("localaddr", addr))
	return b
}

// SetLocalPort specifies the local port to bind to. If not specified, any
// available local port will be used.
//
//	qemu-system-* -chardev udp,localport=port
func (b *UDPBackend) SetLocalPort(port int) *UDPBackend {
	b.properties = append(b.properties, cli.NewProperty("localport", port))
	return b
}

// SetHost specifies the remote host to connect to. If not specified it
// defaults to localhost.
//
//	qemu-system-* -chardev udp,host=host
func (b *UDPBackend) SetHost(host string) *UDPBackend {
	b.properties = append(b.properties, cli.NewProperty("host", host))
	return b
}

// ToggleIPv4 specifies if IPv4 may be used.
//
//	qemu-system-* -chardev udp,ipv4=on|off
func (b *UDPBackend) ToggleIPv4(enabled bool) *UDPBackend {
	b.properties = append(b.properties, cli.NewProperty("ipv4", enabled))
	return b
}

// ToggleIPv6 specifies if IPv6 may be used.
//
//	qemu-system-* -chardev udp,ipv6=on|off
func (b *UDPBackend) ToggleIPv6(enabled bool) *UDPBackend {
	b.properties = append(b.properties, cli.NewProperty("ipv6", enabled))
	return b
}
