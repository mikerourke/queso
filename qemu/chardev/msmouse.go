package chardev

// MSMouseBackend forwards QEMU's emulated msmouse events to the guest. This
// backend does not take any options.
type MSMouseBackend struct {
	*Backend
}

// NewMSMouseBackend returns a new instance of [MSMouseBackend].
// id is the unique ID, which can be any string up to 127 characters long.
//
//	qemu-system-* -chardev msmouse,id=id
func NewMSMouseBackend(id string) *MSMouseBackend {
	return &MSMouseBackend{
		New("msmouse", id),
	}
}
