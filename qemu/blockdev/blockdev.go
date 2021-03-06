// Package blockdev is used to manage block devices for use with QEMU. See
// https://qemu.readthedocs.io/en/latest/system/invocation.html#hxtool-1 for
// more details.
package blockdev

import "github.com/mikerourke/queso"

const (
	// FloppyDiskDriveA represents the 0th index floppy disk drive in the guest.
	FloppyDiskDriveA = "fda"

	// FloppyDiskDriveB represents the 1st index floppy disk drive in the guest.
	FloppyDiskDriveB = "fdb"

	// HardDiskDriveA represents the 0th index hard disk drive in the guest.
	HardDiskDriveA = "hda"

	// HardDiskDriveB represents the 1st index hard disk drive in the guest.
	HardDiskDriveB = "hdb"

	// HardDiskDriveC represents the 2nd index hard disk drive in the guest.
	HardDiskDriveC = "hdc"

	// HardDiskDriveD represents the 3rd index hard disk drive in the guest.
	HardDiskDriveD = "hdd"

	// CDROM represents the CD-ROM image (cannot use this an HardDiskDriveC at the
	// same time). You can use the host CD-ROM by using `/dev/cdrom` as filename.
	CDROM = "cdrom"
)

// DiskDrive uses the specified disk drive name to mount the specified file path.
func DiskDrive(name string, file string) *queso.Option {
	return queso.NewOption(name, file)
}

// FlashMemory uses the specified file as on-board Flash memory image.
func FlashMemory(file string) *queso.Option {
	return queso.NewOption("mtdblock", file)
}

// SecureDigitalCard uses the specified file as a SecureDigital card image.
func SecureDigitalCard(file string) *queso.Option {
	return queso.NewOption("sd", file)
}

// ParallelFlash uses the specified file as a parallel flash image.
func ParallelFlash(file string) *queso.Option {
	return queso.NewOption("pflash", file)
}

// UseSnapshotMode writes to temporary files instead of disk image files. In this
// case, the raw disk image you use is not written back.
func UseSnapshotMode() *queso.Option {
	return queso.NewOption("snapshot", "")
}

// ISCSI configures iSCSI session parameters.
func ISCSI() *queso.Option {
	return queso.NewOption("iscsi", "")
}

// AIOBackend represents the asynchronous I/O backend values that can be
// passed to the WithAIOBackend property.
type AIOBackend string

const (
	// AIOBackendThreads uses threads for asynchronous I/O.
	AIOBackendThreads AIOBackend = "threads"

	// AIOBackendIOURing uses io_uring for asynchronous I/O.
	AIOBackendIOURing AIOBackend = "io_uring"

	// AIOBackendNative uses Linux native asynchronous I/O.
	AIOBackendNative AIOBackend = "native"
)

// IOOperation specifies an operation for I/O (read, write, or all).
type IOOperation int

const (
	IOOperationAll IOOperation = iota
	IOOperationRead
	IOOperationWrite
)
