package chardev

import "github.com/mikerourke/queso"

// StdioBackend connects to standard input and standard output of the QEMU process.
type StdioBackend struct {
	*Backend
}

// NewStdioBackend returns a new instance of [StdioBackend].
// id is the unique ID, which can be any string up to 127 characters long.
//
//	qemu-system-* -chardev stdio,id=id
func NewStdioBackend(id string) *StdioBackend {
	return &StdioBackend{
		NewBackend("stdio", id),
	}
}

// ToggleSignals controls if signals are enabled on the terminal, that includes exiting
// QEMU with the key sequence Control-c. This option is enabled by default, set to
// false to disable it.
//
//	qemu-system-* -chardev stdio,signals=on|off
func (b *StdioBackend) ToggleSignals(enabled bool) *StdioBackend {
	b.properties = append(b.properties, queso.NewProperty("signals", enabled))
	return b
}
