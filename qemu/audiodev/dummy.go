package audiodev

// DummyBackend represents a dummy backend that discards all outputs.
// This backend has no backend specific properties.
type DummyBackend struct {
	*Backend
}

// NewDummyBackend returns a new instance of [DummyBackend].
//
//	qemu-system-* -audiodev none,id=id
func NewDummyBackend(id string) *DummyBackend {
	return &DummyBackend{
		NewBackend("none", id),
	}
}
