package nic

import "github.com/mikerourke/queso/qemu/netdev"

// VDENIC represents a [netdev.VDEBackend] and the corresponding network
// hardware.
type VDENIC struct {
	*netdev.VDEBackend
}

// NewVDENIC returns a new instance of [VDENIC].
//
//	qemu-system-* -nic vde
func NewVDENIC() *VDENIC {
	nic := &VDENIC{netdev.NewVDEBackend("")}

	nic.SetOptionFlag("nic").RemoveProperty("id")

	return nic
}

// SetMACAddress sets the MAC address of the NIC.
//
//	qemu-system-* -nic vde,mac=addr
func (n *VDENIC) SetMACAddress(addr string) *VDENIC {
	n.SetProperty("mac", addr)
	return n
}

// SetModelName sets the model name of the NIC.
//
//	qemu-system-* -nic vde,model=name
func (n *VDENIC) SetModelName(name string) *VDENIC {
	n.SetProperty("model", name)
	return n
}
