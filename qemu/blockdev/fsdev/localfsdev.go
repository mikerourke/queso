package fsdev

import (
	"github.com/mikerourke/queso/qemu/blockdev"
	"github.com/mikerourke/queso/qemu/cli"
)

// LocalFileSystemDevice represents a file system device in which accesses to the
// filesystem are done by QEMU.
type LocalFileSystemDevice struct {
	// ID is the unique identifier for the device.
	ID string

	// Path is the export path for the file system device. Files under this path
	// will be available to the 9P client on the guest.
	Path string

	// SecurityModel is used to specify the security model to be used for the export
	// path in file system devices.
	SecurityModel SecurityModel
	properties    []*cli.Property
}

// NewLocalFileSystemDevice returns a new instance of [LocalFileSystemDevice].
// id is a unique identifier for the device. path is the export path for
// the file system device. Files under this path will be available to the 9P
// client on the guest. model is the [SecurityModel] for the file system
// device.
//
//	qemu-system-* -fsdev local,mount_tag=tag,path=path,security_model=model
func NewLocalFileSystemDevice(id string, path string, model SecurityModel) *LocalFileSystemDevice {
	return &LocalFileSystemDevice{
		ID:            id,
		Path:          path,
		SecurityModel: model,
		properties:    make([]*cli.Property, 0),
	}
}

func (d *LocalFileSystemDevice) option() *cli.Option {
	properties := append(d.properties,
		cli.NewProperty("id", d.ID),
		cli.NewProperty("path", d.Path),
		cli.NewProperty("security_model", d.SecurityModel))
	return cli.NewOption("fsdev", "local", properties...)
}

// EnableWriteOut means that host page cache will be used to read and write data but
// write notification will be sent to the guest only when the data has been reported
// as written by the storage subsystem.
//
//	qemu-system-* -fsdev local,writeout=writeout
func (d *LocalFileSystemDevice) EnableWriteOut() *LocalFileSystemDevice {
	d.properties = append(d.properties,
		// The only supported value is "immediate".
		cli.NewProperty("writeout", "immediate"))
	return d
}

// SetBandwidthBursts specifies bursts in bytes per second for the specified
// operation.
//
//	qemu-system-* -fsdev local,throttling.bps_rd_max=bps
//	qemu-system-* -fsdev local,throttling.bps_wr_max=bps
//	qemu-system-* -fsdev local,throttling.bps_max=bps
func (d *LocalFileSystemDevice) SetBandwidthBursts(
	operation blockdev.IOOperation,
	bps int,
) *LocalFileSystemDevice {
	var property *cli.Property

	switch operation {
	case blockdev.IORead:
		property = cli.NewProperty("throttling.bps_rd_max", bps)

	case blockdev.IOWrite:
		property = cli.NewProperty("throttling.bps_wr_max", bps)

	case blockdev.IOAll:
		property = cli.NewProperty("throttling.bps_max", bps)
	}

	d.properties = append(d.properties, property)
	return d
}

// SetBandwidthThrottling specifies bandwidth throttling limits in bytes per
// second for the specified operation.
//
//	qemu-system-* -fsdev local,throttling.bps_rd=bps
//	qemu-system-* -fsdev local,throttling.bps_wr=bps
//	qemu-system-* -fsdev local,throttling.bps=bps
func (d *LocalFileSystemDevice) SetBandwidthThrottling(
	operation blockdev.IOOperation,
	bps int,
) *LocalFileSystemDevice {
	var property *cli.Property

	switch operation {
	case blockdev.IORead:
		property = cli.NewProperty("throttling.bps_rd", bps)

	case blockdev.IOWrite:
		property = cli.NewProperty("throttling.bps_wr", bps)

	case blockdev.IOAll:
		property = cli.NewProperty("throttling.bps", bps)
	}

	d.properties = append(d.properties, property)
	return d
}

