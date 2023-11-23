package audiodev

// CoreAudioBackend represents an audio backend using Apple's Core Audio. This
// backend is only available on macOS and only supports playback.
type CoreAudioBackend struct {
	*Backend
}

// NewCoreAudioBackend returns a new instance of [CoreAudioBackend].
//
//	qemu-system-* -audiodev coreaudio,id=id
func NewCoreAudioBackend(id string) *CoreAudioBackend {
	return &CoreAudioBackend{
		NewBackend("coreaudio", id),
	}
}

// SetBufferCount sets the count of buffers.
//
//	qemu-system-* -audiodev coreaudio,buffer-count=count
func (b *CoreAudioBackend) SetBufferCount(direction Direction, count string) *CoreAudioBackend {
	b.properties = append(b.properties, newDirectionProperty("buffer-count", direction, count))
	return b
}
