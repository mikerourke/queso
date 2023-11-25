package netdev

import "github.com/mikerourke/queso"

// UserBackend configures a user mode host network backend which requires no
// administrator privilege to run
type UserBackend struct {
	*Backend
}

// NewUserBackend returns a new instance of [UserBackend]. id is the unique
// identifier for the backend.
//
//	qemu-system-* -netdev user,id=id
func NewUserBackend(id string) *UserBackend {
	return &UserBackend{
		New("user").SetProperty("id", id),
	}
}

// AddGuestForwardRule adds a rule to forward guest TCP connections.
//
//	qemu-system-* -netdev user,guestfwd=[tcp]:server:port-dev
func (b *UserBackend) AddGuestForwardRule(rule *GuestForwardRule) *UserBackend {
	b.properties = append(b.properties, queso.NewProperty("guestfwd", rule.Value()))
	return b
}

// AddHostForwardRule adds a rule to forward host TCP connections.
//
//	qemu-system-* -netdev user,hostfwd=[tcp|udp]:[hostaddr]:hostport-[guestaddr]:guestport
func (b *UserBackend) AddHostForwardRule(rule *HostForwardRule) *UserBackend {
	b.properties = append(b.properties, queso.NewProperty("hostfwd", rule.Value()))
	return b
}

// SetBootFile broadcasts file as the BOOTP filename. In conjunction with
// WithTFTP, this can be used to network boot a guest from a local directory.
//
//	qemu-system-* -netdev user,bootfile=file
func (b *UserBackend) SetBootFile(file string) *UserBackend {
	b.properties = append(b.properties, queso.NewProperty("bootfile", file))
	return b
}

// SetDHCPStart specifies the first of the 16 IPs the built-in DHCP server can
// assign. The default is the 15th to 31st IP in the guest network,
// i.e. x.x.x.15 to x.x.x.31.
//
//	qemu-system-* -netdev user,dhcpstart=addr
func (b *UserBackend) SetDHCPStart(addr string) *UserBackend {
	b.properties = append(b.properties, queso.NewProperty("dhcpstart", addr))
	return b
}

// SetDNSIPv4Address specifies the guest-visible address of the IPv4 virtual
// nameserver. The address must be different from the host address. The
// default is the 3rd IP in the guest network, i.e. x.x.x.3.
//
//	qemu-system-* -netdev user,dns=addr
func (b *UserBackend) SetDNSIPv4Address(addr string) *UserBackend {
	b.properties = append(b.properties, queso.NewProperty("dns", addr))
	return b
}

// SetDNSIPv6Address specifies the guest-visible address of the IPv6 virtual
// nameserver. The address must be different from the host address. The
// default is the 3rd IP in the guest network, i.e. xxxx::3.
//
//	qemu-system-* -netdev user,ipv6-dns=addr
func (b *UserBackend) SetDNSIPv6Address(addr string) *UserBackend {
	b.properties = append(b.properties, queso.NewProperty("ipv6-dns", addr))
	return b
}

// SetDNSSearchDomain provides an entry for the domain-search list sent by the
// built-in DHCP server. More than one domain suffix can be transmitted by
// specifying this property multiple times. If supported, this will cause the
// guest to automatically try to append the given domain suffix(es) in case a
// domain name can not be resolved.
//
//	qemu-system-* -netdev user,dnssearch=domain
func (b *UserBackend) SetDNSSearchDomain(domain string) *UserBackend {
	b.properties = append(b.properties, queso.NewProperty("dnssearch", domain))
	return b
}

// SetDomainName specifies the client domain name reported by the built-in DHCP
// server.
//
//	qemu-system-* -netdev user,domainname=domain
func (b *UserBackend) SetDomainName(domain string) *UserBackend {
	b.properties = append(b.properties, queso.NewProperty("domainname", domain))
	return b
}

// SetGuestIPv4Address sets the IP network address the guest will see. Optionally
// specify the netmask, either in the form a.b.c.d or as number of valid top-most
// bits. The default is 10.0.2.0/24.
//
//	qemu-system-* -netdev user,net=addr[/mask]
func (b *UserBackend) SetGuestIPv4Address(addr string) *UserBackend {
	b.properties = append(b.properties, queso.NewProperty("net", addr))
	return b
}

// SetGuestIPv6Address sets IPv6 network address the guest will see (default is
// fec0::/64). The network prefix is given in the usual hexadecimal IPv6 address
// notation. The prefix size is optional, and is given as the number of valid
// top-most bits (default is 64).
//
//	qemu-system-* -netdev user,ipv6-net=addr
func (b *UserBackend) SetGuestIPv6Address(addr string) *UserBackend {
	b.properties = append(b.properties, queso.NewProperty("ipv6-net", addr))
	return b
}