// SetDirectoryMode specifies the default mode for newly created directories on the
// host. Works only with [SecurityModel] set to [SecurityModelMappedXAttr] and
// [SecurityModelMappedFile].
//
//	qemu-system-* -fsdev local,dmode=mode
func (d *LocalFileSystemDevice) SetDirectoryMode(mode string) *LocalFileSystemDevice {
	d.properties = append(d.properties, cli.NewProperty("dmode", mode))
	return d
}

// SetFileMode specifies the default mode for newly created files on the
// host. Works only with [SecurityModel] set to [SecurityModelMappedXAttr] and
// [SecurityModelMappedFile].
//
//	qemu-system-* -fsdev local,fmode=mode
func (d *LocalFileSystemDevice) SetFileMode(mode string) *LocalFileSystemDevice {
	d.properties = append(d.properties, cli.NewProperty("fmode", mode))
	return d
}

// SetID specifies the identifier for this device.
//
//	qemu-system-* -fsdev local,id=id
func (d *LocalFileSystemDevice) SetID(id string) *LocalFileSystemDevice {
	d.ID = id
	return d
}

// SetPath specifies the export path for the file system device. Files under
// this path will be available to the 9P client on the guest.
//
//	qemu-system-* -fsdev local,path=path
func (d *LocalFileSystemDevice) SetPath(path string) *LocalFileSystemDevice {
	d.Path = path
	return d
}

// SetRequestRateBursts specifies bursts in request per second for the specified
// operation. Bursts allow the guest I/O to spike above the limit temporarily.
//
//	qemu-system-* -fsdev local,throttling.iops_rd_max=rps
//	qemu-system-* -fsdev local,throttling.iops_wr_max=rps
//	qemu-system-* -fsdev local,throttling.iops_max=rps
func (d *LocalFileSystemDevice) SetRequestRateBursts(
	operation blockdev.IOOperation,
	rps int,
) *LocalFileSystemDevice {
	var property *cli.Property

	switch operation {
	case blockdev.IORead:
		property = cli.NewProperty("throttling.iops_rd_max", rps)

	case blockdev.IOWrite:
		property = cli.NewProperty("throttling.iops_wr_max", rps)

	case blockdev.IOAll:
		property = cli.NewProperty("throttling.iops_max", rps)
	}

	d.properties = append(d.properties, property)
	return d
}

// SetRequestRateLimits specifies request rate limits in requests per second for
// the specified operation.
//
//	qemu-system-* -fsdev local,throttling.iops_rd=rps
//	qemu-system-* -fsdev local,throttling.iops_wr=rps
//	qemu-system-* -fsdev local,throttling.iops=rps
func (d *LocalFileSystemDevice) SetRequestRateLimits(
	operation blockdev.IOOperation,
	rps int,
) *LocalFileSystemDevice {
	var property *cli.Property

	switch operation {
	case blockdev.IORead:
		property = cli.NewProperty("throttling.iops_rd", rps)

	case blockdev.IOWrite:
		property = cli.NewProperty("throttling.iops_wr", rps)

	case blockdev.IOAll:
		property = cli.NewProperty("throttling.iops", rps)
	}

	d.properties = append(d.properties, property)
	return d
}

// SetRequestSize sets the bytes of a request count as a new request for iops
// throttling purposes.
//
//	qemu-system-* -fsdev local,throttling.iops_size=bytes
func (d *LocalFileSystemDevice) SetRequestSize(bytes int) *LocalFileSystemDevice {
	d.properties = append(d.properties, cli.NewProperty("throttling.iops_size", bytes))
	return d
}

// SetSecurityModel specifies the security model to be used for this export path.
// See [SecurityModel] for additional details.
//
//	qemu-system-* -fsdev local,security_model=model
func (d *LocalFileSystemDevice) SetSecurityModel(model SecurityModel) *LocalFileSystemDevice {
	d.SecurityModel = model
	return d
}

// ToggleReadOnly enables exporting 9P share as a readonly mount for guests.
// By default, read-write access is given.
//
//	qemu-system-* -fsdev local,readonly=on|off
func (d *LocalFileSystemDevice) ToggleReadOnly(enabled bool) *LocalFileSystemDevice {
	d.properties = append(d.properties, cli.NewProperty("readonly", enabled))
	return d
}
