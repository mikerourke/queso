package network

import (
	"fmt"
	"net"

	"github.com/mikerourke/queso"
)

// BackendType represents the network backend type for use with a NIC.
type BackendType string

const (
	BackendTypeNone      BackendType = "none"
	BackendTypeBridge    BackendType = "bridge"
	BackendTypeL2TPv3    BackendType = "l2tpv3"
	BackendTypeNetMap    BackendType = "netmap"
	BackendTypeSocket    BackendType = "socket"
	BackendTypeTAP       BackendType = "tap"
	BackendTypeUser      BackendType = "user"
	BackendTypeVDE       BackendType = "vde"
	BackendTypeVHostUser BackendType = "vhost-user"
)

// NIC is a shortcut for configuring both the on-board (default) guest NIC hardware
// and the host network backend in one go. The host backend options are the same
// as with the corresponding NetworkDevice. The guest NIC model can be set with
// the WithModel property. The hardware MAC address can be set with the
// WithMACAddress property.
func NIC(backendType BackendType, properties ...*Property) *queso.Option {
	props := make([]*queso.Property, 0)

	for _, property := range properties {
		props = append(props, property.Property)
	}
	return queso.NewOption("nic", string(backendType), props...)
}

// newNetworkBackend is a convenience method for creating a new network device
// option to pass to QEMU.
func newNetworkBackend(
	backendType BackendType,
	idField string,
	id string,
	properties ...*Property,
) *queso.Option {
	props := make([]*queso.Property, 0)

	if id != "" {
		props = append(props, queso.NewProperty(idField, id))
	}

	for _, property := range properties {
		props = append(props, property.Property)
	}

	return queso.NewOption("netdev", string(backendType), props...)
}

// UserBackend configures a user mode host network backend which requires no
// administrator privilege to run. The id parameter represents a symbolic name
// for use in monitor commands.
func UserBackend(id string, properties ...*Property) *queso.Option {
	return newNetworkBackend(BackendTypeUser, "id", id, properties...)
}

// TAPBackend configures a host TAP network backend.
func TAPBackend(id string, properties ...*Property) *queso.Option {
	return newNetworkBackend(BackendTypeTAP, "id", id, properties...)
}

// Bridge represents a host bridge device to which a host TAP network
// interface can be connected for a TAPBackend.
//
// Use the network helper WithHelper to configure the TAP interface and attach
// it to the bridge. The default network helper executable is
// `/path/to/qemu-bridge-helper` and the default bridge device is "br0".
func Bridge(id string, properties ...*Property) *queso.Option {
	return newNetworkBackend(BackendTypeBridge, "id", id, properties...)
}

// SocketBackend configures a host network backend can be used to connect the
// guest's network to another QEMU virtual machine using a TCP socket connection
// (by specifying the WithListen and WithConnect properties) or a UDP multicast
// socket (using the WithMulticast property).
func SocketBackend(id string, properties ...*Property) *queso.Option {
	return newNetworkBackend(BackendTypeSocket, "id", id, properties...)
}

// L2TPv3Backend configures a L2TPv3 pseudo-wire host network backend. L2TPv3 (RFC3931)
// is a popular protocol to transport Ethernet (and other Layer 2) data frames
// between two systems. It is present in routers, firewalls and the Linux
// kernel (from version 3.3 onwards).
//
// This transport allows a VM to communicate to another VM, router or firewall directly.
func L2TPv3Backend(
	id string,
	sourceAddr net.IP,
	destAddr net.IP,
	properties ...*Property,
) *queso.Option {
	props := []*Property{
		NewProperty("src", sourceAddr.String()),
		NewProperty("dst", destAddr.String()),
	}

	if properties != nil {
		props = append(props, properties...)
	}

	return newNetworkBackend(BackendTypeL2TPv3, "id", id, props...)
}

// VDEBackend configures a VDE backend. This option is only available if QEMU has
// been compiled with VDE support enabled.
func VDEBackend(id string, properties ...*Property) *queso.Option {
	return newNetworkBackend(BackendTypeVDE, "id", id, properties...)
}

// VHostUserBackend establishes a vhost-user network backend, backed by charDevID.
// The chardev should be a unix domain socket backed one. The vhost-user
// uses a specifically defined protocol to pass vhost ioctl replacement messages
// to an application on the other end of the socket.
func VHostUserBackend(chardev string, properties ...*Property) *queso.Option {
	return newNetworkBackend(BackendTypeVHostUser, "chardev", chardev, properties...)
}

