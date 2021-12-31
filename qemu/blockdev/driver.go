package blockdev

import (
	"fmt"

	"github.com/mikerourke/queso"
)

// Driver defines a new block driver node. Some options apply to all block
// drivers, other options are only accepted for a specific block driver. The
// driver parameter specifies the block driver to use for the given node.
//
// Example
//
//	qemu.New("qemu-system-x86_64").SetOptions(
//		blockdev.FileDriver("disk.img",
//			blockdev.WithNodeName("disk")))
//
// Invocation
//
//	qemu-system-x86_64 -blockdev driver=file,node-name=disk,filename=disk.img
func Driver(driver string, properties ...*DriverProperty) *queso.Option {
	props := []*queso.Property{queso.NewProperty("driver", driver)}

	for _, property := range properties {
		props = append(props, property.Property)
	}

	return queso.NewOption("blockdev", "", props...)
}

// FileDriver defines the protocol-level block Driver for accessing regular files.
func FileDriver(file string, properties ...*DriverProperty) *queso.Option {
	props := []*DriverProperty{NewDriverProperty("filename", file)}

	if properties != nil {
		props = append(props, properties...)
	}

	return Driver("file", props...)
}

// RawDriver is the image format block Driver for raw images. It is usually
// stacked on top of a protocol level block driver such as FileDriver.
//
// Example 1
//
// 	qemu.New("qemu-system-x86_64").SetOptions(
//		blockdev.FileDriver("disk.img",
//			blockdev.WithNodeName("disk_file")),
//		blockdev.RawDriver(
//			blockdev.WithNodeName("disk"),
//			blockdev.WithFile("disk_file")))
//
// Invocation
//
//	qemu-system-x86_64 \
//		-blockdev driver=file,node-name=disk_file,filename=disk.img
//		-blockdev driver=raw,node-name=disk,file=disk_file
//
// Example 2
//
//	qemu.New("qemu-system-x86_64").SetOptions(
//		blockdev.RawDriver(
//			blockdev.WithNodeName("disk"),
//			blockdev.WithDriverProperty("file", blockdev.WithDriverType("file")),
//			blockdev.WithDriverProperty("file", blockdev.WithImageFile("disk.img"))))
//
// Invocation
//
//	qemu-system-x86_64 -blockdev driver=raw,node-name=disk,file.driver=file,file.filename=disk.img
func RawDriver(properties ...*DriverProperty) *queso.Option {
	return Driver("raw", properties...)
}

// QCOW2Driver is the image format block Driver for qcow2 images. It is usually
// stacked on top of a protocol level block driver such as FileDriver.
//
// Example 1
//
//	qemu.New("qemu-system-x86_64").SetOptions(
//		blockdev.FileDriver("/tmp/disk.qcow2",
//			blockdev.WithNodeName("my_file")),
//		blockdev.QCOW2Driver(
//			blockdev.WithNodeName(blockdev.HardDiskDriveA.String()),
//			blockdev.WithFile("my_file"),
//			blockdev.WithOverlapCheck(blockdev.OverlapCheckNone),
//			blockdev.WithTotalCacheSize(16777216)))
//
// Invocation
//
//	qemu-system-x86_64 \
//		-blockdev driver=file,node-name=my_file,filename=/tmp/disk.qcow2
//		-blockdev driver=qcow2,node-name=hda,file=my_file,overlap-check=none,cache-size=16777216
//
// Example 2
//
//	qemu.New("qemu-system-x86_64").SetOptions(
//		blockdev.QCOW2Driver(
//			blockdev.WithNodeName("disk"),
//			blockdev.WithDriverProperty("file", blockdev.WithDriverType("http")),
//			blockdev.WithDriverProperty("file", blockdev.WithImageFile("https://example.com/image.qcow2")))
//
// Invocation
//
//	qemu-system-x86_64 -blockdev driver=qcow2,node-name=disk,file.driver=http,file.filename=https://example.com/image.qcow2
func QCOW2Driver(properties ...*DriverProperty) *queso.Option {
	return Driver("qcow2", properties...)
}

// DriverProperty represents a property that can be passed to a Driver option.
type DriverProperty struct {
	*queso.Property
}

// NewDriverProperty returns a new DriverProperty instance for use with the
// Driver options.
func NewDriverProperty(key string, value interface{}) *DriverProperty {
	return &DriverProperty{
		Property: queso.NewProperty(key, value),
	}
}

// WithDriverProperty is the shorthand for specifying driver details that
// reference another driver node.
func WithDriverProperty(name string, property *DriverProperty) *DriverProperty {
	key := fmt.Sprintf("%s.%s", name, property.Key)

	return NewDriverProperty(key, property.Value)
}

// WithDriverType specifies the type of driver to use with a node.
func WithDriverType(driverType string) *DriverProperty {
	return NewDriverProperty("driver", driverType)
}

