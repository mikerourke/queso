package netdev

// VDEBackend represents a backend using Virtual Distributed Ethernet (VDE). VDE
// is a flexible, virtual network infrastructure system, spanning across multiple
// hosts in a secure way. It enables L2/L3 switching, including spanning-tree
// protocol, VLANs, and WAN emulation.
type VDEBackend struct {
	*Backend
}

// NewVDEBackend returns a new instance of [VDEBackend]. id is a unique
// identifier for the backend.
//
//	qemu-system-* -netdev vde
func NewVDEBackend(id string) *VDEBackend {
	backend := New("vde")

	backend.SetProperty("id", id)

	return &VDEBackend{backend}
}

// SetGroup specifies a group name for changing default ownership and permissions
// for a communication port.
//
//	qemu-system-* -netdev vde,group=name
func (b *VDEBackend) SetGroup(name string) *VDEBackend {
	b.SetProperty("group", name)
	return b
}

// SetMode specifies an octal mode for changing default ownership and permissions
// for a communication port.
//
//	qemu-system-* -netdev vde,mode=mode
func (b *VDEBackend) SetMode(mode int) *VDEBackend {
	b.SetProperty("mode", mode)
	return b
}

// SetPort specifies the port for which to listen for incoming connections.
//
//	qemu-system-* -netdev vde,port=port
func (b *VDEBackend) SetPort(port string) *VDEBackend {
	b.SetProperty("port", port)
	return b
}

// SetSocketPath specifies the socket path for which to listen for incoming
// connections.
//
//	qemu-system-* -netdev vde,sock=path
func (b *VDEBackend) SetSocketPath(path string) *VDEBackend {
	b.SetProperty("sock", path)
	return b
}
