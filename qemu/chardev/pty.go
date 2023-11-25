package chardev

// PTYBackend represents a new pseudo-terminal on the host and connect to it. This
// backend does not take any options and is not available on Windows hosts.
type PTYBackend struct {
	*Backend
}

// NewPTYBackend returns a new instance of [PTYBackend].
// id is the unique ID, which can be any string up to 127 characters long.
//
//	qemu-system-* -chardev pty,id=id
func NewPTYBackend(id string) *PTYBackend {
	return &PTYBackend{
		New("pty", id),
	}
}
