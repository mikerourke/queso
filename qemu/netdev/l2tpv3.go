package netdev

// L2TPv3Backend represents a L2TPv3 pseudowire host network backend. L2TPv3
// (RFC3931) is a popular protocol to transport Ethernet (and other Layer 2)
// data frames between two systems. It is present in routers, firewalls and
// the Linux kernel (from version 3.3 onwards).
type L2TPv3Backend struct {
	*Backend
}

// NewL2TPv3Backend returns a new instance of [L2TPv3Backend]. id is the unique
// identifier for the backend. source is the source address and target is the
// target address.
//
//	qemu-system-* -netdev l2tpv3,id=id,src=srcaddr,dst=dstaddr
func NewL2TPv3Backend(id string, source string, destination string) *L2TPv3Backend {
	backend := New("l2tpv3")

	backend.SetProperty("id", id).
		SetProperty("src", source).
		SetProperty("dst", destination)

	return &L2TPv3Backend{backend}
}

// SetDestinationAddress sets the destination address.
//
//	qemu-system-* -netdev l2tpv3,dst=addr
func (b *L2TPv3Backend) SetDestinationAddress(addr int) *L2TPv3Backend {
	b.SetProperty("dst", addr)
	return b
}

// SetDestinationPort sets the destination UDP port.
//
//	qemu-system-* -netdev l2tpv3,dstport=port
func (b *L2TPv3Backend) SetDestinationPort(port int) *L2TPv3Backend {
	b.SetProperty("srcport", port)
	return b
}

// SetOffset adds an extra offset between header and data.
//
//	qemu-system-* -netdev l2tpv3,offset=offset
func (b *L2TPv3Backend) SetOffset(offset int) *L2TPv3Backend {
	b.SetProperty("offset", offset)
	return b
}

// SetReceiveCookie specifies the receiver cookie. Cookies are a weak form of
// security in the L2TPv3 specification. Their function is mostly to prevent
// misconfiguration. By default, they are 32 bit.
//
//	qemu-system-* -netdev l2tpv3,rxcookie=cookie
func (b *L2TPv3Backend) SetReceiveCookie(cookie string) *L2TPv3Backend {
	b.SetProperty("rxcookie", cookie)
	return b
}

// SetSourceAddress sets the source address.
//
//	qemu-system-* -netdev l2tpv3,src=addr
func (b *L2TPv3Backend) SetSourceAddress(addr int) *L2TPv3Backend {
	b.SetProperty("src", addr)
	return b
}

// SetSourcePort sets the source UDP port.
//
//	qemu-system-* -netdev l2tpv3,srcport=port
func (b *L2TPv3Backend) SetSourcePort(port int) *L2TPv3Backend {
	b.SetProperty("srcport", port)
	return b
}

// SetTransmitCookie specifies the transmission cookie. Cookies are a weak form
// of security in the L2TPv3 specification. Their function is mostly to prevent
// misconfiguration. By default, they are 32 bit.
//
//	qemu-system-* -netdev l2tpv3,txcookie=cookie
func (b *L2TPv3Backend) SetTransmitCookie(cookie string) *L2TPv3Backend {
	b.SetProperty("txcookie", cookie)
	return b
}

// Toggle64BitCookie specifies if the cookie size should be 64-bit instead of the
// default 32-bit.
//
//	qemu-system-* -netdev l2tpv3,cookie64=on|off
func (b *L2TPv3Backend) Toggle64BitCookie(enabled bool) *L2TPv3Backend {
	b.SetProperty("cookie64", enabled)
	return b
}

// ToggleCounter forces a "cut-down" L2TPv3 with no counter as in
// draft-mkonstan-l2tpext-keyed-ipv6-tunnel-00.
//
//	qemu-system-* -netdev l2tpv3,counter=on|off
func (b *L2TPv3Backend) ToggleCounter(enabled bool) *L2TPv3Backend {
	b.SetProperty("counter", enabled)
	return b
}

// ToggleForceIPv6 forces IPv6, otherwise defaults to IPv4.
//
//	qemu-system-* -netdev l2tpv3,ipv6=on|off
func (b *L2TPv3Backend) ToggleForceIPv6(enabled bool) *L2TPv3Backend {
	b.SetProperty("ipv6", enabled)
	return b
}

// TogglePinCounter specifies whether to work around broken counter handling in peer.
// This may also help on networks which have packet reorder.
//
//	qemu-system-* -netdev l2tpv3,pincounter=on|off
func (b *L2TPv3Backend) TogglePinCounter(enabled bool) *L2TPv3Backend {
	b.SetProperty("pincounter", enabled)
	return b
}
