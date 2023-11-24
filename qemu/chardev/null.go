package chardev

// NullBackend represents a void device. This device will not emit any data, and
// will drop any data it receives. This backend does not take any options.
type NullBackend struct {
	*Backend
}

// NewNullBackend returns a new instance of [NullBackend].
// id is the unique ID, which can be any string up to 127 characters long.
//
//	qemu-system-* -chardev null,id=id
func NewNullBackend(id string) *NullBackend {
	return &NullBackend{
		NewBackend("null", id),
	}
}
