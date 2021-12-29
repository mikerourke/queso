package blockdev

import "github.com/mikerourke/queso"

// Driver defines a new block driver node. Some options apply to all block
// drivers, other options are only accepted for a specific block driver. The
// driver parameter specifies the block driver to use for the given node.
func Driver(driver string, properties ...*DriverProperty) *queso.Option {
	props := []*queso.Property{{"driver", driver}}

	for _, property := range properties {
		props = append(props, property.Property)
	}

	return queso.NewOption("blockdev", "", props...)
}

// DriverFile defines the protocol-level block driver for accessing regular files.
func DriverFile(filename string, properties ...*DriverProperty) *queso.Option {
	props := []*DriverProperty{NewDriverProperty("filename", filename)}

	if properties != nil {
		props = append(props, properties...)
	}

	return Driver("file", props...)
}

// DriverRaw is the image format block driver for raw images. It is usually stacked
// on top of a protocol level block driver such as DriverFile.
func DriverRaw(file string, properties ...*DriverProperty) *queso.Option {
	props := []*DriverProperty{NewDriverProperty("file", file)}

	if properties != nil {
		props = append(props, properties...)
	}

	return Driver("raw", props...)
}

// DriverProperty represents a property that can be passed to a driver option.
type DriverProperty struct {
	*queso.Property
}

// NewDriverProperty returns a new DriverProperty instance for use with the
// driver options.
func NewDriverProperty(key string, value interface{}) *DriverProperty {
	return &DriverProperty{
		Property: queso.NewProperty(key, value),
	}
}

// WithNodeName defines the name of the block driver node by which it will be referenced
// later. The name must be unique, i.e. it must not match the name of a different block
// driver node, or (if you use Drive as well) the ID of a drive.
//
// If no node name is specified, it is automatically generated. The generated node
// name is not intended to be predictable and changes between QEMU invocations. For
// the top level, an explicit node name must be specified.
func WithNodeName(name string) *DriverProperty {
	return NewDriverProperty("node-name", name)
}

// IsReadOnly opens the node read-only if enabled is true. Guest write attempts
// will fail.
//
// Note that some block drivers support only read-only access, either generally or
// in certain configurations. In this case, the default value IsReadOnly of false does
// not work and the option must be specified explicitly.
func IsReadOnly(enabled bool) *DriverProperty {
	return NewDriverProperty("read-only", enabled)
}

// IsAutoReadOnly specifies whether QEMU may fall back to read-only usage even when
// IsReadOnly is false is requested, or even switch between modes as needed, e.g.
// depending on whether the image file is writable or whether a writing user is
// attached to the node.
func IsAutoReadOnly(enabled bool) *DriverProperty {
	return NewDriverProperty("auto-read-only", enabled)
}

// IsForceShare specifies whether to override the image locking system of QEMU by forcing
// the node to utilize weaker shared access for permissions where it would normally
// request exclusive access. When there is the potential for multiple instances to have
// the same file open (whether this invocation of QEMU is the first or the second
// instance), both instances must permit shared access for the second instance to succeed
// at opening the file.
//
// Enabling this property requires IsReadOnly to be true.
func IsForceShare(enabled bool) *DriverProperty {
	return NewDriverProperty("force-share", enabled)
}

// IsCacheDirect specifies whether the host page cache can be avoided. If true,
// this will attempt to do disk IO directly to the guest's memory. QEMU may still
// perform an internal copy of the data.
func IsCacheDirect(enabled bool) *DriverProperty {
	return NewDriverProperty("cache.direct", enabled)
}

// IsCacheNoFlush should be enabled if you don't care about data integrity over host
// failures. This option tells QEMU that it never needs to write any data to the disk
// but can instead keep things in cache. If anything goes wrong, like your host
// losing power, the disk storage getting disconnected accidentally, etc. your image
// will most probably be rendered unusable.
func IsCacheNoFlush(enabled bool) *DriverProperty {
	return NewDriverProperty("cache.no-flush", enabled)
}

// DiscardRequestStatus represents the status passed in to the WithDiscardRequestStatus
// property.
type DiscardRequestStatus string

const (
	// DiscardRequestIgnore indicates that discard requests are ignored.
	DiscardRequestIgnore DiscardRequestStatus = "ignore"

	// DiscardRequestUnmap indicates that discard requests are passed to
	// the filesystem.
	DiscardRequestUnmap DiscardRequestStatus = "unmap"
)

// WithDiscardRequestStatus specifies how discard requests are handled. Some machine
// types may not support discard requests. See DiscardRequestStatus for more
// details.
func WithDiscardRequestStatus(status DiscardRequestStatus) *DriverProperty {
	return NewDriverProperty("discard", status)
}

