package spice

import (
	"github.com/mikerourke/queso"
	"github.com/mikerourke/queso/qemu/audiodev"
)

// AudioBackend creates an audio backend for SPICE. This backend has no backend
// specific properties.
func AudioBackend(id string) *queso.Option {
	return audiodev.Backend(audiodev.BackendTypeSpice, id)
}
