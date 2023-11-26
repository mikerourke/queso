package netdev

// HubPortBackend is a network backend that lets you connect a NIC to a QEMU
// emulated hub instead of a single network device. Alternatively, you can also
// connect the hub port to another network device using
// [HubPortBackend.SetNetworkDevice].
type HubPortBackend struct {
	*Backend
}

// NewHubPortBackend returns a new instance of [HubPortBackend]. id is a unique
// identifier for the backend. hubID is the ID of the emulated hub to connect to.
//
//	qemu-system-* -netdev hubport,id=id,hubid=hubid
func NewHubPortBackend(id string, hubID string) *HubPortBackend {
	backend := New("hubport")

	backend.SetProperty("id", id).SetProperty("hubid", hubID)

	return &HubPortBackend{backend}
}

// SetNetworkDevice sets the network device ID that the hub port can connect
// to.
//
//	qemu-system-* ,=
func (b *HubPortBackend) SetNetworkDevice(id string) *HubPortBackend {
	b.SetProperty("netdev", id)
	return b
}