// VHostVDPABackend establishes a vhost-vdpa network device.
//
// vDPA device is a device that uses a data path which complies with the virtio
// specifications with a vendor specific control path. vDPA devices can be both
// physically located on the hardware or emulated by software.
func VHostVDPABackend(vHostDeviceID string) *queso.Option {
	return queso.NewOption("netdev", "vhost-vdpa",
		queso.NewProperty("vhostdev", vHostDeviceID))
}

// HubPort creates a hub port on the emulated hub with ID hubID.
//
// The hub port network device lets you connect a NIC to a QEMU emulated
// hub instead of a single network device. Alternatively, you can also connect
// the hub port to another network device with ID networkDeviceID. Specify an
// empty string for the networkDeviceID parameter to use the emulated hub.
func HubPort(id string, hubID string, networkDeviceID string) *queso.Option {
	props := []*queso.Property{
		queso.NewProperty("id", id),
		queso.NewProperty("hubid", hubID),
	}

	if networkDeviceID != "" {
		props = append(props, queso.NewProperty("netdev", networkDeviceID))
	}

	return queso.NewOption("netdev", "hubport", props...)
}

// Property represents a network property to use for network device options.
type Property struct {
	*queso.Property
}

// NewProperty returns a new instance of Property.
func NewProperty(key string, value interface{}) *Property {
	return &Property{
		Property: queso.NewProperty(key, value),
	}
}

// WithNICModel sets the guest NIC model.
func WithNICModel(model string) *Property {
	return NewProperty("model", model)
}

// WithMACAddress sets the hardware MAC address for the NIC.
func WithMACAddress(addr string) *Property {
	return NewProperty("mac", addr)
}

// WithFileDescriptor specifies the handle of an already opened host TAP interface
// or TCP socket.
func WithFileDescriptor(fd string) *Property {
	return NewProperty("fd", fd)
}

// IsIPv4 is used to specify if IPv4 is enabled for a UserBackend. If this
// property and the IsIPv6 property are omitted, both protocols are enabled.
func IsIPv4(enabled bool) *Property {
	return NewProperty("ipv4", enabled)
}

// IsIPv6 is used to specify if IPv4 is enabled for a UserBackend. If this property
// and the IsIPv4 property are omitted, both protocols are enabled.
func IsIPv6(enabled bool) *Property {
	return NewProperty("ipv6", enabled)
}

// WithIPv4Address sets the IPv4 network address the guest will see for a UserBackend.
// The mask parameter represents the netmask. To use the default mask, specify -1.
// The default is "10.0.2.0/24".
func WithIPv4Address(addr net.IP, mask int) *Property {
	value := addr.String()

	if mask != -1 {
		value = fmt.Sprintf("%s/%d", value, mask)
	}

	return NewProperty("net", value)
}

// WithIPv4Host specifies the guest-visible IPv4 address of the host for a UserBackend.
// The default is the 2nd IP in the guest network, i.e. x.x.x.2.
func WithIPv4Host(addr net.IP) *Property {
	return NewProperty("host", addr.String())
}

// WithIPv6Address sets the IPv6 network address the guest will see for a UserBackend.
// The network prefix is given in the usual hexadecimal IPv6 address notation.
// The prefixSize is optional, and is given as the number of valid top-most
// bits. To use the default, specify -1 for prefixSize. The default is
// "fec0::/64".
func WithIPv6Address(addr net.IP, prefixSize int) *Property {
	value := addr.String()

	if prefixSize != -1 {
		value = fmt.Sprintf("%s/%d", value, prefixSize)
	}

	return NewProperty("ipv6-net", value)
}

// WithIPv6Host specifies the guest-visible IPv6 address of the host for a UserBackend.
// The default is the 2nd IPv6 in the guest network, i.e. xxxx::2.
func WithIPv6Host(addr net.IP) *Property {
	return NewProperty("ipv6-host", addr.String())
}

// IsRestricted specifies whether the guest will be isolated for a UserBackend,
// i.e. it will not be able to contact the host and no guest IP packets will be
// routed over the host to the outside. This option does not affect any
// explicitly set forwarding rules.
func IsRestricted(restricted bool) *Property {
	return NewProperty("restrict", restricted)
}

// WithHostName specifies the client hostname reported by the built-in DHCP server
// for a UserBackend.
func WithHostName(name string) *Property {
	return NewProperty("hostname", name)
}

// WithDHCPStart specifies the first of the 16 IPs the built-in DHCP server can
// assign for a UserBackend. The default is the 15th to 31st IP in the guest
// network, i.e. x.x.x.15 to x.x.x.31.
func WithDHCPStart(addr net.IP) *Property {
	return NewProperty("dhcpstart", addr.String())
}

// WithIPv4DNS specifies the guest-visible address of the IPv4 virtual nameserver
// for a UserBackend. The address must be different from the host address. The
// default is the 3rd IP in the guest network, i.e. x.x.x.3.
func WithIPv4DNS(addr net.IP) *Property {
	return NewProperty("dns", addr.String())
}

