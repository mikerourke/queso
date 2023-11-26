package nic

import "github.com/mikerourke/queso/qemu/netdev"

// BridgeNIC represents a [netdev.BridgeBackend] and the corresponding network
// hardware.
type BridgeNIC struct {
	*netdev.BridgeBackend
}

// NewBridgeNIC returns a new instance of [BridgeNIC].
//
//	qemu-system-* -nic bridge
func NewBridgeNIC() *BridgeNIC {
	nic := &BridgeNIC{netdev.NewBridgeBackend("")}

	nic.SetOptionFlag("nic").RemoveProperty("id")

	return nic
}

// SetMACAddress sets the MAC address of the NIC.
//
//	qemu-system-* -nic bridge,mac=addr
func (n *BridgeNIC) SetMACAddress(addr string) *BridgeNIC {
	n.SetProperty("mac", addr)
	return n
}

// SetModelName sets the model name of the NIC.
//
//	qemu-system-* -nic bridge,model=name
func (n *BridgeNIC) SetModelName(name string) *BridgeNIC {
	n.SetProperty("model", name)
	return n
}