// SetHostIPv4Address specifies the guest-visible address of the host. The default is
// the 2nd IP in the guest network, i.e. x.x.x.2.
//
//	qemu-system-* -netdev user,host=addr
func (b *UserBackend) SetHostIPv4Address(addr string) *UserBackend {
	b.properties = append(b.properties, queso.NewProperty("host", addr))
	return b
}

// SetHostIPv6Address specifies the guest-visible IPv6 address of the host. The
// default is the 2nd IPv6 in the guest network, i.e. xxxx::2.
//
//	qemu-system-* -netdev user,ipv6-host=addr
func (b *UserBackend) SetHostIPv6Address(addr string) *UserBackend {
	b.properties = append(b.properties, queso.NewProperty("ipv6-host", addr))
	return b
}

// SetHostName specifies the client hostname reported by the built-in DHCP server.
//
//	qemu-system-* -netdev user,hostname=name
func (b *UserBackend) SetHostName(name string) *UserBackend {
	b.properties = append(b.properties, queso.NewProperty("hostname", name))
	return b
}

// SetSMBDirectory activates a built-in SMB server so that Windows OSes can access to the
// host files in the specified dir transparently. The server address can be set
// with the [UserBackend.SetSMBServerAddress] method.
//
// In the guest Windows OS, the line:
//
//	10.0.2.4 smbserver
//
// Must be added in the file `C:\WINDOWS\LMHOSTS` (for Windows 9x/ME) or
// `C:\WINNT\SYSTEM32\DRIVERS\ETC\LMHOSTS` (Windows NT/2000).
//
// Then directory can be accessed in `\\smbserver\qemu`.
//
// Note that a SAMBA server must be installed on the host OS.
//
//	qemu-system-* -netdev user,smb=dir
func (b *UserBackend) SetSMBDirectory(dir string) *UserBackend {
	b.properties = append(b.properties, queso.NewProperty("smb", dir))
	return b
}

// SetSMBServerAddress sets the address of the SMB server. By default, the 4th
// IP in the guest network is used, i.e. x.x.x.4. This must be used in
// conjunction with the [UserBackend.SetSMBDirectory] method.
//
//	qemu-system-* -netdev user,smbserver=addr
func (b *UserBackend) SetSMBServerAddress(addr string) *UserBackend {
	b.properties = append(b.properties, queso.NewProperty("smbserver", addr))
	return b
}

// SetTFTPDirectory activates a built-in TFTP server. The files in the
// specified dir parameter will be exposed as the root of a TFTP server. The
// TFTP client on the guest must be configured in binary mode (use the command
// bin of the Unix TFTP client).
//
//	qemu-system-* -netdev user,tftp=dir
func (b *UserBackend) SetTFTPDirectory(dir string) *UserBackend {
	b.properties = append(b.properties, queso.NewProperty("tftp", dir))
	return b
}

// SetTFTPServerName broadcasts the specified name as the "TFTP server name"
// (RFC2132 option 66) in BOOTP reply. This can be used to advise the guest to
// load boot files or configurations from a different server than the host
// address.
//
//	qemu-system-* -netdev user,tftp-server-name=name
func (b *UserBackend) SetTFTPServerName(name string) *UserBackend {
	b.properties = append(b.properties, queso.NewProperty("tftp-server-name", name))
	return b
}

// ToggleIPv4 specifies if IPv4 is enabled.
//
//	qemu-system-* -netdev user,ipv4=on|off
func (b *UserBackend) ToggleIPv4(enabled bool) *UserBackend {
	b.properties = append(b.properties, queso.NewProperty("ipv4", enabled))
	return b
}

// ToggleIPv6 specifies if IPv6 is enabled.
//
//	qemu-system-* -netdev user,ipv6=on|off
func (b *UserBackend) ToggleIPv6(enabled bool) *UserBackend {
	b.properties = append(b.properties, queso.NewProperty("ipv6", enabled))
	return b
}

// ToggleRestricted specifies whether the guest will be isolated i.e. it will not
// be able to contact the host and no guest IP packets will be routed over the
// host to the outside. This option does not affect any explicitly set forwarding
// rules.
//
//	qemu-system-* -netdev user,restrict=on|off
func (b *UserBackend) ToggleRestricted(enabled bool) *UserBackend {
	b.properties = append(b.properties, queso.NewProperty("restrict", enabled))
	return b
}
