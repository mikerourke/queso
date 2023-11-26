// Package display is used to manage displays for the QEMU guest.
// See https://qemu.readthedocs.io/en/latest/system/invocation.html#hxtool-3 for more details.
package display

import (
	"fmt"
	"strconv"

	"github.com/mikerourke/queso"
)

// Display represents a generic display of any allowable type.
type Display struct {
	// Type is the type of the display.
	Type       string
	properties []*queso.Property
}

// New returns a new instance of a generic [Display].
func New(displayType string) *Display {
	return &Display{
		Type:       displayType,
		properties: make([]*queso.Property, 0),
	}
}

// Option returns the invoked option that gets converted to an argument when
// passed to QEMU.
func (d *Display) Option() *queso.Option {
	return queso.NewOption("display", d.Type, d.properties...)
}

// SetProperty sets arbitrary properties on the [Display].
func (d *Display) SetProperty(key string, value interface{}) *Display {
	d.properties = append(d.properties, queso.NewProperty(key, value))
	return d
}

// WithPortraitMode rotates graphical output 90 degrees left (only PXA LCD).
//
//	qemu-system-* -portrait
func WithPortraitMode() *queso.Option {
	return queso.NewOption("portrait", "")
}

// WithRotation rotates graphical output by the specified degrees (only PXA LCD).
//
//	qemu-system-* -rotate degrees
func WithRotation(degrees int) *queso.Option {
	return queso.NewOption("rotate", strconv.Itoa(degrees))
}

// WithFullScreen starts in full screen.
//
//	qemu-system-* -full-screen
func WithFullScreen() *queso.Option {
	return queso.NewOption("full-screen", "")
}

// WithScreenResolution sets the initial graphical resolution, and depth
// (PPC, SPARC only). If depth is 0, uses the default.
//
//	qemu-system-* -g <width>x<height>[x<depth>]
func WithScreenResolution(width int, height int, depth int) *queso.Option {
	resolution := fmt.Sprintf("%dx%d", width, height)

	if depth != 0 {
		resolution = fmt.Sprintf("%sx%d", resolution, depth)
	}

	return queso.NewOption("g", resolution)
}
