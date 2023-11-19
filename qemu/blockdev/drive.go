package blockdev

import (
	"github.com/mikerourke/queso"
	"github.com/mikerourke/queso/diskimage"
)

// Drive defines a new drive. This includes creating a block driver node (the
// backend) as well as a guest device, and is mostly a shortcut for defining the
// corresponding Driver and device options.
func Drive(properties ...*DriveProperty) *queso.Option {
	props := make([]*queso.Property, 0)

	for _, property := range properties {
		props = append(props, property.Property)
	}

	return queso.NewOption("drive", "", props...)
}

// DriveProperty represents a property to use with the Drive option.
type DriveProperty struct {
	*queso.Property
}

// NewDriveProperty returns a new instance of DriveProperty.
func NewDriveProperty(key string, value interface{}) *DriveProperty {
	return &DriveProperty{
		Property: queso.NewProperty(key, value),
	}
}

// WithDiskImageFile defines which disk image to use with the Drive. See
// https://qemu.readthedocs.io/en/latest/system/images.html for more details.
//
// If the file contains commas, you must double it. For example. to use file "my,file":
// 	WithDiskImageFile("my,,file")
//
// Special files such as iSCSI devices can be specified using protocol specific URLs.
// See https://www.qemu.org/docs/master/system/invocation.html#device-url-syntax for
// more details.
func WithDiskImageFile(file string) *DriveProperty {
	return NewDriveProperty("file", file)
}

// WithDiskImageFormat defines the format of the associated image file.
func WithDiskImageFormat(format diskimage.FileFormat) *DriveProperty {
	return NewDriveProperty("file", format)
}

// DriveInterface represents an interface that can be used with a Drive and is
// passed to the WithDriveInterface property.
type DriveInterface string

const (
	DriveInterfaceNone          DriveInterface = "none"
	DriveInterfaceFlashMemory   DriveInterface = "mtd"
	DriveInterfaceFloppy        DriveInterface = "floppy"
	DriveInterfaceIDE           DriveInterface = "ide"
	DriveInterfaceParallelFlash DriveInterface = "pflash"
	DriveInterfaceSCSI          DriveInterface = "scsi"
	DriveInterfaceSDCard        DriveInterface = "sd"
	DriveInterfaceVirtio        DriveInterface = "virtio"
)

// WithDriveInterface defines on which type on interface the Drive is connected.
func WithDriveInterface(driveInterface DriveInterface) *DriveProperty {
	return NewDriveProperty("if", driveInterface)
}

// WithBus defines the bus number for the connected Drive.
func WithBus(bus string) *DriveProperty {
	return NewDriveProperty("bus", bus)
}

// WithUnit defines the unit ID for the connected Drive.
func WithUnit(unit string) *DriveProperty {
	return NewDriveProperty("unit", unit)
}

// WithIndex defines where the Drive is connected by using an index in the list
// of available connectors of a given interface type.
func WithIndex(index int) *DriveProperty {
	return NewDriveProperty("index", index)
}

// DriveMedia represents the type of Drive media and is passed to the
// WithDriveMedia property.
type DriveMedia string

const (
	// DriveMediaCDROM indicates that the Drive is a CD-ROM drive.
	DriveMediaCDROM DriveMedia = "cdrom"

	// DriveMediaDisk indicates that the Drive is a disk drive.
	DriveMediaDisk DriveMedia = "disk"
)

// WithDriveMedia defines the type of the media for the Drive.
func WithDriveMedia(media DriveMedia) *DriveProperty {
	return NewDriveProperty("media", media)
}

// IsSnapshotMode specifies whether the Drive should use snapshot mode. If
// enabled, no changes are persisted to the drive. This is useful for
// debugging purposes.
func IsSnapshotMode(enabled bool) *DriveProperty {
	return NewDriveProperty("snapshot", enabled)
}

// CacheAccessMode controls how the host cache is used to access block data on
// the Drive and is passed to the WithCacheAccessMode property.
type CacheAccessMode string

const (
	CacheAccessNone         CacheAccessMode = "none"
	CacheAccessDirectSync   CacheAccessMode = "directsync"
	CacheAccessUnsafe       CacheAccessMode = "unsafe"
	CacheAccessWriteBack    CacheAccessMode = "writeback"
	CacheAccessWriteThrough CacheAccessMode = "writethrough"
)

// WithCacheAccessMode controls how the host cache is used to access block data
// on the Drive. The default mode is CacheAccessWriteBack.
func WithCacheAccessMode(mode CacheAccessMode) *DriveProperty {
	return NewDriveProperty("cache", mode)
}

