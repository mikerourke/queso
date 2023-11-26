package nic

import "github.com/mikerourke/queso/qemu/netdev"

// VHostUserNIC represents a [netdev.VHostUserBackend] and the corresponding
// network hardware.
type VHostUserNIC struct {
	*netdev.VHostUserBackend
}

// NewVHostUserNIC returns a new instance of [VHostUserNIC].
//
//	qemu-system-* -nic vhost-user
func NewVHostUserNIC() *VHostUserNIC {
	// TODO: Find out if we still need the chardev property.
	nic := &VHostUserNIC{netdev.NewVHostUserBackend("")}

	nic.SetOptionFlag("nic").RemoveProperty("chardev")

	return nic
}

// SetMACAddress sets the MAC address of the NIC.
//
//	qemu-system-* -nic vhost-user,mac=addr
func (n *VHostUserNIC) SetMACAddress(addr string) *VHostUserNIC {
	n.SetProperty("mac", addr)
	return n
}

// SetModelName sets the model name of the NIC.
//
//	qemu-system-* -nic vhost-user,model=name
func (n *VHostUserNIC) SetModelName(name string) *VHostUserNIC {
	n.SetProperty("model", name)
	return n
}
