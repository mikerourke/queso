package chardev

import "github.com/mikerourke/queso/qemu/cli"

// VirtualConsoleBackend connects to a QEMU text console.
type VirtualConsoleBackend struct {
	*Backend
}

// NewVirtualConsoleBackend returns a new instance of [VirtualConsoleBackend].
// id is the unique ID, which can be any string up to 127 characters long.
//
//	qemu-system-* -chardev vc,id=id
func NewVirtualConsoleBackend(id string) *VirtualConsoleBackend {
	return &VirtualConsoleBackend{
		NewBackend("vc", id),
	}
}

// SetColumns specifies that the console be sized to fit a text console with
// the given column count.
//
//	qemu-system-* -chardev vc,cols=count
func (b *VirtualConsoleBackend) SetColumns(count int) *VirtualConsoleBackend {
	b.properties = append(b.properties, cli.NewProperty("cols", count))
	return b
}

// SetHeight specifies the height of the console in pixels.
//
//	qemu-system-* -chardev vc,height=pixels
func (b *VirtualConsoleBackend) SetHeight(pixels int) *VirtualConsoleBackend {
	b.properties = append(b.properties, cli.NewProperty("width", pixels))
	return b
}

// SetRows specifies that the console be sized to fit a text console with
// the given row count.
//
//	qemu-system-* -chardev vc,rows=count
func (b *VirtualConsoleBackend) SetRows(count int) *VirtualConsoleBackend {
	b.properties = append(b.properties, cli.NewProperty("rows", count))
	return b
}

// SetWidth specifies the width of the console in pixels.
//
//	qemu-system-* -chardev vc,width=pixels
func (b *VirtualConsoleBackend) SetWidth(pixels int) *VirtualConsoleBackend {
	b.properties = append(b.properties, cli.NewProperty("width", pixels))
	return b
}
