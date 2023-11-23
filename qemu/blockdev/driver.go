package blockdev

import "github.com/mikerourke/queso/qemu/cli"

// Driver defines a new block driver node.
type Driver struct {
	properties []*cli.Property
}

// NewDriver returns a new instance of [Driver].
//
//	qemu-system-* -blockdev driver=file
func NewDriver(name string) *Driver {
	return &Driver{
		properties: []*cli.Property{
			cli.NewProperty("driver", name),
		},
	}
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
//	qemu-system-* -blockdev driver=file,detect-zeroes=detect-zeroes
func (d *Driver) SetDetectZeroes(status DetectZeroesStatus) *Driver {
	d.properties = append(d.properties, cli.NewProperty("detect-zeroes", status))
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
//	qemu-system-* -blockdev driver=file,node-name=name
func (d *Driver) SetNodeName(name string) *Driver {
	d.properties = append(d.properties, cli.NewProperty("node-name", name))
	return d
}

// ToggleAutoReadOnly may cause QEMU to fall back to read-only usage
// even when [Driver.ToggleReadOnly] is called with false, or even switch between
// modes as needed, e.g. depending on whether the image file is writable or whether
// a writing user is attached to the node.
//
//	qemu-system-* -blockdev driver=file,auto-read-only=on|off
func (d *Driver) ToggleAutoReadOnly(enabled bool) *Driver {
	d.properties = append(d.properties, cli.NewProperty("auto-read-only", enabled))
	return d
}

// ToggleDirectCache specifies whether the host page cache for a driver can be avoided.
// If true, this will attempt to do disk IO directly to the guest's memory. QEMU
// may still perform an internal copy of the data.
//
//	qemu-system-* -blockdev driver=file,cache.direct=on|off
func (d *Driver) ToggleDirectCache(enabled bool) *Driver {
	d.properties = append(d.properties, cli.NewProperty("cache.direct", enabled))
	return d
}

// ToggleDiscard controls whether discard (also known as trim or unmap) requests are
// ignored (false) or passed to the filesystem (true). Some machine types may not
// support discard requests.
//
//	qemu-system-* -blockdev driver=file,discard=on|off
func (d *Driver) ToggleDiscard(enabled bool) *Driver {
	d.properties = append(d.properties, cli.NewProperty("discard", enabled))
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
//	qemu-system-* -blockdev driver=file,force-share=on|off
func (d *Driver) ToggleForceShare(enabled bool) *Driver {
	d.properties = append(d.properties, cli.NewProperty("force-share", enabled))
	return d
}

// ToggleNoCacheFlushing should be enabled if you don't care about data integrity over
// host failures for a Driver. This option tells QEMU that it never needs to write
// any data to the disk but can instead keep things in cache. If anything goes
// wrong, like your host losing power, the disk storage getting disconnected
// accidentally, etc. your image will most probably be rendered unusable.
//
//	qemu-system-* -blockdev driver=file,cache.no-flush=on|off
func (d *Driver) ToggleNoCacheFlushing(enabled bool) *Driver {
	d.properties = append(d.properties, cli.NewProperty("cache.no-flush", enabled))
	return d
}

// ToggleReadOnly opens the node read-only when enabled. Guest write attempts will fail.
//
// Note that some block drivers support only read-only access, either generally
// or in certain configurations. In this case, the default value of false does not
// work and the option must be specified explicitly.
//
//	qemu-system-* -blockdev driver=file,read-only=on|off
func (d *Driver) ToggleReadOnly(enabled bool) *Driver {
	d.properties = append(d.properties, cli.NewProperty("read-only", enabled))
	return d
}
