// Package chardev is used to define character devices for use with QEMU.
// See https://qemu.readthedocs.io/en/latest/system/invocation.html#hxtool-6
// for more details.
package chardev

import "github.com/mikerourke/queso"

const (
	BackendTypeNull           = "null"
	BackendTypeBraille        = "braille"
	BackendTypeConsole        = "console"
	BackendTypeFile           = "file"
	BackendTypeMSMouse        = "msmouse"
	BackendTypeParallel       = "parallel"
	BackendTypePipe           = "pipe"
	BackendTypePTY            = "pty"
	BackendTypeRingBuffer     = "ringbuf"
	BackendTypeSerial         = "serial"
	BackendTypeSocket         = "socket"
	BackendTypeSpicePort      = "spiceport"
	BackendTypeSpiceVMC       = "spicevmc"
	BackendTypeStdio          = "stdio"
	BackendTypeTTY            = "tty"
	BackendTypeUDP            = "udp"
	BackendTypeVirtualConsole = "vc"
)

// Backend returns a character device of the specified backendType with the
// specified id and properties.
func Backend(backendType string, id string, properties ...*Property) *queso.Option {
	props := []*queso.Property{queso.NewProperty("id", id)}

	for _, property := range properties {
		props = append(props, property.Property)
	}

	return queso.NewOption("chardev", backendType, props...)
}

// NullBackend represents a void device. This device will not emit any data, and
// will drop any data it receives. The NullBackend does not take any options.
func NullBackend(id string) *queso.Option {
	return Backend(BackendTypeNull, id)
}

// BrailleBackend connects to a local BrlAPI server. The BrailleBackend does not
// take any options.
func BrailleBackend(id string) *queso.Option {
	return Backend(BackendTypeBraille, id)
}

// ConsoleBackend sends traffic from the guest to QEMU's standard output.
// The ConsoleBackend does not take any options.
func ConsoleBackend(id string) *queso.Option {
	return Backend(BackendTypeConsole, id)
}

// FileBackend logs all traffic received from the guest to a file.
//
// The path parameter specifies the path of the file to be opened. This file will
// be created if it does not already exist, and overwritten if it does.
func FileBackend(id string, path string) *queso.Option {
	return Backend(BackendTypeFile, id, NewProperty("path", path))
}

// MSMouseBackend forwards QEMU's emulated msmouse events to the guest.
// The MSMouseBackend does not take any options.
func MSMouseBackend(id string) *queso.Option {
	return Backend(BackendTypeMSMouse, id)
}

// ParallelBackend connects to a local parallel port and is only available on Linux,
// FreeBSD and DragonFlyBSD hosts.
//
// The path parameter specifies the path to the parallel port device.
func ParallelBackend(id string, path string) *queso.Option {
	return Backend(BackendTypeParallel, id, NewProperty("path", path))
}

// PipeBackend creates a two-way connection to the guest. The behaviour differs
// slightly between Windows hosts and other hosts.
//
// On Windows, a single duplex pipe will be created at `\\.pipe\path`.
//
// On other hosts, 2 pipes will be created called path.in and path.out. Data
// written to path.in will be received by the guest. Data written by the guest can
// be read from path.out. QEMU will not create these FIFOs, and requires them to
// be present.
//
// The path parameter forms part of the pipe path as described above.
func PipeBackend(id string, path string) *queso.Option {
	return Backend(BackendTypePipe, id, NewProperty("path", path))
}

// PTYBackend creates a new pseudo-terminal on the host and connects to it.
// The PTYBackend does not take any options and is not available on Windows hosts.
func PTYBackend(id string) *queso.Option {
	return Backend(BackendTypePTY, id)
}

// RingBufferBackend creates a ring buffer with the specified fixed size.
// If specified, the size parameter must be a power of two. If the size parameter
// is an empty string, 64K is used.
func RingBufferBackend(id string, size string) *queso.Option {
	props := make([]*Property, 0)

	if size != "" {
		props = append(props, NewProperty("size", size))
	}

	return Backend(BackendTypeRingBuffer, id, props...)
}

// SerialBackend sends traffic from the guest to a serial device on the host.
//
// On Unix hosts serial will actually accept any tty device, not only serial lines.
//
// The path parameter specifies the name of the serial device to open.
func SerialBackend(id string, path string) *queso.Option {
	return Backend(BackendTypeSerial, id, NewProperty("path", path))
}