// WithFile specifies the path to the image file in the local filesystem or
// a FileDriver node name.
func WithFile(file string) *DriverProperty {
	return NewDriverProperty("file", file)
}

// WithImageFile specifies the path to a disk image file (i.e. raw or qcow2).
func WithImageFile(file string) *DriverProperty {
	return NewDriverProperty("filename", file)
}

// WithNodeName defines the name of the block Driver node by which it will be
// referenced later. The name must be unique, i.e. it must not match the name of
// a different block driver node, or (if you use Drive as well) the ID of a drive.
//
// If no node name is specified, it is automatically generated. The generated node
// name is not intended to be predictable and changes between QEMU invocations.
// For the top level, an explicit node name must be specified.
func WithNodeName(name string) *DriverProperty {
	return NewDriverProperty("node-name", name)
}

// IsReadOnly opens the node read-only if enabled is true. Guest write attempts
// will fail.
//
// Note that some block Driver instances support only read-only access, either
// generally or in certain configurations. In this case, the default value
// IsReadOnly of false does not work and the option must be specified explicitly.
func IsReadOnly(enabled bool) *DriverProperty {
	return NewDriverProperty("read-only", enabled)
}

// IsAutoReadOnly specifies whether QEMU may fall back to read-only usage even when
// IsReadOnly is false is requested, or even switch between modes as needed, e.g.
// depending on whether the image file is writable or whether a writing user is
// attached to the node for a Driver.
func IsAutoReadOnly(enabled bool) *DriverProperty {
	return NewDriverProperty("auto-read-only", enabled)
}

// IsForceShare specifies whether to override the Driver image locking system of
// QEMU by forcing the node to utilize weaker shared access for permissions where
// it would normally request exclusive access. When there is the potential for
// multiple instances to have the same file open (whether this invocation of QEMU
// is the first or the second instance), both instances must permit shared access
// for the second instance to succeed at opening the file.
//
// Enabling this property requires IsReadOnly to be true.
func IsForceShare(enabled bool) *DriverProperty {
	return NewDriverProperty("force-share", enabled)
}

// IsCacheDirect specifies whether the host page cache for a Driver can be avoided.
// If true, this will attempt to do disk IO directly to the guest's memory. QEMU
// may still perform an internal copy of the data.
func IsCacheDirect(enabled bool) *DriverProperty {
	return NewDriverProperty("cache.direct", enabled)
}

// IsCacheNoFlush should be enabled if you don't care about data integrity over
// host failures for a Driver. This option tells QEMU that it never needs to write
// any data to the disk but can instead keep things in cache. If anything goes
// wrong, like your host losing power, the disk storage getting disconnected
// accidentally, etc. your image will most probably be rendered unusable.
func IsCacheNoFlush(enabled bool) *DriverProperty {
	return NewDriverProperty("cache.no-flush", enabled)
}

// DiscardRequestStatus represents the status passed in to the
// WithDiscardRequestStatus property for a Driver.
type DiscardRequestStatus string

const (
	// DiscardRequestIgnore indicates that discard requests are ignored.
	DiscardRequestIgnore DiscardRequestStatus = "ignore"

	// DiscardRequestUnmap indicates that discard requests are passed to
	// the filesystem.
	DiscardRequestUnmap DiscardRequestStatus = "unmap"
)

// WithDiscardRequestStatus specifies how discard requests are handled for a Driver.
// Some machine types may not support discard requests. See DiscardRequestStatus
// for more details.
func WithDiscardRequestStatus(status DiscardRequestStatus) *DriverProperty {
	return NewDriverProperty("discard", status)
}

// DetectZeroesStatus represents the status passed in to the
// WithDetectZeroesStatus property for a Driver.
type DetectZeroesStatus string

const (
	// DetectZeroesOff disables the automatic conversion of plain zero writes.
	DetectZeroesOff DetectZeroesStatus = "off"

	// DetectZeroesOn enables the automatic conversion of plain zero writes.
	DetectZeroesOn DetectZeroesStatus = "on"

	// DetectZeroesUnmap allows a zero write to be converted to an unmap operation
	// if WithDiscardRequestStatus is set to DiscardRequestUnmap.
	DetectZeroesUnmap DetectZeroesStatus = "unmap"
)

// WithDetectZeroesStatus enables the automatic conversion of plain zero writes
// by the OS to Driver specific optimized zero write commands.
func WithDetectZeroesStatus(status DetectZeroesStatus) *DriverProperty {
	return NewDriverProperty("detect-zeroes", status)
}

// WithAIOBackend specifies the AIO backend for a FileDriver. The Default is
// AIOBackendThreads.
func WithAIOBackend(backend AIOBackend) *DriverProperty {
	return NewDriverProperty("aio", backend)
}

