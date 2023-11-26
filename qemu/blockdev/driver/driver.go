// Package driver is used to define and manage block driver nodes.
package driver

import "github.com/mikerourke/queso"

// Driver defines a new block driver node.
type Driver struct {
	*queso.Entity
}

// New returns a new instance of [Driver]. name is the name of the driver type.
//
//	qemu-system-* -blockdev driver=name
func New(name string) *Driver {
	var entity *queso.Entity

	// For a Drive, the name gets set to an empty string. In that case, we
	// don't want to add a "driver" property with that name. Drives are defined
	// with the -drive flag, rather than -blockdev, so we specify the flag
	// to ensure QEMU is called correctly.
	if name == "" {
		entity = queso.NewEntity("drive", "")
	} else {
		entity = queso.NewEntity("blockdev", "")
		entity.SetProperty("driver", name)
	}

	return &Driver{entity}
}

// DetectZeroesStatus represents the status passed in to the [Driver.SetDetectZeroes]
// method for a driver.
type DetectZeroesStatus string

const (
	// DetectZeroesOff disables the automatic conversion of plain zero writes.
	DetectZeroesOff DetectZeroesStatus = "off"

	// DetectZeroesOn enables the automatic conversion of plain zero writes.
	DetectZeroesOn DetectZeroesStatus = "on"

	// DetectZeroesUnmap allows a zero write to be converted to an unmap operation
	// if ToggleDiscard is called with false.
	DetectZeroesUnmap DetectZeroesStatus = "unmap"
)

// SetDetectZeroes enables the automatic conversion of plain zero writes
// by the OS to driver specific optimized zero write commands. You may even choose
// [DetectZeroesUnmap] if discard is set to “unmap” to allow a zero write to be
// converted to an unmap operation.
//
//	qemu-system-* -blockdev driver=<name>,detect-zeroes=detect-zeroes
func (d *Driver) SetDetectZeroes(status DetectZeroesStatus) *Driver {
	d.SetProperty("detect-zeroes", status)
	return d
}

// SetNodeName defines the name of the block driver node by which it will be
// referenced later. The name must be unique, i.e. it must not match the name of a
// different block driver node, or (if you use -drive as well) the ID of a drive.
//
// If no node name is specified, it is automatically generated. The generated node
// name is not intended to be predictable and changes between QEMU invocations.
// For the top level, an explicit node name must be specified.
//
//	qemu-system-* -blockdev driver=<name>,node-name=name
func (d *Driver) SetNodeName(name string) *Driver {
	d.SetProperty("node-name", name)
	return d
}

// ToggleAutoReadOnly may cause QEMU to fall back to read-only usage
// even when [Driver.ToggleReadOnly] is called with false, or even switch between
// modes as needed, e.g. depending on whether the image file is writable or whether
// a writing user is attached to the node.
//
//	qemu-system-* -blockdev driver=<name>,auto-read-only=on|off
func (d *Driver) ToggleAutoReadOnly(enabled bool) *Driver {
	d.SetProperty("auto-read-only", enabled)
	return d
}

// ToggleDirectCache specifies whether the host page cache for a driver can be avoided.
// If true, this will attempt to do disk IO directly to the guest's memory. QEMU
// may still perform an internal copy of the data.
//
//	qemu-system-* -blockdev driver=<name>,cache.direct=on|off
func (d *Driver) ToggleDirectCache(enabled bool) *Driver {
	d.SetProperty("cache.direct", enabled)
	return d
}

// ToggleDiscard controls whether discard (also known as trim or unmap) requests are
// ignored (false) or passed to the filesystem (true). Some machine types may not
// support discard requests.
//
//	qemu-system-* -blockdev driver=<name>,discard=on|off
func (d *Driver) ToggleDiscard(enabled bool) *Driver {
	d.SetProperty("discard", enabled)
	return d
}

// ToggleForceShare overrides the image locking system of QEMU by forcing the node
// to utilize weaker shared access for permissions where it would normally request
// exclusive access. When there is the potential for multiple instances to have the
// same file open (whether this invocation of QEMU is the first or the second instance),
// both instances must permit shared access for the second instance to succeed at
// opening the file.
//
// Calling this with true requires [Driver.ToggleReadOnly] to be called with true.
//
//	qemu-system-* -blockdev driver=<name>,force-share=on|off
func (d *Driver) ToggleForceShare(enabled bool) *Driver {
	d.SetProperty("force-share", enabled)
	return d
}

// ToggleNoCacheFlushing should be enabled if you don't care about data integrity over
// host failures for a [Driver]. This option tells QEMU that it never needs to write
// any data to the disk but can instead keep things in cache. If anything goes
// wrong, like your host losing power, the disk storage getting disconnected
// accidentally, etc. your image will most probably be rendered unusable.
//
//	qemu-system-* -blockdev driver=<name>,cache.no-flush=on|off
func (d *Driver) ToggleNoCacheFlushing(enabled bool) *Driver {
	d.SetProperty("cache.no-flush", enabled)
	return d
}

// ToggleReadOnly opens the node read-only when enabled. Guest write attempts will fail.
//
// Note that some block drivers support only read-only access, either generally
// or in certain configurations. In this case, the default value of false does not
// work and the option must be specified explicitly.
//
//	qemu-system-* -blockdev driver=<name>,read-only=on|off
func (d *Driver) ToggleReadOnly(enabled bool) *Driver {
	d.SetProperty("read-only", enabled)
	return d
}
