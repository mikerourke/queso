package chardev

import "github.com/mikerourke/queso/qemu/cli"

// RingBufferBackend represents a ring buffer with fixed size.
type RingBufferBackend struct {
	*Backend
}

// NewRingBufferBackend returns a new instance of [RingBufferBackend].
// id is the unique ID, which can be any string up to 127 characters long.
//
//	qemu-system-* -chardev ringbuf,id=id
func NewRingBufferBackend(id string) *RingBufferBackend {
	return &RingBufferBackend{
		NewBackend("ringbuf", id),
	}
}

// SetSize sets the size of the ring buffer. It must be a power of two and defaults
// to 64K if not specified.
//
//	qemu-system-* -chardev ringbuf,size=size
func (b *RingBufferBackend) SetSize(size string) *RingBufferBackend {
	b.properties = append(b.properties, cli.NewProperty("size", size))
	return b
}
