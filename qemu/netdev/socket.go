package netdev

import (
	"fmt"

	"github.com/mikerourke/queso"
)

// TCPSocketBackend can be used to connect the guest’s network to another QEMU
// virtual machine using a TCP socket connection.
type TCPSocketBackend struct {
	*Backend
}

func NewTCPSocketBackend(id string) *TCPSocketBackend {
	return &TCPSocketBackend{
		New("socket").SetProperty("id", id),
	}
}

// SetFileDescriptor specifies the handle of an already opened host TCP socket.
//
//	qemu-system-* -netdev socket,fd=fd
func (b *TCPSocketBackend) SetFileDescriptor(fd int) *TCPSocketBackend {
	b.properties = append(b.properties, queso.NewProperty("fd", fd))
	return b
}

// SetListeningAddress sets the port and host that QEMU waits for incoming
// connections on. If no specific host is required, use an empty string for host.
//
//	qemu-system-* -netdev socket,listen=[host]:port
func (b *TCPSocketBackend) SetListeningAddress(port int, host string) *TCPSocketBackend {
	b.properties = append(b.properties,
		queso.NewProperty("listen", fmt.Sprintf("%s:%d", host, port)))
	return b
}

// UDPSocketBackend can be used to connect the guest’s network to another QEMU
// virtual machine using a UDP socket connection.
type UDPSocketBackend struct {
	*Backend
}

func NewUDPSocketBackend(id string) *UDPSocketBackend {
	return &UDPSocketBackend{
		New("socket").SetProperty("id", id),
	}
}

// SetFileDescriptor specifies the handle of an already opened host TCP socket.
//
//	qemu-system-* -netdev socket,fd=fd
func (b *UDPSocketBackend) SetFileDescriptor(fd int) *UDPSocketBackend {
	b.properties = append(b.properties, queso.NewProperty("fd", fd))
	return b
}

// SetLocalAddress specifies the host address to send packets on for a backend
// with multicast configured.
//
//	qemu-system-* -netdev socket,localaddr=addr
func (b *UDPSocketBackend) SetLocalAddress(addr string) *UDPSocketBackend {
	b.properties = append(b.properties, queso.NewProperty("localaddr", addr))
	return b
}

// SetMulticast configures a backend to share the guest's network traffic
// with another QEMU virtual machine using a UDP multicast socket, effectively
// making a bus for every QEMU with same multicast address addr and port.
//
//	qemu-system-* -netdev socket,mcast=maddr:port
//
// # Notes
//
//  1. Several QEMU can be running on different hosts and share same bus (assuming
//     correct multicast setup for these hosts).
//
//  2. Multicast support is compatible with User Mode Linux (argument ethN=mcast).
//     See http://user-mode-linux.sf.net for more details.
//
//  3. Use [UDPSocketBackend.SetFileDescriptor] with a value of "h" to specify an
//     already opened UDP multicast socket.
func (b *UDPSocketBackend) SetMulticast(addr string, port int) *UDPSocketBackend {
	b.properties = append(b.properties,
		queso.NewProperty("mcast", fmt.Sprintf("%s:%d", addr, port)))
	return b
}
