package audiodev

import "github.com/mikerourke/queso/internal/cli"

// WAVBackend represents a backend that writes audio to a WAV file.
type WAVBackend struct {
	*Backend
}

// NewWAVBackend returns a new instance of [WAVBackend].
//
//	qemu-system-* -audiodev spice,id=id
func NewWAVBackend(id string) *WAVBackend {
	return &WAVBackend{
		NewBackend("wav", id),
	}
}

// SetPath specifies the file path to which audio is written. The default is "qemu.wav".
//
//	qemu-system-* -audiodev wav,path=path
func (b *WAVBackend) SetPath(path string) *WAVBackend {
	b.properties = append(b.properties, cli.NewProperty("path", path))
	return b
}
