package nic

import "github.com/mikerourke/queso/qemu/netdev"

// AFXDPNIC represents a [netdev.AFXDPBackend] and the corresponding network
// hardware.
type AFXDPNIC struct {
	*netdev.AFXDPBackend
}

// NewAFXDPNIC returns a new instance of [AFXDPNIC].
//
//	qemu-system-* -nic af-xdp
func NewAFXDPNIC() *AFXDPNIC {
	nic := &AFXDPNIC{netdev.NewAFXDPBackend("", "")}

	nic.SetOptionFlag("nic").RemoveProperty("id")

	return nic
}

// SetMACAddress sets the MAC address of the NIC.
//
//	qemu-system-* -nic af-xdp,mac=addr
func (n *AFXDPNIC) SetMACAddress(addr string) *AFXDPNIC {
	n.SetProperty("mac", addr)
	return n
}

// SetModelName sets the model name of the NIC.
//
//	qemu-system-* -nic af-xdp,model=name
func (n *AFXDPNIC) SetModelName(name string) *AFXDPNIC {
	n.SetProperty("model", name)
	return n
}
