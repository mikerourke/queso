package chardev

// ConsoleBackend sends traffic from the guest to QEMU's standard output.
// This backend does not take any options and is only available on Windows.
type ConsoleBackend struct {
	*Backend
}

// NewConsoleBackend returns a new instance of [ConsoleBackend].
// id is the unique ID, which can be any string up to 127 characters long.
//
//	qemu-system-* -chardev console,id=id
func NewConsoleBackend(id string) *ConsoleBackend {
	return &ConsoleBackend{
		New("console", id),
	}
}
