package network

import "fmt"

// ForwardRule are used in a UserBackend to redirect TCP/UDP ports to the
// guest or host.
type ForwardRule interface {
	PropertyKey() string
	PropertyValue() string
}

// PortType describes the port type specified for a HostForwardRule.
type PortType string

const (
	PortTypeTCP PortType = "tcp"
	PortTypeUDP PortType = "udp"
)

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
// guest server address can be specified with WithHostIP and WithGuestIP
// respectively.
//
// Examples
//
// To redirect host X11 connection from screen 1 to guest screen 0, pass the
// following into WithForwardRule:
//	NewHostForwardRule(PortTypeTCP, 6001, 6000).WithHostIP("127.0.0.1")
//
// To redirect telnet connections from host port 555 to telnet port on the guest,
// pass the following into WithForwardRule:
//	NewHostForwardRule(PortTypeTCP, 5555, 23)
func NewHostForwardRule(portType PortType, hostPort int, guestPort int) HostForwardRule {
	return HostForwardRule{
		portType:  portType,
		hostIP:    "",
		hostPort:  hostPort,
		guestIP:   "",
		guestPort: guestPort,
	}
}

// WithHostIP sets the host IP address for the host forward rule. By specifying
// this value, the rule can be bound to a specific host interface.
func (hfr HostForwardRule) WithHostIP(ip string) HostForwardRule {
	hfr.hostIP = ip

	return hfr
}

// WithGuestIP sets the guest IP address for the host forward rule. If not
// specified, its value is x.x.x.15 (default first address given by the built-in
// DHCP server).
func (hfr HostForwardRule) WithGuestIP(ip string) HostForwardRule {
	hfr.guestIP = ip

	return hfr
}

// PropertyKey returns the key of the property to pass to the WithForwardRule
// option (left side of the "=").
func (hfr HostForwardRule) PropertyKey() string {
	return "hostfwd"
}

// PropertyValue returns the string representation of the rule to pass to the
// WithForwardRule property (right side of the "=").
func (hfr HostForwardRule) PropertyValue() string {
	return fmt.Sprintf("%s:%s:%d-%s:%d",
		hfr.portType, hfr.hostIP, hfr.hostPort, hfr.guestIP, hfr.guestPort)
}

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
// Examples
//
// To open 10.10.1.1:4321 on boot and connect 10.0.2.100:1234 to it whenever the
// guest accesses it, pass the following into WithForwardRule:
//	NewGuestForwardRule("10.0.2.100", 1234, "10.10.1.1:4321")
//
// To execute a command on every TCP connection established by guest, pass the
// following into WithForwardRule:
//	NewGuestForwardRule("10.0.2.100", 1234, "cmd:netcat 10.10.1.1 4321")
func NewGuestForwardRule(serverIP string, serverPort int, target string) GuestForwardRule {
	return GuestForwardRule{
		serverIP:   serverIP,
		serverPort: serverPort,
		target:     target,
	}
}

// PropertyKey returns the key of the property to pass to the WithForwardRule
// option (left side of the "=").
func (gfr GuestForwardRule) PropertyKey() string {
	return "guestfwd"
}

// PropertyValue returns the string representation of the rule to pass to the
// WithForwardRule property (right side of the "=").
func (gfr GuestForwardRule) PropertyValue() string {
	return fmt.Sprintf("tcp:%s:%d-%s", gfr.serverIP, gfr.serverPort, gfr.target)
}
