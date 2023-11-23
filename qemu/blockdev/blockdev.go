package blockdev

import "github.com/mikerourke/queso/qemu/cli"

// IOOperation specifies an operation for I/O (read, write, or all).
type IOOperation int

const (
	// IOAll represents all I/O operations.
	IOAll IOOperation = iota

	// IORead represents read I/O operation.
	IORead

	// IOWrite represents write I/O operation.
	IOWrite
)

// EnableSnapshotMode writes to temporary files instead of disk image files. In this
// case, the raw disk image you use is not written back.
//
// Warning: EnableSnapshotMode is incompatible with -blockdev (instead use diskimage.Create
// to manually create snapshot images to attach to your blockdev). If you have
// mixed -blockdev and Drive declarations you can use [Drive.ToggleSnapshotMode] on
// your drive declarations instead of this global option.
//
//	qemu-system-* -snapshot
func EnableSnapshotMode() *cli.Option {
	return cli.NewOption("snapshot", "")
}

// FlashMemory uses the specified file as on-board Flash memory image.
//
//	qemu-system-* -mtdblock file
func FlashMemory(file string) *cli.Option {
	return cli.NewOption("mtdblock", file)
}

// ISCSI configures iSCSI session parameters.
//
//	qemu-system-* -iscsi
func ISCSI() *cli.Option {
	return cli.NewOption("iscsi", "")
}

// SecureDigitalCard uses the specified file as a SecureDigital card image.
//
//	qemu-system-* -sd file
func SecureDigitalCard(file string) *cli.Option {
	return cli.NewOption("sd", file)
}
