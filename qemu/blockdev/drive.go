package blockdev

import (
	"github.com/mikerourke/queso/diskimage"
	"github.com/mikerourke/queso/qemu/cli"
)

// Drive represents a new drive. This includes creating a block driver node (the backend)
// as well as a guest device, and is mostly a shortcut for defining the corresponding
// [Driver] and Device options.
type Drive struct {
	*Driver
}

// NewDrive returns a new instance of a [Drive].
//
//	qemu-system-* -drive
func NewDrive() *Drive {
	return &Drive{
		&Driver{
			properties: make([]*cli.Property, 0),
		},
	}
}

// SetAIOBackend specified the AIO backend. The default value is [AIOBackendThreads].
//
//	qemu-system-* -drive aio=backend
func (d *Drive) SetAIOBackend(backend string) *Drive {
	d.properties = append(d.properties, cli.NewProperty("aio", backend))
	return d
}

// SetBandwidthBursts specifies bursts in bytes per second for the specified
// drive operation. Bursts allow the guest I/O to spike above the limit
// temporarily.
//
//	qemu-system-* -drive bps_rd_max=bps
//	qemu-system-* -drive bps_wr_max=bps
//	qemu-system-* -drive bps_max=bps
func (d *Drive) SetBandwidthBursts(operation IOOperation, bps int) *Drive {
	var property *cli.Property

	switch operation {
	case IORead:
		property = cli.NewProperty("bps_rd_max", bps)

	case IOWrite:
		property = cli.NewProperty("bps_wr_max", bps)

	case IOAll:
		property = cli.NewProperty("bps_max", bps)
	}

	d.properties = append(d.properties, property)
	return d
}

// SetBandwidthThrottling specifies bandwidth throttling limits in bytes per
// second for the specified drive operation. Small values can lead to timeouts or
// hangs inside the guest. A safe minimum for disks is 2 MB/s.
//
//	qemu-system-* -drive bps_rd=bps
//	qemu-system-* -drive bps_wr=bps
//	qemu-system-* -drive bps=bps
func (d *Drive) SetBandwidthThrottling(operation IOOperation, bps int) *Drive {
	var property *cli.Property

	switch operation {
	case IORead:
		property = cli.NewProperty("bps_rd", bps)

	case IOWrite:
		property = cli.NewProperty("bps_wr", bps)

	case IOAll:
		property = cli.NewProperty("bps", bps)
	}

	d.properties = append(d.properties, property)
	return d
}

// SetBus defines the bus number for the connected drive.
//
//	qemu-system-* -drive bus=bus
func (d *Drive) SetBus(bus string) *Drive {
	d.properties = append(d.properties, cli.NewProperty("bus", bus))
	return d
}

// CacheAccess controls how the host cache is used to access block data on
// the drive and is passed to the [Drive.SetCacheAccess] property.
type CacheAccess string

const (
	CacheAccessNone         CacheAccess = "none"
	CacheAccessDirectSync   CacheAccess = "directsync"
	CacheAccessUnsafe       CacheAccess = "unsafe"
	CacheAccessWriteBack    CacheAccess = "writeback"
	CacheAccessWriteThrough CacheAccess = "writethrough"
)

// SetCacheAccess controls how the host cache is used to access block data
// on the drive. The default mode is [CacheAccessWriteBack].
//
//	qemu-system-* -drive cache=cache
func (d *Drive) SetCacheAccess(access CacheAccess) *Drive {
	d.properties = append(d.properties, cli.NewProperty("cache", access))
	return d
}

// SetFile defines which disk image to use with the drive. See
// https://qemu.readthedocs.io/en/latest/system/images.html for more details.
//
// If the file contains commas, you must double it. For example. to use file "my,file":
//
//	SetFile(`"my,,file"`)
//
// Special files such as iSCSI devices can be specified using protocol specific URLs.
// See https://www.qemu.org/docs/master/system/invocation.html#device-url-syntax for
// more details.
//
//	qemu-system-* -drive file=file
func (d *Drive) SetFile(file string) *Drive {
	d.properties = append(d.properties, cli.NewProperty("file", file))
	return d
}

// SetFormat specifies which disk format will be used rather than detecting the format.
// Can be used to specify [diskimage.FileFormatRaw] to avoid interpreting an untrusted
// format header.
//
//	qemu-system-* -drive format=format
func (d *Drive) SetFormat(format diskimage.FileFormat) *Drive {
	d.properties = append(d.properties, cli.NewProperty("format", format))
	return d
}

// SetGroupName specifies a throttling quota group with the specified name for
// a drive. All drives that are members of the same group are accounted for together.
// Use this option to prevent guests from circumventing throttling limits by using
// many small disks instead of a single larger disk.
//
//	qemu-system-* -drive group=name
func (d *Drive) SetGroupName(name string) *Drive {
	d.properties = append(d.properties, cli.NewProperty("group", name))
	return d
}

// SetIndex defines where the drive is connected by using an index in the list
// of available connectors of a given interface type.
//
//	qemu-system-* -drive index=index
func (d *Drive) SetIndex(index int) *Drive {
	d.properties = append(d.properties, cli.NewProperty("index", index))
	return d
}

// DriveInterface represents an interface that can be used with a Drive and is
// passed to the [Drive.SetInterface] method.
type DriveInterface string

