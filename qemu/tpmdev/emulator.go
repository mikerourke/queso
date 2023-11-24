package tpmdev

import "github.com/mikerourke/queso"

// EmulatorBackend enables access to a TPM emulator using Unix domain socket-based
// chardev backend and is only available on Linux hosts.
type EmulatorBackend struct {
	*Backend
}

// NewEmulatorBackend returns a new instance of [EmulatorBackend]. id
// is the unique identifier for the TPM device backend. chardev specifies the unique
// ID of a character device backend that provides connection to the software
// TPM server.
//
//	qemu-system-* -tpmdev emulator,id=id,chardev=chardev
func NewEmulatorBackend(id string, chardev string) *EmulatorBackend {
	backend := New("emulator")

	backend.properties = append(backend.properties,
		queso.NewProperty("id", id),
		queso.NewProperty("chardev", chardev))

	return &EmulatorBackend{backend}
}
