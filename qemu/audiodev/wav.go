package audiodev

// WAVBackend represents a backend that writes audio to a WAV file.
type WAVBackend struct {
	*Backend
}

// NewWAVBackend returns a new instance of [WAVBackend].
//
//	qemu-system-* -audiodev spice,id=id
func NewWAVBackend(id string) *WAVBackend {
	return &WAVBackend{
		New("wav", id),
	}
}

// SetPath specifies the file path to which audio is written. The default is "qemu.wav".
//
//	qemu-system-* -audiodev wav,path=path
func (b *WAVBackend) SetPath(path string) *WAVBackend {
	b.SetProperty("path", path)
	return b
}
