package qemu

import (
	"strings"

	"github.com/mikerourke/queso"
)

// BootOrderProperty represents a property that can be used with the BootOrder
// option.
type BootOrderProperty struct {
	*queso.Property
}

func newBootOrderProperty(key string, value interface{}) *BootOrderProperty {
	return &BootOrderProperty{
		Property: &queso.Property{key, value},
	}
}

// BootOrder is used to specify options for the boot order of drives.
func BootOrder(properties ...*BootOrderProperty) *queso.Option {
	if len(properties) == 0 {
		panic("at least one property is specified for BootOrder")
	}

	props := make([]*queso.Property, 0)

	for _, property := range properties {
		props = append(props, property.Property)
	}

	return queso.NewOption("boot", "", props...)
}

// WithOrder specifies boot order drives as a string of drive letters.
// Valid drive letters depend on the target architecture.
// The x86 PC uses: a, b (floppy 1 and 2), c (first hard disk), d (first CD-ROM),
// n-p (Etherboot from network adapter 1-4), hard disk boot is the default.
//
// This should not be used together with the WithBootIndex property of devices,
// since the firmware implementations normally do not support both at the same
// time.
func WithOrder(drives ...string) *BootOrderProperty {
	order := strings.Join(drives, "")

	return newBootOrderProperty("order", order)
}

// WithOnce specifies boot order drives as a string of drive letters. It applies
// the boot order only on the first startup. See WithOrder for more information.
//
// This should not be used together with the WithBootIndex property of devices,
// since the firmware implementations normally do not support both at the same
// time.
func WithOnce(drives ...string) *BootOrderProperty {
	order := strings.Join(drives, "")

	return newBootOrderProperty("once", order)
}

// IsMenu specifies whether interactive boot menus/prompts should be enabled as
// far as firmware/BIOS supports them. The default is non-interactive boot.
func IsMenu(enabled bool) *BootOrderProperty {
	return newBootOrderProperty("menu", enabled)
}

// WithSplashImage specifies a splash picture that could be passed to BIOS, enabling
// user to show it as logo, when used with IsMenu (if firmware/BIOS supports them).
// Currently, Seabios for X86 system supports it.
//
// Limitation: The splash file could be a JPEG file or a BMP file in 24 BPP
// format (true color). The resolution should be supported by the SVGA mode, so
// the recommended is 320x240, 640x480, or 800x640.
func WithSplashImage(filename string) *BootOrderProperty {
	return newBootOrderProperty("splash", filename)
}

// WithSplashTime specifies the amount of time to show the image specified with
// the WithSplashImage property (in milliseconds).
func WithSplashTime(duration int) *BootOrderProperty {
	return newBootOrderProperty("splash-time", duration)
}

// WithRebootTimeout causes the guest to pause for the specified timeout (in
// milliseconds) when boot failed, then reboot. If the timeout is -1, guest will
// not reboot and QEMU passes -1 to BIOS by default. Currently, Seabios for X86
// system supports it.
func WithRebootTimeout(timeout int) *BootOrderProperty {
	return newBootOrderProperty("reboot-timeout", timeout)
}

// IsStrict specifies that strict boot should be used.
func IsStrict(enabled bool) *BootOrderProperty {
	return newBootOrderProperty("strict", enabled)
}
