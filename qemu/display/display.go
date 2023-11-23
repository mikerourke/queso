// Package display is used to manage displays for the QEMU guest.
// See https://qemu.readthedocs.io/en/latest/system/invocation.html#hxtool-3 for more details.
package display

import (
	"fmt"
	"strconv"

	"github.com/mikerourke/queso/internal/cli"
)

type Type string

const (
	TypeSpiceApp    Type = "spice-app"
	TypeDBus        Type = "dbus"
	TypeSDL         Type = "sdl"
	TypeGTK         Type = "gtk"
	TypeCurses      Type = "curses"
	TypeEGLHeadless Type = "egl-headless"
	TypeVNC         Type = "vnc"
)

// NoGraphic totally disables graphical output so that QEMU is a simple command
// line application. The emulated serial port is redirected on the console and
// muxed with the monitor (unless redirected elsewhere explicitly). Therefore, you
// can still use QEMU to debug a Linux kernel with a serial console.
func NoGraphic() *cli.Option {
	return cli.NewOption("nographic", "")
}

// Curses displays the VGA output when in text mode using a curses/ncurses interface.
// Nothing is displayed in graphical mode.
func Curses() *cli.Option {
	return cli.NewOption("curses", "")
}

// PortraitMode rotates graphical output 90 degrees left (only PXA LCD).
func PortraitMode() *cli.Option {
	return cli.NewOption("portrait", "")
}

// Rotate rotates graphical output by the specified degrees (only PXA LCD).
func Rotate(degrees int) *cli.Option {
	return cli.NewOption("rotate", strconv.Itoa(degrees))
}

// FullScreen starts in full screen.
func FullScreen() *cli.Option {
	return cli.NewOption("full-screen", "")
}

// ScreenResolution sets the initial graphical resolution, and depth (PPC, SPARC only).
// If depth is 0, uses the default.
func ScreenResolution(width int, height int, depth int) *cli.Option {
	resolution := fmt.Sprintf("%dx%d", width, height)

	if depth != 0 {
		resolution = fmt.Sprintf("%sx%d", resolution, depth)
	}

	return cli.NewOption("g", resolution)
}
