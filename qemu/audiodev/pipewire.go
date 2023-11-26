package audiodev

// PipeWireBackend represents an audio backend using PipeWire. This backend is
// available on most systems.
type PipeWireBackend struct {
	*Backend
}

// NewPipeWireBackend returns a new instance of [PipeWireBackend].
//
//	qemu-system-* -audiodev pipewire,id=id
func NewPipeWireBackend(id string) *PipeWireBackend {
	return &PipeWireBackend{
		New("pipewire", id),
	}
}

// SetLatency sets the desired latency in microseconds.
//
//	qemu-system-* -audiodev pipewire,in|out.latency=usecs
func (b *PipeWireBackend) SetLatency(direction Direction, latency int) *PipeWireBackend {
	b.SetDirectionProperty(direction, "latency", latency)
	return b
}

// SetSink specified the source/sink to use for recording/playback.
//
//	qemu-system-* -audiodev pipewire,in|out.name=sink
func (b *PipeWireBackend) SetSink(direction Direction, sink string) *PipeWireBackend {
	b.SetDirectionProperty(direction, "name", sink)
	return b
}

// SetStreamName sets the name of the PipeWire stream.
//
//	qemu-system-* -audiodev pipewire,in|out.stream=name
func (b *PipeWireBackend) SetStreamName(direction Direction, name string) *PipeWireBackend {
	b.SetDirectionProperty(direction, "stream", name)
	return b
}
