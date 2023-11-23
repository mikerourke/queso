package audiodev

import (
	"github.com/mikerourke/queso/qemu/cli"
)

// DirectSoundBackend represents an audio backend using Microsoft’s DirectSound.
// This backend is only available on Windows and only supports playback.
type DirectSoundBackend struct {
	*Backend
}

// NewDirectSoundBackend returns a new instance of [DirectSoundBackend].
//
//	qemu-system-* -audiodev dsound,id=id
func NewDirectSoundBackend(id string) *DirectSoundBackend {
	return &DirectSoundBackend{
		NewBackend("dsound", id),
	}
}

// SetLatency adds extra microseconds of latency to playback. The
// default is 10,000 (10 us).
//
//	qemu-system-* -audiodev dsound,latency=usecs
func (b *DirectSoundBackend) SetLatency(usecs int) *DirectSoundBackend {
	b.properties = append(b.properties, cli.NewProperty("latency", usecs))
	return b
}
