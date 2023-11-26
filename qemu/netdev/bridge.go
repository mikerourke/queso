package netdev

// BridgeBackend connects a host TAP network interface to a host bridge device.
type BridgeBackend struct {
	*Backend
}

// NewBridgeBackend returns a new instance of [BridgeBackend]. id is the unique
// identifier for the backend.
//
//	qemu-system-* -netdev bridge,id=id
func NewBridgeBackend(id string) *BridgeBackend {
	backend := New("bridge")

	backend.SetProperty("id", id)

	return &BridgeBackend{backend}
}

// SetBridge specifies the bridge device.
//
//	qemu-system-* -netdev bridge,br=bridge
func (b *BridgeBackend) SetBridge(bridge string) *BridgeBackend {
	b.SetProperty("br", bridge)
	return b
}

// SetHelper specifies an executable path to configure the TAP interface and
// attach it to the [BridgeBackend].
//
//	qemu-system-* -netdev bridge,helper=helper
func (b *BridgeBackend) SetHelper(helper string) *BridgeBackend {
	b.SetProperty("helper", helper)
	return b
}
