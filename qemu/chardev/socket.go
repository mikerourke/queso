package chardev

import "github.com/mikerourke/queso"

// SocketBackend creates a two-way stream socket, which can be either a TCP or a
// Unix socket.
type SocketBackend struct {
	*Backend
}

// NewSocketBackend returns a new instance of [SocketBackend].
// id is the unique ID, which can be any string up to 127 characters long.
//
//	qemu-system-* -chardev socket,id=id
func NewSocketBackend(id string) *SocketBackend {
	return &SocketBackend{
		New("socket", id),
	}
}

// SetReconnectTimeout sets the timeout for reconnecting on non-server sockets when
// the remote end goes away. QEMU will delay this many seconds and then attempt
// to reconnect. A value of 0 disables reconnecting, and is the default.
//
//	qemu-system-* -chardev socket,reconnect=seconds
func (b *SocketBackend) SetReconnectTimeout(seconds int) *SocketBackend {
	b.properties = append(b.properties, queso.NewProperty("reconnect", seconds))
	return b
}

// SetTLSAuth provides the ID of the QAuthZ authorization object against which
// the client's x509 distinguished name will be validated. This object is only
// resolved at time of use, so can be deleted and recreated on the fly while the
// server is active. If missing, it will default to denying access.
//
//	qemu-system-* -chardev socket,tls-auth=id
func (b *SocketBackend) SetTLSAuth(id string) *SocketBackend {
	b.properties = append(b.properties, queso.NewProperty("tls-auth", id))
	return b
}

// SetTLSCredentials requests enablement of the TLS protocol for encryption, and
// specifies the id of the TLS credentials to use for the handshake.
// The credentials must be previously created with object.TLSCredentials*
// (see object/tls.go).
//
//	qemu-system-* -chardev socket,tls-creds=id
func (b *SocketBackend) SetTLSCredentials(id string) *SocketBackend {
	b.properties = append(b.properties, queso.NewProperty("tls-creds", id))
	return b
}

// ToggleBlockWaitingForClient specifies whether QEMU should block waiting for a client
// to connect to a listening socket.
//
//	qemu-system-* -chardev socket,wait=on|off
func (b *SocketBackend) ToggleBlockWaitingForClient(enabled bool) *SocketBackend {
	b.properties = append(b.properties, queso.NewProperty("server", enabled))
	return b
}

// ToggleListening specifies if the socket shall be a listening socket.
//
//	qemu-system-* -chardev socket,server=on|off
func (b *SocketBackend) ToggleListening(enabled bool) *SocketBackend {
	b.properties = append(b.properties, queso.NewProperty("server", enabled))
	return b
}

// ToggleTelnetEscapeSequences specifies if traffic on the socket should interpret
// telnet escape sequences.
//
//	qemu-system-* -chardev socket,telnet=on|off
func (b *SocketBackend) ToggleTelnetEscapeSequences(enabled bool) *SocketBackend {
	b.properties = append(b.properties, queso.NewProperty("telnet", enabled))
	return b
}

// ToggleWebSocketProtocol specifies whether the socket uses WebSocket protocol for communication.
//
//	qemu-system-* -chardev socket,websocket=on|off
func (b *SocketBackend) ToggleWebSocketProtocol(enabled bool) *SocketBackend {
	b.properties = append(b.properties, queso.NewProperty("websocket", enabled))
	return b
}

// TCPSocketBackend represents a [SocketBackend] that uses a TCP socket.
type TCPSocketBackend struct {
	*SocketBackend
	// Port is the port number or service name of the TCP socket.
	Port string
}

// NewTCPSocketBackend returns a new instance of [TCPSocketBackend]. id is the
// unique ID, which can be any string up to 127 characters long. port is the local
// port to be bound. For a connecting socket specifies the port on the remote
// host to connect to. port can be given as either a port number or a service name.
//
//	qemu-system-* -chardev socket,port=port
func NewTCPSocketBackend(id string, port string) *TCPSocketBackend {
	return &TCPSocketBackend{
		SocketBackend: NewSocketBackend(id),
		Port:          port,
	}
}

// SetHost specifies the local address to be bound. For a connecting socket
// specifies the remote host to connect to. This value is optional for listening
// sockets. If not specified, it defaults to 0.0.0.0.
//
//	qemu-system-* -chardev socket,host=host
func (b *TCPSocketBackend) SetHost(host string) *TCPSocketBackend {
	b.properties = append(b.properties, queso.NewProperty("host", host))
	return b
}

// SetToPort sets the port number that QEMU will attempt to bind to subsequent ports
// up to and including the specified port until it succeeds. This can only be
// specified for listening sockets.
//
//	qemu-system-* -chardev socket,to=to
func (b *TCPSocketBackend) SetToPort(to int) *TCPSocketBackend {
	b.properties = append(b.properties, queso.NewProperty("to", to))
	return b
}

// ToggleIPv4 specifies if IPv4 may be used.
//
//	qemu-system-* -chardev socket,ipv4=on|off
func (b *TCPSocketBackend) ToggleIPv4(enabled bool) *TCPSocketBackend {
	b.properties = append(b.properties, queso.NewProperty("ipv4", enabled))
	return b
}

// ToggleIPv6 specifies if IPv6 may be used.
//
//	qemu-system-* -chardev socket,ipv6=on|off
func (b *TCPSocketBackend) ToggleIPv6(enabled bool) *TCPSocketBackend {
	b.properties = append(b.properties, queso.NewProperty("ipv6", enabled))
	return b
}

// ToggleNoDelay enables or disables the Nagle algorithm.
//
//	qemu-system-* -chardev socket,nodelay=on|off
func (b *TCPSocketBackend) ToggleNoDelay(enabled bool) *TCPSocketBackend {
	b.properties = append(b.properties, queso.NewProperty("nodelay", enabled))
	return b
}

// UnixSocketBackend represents a [SocketBackend] that uses a Unix socket.
type UnixSocketBackend struct {
	*SocketBackend
	// Path specifies the local path of the Unix socket.
	Path string
}

// NewUnixSocketBackend returns a new instance of [UnixSocketBackend]. id is the
// unique ID, which can be any string up to 127 characters long. path specifies
// the local path of the Unix socket.
//
//	qemu-system-* -chardev socket,port=port
func NewUnixSocketBackend(id string, path string) *UnixSocketBackend {
	return &UnixSocketBackend{
		SocketBackend: NewSocketBackend(id),
		Path:          path,
	}
}

// ToggleAbstractNamespace enables or disables the use of the abstract socket
// namespace, rather than the file system. The default value is false.
//
//	qemu-system-* -chardev socket,abstract=on|off
func (b *UnixSocketBackend) ToggleAbstractNamespace(enabled bool) *UnixSocketBackend {
	b.properties = append(b.properties, queso.NewProperty("abstract", enabled))
	return b
}

// ToggleTight specifies whether to set the socket length of abstract sockets to their
// minimum, rather than the full sun_path length. The default value is true.
//
//	qemu-system-* -chardev socket,tight=on|off
func (b *UnixSocketBackend) ToggleTight(enabled bool) *UnixSocketBackend {
	b.properties = append(b.properties, queso.NewProperty("tight", enabled))
	return b
}
