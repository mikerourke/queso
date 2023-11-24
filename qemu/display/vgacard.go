package display

import "github.com/mikerourke/queso"

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

// WithEmulatedVGACard emulates the specified VGA card. See [VGACard] for more
// details.
//
//	qemu-system-* -vga type
func WithEmulatedVGACard(card VGACard) *queso.Option {
	return queso.NewOption("vga", string(card))
}

// WithNoVGA is a shortcut for using `-vga` with the "none" option.
//
//	qemu-system-* -vga none
func WithNoVGA() *queso.Option {
	return queso.NewOption("vga", string(VGACardNone))
}
