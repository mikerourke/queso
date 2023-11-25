package chardev

// SerialBackend sends traffic from the guest to a serial device on the host.
// On Unix hosts serial will actually accept any TTY device, not only serial lines.
type SerialBackend struct {
	*Backend
	// Path specifies the name of the serial device to open.
	Path string
}

// NewSerialBackend returns a new instance of [SerialBackend]. id is the unique ID,
// which can be any string up to 127 characters long. path specifies the name of
// the serial device to open.
//
//	qemu-system-* -chardev serial,id=id,path=path
func NewSerialBackend(id string, path string) *SerialBackend {
	return &SerialBackend{
		Backend: New("serial", id),
		Path:    path,
	}
}