// SpicePortBackend connects to a spice port, allowing a Spice client to handle the
// traffic identified by a name (preferably a fqdn)., and is only available when
// spice support is built in. The name parameter is the name of spice channel to
// connect to.
func SpicePortBackend(id string, name string, properties ...*Property) *queso.Option {
	props := []*Property{NewProperty("name", name)}

	if properties != nil {
		props = append(props, properties...)
	}

	return Backend(BackendTypeSpicePort, id, props...)
}

// SpiceVMCBackend connects to a spice virtual machine channel, such as `vdiport`,
// and is only available when spice support is built in. The name parameter is
// the name of spice channel to connect to.
func SpiceVMCBackend(id string, name string, properties ...*Property) *queso.Option {
	props := []*Property{NewProperty("name", name)}

	if properties != nil {
		props = append(props, properties...)
	}

	return Backend(BackendTypeSpiceVMC, id, props...)
}

// WithDebugLevel returns the debug flag to set for Spice port and VMC.
// TODO: Find out if this is supposed to be a string or a number or something else.
func WithDebugLevel(value string) *Property {
	return NewProperty("debug", value)
}

// StdioBackend connects to standard input and standard output of the QEMU process.
//
// The signal parameter controls if signals are enabled on the terminal, that
// includes exiting QEMU with the key sequence Control-c. This option is enabled
// by default.
func StdioBackend(id string, signal bool) *queso.Option {
	return Backend(BackendTypeStdio, id, NewProperty("signal", signal))
}

// TCPSocketBackend creates a SocketBackend for a two-way stream TCP socket. The
// port parameter specifies the local port to be bound. For a connecting socket
// specifies the port on the remote host to connect to. It can be given as either a
// port number or a service name.
func TCPSocketBackend(id string, port interface{}, properties ...*Property) *queso.Option {
	props := []*Property{NewProperty("port", port)}

	if properties != nil {
		props = append(props, properties...)
	}

	return Backend(BackendTypeSocket, id, props...)
}

// TTYBackend is only available on Linux, Sun, FreeBSD, NetBSD, OpenBSD and
// DragonFlyBSD hosts. It is an alias for SerialBackend.
//
// The path parameter specifies the path to the tty.
func TTYBackend(id string, path string) *queso.Option {
	return Backend(BackendTypeTTY, id, NewProperty("path", path))
}

// UDPBackend sends all traffic from the guest to a remote host over UDP.
func UDPBackend(id string, port int, properties ...*Property) *queso.Option {
	props := []*Property{NewProperty("port", port)}

	if properties != nil {
		props = append(props, properties...)
	}

	return Backend(BackendTypeUDP, id, props...)
}

// UnixSocketBackend creates a SocketBackend for a two-way stream Unix socket.
// The path parameter specifies the local path of the Unix socket.
func UnixSocketBackend(id string, path string, properties ...*Property) *queso.Option {
	props := []*Property{NewProperty("path", path)}

	if properties != nil {
		props = append(props, properties...)
	}

	return Backend(BackendTypeSocket, id, props...)
}

// VirtualConsoleBackend connects to a QEMU text console.
func VirtualConsoleBackend(id string, properties ...*Property) *queso.Option {
	return Backend(BackendTypeVirtualConsole, id, properties...)
}

// Property represents a property to associate with a character device Backend.
type Property struct {
	*queso.Property
}

// NewProperty returns a new instance of Property.
func NewProperty(key string, value interface{}) *Property {
	return &Property{
		Property: queso.NewProperty(key, value),
	}
}

// IsListeningSocket specifies that the socket shall be a listening socket for
// a TCPSocketBackend or UnixSocketBackend.
func IsListeningSocket(listening bool) *Property {
	return NewProperty("server", listening)
}

// IsBlockWaitingForClient specifies whether QEMU should block waiting for a client
// to connect to a listening socket in a TCPSocketBackend or UnixSocketBackend.
func IsBlockWaitingForClient(block bool) *Property {
	return NewProperty("wait", !block)
}

// IsTelnet specifies if traffic on the socket should interpret telnet escape
// sequences for a TCPSocketBackend or UnixSocketBackend.
func IsTelnet(enabled bool) *Property {
	return NewProperty("telnet", enabled)
}