// WithIPv6DNS specifies the guest-visible address of the IPv6 virtual nameserver
// for a UserBackend. The address must be different from the host address. The
// default is the 3rd IP in the guest network, i.e. xxxx::3.
func WithIPv6DNS(addr net.IP) *Property {
	return NewProperty("ipv6-dns", addr.String())
}

// WithDNSSearch provides an entry for the domain-search list sent by the built-in
// DHCP server for a UserBackend. More than one domain suffix can be transmitted
// by specifying this property multiple times. If supported, this will cause the
// guest to automatically try to append the given domain suffix(es) in case a domain
// name can not be resolved.
func WithDNSSearch(domain string) *Property {
	return NewProperty("dnssearch", domain)
}

// WithDomainName specifies the client domain name reported by the built-in DHCP
// server for a UserBackend.
func WithDomainName(domain string) *Property {
	return NewProperty("domainname", domain)
}

// WithTFTP activates a built-in TFTP server for a UserBackend. The files in the
// specified dir parameter will be exposed as the root of a TFTP server. The
// TFTP client on the guest must be configured in binary mode (use the command
// bin of the Unix TFTP client).
func WithTFTP(dir string) *Property {
	return NewProperty("tftp", dir)
}

// WithTFTPServerName broadcasts the specified name as the "TFTP server name"
// (RFC2132 option 66) in BOOTP reply for a UserBackend. This can be used to advise
// the guest to load boot files or configurations from a different server than
// the host address.
func WithTFTPServerName(name string) *Property {
	return NewProperty("tftp-server-name", name)
}

// WithBootFile broadcasts file as the BOOTP filename for a UserBackend. In
// conjunction with WithTFTP, this can be used to network boot a guest from a
// local directory.
func WithBootFile(file string) *Property {
	return NewProperty("bootfile", file)
}

// WithSMB activates a built-in SMB server so that Windows OSes can access to the
// host files in the specified dir transparently for a UserBackend. The server
// address can be set with the WithSMBServerAddress property.
//
// In the guest Windows OS, the line:
//	10.0.2.4 smbserver
// Must be added in the file `C:\WINDOWS\LMHOSTS` (for Windows 9x/ME) or
// `C:\WINNT\SYSTEM32\DRIVERS\ETC\LMHOSTS` (Windows NT/2000).
//
// Then directory can be accessed in `\\smbserver\qemu`.
//
// Note that a SAMBA server must be installed on the host OS.
func WithSMB(dir string) *Property {
	return NewProperty("smb", dir)
}

// WithSMBServerAddress sets the address of the SMB server for a UserBackend.
// By default, the 4th IP in the guest network is used, i.e. x.x.x.4. This must
// be used in conjunction with the WithSMB property.
func WithSMBServerAddress(addr net.IP) *Property {
	return NewProperty("smbserver", addr.String())
}

// WithForwardRule defines a forwarding rule to use for the network (either a
// HostForwardRule or GuestForwardRule instance) for a UserBackend.
func WithForwardRule(rule ForwardRule) *Property {
	return NewProperty(rule.PropertyKey(), rule.PropertyValue())
}

// WithInterfaceName defines the interface name for a TAPBackend.
func WithInterfaceName(name string) *Property {
	return NewProperty("ifname", name)
}

// WithUpScript is the network script file to configure the TAPBackend.
// The default network configure script is `/etc/qemu-ifup`.
func WithUpScript(file string) *Property {
	return NewProperty("script", file)
}

// WithDownScript is the network script file to de-configure the TAPBackend.
// The default network de-configure script is `/etc/qemu-ifdown`.
func WithDownScript(file string) *Property {
	return NewProperty("downscript", file)
}

// WithBridge specifies the bridge device for a TAPBackend.
func WithBridge(bridge string) *Property {
	return NewProperty("br", bridge)
}

// WithHelper specifies an executable path to configure the TAP interface and
// attach it to the Bridge for a TAPBackend.
func WithHelper(helper string) *Property {
	return NewProperty("helper", helper)
}

// WithListen specifies that QEMU wait for incoming connections on port for a
// SocketBackend. The host parameter can be nil if not required.
func WithListen(host net.IP, port int) *Property {
	value := fmt.Sprintf(":%d", port)

	if host != nil {
		value = fmt.Sprintf("%s:%d", host.String(), port)
	}

	return NewProperty("listen", value)
}

// WithConnect is used to connect to another QEMU instance using the WithListen
// property for a SocketBackend. WithFileDescriptor("h") specifies an already
// opened TCP socket.
func WithConnect(host net.IP, port int) *Property {
	return NewProperty("connect", fmt.Sprintf("%s:%d", host.String(), port))
}

