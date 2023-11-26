package netdev

import "fmt"

// TCPSocketBackend can be used to connect the guest’s network to another QEMU
// virtual machine using a TCP socket connection.
type TCPSocketBackend struct {
	*Backend
}

// NewTCPSocketBackend returns a new instance of [TCPSocketBackend]. id is a
// unique identifier for the backend.
//
//	qemu-system-* -netdev socket,id=id
func NewTCPSocketBackend(id string) *TCPSocketBackend {
	backend := New("socket")

	backend.SetProperty("id", id)

	return &TCPSocketBackend{backend}
}

// SetFileDescriptor specifies the handle of an already opened host TCP socket.
//
//	qemu-system-* -netdev socket,fd=fd
func (b *TCPSocketBackend) SetFileDescriptor(fd int) *TCPSocketBackend {
	b.SetProperty("fd", fd)
	return b
}

// SetListeningAddress sets the port and host that QEMU waits for incoming
// connections on. If no specific host is required, use an empty string for host.
//
//	qemu-system-* -netdev socket,listen=[host]:port
func (b *TCPSocketBackend) SetListeningAddress(port int, host string) *TCPSocketBackend {
	b.SetProperty("listen", fmt.Sprintf("%s:%d", host, port))
	return b
}

// UDPSocketBackend can be used to connect the guest’s network to another QEMU
// virtual machine using a UDP socket connection.
type UDPSocketBackend struct {
	*Backend
}

// NewUDPSocketBackend returns a new instance of [UDPSocketBackend]. id is a
// unique identifier for the backend.
//
//	qemu-system-* -netdev socket,id=id
func NewUDPSocketBackend(id string) *UDPSocketBackend {
	backend := New("socket")

	backend.SetProperty("id", id)

	return &UDPSocketBackend{backend}
}

// SetFileDescriptor specifies the handle of an already opened host TCP socket.
//
//	qemu-system-* -netdev socket,fd=fd
func (b *UDPSocketBackend) SetFileDescriptor(fd int) *UDPSocketBackend {
	b.SetProperty("fd", fd)
	return b
}

// SetLocalAddress specifies the host address to send packets on for a backend
// with multicast configured.
//
//	qemu-system-* -netdev socket,localaddr=addr
func (b *UDPSocketBackend) SetLocalAddress(addr string) *UDPSocketBackend {
	b.SetProperty("localaddr", addr)
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
	b.SetProperty("mcast", fmt.Sprintf("%s:%d", addr, port))
	return b
}
