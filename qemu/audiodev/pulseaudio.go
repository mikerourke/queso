package audiodev

import "github.com/mikerourke/queso"

// PulseAudioBackend represents an audio backend using PulseAudio.
// This backend is available on most systems.
type PulseAudioBackend struct {
	*Backend
}

// NewPulseAudioBackend returns a new instance of [PulseAudioBackend].
//
//	qemu-system-* -audiodev pa,id=id
func NewPulseAudioBackend(id string) *PulseAudioBackend {
	return &PulseAudioBackend{
		NewBackend("pa", id),
	}
}

// SetLatency sets the desired latency in microseconds. The PulseAudio server
// will try to honor this value but actual latencies may be lower or higher.
//
//	qemu-system-* -audiodev pa,in|out.latency=usecs
func (b *PulseAudioBackend) SetLatency(direction Direction, usecs int) *PulseAudioBackend {
	b.properties = append(b.properties, newDirectionProperty("latency", direction, usecs))
	return b
}

// SetServer sets the PulseAudio server to connect to.
//
//	qemu-system-* -audiodev pa,server=server
func (b *PulseAudioBackend) SetServer(server string) *PulseAudioBackend {
	b.properties = append(b.properties, queso.NewProperty("server", server))
	return b
}

// SetSink specified the source/sink to use for recording/playback.
//
//	qemu-system-* -audiodev pa,in|out.name=sink
func (b *PulseAudioBackend) SetSink(direction Direction, name string) *PulseAudioBackend {
	b.properties = append(b.properties, newDirectionProperty("name", direction, name))
	return b
}