const (
	InterfaceNone          DriveInterface = "none"
	InterfaceFlashMemory   DriveInterface = "mtd"
	InterfaceFloppy        DriveInterface = "floppy"
	InterfaceIDE           DriveInterface = "ide"
	InterfaceParallelFlash DriveInterface = "pflash"
	InterfaceSCSI          DriveInterface = "scsi"
	InterfaceSDCard        DriveInterface = "sd"
	InterfaceVirtio        DriveInterface = "virtio"
)

// SetInterface defines on which type on interface the drive is connected.
//
//	qemu-system-* -drive if=interface
func (d *Drive) SetInterface(driveInterface DriveInterface) *Drive {
	d.properties = append(d.properties, cli.NewProperty("if", driveInterface))
	return d
}

// DriveMedia represents the type of Drive media and is passed to the
// [Drive.SetMedia] method.
type DriveMedia string

const (
	// DriveMediaCDROM indicates that the drive is a CD-ROM drive.
	DriveMediaCDROM DriveMedia = "cdrom"

	// DriveMediaDisk indicates that the drive is a disk drive.
	DriveMediaDisk DriveMedia = "disk"
)

// SetMedia defines the type of the media for the drive.
//
//	qemu-system-* -drive media=media
func (d *Drive) SetMedia(media DriveMedia) *Drive {
	d.properties = append(d.properties, cli.NewProperty("media", media))
	return d
}

// SetRequestRateBursts specifies bursts in request per second for the specified
// drive operation. Bursts allow the guest I/O to spike above the limit temporarily.
//
//	qemu-system-* -drive iops_rd_max=rps
//	qemu-system-* -drive iops_wr_max=rps
//	qemu-system-* -drive iops_max=rps
func (d *Drive) SetRequestRateBursts(operation IOOperation, rps int) *Drive {
	var property *cli.Property

	switch operation {
	case IORead:
		property = cli.NewProperty("iops_rd_max", rps)

	case IOWrite:
		property = cli.NewProperty("iops_wr_max", rps)

	case IOAll:
		property = cli.NewProperty("iops_max", rps)
	}

	d.properties = append(d.properties, property)
	return d
}

// SetRequestRateLimits specifies request rate limits in requests per second for
// the specified drive operation.
//
//	qemu-system-* -drive iops_rd=rps
//	qemu-system-* -drive iops_wr=rps
//	qemu-system-* -drive iops=rps
func (d *Drive) SetRequestRateLimits(operation IOOperation, rps int) *Drive {
	var property *cli.Property

	switch operation {
	case IORead:
		property = cli.NewProperty("iops_rd", rps)

	case IOWrite:
		property = cli.NewProperty("iops_wr", rps)

	case IOAll:
		property = cli.NewProperty("iops", rps)
	}

	d.properties = append(d.properties, property)
	return d
}

// SetRequestSize sets the bytes of a request count as a new request for iops
// throttling purposes on the drive. Use this option to prevent guests from
// circumventing iops limits by sending fewer but larger requests.
//
//	qemu-system-* -drive iops_size=bytes
func (d *Drive) SetRequestSize(bytes int) *Drive {
	d.properties = append(d.properties, cli.NewProperty("iops_size", bytes))
	return d
}

// SetUnit defines the unit ID for the connected drive.
//
//	qemu-system-* -drive unit=unit
func (d *Drive) SetUnit(unit string) *Drive {
	d.properties = append(d.properties, cli.NewProperty("unit", unit))
	return d
}

// DriveErrorAction specifies which action to take on drive write and read errors.
// The value is passed to the [Drive.SetWriteErrorAction] and SetReadErrorAction
// methods.
type DriveErrorAction string

const (
	// DriveErrorActionIgnore ignores the error and tries to continue.
	DriveErrorActionIgnore DriveErrorAction = "ignore"

	// DriveErrorActionStop pauses QEMU.
	DriveErrorActionStop DriveErrorAction = "stop"

	// DriveErrorActionReport reports the error to the guest.
	DriveErrorActionReport DriveErrorAction = "report"

	// DriveErrorActionENOSPC pauses QEMU only if the host disk is full; reports the
	// error to the guest otherwise.
	DriveErrorActionENOSPC DriveErrorAction = "enospc"
)

// SetReadErrorAction specifies which action to take on drive read errors. The
// default is [DriveErrorActionReport].
//
//	qemu-system-* -drive rerror=action
func (d *Drive) SetReadErrorAction(action DriveErrorAction) *Drive {
	d.properties = append(d.properties, cli.NewProperty("rerror", action))
	return d
}

// SetWriteErrorAction specifies which action to take on drive write errors.
// The default is [DriveErrorActionENOSPC].
//
//	qemu-system-* -drive werror=action
func (d *Drive) SetWriteErrorAction(action DriveErrorAction) *Drive {
	d.properties = append(d.properties, cli.NewProperty("werror", action))
	return d
}

// ToggleCopyOnRead specifies whether to copy read backing file sectors into the
// image file.
//
//	qemu-system-* -drive copy-on-read=on|off
func (d *Drive) ToggleCopyOnRead(enabled bool) *Drive {
	d.properties = append(d.properties, cli.NewProperty("copy-on-read", enabled))
	return d
}

// ToggleSnapshotMode specifies whether the drive should use snapshot mode. If
// enabled, no changes are persisted to the drive. This is useful for
// debugging purposes.
//
//	qemu-system-* -drive snapshot=on|off
func (d *Drive) ToggleSnapshotMode(enabled bool) *Drive {
	d.properties = append(d.properties, cli.NewProperty("snapshot", enabled))
	return d
}