// DetectZeroesStatus represents the status passed in to the WithDetectZeroesStatus
// property.
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

// WithDetectZeroesStatus enables the automatic conversion of plain zero writes by
// the OS to driver specific optimized zero write commands.
func WithDetectZeroesStatus(status DetectZeroesStatus) *DriverProperty {
	return NewDriverProperty("detect-zeroes", status)
}

// WithAIOBackend specifies the AIO backend. Default is AIOBackendThreads.
//
// This property can only be used with DriverFile.
func WithAIOBackend(backend AIOBackend) *DriverProperty {
	return NewDriverProperty("aio", backend)
}

// OFDLockingStatus represents the value passed to the WithOFDLocking property.
type OFDLockingStatus string

const (
	OFDLockingOn   OFDLockingStatus = "on"
	OFDLockingOff  OFDLockingStatus = "off"
	OFDLockingAuto OFDLockingStatus = "auto"
)

// WithOFDLocking specifies whether the image file is protected with Linux
// OFD / POSIX locks. The default is to use the Linux Open File Descriptor API if
// available (OFDLockingAuto), otherwise no lock is applied (OFDLockingOff).
//
// This property can only be used with DriverFile.
func WithOFDLocking(status OFDLockingStatus) *DriverProperty {
	return NewDriverProperty("locking", status)
}

// WithBackingFile represents the reference to or definition of the backing file block
// device (default is taken from the image file).
func WithBackingFile(file string) *DriverProperty {
	return NewDriverProperty("backing", file)
}

// IsLazyRefCounts enables/disables the lazy refcounts feature. The default is
// taken from the image file.
func IsLazyRefCounts(enabled bool) *DriverProperty {
	return NewDriverProperty("lazy-refcounts", enabled)
}

// WithTotalCacheSize represents the maximum total size of the L2 table and refcount
// block caches in bytes. The default is the sum of WithL2CacheSize and
// WithRefCountCacheSize values.
func WithTotalCacheSize(bytes int) *DriverProperty {
	return NewDriverProperty("cache-size", bytes)
}

// WithL2CacheSize represents the maximum size of the L2 table cache in bytes.
// If WithTotalCacheSize is not specified, 32M is used on Linux platforms, and 8M
// is used on non-Linux platforms; otherwise, as large as possible within the
// WithTotalCacheSize value, while permitting the requested or the minimal
// WithRefCountCacheSize.
func WithL2CacheSize(bytes string) *DriverProperty {
	return NewDriverProperty("l2-cache-size", bytes)
}

// WithRefCountCacheSize represents the maximum size of the refcount block cache in
// bytes. The default is 4 times the cluster size. or if WithTotalCacheSize is
// specified, the part of it which is not used for the L2 cache.
func WithRefCountCacheSize(bytes int) *DriverProperty {
	return NewDriverProperty("refcount-cache-size", bytes)
}

// WithCacheCleanInterval specifies the interval to clean unused entries in the
// L2 and refcount caches. The interval is in seconds. The default value is 600 on
// supporting platforms, and 0 on other platforms. Setting it to 0 disables this feature.
func WithCacheCleanInterval(seconds int) *DriverProperty {
	return NewDriverProperty("cache-clean-interval", seconds)
}

// IsPassDiscardRequest specifies whether discard requests to the qcow2 device
// should be forwarded to the data source. By default, this is enabled if
// WithDiscardRequestStatus is DiscardRequestUnmap, otherwise it is disabled.
func IsPassDiscardRequest(enabled bool) *DriverProperty {
	return NewDriverProperty("pass-discard-request", enabled)
}

// IsPassDiscardSnapshot specifies whether discard requests for the data source should
// be issued when a snapshot operation (e.g. deleting a snapshot) frees clusters in
// the qcow2 file. It is enabled by default.
func IsPassDiscardSnapshot(enabled bool) *DriverProperty {
	return NewDriverProperty("pass-discard-snapshot", enabled)
}

// IsPassDiscardOther specifies whether discard requests for the data source should
// be issued on other occasions where a cluster gets freed. It is disabled by
// default.
func IsPassDiscardOther(enabled bool) *DriverProperty {
	return NewDriverProperty("pass-discard-other", enabled)
}

// OverlapCheckType represents the overlap check to perform that is passed to
// the WithOverlapCheck property.
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
// image. The default value is OverlapCheckCached.
func WithOverlapCheck(checkType OverlapCheckType) *DriverProperty {
	return NewDriverProperty("overlap-check", checkType)
}
