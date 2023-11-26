package audiodev

// SNDIOBackend represents an audio backend using SNDIO. This backend is
// available on OpenBSD and most other Unix-like systems.
type SNDIOBackend struct {
	*Backend
}

// NewSNDIOBackend returns a new instance of [SNDIOBackend].
//
//	qemu-system-* -audiodev sndio,id=id
func NewSNDIOBackend(id string) *SNDIOBackend {
	return &SNDIOBackend{
		New("sndio", id),
	}
}

// SetDevice specifies the SNDIO device to use for input and/or output.
// The default is "default".
//
//	qemu-system-* -audiodev sndio,in|out.dev=device
func (b *SNDIOBackend) SetDevice(direction Direction, device string) *SNDIOBackend {
	b.SetDirectionProperty(direction, "dev", device)
	return b
}

// SetLatency sets the desired latency in microseconds.
//
//	qemu-system-* -audiodev sndio,in|out.latency=usecs
func (b *SNDIOBackend) SetLatency(direction Direction, latency int) *SNDIOBackend {
	b.SetDirectionProperty(direction, "latency", latency)
	return b
}
