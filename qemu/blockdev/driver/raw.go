package driver

import "github.com/mikerourke/queso"

// RawDriver is the image format block driver for raw images. It is usually stacked
// on top of a protocol level block driver such as [FileDriver].
type RawDriver struct {
	*Driver
}

// NewRawDriver returns a new instance of [RawDriver].
//
//	qemu-system-* -blockdev driver=raw
func NewRawDriver() *RawDriver {
	return &RawDriver{
		New("file"),
	}
}

// SetFile sets the reference to or definition of the data source block driver
// node (e.g. a [FileDriver] node).
//
//	qemu-system-* -blockdev driver=raw,file=name
func (d *RawDriver) SetFile(name string) *RawDriver {
	d.properties = append(d.properties, queso.NewProperty("file", name))
	return d
}