// IsWebSocketUsed specifies whether the socket uses WebSocket protocol for
// communication for a TCPSocketBackend or UnixSocketBackend.
func IsWebSocketUsed(used bool) *Property {
	return NewProperty("websocket", used)
}

// WithReconnect sets the timeout for reconnecting on non-server sockets when the
// remote end goes away for a TCPSocketBackend or UnixSocketBackend. QEMU will
// delay this many seconds and then attempt to reconnect. A value of 0 disables
// reconnecting, and is the default.
func WithReconnect(seconds int) *Property {
	return NewProperty("reconnect", seconds)
}

// WithTLSCredentials requests enablement of the TLS protocol for encryption, and
// specifies the id of the TLS credentials to use for the handshake in a TCPSocketBackend
// or UnixSocketBackend. The credentials must be previously created with the
// object.TLSCredentials* option (see object/tls.go).
func WithTLSCredentials(id string) *Property {
	return NewProperty("tls-creds", id)
}

// WithTLSAuthZ provides the ID of the QAuthZ authorization object against which
// the client's x509 distinguished name will be validated in a TCPSocketBackend
// or UnixSocketBackend. This object is only resolved at time of use, so can be
// deleted and recreated on the fly while the TCPSocketBackend or UnixSocketBackend
// server is active. If missing, it will default to denying access.
func WithTLSAuthZ(id string) *Property {
	return NewProperty("tls-authz", id)
}

// WithHost specifies the local address to be bound for a TCPSocketBackend or
// a remote host to connect to for a UDPBackend.
func WithHost(addr string) *Property {
	return NewProperty("host", addr)
}

// WithToPort is only relevant to listening sockets on a TCPSocketBackend. If it
// is specified, and port cannot be bound, QEMU will attempt to bind to subsequent
// ports up to and including to until it succeeds. It must be specified as a port
// number.
func WithToPort(port int) *Property {
	return NewProperty("to", port)
}

// IsIPv4 is used to specify if IPv4 is enabled for a TCPSocketBackend or UDPBackend.
// If this property and the IsIPv6 property are omitted, the socket may use either
// protocol.
func IsIPv4(enabled bool) *Property {
	return NewProperty("ipv4", enabled)
}

// IsIPv6 is used to specify if IPv4 is enabled for a TCPSocketBackend or UDPBackend.
// If this property and the IsIPv4 property are omitted, the socket may use
// either protocol.
func IsIPv6(enabled bool) *Property {
	return NewProperty("ipv6", enabled)
}

// IsNoDelay enables/disables the Nagle algorithm in a TCPSocketBackend.
func IsNoDelay(disabled bool) *Property {
	return NewProperty("nodelay", disabled)
}

// IsAbstractNamespace specifies the use of the abstract socket namespace, rather
// than the filesystem for a UnixSocketBackend. The default is false.
func IsAbstractNamespace(abstract bool) *Property {
	return NewProperty("abstract", abstract)
}

// IsTight specifies whether to set the socket length of abstract sockets to their
// minimum, rather than the full sun_path length in a UnixSocketBackend.
// The default is true.
func IsTight(tight bool) *Property {
	return NewProperty("tight", tight)
}

// WithLocalAddress specifies the local address to bind to for a UDPBackend.
// If not specified, it defaults to 0.0.0.0.
func WithLocalAddress(addr string) *Property {
	return NewProperty("localaddr", addr)
}

// WithLocalPort specifies the local port to bind to for a UDPBackend. If not
// specified any available local port will be used.
func WithLocalPort(port int) *Property {
	return NewProperty("localport", port)
}

// WithConsoleWidth specifies the width in pixels for a VirtualConsoleBackend.
func WithConsoleWidth(pixels int) *Property {
	return NewProperty("width", pixels)
}

// WithConsoleHeight specifies the height in pixels for a VirtualConsoleBackend.
func WithConsoleHeight(pixels int) *Property {
	return NewProperty("height", pixels)
}

// WithConsoleRows specifies that the console be sized to fit a text console
// with the given row count for a VirtualConsoleBackend.
func WithConsoleRows(count int) *Property {
	return NewProperty("rows", count)
}

// WithConsoleColumns specifies that the console be sized to fit a text console
// with the given column count for a VirtualConsoleBackend.
func WithConsoleColumns(count int) *Property {
	return NewProperty("cols", count)
}
