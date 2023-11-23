package fsdev

import (
	"github.com/mikerourke/queso/internal/cli"
	"github.com/mikerourke/queso/qemu/blockdev2"
)

type anyLocalDevice struct {
	Identifier    string
	Path          string
	SecurityModel SecurityModel
	properties    []*cli.Property
}

// EnableWriteOut means that host page cache will be used to read and write data but
// write notification will be sent to the guest only when the data has been reported
// as written by the storage subsystem.
//
//	qemu-system-* -fsdev local,writeout=writeout
func (d *anyLocalDevice) EnableWriteOut() *anyLocalDevice {
	d.properties = append(d.properties,
		// The only supported value is "immediate".
		cli.NewProperty("writeout", "immediate"))
	return d
}

func (d *anyLocalDevice) SetDirectoryMode(mode string) *anyLocalDevice {
	d.properties = append(d.properties, cli.NewProperty("dmode", mode))
	return d
}

func (d *anyLocalDevice) SetFileMode(mode string) *anyLocalDevice {
	d.properties = append(d.properties, cli.NewProperty("fmode", mode))
	return d
}

func (d *anyLocalDevice) SetPath(path string) *anyLocalDevice {
	d.Path = path
	return d
}

func (d *anyLocalDevice) SetSecurityModel(model SecurityModel) *anyLocalDevice {
	d.SecurityModel = model
	return d
}

func (d *anyLocalDevice) ToggleReadOnly(enabled bool) *anyLocalDevice {
	d.properties = append(d.properties, cli.NewProperty("readonly", enabled))
	return d
}

type LocalFileSystemDevice struct {
	*anyLocalDevice
}

func NewLocalFileSystemDevice(id string, path string, model SecurityModel) *LocalFileSystemDevice {
	return &LocalFileSystemDevice{
		&anyLocalDevice{
			Identifier:    id,
			Path:          path,
			SecurityModel: model,
			properties:    make([]*cli.Property, 0),
		},
	}
}

func (d *LocalFileSystemDevice) option() *cli.Option {
	properties := append(d.properties,
		cli.NewProperty("id", d.Identifier),
		cli.NewProperty("path", d.Path),
		cli.NewProperty("security_model", d.SecurityModel))
	return cli.NewOption("fsdev", "local", properties...)
}

func (d *LocalFileSystemDevice) SetID(id string) *LocalFileSystemDevice {
	d.Identifier = id
	return d
}

// SetBandwidthBursts specifies bursts in bytes per second for the specified
// operation.
//
//	qemu-system-* -fsdev local,throttling.bps_rd_max=bps
//	qemu-system-* -fsdev local,throttling.bps_wr_max=bps
//	qemu-system-* -fsdev local,throttling.bps_max=bps
func (d *LocalFileSystemDevice) SetBandwidthBursts(operation blockdev2.IOOperation, bps int) *LocalFileSystemDevice {
	var property *cli.Property

	switch operation {
	case blockdev2.IORead:
		property = cli.NewProperty("throttling.bps_rd_max", bps)

	case blockdev2.IOWrite:
		property = cli.NewProperty("throttling.bps_wr_max", bps)

	case blockdev2.IOAll:
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
func (d *LocalFileSystemDevice) SetBandwidthThrottling(operation blockdev2.IOOperation, bps int) *LocalFileSystemDevice {
	var property *cli.Property

	switch operation {
	case blockdev2.IORead:
		property = cli.NewProperty("throttling.bps_rd", bps)

	case blockdev2.IOWrite:
		property = cli.NewProperty("throttling.bps_wr", bps)

	case blockdev2.IOAll:
		property = cli.NewProperty("throttling.bps", bps)
	}

	d.properties = append(d.properties, property)
	return d
}

// SetRequestRateBursts specifies bursts in request per second for the specified
// drive operation. Bursts allow the guest I/O to spike above the limit temporarily.
//
//	qemu-system-* -fsdev local,throttling.iops_rd_max=rps
//	qemu-system-* -fsdev local,throttling.iops_wr_max=rps
//	qemu-system-* -fsdev local,throttling.iops_max=rps
func (d *LocalFileSystemDevice) SetRequestRateBursts(operation blockdev2.IOOperation, rps int) *LocalFileSystemDevice {
	var property *cli.Property

	switch operation {
	case blockdev2.IORead:
		property = cli.NewProperty("throttling.iops_rd_max", rps)

	case blockdev2.IOWrite:
		property = cli.NewProperty("throttling.iops_wr_max", rps)

	case blockdev2.IOAll:
		property = cli.NewProperty("throttling.iops_max", rps)
	}

	d.properties = append(d.properties, property)
	return d
}

// SetRequestRateLimits specifies request rate limits in requests per second for
// the specified drive operation.
//
//	qemu-system-* -fsdev local,throttling.iops_rd=rps
//	qemu-system-* -fsdev local,throttling.iops_wr=rps
//	qemu-system-* -fsdev local,throttling.iops=rps
func (d *LocalFileSystemDevice) SetRequestRateLimits(operation blockdev2.IOOperation, rps int) *LocalFileSystemDevice {
	var property *cli.Property

	switch operation {
	case blockdev2.IORead:
		property = cli.NewProperty("throttling.iops_rd", rps)

	case blockdev2.IOWrite:
		property = cli.NewProperty("throttling.iops_wr", rps)

	case blockdev2.IOAll:
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

type VirtualLocalFileSystemDevice struct {
	*anyLocalDevice
}

func NewVirtualLocalFileSystemDevice(
	mountTag string,
	path string,
	model SecurityModel,
) *VirtualLocalFileSystemDevice {
	return &VirtualLocalFileSystemDevice{
		&anyLocalDevice{
			Identifier:    mountTag,
			Path:          path,
			SecurityModel: model,
			properties:    make([]*cli.Property, 0),
		},
	}
}

func (d *VirtualLocalFileSystemDevice) option() *cli.Option {
	properties := append(d.properties,
		cli.NewProperty("mount_tag", d.Identifier),
		cli.NewProperty("path", d.Path),
		cli.NewProperty("security_model", d.SecurityModel))
	return cli.NewOption("virtfs", "local", properties...)
}

func (d *VirtualLocalFileSystemDevice) SetMountTag(tag string) *VirtualLocalFileSystemDevice {
	d.properties = append(d.properties, cli.NewProperty("mount_tag", tag))
	return d
}

type MultiDeviceSharing string

const (
	MultiDeviceSharingRemap  MultiDeviceSharing = "remap"
	MultiDeviceSharingForbid MultiDeviceSharing = "forbid"
	MultiDeviceSharingWarn   MultiDeviceSharing = "warn"
)

func (d *VirtualLocalFileSystemDevice) SetMultiDeviceSharing(sharing MultiDeviceSharing) *VirtualLocalFileSystemDevice {
	d.properties = append(d.properties, cli.NewProperty("multidevs", sharing))
	return d
}
