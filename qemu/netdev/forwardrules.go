package netdev

import "fmt"

// ForwardRule are used in a [UserBackend] to redirect TCP/UDP ports to the
// guest or host.
type ForwardRule interface {
	Value() string
}

// PortType describes the port type specified for a [HostForwardRule].
type PortType string

const (
	// PortTCP represents a TCP port.
	PortTCP PortType = "tcp"

	// PortUDP represents a UDP port.
	PortUDP PortType = "udp"
)

// GuestForwardRule is used to forward guest TCP connections to the host.
type GuestForwardRule struct {
	serverIP   string
	serverPort int
	target     string
}

// NewGuestForwardRule returns a forward rule that forwards guest TCP connections
// to the IP address serverIP on port serverPort to the specified target, which
// can be a character device or to a program executed by cmd:command which gets
// spawned for each connection.
//
// # Examples
//
// To open 10.10.1.1:4321 on boot and connect 10.0.2.100:1234 to it whenever the
// guest accesses it, pass the following into WithForwardRule:
//
//	NewGuestForwardRule("10.0.2.100", 1234, "10.10.1.1:4321")
//
// To execute a command on every TCP connection established by guest, pass the
// following into WithForwardRule:
//
//	NewGuestForwardRule("10.0.2.100", 1234, "cmd:netcat 10.10.1.1 4321")
func NewGuestForwardRule(serverIP string, serverPort int, target string) *GuestForwardRule {
	return &GuestForwardRule{
		serverIP:   serverIP,
		serverPort: serverPort,
		target:     target,
	}
}

// Value returns the string representation of the forward rule.
func (r *GuestForwardRule) Value() string {
	return fmt.Sprintf("tcp:%s:%d-%s", r.serverIP, r.serverPort, r.target)
}

// HostForwardRule is used to redirect incoming TCP or UDP connections to the host.
type HostForwardRule struct {
	portType  PortType
	hostIP    string
	hostPort  int
	guestIP   string
	guestPort int
}

// NewHostForwardRule returns a forward rule that redirects incoming TCP or UDP
// connections to the host port hostPort to the guest port guestPort. The host and
// guest server address can be specified with SetHostIP and SetGuestIP
// respectively.
//
// # Examples
//
// To redirect host X11 connection from screen 1 to guest screen 0, pass the
// following into WithForwardRule:
//
//	NewHostForwardRule(PortTCP, 6001, 6000).SetHostIP("127.0.0.1")
//
// To redirect telnet connections from host port 555 to telnet port on the guest,
// pass the following into WithForwardRule:
//
//	NewHostForwardRule(PortTCP, 5555, 23)
func NewHostForwardRule(portType PortType, hostPort int, guestPort int) *HostForwardRule {
	return &HostForwardRule{
		portType:  portType,
		hostIP:    "",
		hostPort:  hostPort,
		guestIP:   "",
		guestPort: guestPort,
	}
}

// SetHostIP sets the host IP address for the host forward rule. By specifying
// this value, the rule can be bound to a specific host interface.
func (r *HostForwardRule) SetHostIP(ip string) *HostForwardRule {
	r.hostIP = ip
	return r
}

// SetGuestIP sets the guest IP address for the host forward rule. If not
// specified, its value is x.x.x.15 (default first address given by the built-in
// DHCP server).
func (r *HostForwardRule) SetGuestIP(ip string) *HostForwardRule {
	r.guestIP = ip
	return r
}

// Value returns the string representation of the forward rule.
func (r *HostForwardRule) Value() string {
	return fmt.Sprintf("%s:%s:%d-%s:%d",
		r.portType, r.hostIP, r.hostPort, r.guestIP, r.guestPort)
}
