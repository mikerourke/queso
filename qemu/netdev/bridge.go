package netdev

import "github.com/mikerourke/queso"

// BridgeBackend connects a host TAP network interface to a host bridge device.
type BridgeBackend struct {
	*Backend
}

// NewBridgeBackend returns a new instance of [BridgeBackend]. id is the unique
// identifier for the backend.
//
//	qemu-system-* -netdev bridge,id=id
func NewBridgeBackend(id string) *BridgeBackend {
	return &BridgeBackend{
		New("bridge").SetProperty("id", id),
	}
}

// SetBridge specifies the bridge device.
//
//	qemu-system-* -netdev bridge,br=bridge
func (b *BridgeBackend) SetBridge(bridge string) *BridgeBackend {
	b.properties = append(b.properties, queso.NewProperty("br", bridge))
	return b
}

// SetHelper specifies an executable path to configure the TAP interface and
// attach it to the [BridgeBackend].
//
//	qemu-system-* -netdev bridge,helper=helper
func (b *BridgeBackend) SetHelper(helper string) *BridgeBackend {
	b.properties = append(b.properties, queso.NewProperty("helper", helper))
	return b
}
