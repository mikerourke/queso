package nic

import "github.com/mikerourke/queso/qemu/netdev"

// TAPNIC represents a [netdev.TAPBackend] and the corresponding network
// hardware.
type TAPNIC struct {
	*netdev.TAPBackend
}

// NewTAPNIC returns a new instance of [TAPNIC].
//
//	qemu-system-* -nic tap
func NewTAPNIC() *TAPNIC {
	nic := &TAPNIC{netdev.NewTAPBackend("")}

	nic.SetOptionFlag("nic").RemoveProperty("id")

	return nic
}

// SetMACAddress sets the MAC address of the NIC.
//
//	qemu-system-* -nic tap,mac=addr
func (n *TAPNIC) SetMACAddress(addr string) *TAPNIC {
	n.SetProperty("mac", addr)
	return n
}

// SetModelName sets the model name of the NIC.
//
//	qemu-system-* -nic tap,model=name
func (n *TAPNIC) SetModelName(name string) *TAPNIC {
	n.SetProperty("model", name)
	return n
}
