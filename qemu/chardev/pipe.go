package chardev

// PipeBackend represents a two-way connection to the guest. The behaviour differs
// slightly between Windows hosts and other hosts.
//
// On Windows, a single duplex pipe will be created at "\\.pipe\path".
//
// On other hosts, 2 pipes will be created called "path.in" and "path.out". Data
// written to "path.in" will be received by the guest. Data written by the guest
// can be read from "path.out". QEMU will not create these fifos, and requires
// them to be present.
type PipeBackend struct {
	*Backend
	// Path forms part of the pipe path as described above.
	Path string
}

// NewPipeBackend returns a new instance of [PipeBackend]. id is the unique ID,
// which can be any string up to 127 characters long. path forms part of the pipe
// path as described in [PipeBackend].
//
//	qemu-system-* -chardev pipe,id=id,path=path
func NewPipeBackend(id string, path string) *PipeBackend {
	return &PipeBackend{
		Backend: NewBackend("pipe", id),
		Path:    path,
	}
}