// WithDriverAIOBackend specifies the AIO backend for the Drive. The default
// is AIOBackendThreads.
func WithDriverAIOBackend(backend AIOBackend) *DriverProperty {
	return NewDriverProperty("aio", backend)
}

// IOErrorAction specifies which action to take on Drive write and read errors.
// The value is passed to the WithWriteErrorAction and WithReadErrorAction
// properties.
type IOErrorAction string

const (
	// IOErrorActionIgnore ignores the error and tries to continue.
	IOErrorActionIgnore IOErrorAction = "ignore"

	// IOErrorActionStop pauses QEMU.
	IOErrorActionStop IOErrorAction = "stop"

	// IOErrorActionReport reports the error to the guest.
	IOErrorActionReport IOErrorAction = "report"

	// IOErrorActionENOSPC pauses QEMU only if the host disk is full; reports the
	// error to the guest otherwise.
	IOErrorActionENOSPC IOErrorAction = "enospc"
)

// WithReadErrorAction specifies which action to take on Drive read errors. The
// default is IOErrorActionReport.
func WithReadErrorAction(action IOErrorAction) *DriveProperty {
	return NewDriveProperty("rerror", action)
}

// WithWriteErrorAction specifies which action to take on Drive write errors.
// The default is IOErrorActionENOSPC.
func WithWriteErrorAction(action IOErrorAction) *DriveProperty {
	return NewDriveProperty("werror", action)
}

// IsCopyOnRead specifies whether to copy read backing file sectors into the
// image file for a Drive.
func IsCopyOnRead(enabled bool) *DriveProperty {
	return NewDriveProperty("copy-on-read", enabled)
}

// WithBandwidthThrottling specifies bandwidth throttling limits in bytes per
// second for the specified Drive operation. Small values can lead to timeouts or
// hangs inside the guest. A safe minimum for disks is 2 MB/s.
func WithBandwidthThrottling(operation IOOperation, bytesPerSecond int) *DriveProperty {
	switch operation {
	case IOOperationRead:
		return NewDriveProperty("bps_rd", bytesPerSecond)

	case IOOperationWrite:
		return NewDriveProperty("bps_wr", bytesPerSecond)

	default:
		return NewDriveProperty("bps", bytesPerSecond)
	}
}

// WithBandwidthBursts specifies bursts in bytes per second for the specified
// Drive operation. Bursts allow the guest I/O to spike above the limit
// temporarily.
func WithBandwidthBursts(operation IOOperation, bytesPerSecond int) *DriveProperty {
	switch operation {
	case IOOperationRead:
		return NewDriveProperty("bps_rd_max", bytesPerSecond)

	case IOOperationWrite:
		return NewDriveProperty("bps_wr_max", bytesPerSecond)

	default:
		return NewDriveProperty("bps_max", bytesPerSecond)
	}
}

// WithRequestRateLimits specifies request rate limits in requests per second for
// the specified Drive operation.
func WithRequestRateLimits(operation IOOperation, requestsPerSecond int) *DriveProperty {
	switch operation {
	case IOOperationRead:
		return NewDriveProperty("iops_rd", requestsPerSecond)

	case IOOperationWrite:
		return NewDriveProperty("iops_wr", requestsPerSecond)

	default:
		return NewDriveProperty("iops", requestsPerSecond)
	}
}

// WithRequestRateBursts specifies bursts in request per second for the specified
// Drive operation. Bursts allow the guest I/O to spike above the limit temporarily.
func WithRequestRateBursts(operation IOOperation, requestsPerSecond int) *DriveProperty {
	switch operation {
	case IOOperationRead:
		return NewDriveProperty("iops_rd_max", requestsPerSecond)

	case IOOperationWrite:
		return NewDriveProperty("iops_wr_max", requestsPerSecond)

	default:
		return NewDriveProperty("iops_max", requestsPerSecond)
	}
}

// WithRequestSize sets the bytes of a request count as a new request for iops
// throttling purposes on the Drive. Use this option to prevent guests from
// circumventing iops limits by sending fewer but larger requests.
func WithRequestSize(bytes int) *DriveProperty {
	return NewDriveProperty("iops_size", bytes)
}

// WithGroupName specifies a throttling quota group with the specified name for
// a Drive. All drives that are members of the same group are accounted for together.
// Use this option to prevent guests from circumventing throttling limits by using
// many small disks instead of a single larger disk.
func WithGroupName(name string) *DriveProperty {
	return NewDriveProperty("group", name)
}
