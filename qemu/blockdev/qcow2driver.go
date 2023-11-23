package blockdev

import "github.com/mikerourke/queso/qemu/cli"

// QCOW2Driver is the image format block driver for raw images. It is usually stacked
// on top of a protocol level block driver such as [FileDriver].
type QCOW2Driver struct {
	*Driver
}

// NewQCOW2Driver returns a new instance of [QCOW2Driver].
//
//	qemu-system-* -blockdev driver=qcow2
func NewQCOW2Driver() *QCOW2Driver {
	return &QCOW2Driver{
		NewDriver("qcow2"),
	}
}

// SetBackingFile sets the reference to or definition of the backing file block device
// (default is taken from the image file). It is allowed to pass null here in order
// to disable the default backing file.
//
//	qemu-system-* -blockdev driver=qcow2,backing=name
func (d *QCOW2Driver) SetBackingFile(name string) *QCOW2Driver {
	d.properties = append(d.properties, cli.NewProperty("backing", name))
	return d
}

// SetCacheCleanInterval sets the interval to clean unused entries in the L2 and
// refcount caches. The interval is in seconds. The default value is 600 on supporting
// platforms, and 0 on other platforms. Setting it to 0 disables this feature.
//
//	qemu-system-* -blockdev driver=qcow2,cache-clean-interval=seconds
func (d *QCOW2Driver) SetCacheCleanInterval(seconds int) *QCOW2Driver {
	d.properties = append(d.properties, cli.NewProperty("cache-clean-interval", seconds))
	return d
}

// SetFile sets the reference to or definition of the data source block driver
// node (e.g. a [FileDriver] node).
//
//	qemu-system-* -blockdev driver=qcow2,file=name
func (d *QCOW2Driver) SetFile(name string) *QCOW2Driver {
	d.properties = append(d.properties, cli.NewProperty("file", name))
	return d
}

// SetL2CacheSize sets the maximum size of the L2 table cache in bytes.
// The default if [QCOW2Driver.SetTotalCacheSize] is not specified is 32M on Linux
// platforms, and 8M on non-Linux platforms.
//
// Ootherwise, as large as possible within the [QCOW2Driver.SetTotalCacheSize], while
// permitting the requested or the minimal [QCOW2Driver.SetRefcountCacheSize].
//
//	qemu-system-* -blockdev driver=qcow2,l2-cache-size=bytes
func (d *QCOW2Driver) SetL2CacheSize(bytes int) *QCOW2Driver {
	d.properties = append(d.properties, cli.NewProperty("l2-cache-size", bytes))
	return d
}

// OverlapCheck represents the overlap check to perform that is passed to
// the [QCOW2Driver.SetOverlapCheck] method.
type OverlapCheck string

const (
	// OverlapNone indicates that no overlap checks should be performed.
	OverlapNone OverlapCheck = "none"

	// OverlapConstant indicates that only checks which can be done in constant
	// time and without reading anything from disk should be performed.
	OverlapConstant OverlapCheck = "constant"

	// OverlapCached indicates that only checks which can be done without
	// reading anything from disk should be performed.
	OverlapCached OverlapCheck = "cached"

	// OverlapAll indicates that all overlap checks should be performed.
	OverlapAll OverlapCheck = "all"
)

// SetOverlapCheck specified which overlap checks to perform for writes to the
// image. The default value is OverlapCached. For details or finer granularity
// control refer to the QAPI documentation of blockdev-add.
//
//	qemu-system-* -blockdev driver=qcow2,overlap-check=check
func (d *QCOW2Driver) SetOverlapCheck(check OverlapCheck) *QCOW2Driver {
	d.properties = append(d.properties, cli.NewProperty("overlap-check", check))
	return d
}

// SetRefcountCacheSize sets the maximum size of the refcount block cache in bytes.
// The default is 4 times the cluster size; or if [QCOW2Driver.SetTotalCacheSize] is
// specified, the part of it which is not used for the L2 cache.
//
//	qemu-system-* -blockdev driver=qcow2,refcount-cache-size=bytes
func (d *QCOW2Driver) SetRefcountCacheSize(bytes int) *QCOW2Driver {
	d.properties = append(d.properties, cli.NewProperty("refcount-cache-size", bytes))
	return d
}

// SetTotalCacheSize sets the maximum total size of the L2 table and refcount block
// caches in bytes. The default is the sum of [QCOW2Driver.SetL2CacheSize] and
// [QCOW2Driver.SetRefcountCacheSize]).
//
//	qemu-system-* -blockdev driver=qcow2,cache-size=bytes
func (d *QCOW2Driver) SetTotalCacheSize(bytes int) *QCOW2Driver {
	d.properties = append(d.properties, cli.NewProperty("cache-size", bytes))
	return d
}

// ToggleDiscardNoUnref enables or disables data clusters to remain preallocated when
// they are no longer used, e.g. because they are discarded or converted to zero
// clusters. As usual, whether the old data is discarded or kept on the protocol
// level (i.e. in the image file) depends on the setting of the pass-discard-request
// option. Keeping the clusters preallocated prevents QCOW2 fragmentation that
// would otherwise be caused by freeing and re-allocating them later.
//
// Besides potential performance degradation, such fragmentation can lead to increased
// allocation of clusters past the end of the image file, resulting in image files
// whose file length can grow much larger than their guest disk size would suggest.
//
// If image file length is of concern (e.g. when storing QCOW2 images directly on
// block devices), you should consider enabling this option.
func (d *Driver) ToggleDiscardNoUnref(enabled bool) *Driver {
	d.properties = append(d.properties, cli.NewProperty("discard-no-unref", enabled))
	return d
}

// ToggleLazyRefcounts specifies whether to enable the lazy refcounts feature.
// The default is taken from the image file.
//
//	qemu-system-* -blockdev driver=qcow2,lazy-refcounts=on|off
func (d *Driver) ToggleLazyRefcounts(enabled bool) *Driver {
	d.properties = append(d.properties, cli.NewProperty("lazy-refcounts", enabled))
	return d
}

// TogglePassDiscardOther specified whether discard requests for the data source should
// be issued on other occasions where a cluster gets freed. The default is false.
//
//	qemu-system-* -blockdev driver=qcow2,pass-discard-other=on|off
func (d *Driver) TogglePassDiscardOther(enabled bool) *Driver {
	d.properties = append(d.properties, cli.NewProperty("pass-discard-other", enabled))
	return d
}

// TogglePassDiscardRequests specifies whether discard requests to the QCOW2 device should
// be forwarded to the data source. The default is true if [Driver.ToggleDiscard] is
// specified, false otherwise.
//
//	qemu-system-* -blockdev driver=qcow2,pass-discard-request=on|off
func (d *Driver) TogglePassDiscardRequests(enabled bool) *Driver {
	d.properties = append(d.properties, cli.NewProperty("pass-discard-request", enabled))
	return d
}

// TogglePassDiscardSnapshots specifies whether discard requests for the data source
// should be issued when a snapshot operation (e.g. deleting a snapshot) frees
// clusters in the QCOW2 file. The default is true.
//
//	qemu-system-* -blockdev driver=qcow2,pass-discard-snapshot=on|off
func (d *Driver) TogglePassDiscardSnapshots(enabled bool) *Driver {
	d.properties = append(d.properties, cli.NewProperty("pass-discard-snapshot", enabled))
	return d
}
