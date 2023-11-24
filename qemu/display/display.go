// Package display is used to manage displays for the QEMU guest.
// See https://qemu.readthedocs.io/en/latest/system/invocation.html#hxtool-3 for more details.
package display

import (
	"fmt"
	"strconv"

	"github.com/mikerourke/queso/qemu/cli"
)

// Display represents a generic display of any allowable type.
type Display struct {
	// Type is the type of the display.
	Type       string
	properties []*cli.Property
}

// New returns a new instance of a generic [Display].
func New(displayType string) *Display {
	return &Display{
		Type:       displayType,
		properties: make([]*cli.Property, 0),
	}
}

func (d *Display) option() *cli.Option {
	return cli.NewOption("display", d.Type, d.properties...)
}

// SetProperty sets arbitrary properties on the [Display].
func (d *Display) SetProperty(key string, value interface{}) *Display {
	d.properties = append(d.properties, cli.NewProperty(key, value))
	return d
}

// WithPortraitMode rotates graphical output 90 degrees left (only PXA LCD).
//
//	qemu-system-* -portrait
func WithPortraitMode() *cli.Option {
	return cli.NewOption("portrait", "")
}

// WithRotation rotates graphical output by the specified degrees (only PXA LCD).
//
//	qemu-system-* -rotate degrees
func WithRotation(degrees int) *cli.Option {
	return cli.NewOption("rotate", strconv.Itoa(degrees))
}

// WithFullScreen starts in full screen.
//
//	qemu-system-* -full-screen
func WithFullScreen() *cli.Option {
	return cli.NewOption("full-screen", "")
}

// WithScreenResolution sets the initial graphical resolution, and depth
// (PPC, SPARC only). If depth is 0, uses the default.
//
//	qemu-system-* -g <width>x<height>[x<depth>]
func WithScreenResolution(width int, height int, depth int) *cli.Option {
	resolution := fmt.Sprintf("%dx%d", width, height)

	if depth != 0 {
		resolution = fmt.Sprintf("%sx%d", resolution, depth)
	}

	return cli.NewOption("g", resolution)
}
