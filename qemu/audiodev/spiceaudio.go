package audiodev

// SpiceBackend represents a backend that sends audio through Spice. This backend
// requires the [qemu.SPICE] and automatically selected in that case, so usually
// you can ignore this option. This backend has no backend specific properties.
type SpiceBackend struct {
	*Backend
}

// NewSpiceBackend returns a new instance of [SpiceBackend].
//
//	qemu-system-* -audiodev spice,id=id
func NewSpiceBackend(id string) *SpiceBackend {
	return &SpiceBackend{
		New("spice", id),
	}
}
