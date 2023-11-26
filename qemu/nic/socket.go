package nic

import "github.com/mikerourke/queso/qemu/netdev"

// TCPSocketNIC represents a [netdev.TCPSocketBackend] and the corresponding
// network hardware.
type TCPSocketNIC struct {
	*netdev.TCPSocketBackend
}

// NewTCPSocketNIC returns a new instance of [TCPSocketNIC].
//
//	qemu-system-* -nic socket
func NewTCPSocketNIC() *TCPSocketNIC {
	nic := &TCPSocketNIC{netdev.NewTCPSocketBackend("")}

	nic.SetOptionFlag("nic").RemoveProperty("id")

	return nic
}

// SetMACAddress sets the MAC address of the NIC.
//
//	qemu-system-* -nic socket,mac=addr
func (n *TCPSocketNIC) SetMACAddress(addr string) *TCPSocketNIC {
	n.SetProperty("mac", addr)
	return n
}

// SetModelName sets the model name of the NIC.
//
//	qemu-system-* -nic socket,model=name
func (n *TCPSocketNIC) SetModelName(name string) *TCPSocketNIC {
	n.SetProperty("model", name)
	return n
}

// UDPSocketNIC represents a [netdev.UDPSocketBackend] and the corresponding
// network hardware.
type UDPSocketNIC struct {
	*netdev.UDPSocketBackend
}

// NewUDPSocketNIC returns a new instance of [UDPSocketNIC].
//
//	qemu-system-* -nic socket
func NewUDPSocketNIC() *UDPSocketNIC {
	nic := &UDPSocketNIC{netdev.NewUDPSocketBackend("")}

	nic.SetOptionFlag("nic").RemoveProperty("id")

	return nic
}

// SetMACAddress sets the MAC address of the NIC.
//
//	qemu-system-* -nic socket,mac=addr
func (n *UDPSocketNIC) SetMACAddress(addr string) *UDPSocketNIC {
	n.SetProperty("mac", addr)
	return n
}

// SetModelName sets the model name of the NIC.
//
//	qemu-system-* -nic socket,model=name
func (n *UDPSocketNIC) SetModelName(name string) *UDPSocketNIC {
	n.SetProperty("model", name)
	return n
}
