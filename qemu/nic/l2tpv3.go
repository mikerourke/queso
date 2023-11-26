package nic

import "github.com/mikerourke/queso/qemu/netdev"

// L2TPv3NIC represents a [netdev.L2TPv3Backend] and the corresponding network
// hardware.
type L2TPv3NIC struct {
	*netdev.L2TPv3Backend
}

// NewL2TPv3NIC returns a new instance of [L2TPv3NIC].
//
//	qemu-system-* -nic af-xdp
func NewL2TPv3NIC() *L2TPv3NIC {
	// TODO: Find out if I still need src and dst here.
	nic := &L2TPv3NIC{netdev.NewL2TPv3Backend("", "", "")}

	nic.SetOptionFlag("nic").RemoveProperty("id")

	return nic
}

// SetMACAddress sets the MAC address of the NIC.
//
//	qemu-system-* -nic af-xdp,mac=addr
func (n *L2TPv3NIC) SetMACAddress(addr string) *L2TPv3NIC {
	n.SetProperty("mac", addr)
	return n
}

// SetModelName sets the model name of the NIC.
//
//	qemu-system-* -nic af-xdp,model=name
func (n *L2TPv3NIC) SetModelName(name string) *L2TPv3NIC {
	n.SetProperty("model", name)
	return n
}