// OFDLockingStatus represents the value passed to the WithOFDLocking property
// for a FileDriver.
type OFDLockingStatus string

const (
	OFDLockingOn   OFDLockingStatus = "on"
	OFDLockingOff  OFDLockingStatus = "off"
	OFDLockingAuto OFDLockingStatus = "auto"
)

// WithOFDLocking specifies whether the image file is protected with Linux
// OFD/POSIX locks for a FileDriver. The default is to use the Linux Open File
// Descriptor API if available (OFDLockingAuto), otherwise no lock is
// applied (OFDLockingOff).
func WithOFDLocking(status OFDLockingStatus) *DriverProperty {
	return NewDriverProperty("locking", status)
}

// WithBackingFile represents the reference to or definition of the backing file
// block device for a QCOW2Driver. The default is taken from the image file.
func WithBackingFile(file string) *DriverProperty {
	return NewDriverProperty("backing", file)
}

// IsLazyRefCounts enables/disables the lazy refcounts feature for a
// QCOW2Driver. The default is taken from the image file.
func IsLazyRefCounts(enabled bool) *DriverProperty {
	return NewDriverProperty("lazy-refcounts", enabled)
}

// WithTotalCacheSize represents the maximum total size of the L2 table and
// refcount block caches in bytes for a QCOW2Driver. The default is the sum
// of WithL2CacheSize and WithRefCountCacheSize values.
func WithTotalCacheSize(bytes int) *DriverProperty {
	return NewDriverProperty("cache-size", bytes)
}

// WithL2CacheSize represents the maximum size of the L2 table cache in bytes
// for a QCOW2Driver. If WithTotalCacheSize is not specified, 32M is used on
// Linux platforms, and 8M is used on non-Linux platforms; otherwise, as large
// as possible within the WithTotalCacheSize value, while permitting the requested
// or the minimal WithRefCountCacheSize.
func WithL2CacheSize(bytes string) *DriverProperty {
	return NewDriverProperty("l2-cache-size", bytes)
}

// WithRefCountCacheSize represents the maximum size of the refcount block cache
// in bytes for a QCOW2Driver. The default is 4 times the cluster size. or
// if WithTotalCacheSize is specified, the part of it which is not used for the
// L2 cache.
func WithRefCountCacheSize(bytes int) *DriverProperty {
	return NewDriverProperty("refcount-cache-size", bytes)
}

// WithCacheCleanInterval specifies the interval to clean unused entries in the
// L2 and refcount caches for a QCOW2Driver. The interval is in seconds. The
// default value is 600 on supporting platforms, and 0 on other platforms.
// Setting it to 0 disables this feature.
func WithCacheCleanInterval(seconds int) *DriverProperty {
	return NewDriverProperty("cache-clean-interval", seconds)
}

// IsPassDiscardRequest specifies whether discard requests to the qcow2 device
// should be forwarded to the data source for a QCOW2Driver. By default,
// this property is enabled if WithDiscardRequestStatus is DiscardRequestUnmap,
// otherwise it is disabled.
func IsPassDiscardRequest(enabled bool) *DriverProperty {
	return NewDriverProperty("pass-discard-request", enabled)
}

// IsPassDiscardSnapshot specifies whether discard requests for the data source
// should be issued when a snapshot operation (e.g. deleting a snapshot) frees
// clusters in the qcow2 file for a QCOW2Driver. It is enabled by default.
func IsPassDiscardSnapshot(enabled bool) *DriverProperty {
	return NewDriverProperty("pass-discard-snapshot", enabled)
}

// IsPassDiscardOther specifies whether discard requests for the data source
// should be issued on other occasions where a cluster gets freed for a
// QCOW2Driver. This property is disabled by default.
func IsPassDiscardOther(enabled bool) *DriverProperty {
	return NewDriverProperty("pass-discard-other", enabled)
}

// OverlapCheckType represents the overlap check to perform that is passed to
// the WithOverlapCheck property for a QCOW2Driver.
type OverlapCheckType string

const (
	// OverlapCheckNone indicates that no overlap checks should be performed.
	OverlapCheckNone OverlapCheckType = "none"

	// OverlapCheckConstant indicates that only checks which can be done in constant
	// time and without reading anything from disk should be performed.
	OverlapCheckConstant OverlapCheckType = "constant"

	// OverlapCheckCached indicates that only checks which can be done without
	// reading anything from disk should be performed.
	OverlapCheckCached OverlapCheckType = "cached"

	// OverlapCheckAll indicates that all overlap checks should be performed.
	OverlapCheckAll OverlapCheckType = "all"
)

// WithOverlapCheck represents which overlap checks to perform for writes to the
// image for a QCOW2Driver. The default value is OverlapCheckCached.
func WithOverlapCheck(checkType OverlapCheckType) *DriverProperty {
	return NewDriverProperty("overlap-check", checkType)
}
