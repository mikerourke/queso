package chardev

// ParallelBackend connects to a local parallel port. This backend is only available
// on Linux, FreeBSD, and DragonFlyBSD hosts.
type ParallelBackend struct {
	*Backend
	// Path specifies the path to the parallel port device.
	Path string
}

// NewParallelBackend returns a new instance of [ParallelBackend]. id is the unique ID,
// which can be any string up to 127 characters long. path specifies the path to
// the parallel port device.
//
//	qemu-system-* -chardev parallel,id=id,path=path
func NewParallelBackend(id string, path string) *ParallelBackend {
	return &ParallelBackend{
		Backend: NewBackend("parallel", id),
		Path:    path,
	}
}
