package driver

import (
	"github.com/mikerourke/queso/internal/vals"
	"github.com/mikerourke/queso/qemu/cli"
)

// FileDriver is the protocol-level block driver for accessing regular files.
type FileDriver struct {
	*Driver
}

// NewFileDriver returns a new instance of [FileDriver].
//
//	qemu-system-* -blockdev driver=file
func NewFileDriver() *FileDriver {
	return &FileDriver{
		NewDriver("file"),
	}
}

// AIOBackend represents the asynchronous I/O backend values that can be
// passed to the SetAIOBackend method.
type AIOBackend string

const (
	// AIOBackendThreads uses threads for asynchronous I/O.
	AIOBackendThreads AIOBackend = "threads"

	// AIOBackendIOURing uses io_uring for asynchronous I/O.
	AIOBackendIOURing AIOBackend = "io_uring"

	// AIOBackendNative uses Linux native asynchronous I/O.
	AIOBackendNative AIOBackend = "native"
)

// SetAIOBackend specified the AIO backend. The default value is [AIOBackendThreads].
//
//	qemu-system-* -blockdev driver=file,filename=name
func (d *FileDriver) SetAIOBackend(backend string) *FileDriver {
	d.properties = append(d.properties, cli.NewProperty("aio", backend))
	return d
}

// SetFileName sets the path to the image file in the local filesystem.
//
//	qemu-system-* -blockdev driver=file,filename=name
func (d *FileDriver) SetFileName(name string) *FileDriver {
	d.properties = append(d.properties, cli.NewProperty("filename", name))
	return d
}

// SetOFDLockingStatus specifies whether the image file is protected with Linux
// OFD/POSIX locks. The default is to use the Linux Open File
// Descriptor API if available ([vals.Status.Auto]), otherwise no lock is
// applied ([vals.Status.Off]).
//
//	qemu-system-* -blockdev driver=file,locking=on|off|auto
func (d *FileDriver) SetOFDLockingStatus(status vals.Status) *FileDriver {
	d.properties = append(d.properties, cli.NewProperty("locking", status.String()))
	return d
}
