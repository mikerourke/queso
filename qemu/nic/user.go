package nic

import "github.com/mikerourke/queso/qemu/netdev"

// UserNIC represents a [netdev.UserBackend] and the corresponding network
// hardware.
type UserNIC struct {
	*netdev.UserBackend
}

// NewUserNIC returns a new instance of [UserNIC].
//
//	qemu-system-* -nic user
func NewUserNIC() *UserNIC {
	nic := &UserNIC{netdev.NewUserBackend("")}

	nic.SetOptionFlag("nic").RemoveProperty("id")

	return nic
}

// SetMACAddress sets the MAC address of the NIC.
//
//	qemu-system-* -nic user,mac=addr
func (n *UserNIC) SetMACAddress(addr string) *UserNIC {
	n.SetProperty("mac", addr)
	return n
}

// SetModelName sets the model name of the NIC.
//
//	qemu-system-* -nic user,model=name
func (n *UserNIC) SetModelName(name string) *UserNIC {
	n.SetProperty("model", name)
	return n
}
