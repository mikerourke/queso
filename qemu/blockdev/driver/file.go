package driver

// FileDriver is the protocol-level block driver for accessing regular files.
type FileDriver struct {
	*Driver
}

// NewFileDriver returns a new instance of [FileDriver].
//
//	qemu-system-* -blockdev driver=file
func NewFileDriver() *FileDriver {
	return &FileDriver{
		New("file"),
	}
}

// AIOBackend represents the asynchronous I/O backend values that can be
// passed to the [FileDriver.SetAIOBackend] method.
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
//	qemu-system-* -blockdev driver=file,aio=backend
func (d *FileDriver) SetAIOBackend(backend string) *FileDriver {
	d.SetProperty("aio", backend)
	return d
}

// SetFileName sets the path to the image file in the local filesystem.
//
//	qemu-system-* -blockdev driver=file,filename=name
func (d *FileDriver) SetFileName(name string) *FileDriver {
	d.SetProperty("filename", name)
	return d
}

// OFDLockingStatus represents the possible status options for
// [FileDriver.SetOFDLockingStatus].
type OFDLockingStatus string

const (
	// OFDLockingOn indicates that OFD/POSIX locks are applied.
	OFDLockingOn = "on"

	// OFDLockingOff indicates that no OFD/POSIX locks are applied.
	OFDLockingOff = "off"

	// OFDLockingAuto indicates that QEMU should try to use the Linux Open File
	// Descriptor API if available.
	OFDLockingAuto = "auto"
)

// SetOFDLockingStatus specifies whether the image file is protected with Linux
// OFD/POSIX locks. The default is to use the Linux Open File
// Descriptor API if available ([OFDLockingAuto]), otherwise no lock is
// applied ([OFDLockingOff]).
//
//	qemu-system-* -blockdev driver=file,locking=on|off|auto
func (d *FileDriver) SetOFDLockingStatus(status OFDLockingStatus) *FileDriver {
	d.SetProperty("locking", string(status))
	return d
}