// WithMulticast configures a SocketBackend to share the guest's network traffic
// with another QEMU virtual machines using a UDP multicast socket, effectively
// making a bus for every QEMU with same multicast address addr and port.
//
// Notes
//
// 1. Several QEMU can be running on different hosts and share same bus (assuming
// correct multicast setup for these hosts).
//
// 2. WithMulticast support is compatible with User Mode Linux (argument ethN=mcast).
// See http://user-mode-linux.sf.net for more details.
//
// 3. Use WithFileDescriptor("h") to specify an already opened UDP multicast socket.
func WithMulticast(addr net.IP, port int) *Property {
	return NewProperty("mcast", fmt.Sprintf("%s:%d", addr.String(), port))
}

// WithLocalAddress specifies the host address to send packets on for a SocketBackend
// using the WithMulticast property.
func WithLocalAddress(addr net.IP) *Property {
	return NewProperty("localaddr", addr.String())
}

// WithSourcePort represents the source UDP port for a L2TPv3Backend.
func WithSourcePort(port int) *Property {
	return NewProperty("srcport", port)
}

// WithDestPort represents the destination UDP port for a L2TPv3Backend.
func WithDestPort(port int) *Property {
	return NewProperty("srcport", port)
}

// WithReceiveSession is the receiver session tunnel for a L2TPv3Backend.
func WithReceiveSession(session string) *Property {
	return NewProperty("rxsession", session)
}

// WithTransmitSession is the transmission session tunnel for a L2TPv3Backend.
func WithTransmitSession(session string) *Property {
	return NewProperty("txsession", session)
}

// IsForceIPv6 forces IPv6 for a L2TPv3Backend, otherwise defaults to IPv4.
func IsForceIPv6(force bool) *Property {
	return NewProperty("ipv6", force)
}

// WithReceiveCookie specifies the receiver cookie for a L2TPv3Backend.
// Cookies are a weak form of security in the L2TPv3 specification. Their function
// is mostly to prevent misconfiguration. By default, they are 32 bit.
func WithReceiveCookie(cookie string) *Property {
	return NewProperty("rxcookie", cookie)
}

// WithTransmitCookie specifies the transmission cookie for a L2TPv3Backend.
// Cookies are a weak form of security in the L2TPv3 specification. Their function
// is mostly to prevent misconfiguration. By default, they are 32 bit.
func WithTransmitCookie(cookie string) *Property {
	return NewProperty("txcookie", cookie)
}

// IsCookie64Bit specifies if the cookie size should be 64-bit instead of the
// default 32-bit for a L2TPv3Backend.
func IsCookie64Bit(is64bit bool) *Property {
	return NewProperty("cookie64", is64bit)
}

// IsCounter specifies whether to force a "cut-down" L2TPv3 with no counter as
// in "draft-mkonstan-l2tpext-keyed-ipv6-tunnel-00" for a L2TPv3Backend.
func IsCounter(enabled bool) *Property {
	return NewProperty("counter", enabled)
}

// IsPinCounter specifies whether to work around broken counter handling in peer
// for a L2TPv3Backend. This may also help on networks which have packet reorder.
func IsPinCounter(enabled bool) *Property {
	return NewProperty("pincounter", enabled)
}

// WithOffset adds an extra offset between header and data for a L2TPv3Backend.
func WithOffset(offset int) *Property {
	return NewProperty("offset", offset)
}

// WithSocketPath specifies the socket path for which to listen for incoming
// connections on a VDEBackend.
func WithSocketPath(path string) *Property {
	return NewProperty("sock", path)
}

// WithPort specifies the port for which to listen for incoming connections on
// a VDEBackend.
func WithPort(port int) *Property {
	return NewProperty("port", port)
}

// WithGroup specifies a group name for changing default ownership and permissions
// for a communication port on a VDEBackend.
func WithGroup(name string) *Property {
	return NewProperty("group", name)
}

// WithMode specifies an octal mode for changing default ownership and permissions
// for a communication port on a VDEBackend.
func WithMode(mode int) *Property {
	return NewProperty("mode", mode)
}

// IsVHostForce specifies whether to forcefully use a specifically defined protocol
// to pass  vhost ioctl replacement messages to an application on the other end of the
// socket for a VHostUserBackend. This feature can only be enabled on non-MSIX guests.
func IsVHostForce(force bool) *Property {
	return NewProperty("vhostforce", force)
}

// WithQueues specifies the number of queues to be created for multi-queue
// vhost-user for a VHostUserBackend.
func WithQueues(count int) *Property {
	return NewProperty("queues", count)
}
