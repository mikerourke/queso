// Package display is used to define display options for use with QEMU.
// See https://qemu.readthedocs.io/en/latest/system/invocation.html#hxtool-3
// for more details.
package display

import (
	"fmt"
	"strconv"

	"github.com/mikerourke/queso"
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
func NoGraphic() *queso.Option {
	return queso.NewOption("nographic", "")
}

// Curses displays the VGA output when in text mode using a curses/ncurses interface.
// Nothing is displayed in graphical mode.
func Curses() *queso.Option {
	return queso.NewOption("curses", "")
}

// PortraitMode rotates graphical output 90 degrees left (only PXA LCD).
func PortraitMode() *queso.Option {
	return queso.NewOption("portrait", "")
}

// Rotate rotates graphical output by the specified degrees (only PXA LCD).
func Rotate(degrees int) *queso.Option {
	return queso.NewOption("rotate", strconv.Itoa(degrees))
}

// VGACard represents the VGA card to emulate.
type VGACard string

const (
	// VGACardNone is used to disable the VGA card.
	VGACardNone VGACard = "none"

	// VGACardCirrus represents the Cirrus Logic GD5446 Video card. All Windows versions
	// starting from Windows 95 should recognize and use this graphic card. For optimal
	// performances, use 16-bit color depth in the guest and the host OS.
	VGACardCirrus VGACard = "cirrus"

	// VGACardStandard represents the Standard VGA card with Bochs VBE extensions.
	// If your guest OS supports the VESA 2.0 VBE extensions (e.g. Windows XP) and
	// if you want to use high resolution modes (>= 1280 x 1024 x 16) then you should
	// use this option.
	VGACardStandard VGACard = "std"

	// VGACardVMWare represents the VMWare SVGA-II compatible adapter. Use it if
	// you have sufficiently recent XFree86/XOrg server or Windows guest with a
	// driver for this card.
	VGACardVMWare VGACard = "vmware"

	// VGACardQXL represents the QXL para-virtual graphic card. It is VGA compatible
	// (including VESA 2.0 VBE support). Works best with qxl guest drivers installed
	// though. Recommended choice when using the spice protocol.
	VGACardQXL VGACard = "qxl"

	// VGACardTCX represents the Sun TCX framebuffer. This is the default framebuffer
	// for sun4m machines and offers both 8-bit and 24-bit color depths at a fixed
	// resolution of 1024 x 768.
	VGACardTCX VGACard = "tcx"

	// VGACardCG3 represents the un cg-three framebuffer. This is a simple 8-bit
	// framebuffer for sun4m machines available in both 1024 x 768 (OpenBIOS) and
	// 1152 x 900 (OBP) resolutions aimed at people wishing to run older Solaris versions.
	VGACardCG3 VGACard = "cg3"

	// VGACardVirtio represents the Virtio VGA card.
	VGACardVirtio VGACard = "virtio"
)

// EmulateVGACard emulates the specified VGA card.
func EmulateVGACard(card VGACard) *queso.Option {
	return queso.NewOption("vga", string(card))
}

// WithNoVGA is a shortcut for using `-vga` with the "none" option.
func WithNoVGA() *queso.Option {
	return queso.NewOption("vga", string(VGACardNone))
}

// FullScreen starts in full screen.
func FullScreen() *queso.Option {
	return queso.NewOption("full-screen", "")
}

// ScreenResolution sets the initial graphical resolution, and depth (PPC, SPARC only).
// If depth is 0, uses the default.
func ScreenResolution(width int, height int, depth int) *queso.Option {
	resolution := fmt.Sprintf("%dx%d", width, height)

	if depth != 0 {
		resolution = fmt.Sprintf("%sx%d", resolution, depth)
	}

	return queso.NewOption("g", resolution)
}
