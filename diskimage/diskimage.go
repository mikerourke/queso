// Package diskimage is used to manage QEMU disk images. See
// https://qemu.readthedocs.io/en/latest/system/images.html for more details.
package diskimage

// FileFormat represents disk image file formats that can be mounted/edited with
// QEMU.
type FileFormat string

const (
	// FileFormatRaw represents the raw disk image format. This format has the
	// advantage of being simple and easily exportable to all other emulators.
	// If your file system supports holes (for example in ext2 or ext3 on Linux
	// or NTFS on Windows), then only the written sectors will reserve space.
	FileFormatRaw FileFormat = "raw"

	// FileFormatQCOW2 represents the QEMU image format, the most versatile format.
	// Use it to have smaller images (useful if your filesystem does not support
	// holes, for example on Windows), zlib based compression and support of
	// multiple VM snapshots.
	FileFormatQCOW2 FileFormat = "qcow2"

	// FileFormatQED represents the old QEMU image format with support for
	// backing files and compact image files (when your filesystem or transport
	// medium does not support holes).
	FileFormatQED FileFormat = "qed"

	// FileFormatQCOW represents the old QEMU image format with support for backing
	// files, compact image files, encryption and compression.
	FileFormatQCOW FileFormat = "qcow"

	// FileFormatLUKS represents the LUKS v1 encryption format, compatible with
	// Linux dm-crypt/cryptsetup.
	FileFormatLUKS FileFormat = "luks"

	// FileFormatVDI represents the VirtualBox 1.1 compatible image format.
	FileFormatVDI FileFormat = "vdi"

	// FileFormatVMDK represents the VMware 3 and 4 compatible image format.
	FileFormatVMDK FileFormat = "vmdk"

	// FileFormatVPC represents the VirtualPC compatible image format (VHD).
	FileFormatVPC FileFormat = "vmdk"

	// FileFormatVHDX represents the Hyper-V compatible image format (VHDX).
	FileFormatVHDX FileFormat = "VHDX"
)

// ReadOnlyFormat represents disk image file formats that are supported in a
// read-only mode.
type ReadOnlyFormat string

const (
	// ReadOnlyFormatBochs represents the Bochs images of growing type.
	ReadOnlyFormatBochs ReadOnlyFormat = "bochs"

	// ReadOnlyFormatCloop represents the Linux Compressed Loop image, useful
	// only to reuse directly compressed CD-ROM images present for example in the
	// Knoppix CD-ROMs.
	ReadOnlyFormatCloop ReadOnlyFormat = "cloop"

	// ReadOnlyFormatDMG represents the Apple disk image.
	ReadOnlyFormatDMG ReadOnlyFormat = "dmg"

	// ReadOnlyFormatParallels represents the Parallels disk image format.
	ReadOnlyFormatParallels ReadOnlyFormat = "parallels"
)
