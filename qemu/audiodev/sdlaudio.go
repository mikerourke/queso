package audiodev

// SDLBackend represents an audio backend using SDL. This backend is available on
// most systems, but you should use your platform’s native backend if possible.
type SDLBackend struct {
	*Backend
}

// NewSDLBackend returns a new instance of [SDLBackend].
//
//	qemu-system-* -audiodev sdl,id=id
func NewSDLBackend(id string) *SDLBackend {
	return &SDLBackend{
		New("sdl", id),
	}
}

// SetBufferCount sets the count of buffers.
//
//	qemu-system-* -audiodev sdl,buffer-count=count
func (b *SDLBackend) SetBufferCount(direction Direction, count string) *SDLBackend {
	b.SetDirectionProperty(direction, "buffer-count", count)
	return b
}
