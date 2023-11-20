package qemu

import (
	"strings"

	"github.com/mikerourke/queso"
)

// BootOptions are used to specify options for booting, such as the boot order
// of drives.
//
// # Example 1
//
// Try to boot from network first, then from hard disk:
//
//	qemu.New("qemu-system-x86_64").Use(
//		qemu.NewBootOptions().)
//
// Invocation
//
//	qemu-system-x86_64 -boot order=nc
//
// # Example 2
//
// Boot from CD-ROM first, switch back to default order after reboot:
//
//	qemu.New("qemu-system-x86_64").SetOptions(
//		qemu.Boot(qemu.WithBootOnce("d")))
//
// Invocation
//
//	qemu-system-x86_64 -boot once=d
//
// # Example 3
//
// Boot with a splash picture for 5 seconds:
//
//	qemu.New("qemu-system-x86_64").SetOptions(
//		qemu.Boot(
//			qemu.IsInteractive(true),
//			qemu.WithSplashImage("/root/boot.bmp"),
//			qemu.WithSplashTime(5000)))
//
// Invocation
//
//	qemu-system-x86_64 -boot menu=on,splash=/root/boot.bmp,splash-time=5000
type BootOptions struct {
	properties []*queso.Property
}

func NewBootOptions() *BootOptions {
	return &BootOptions{}
}

func (bo *BootOptions) Option() *queso.Option {
	if len(bo.properties) == 0 {
		panic("at least one property is specified for BootOptions")
	}

	return queso.NewOption("boot", "", bo.properties...)
}

// SetBootOrder specifies boot order drives as a string of drive letters for BootOptions.
// Valid drive letters depend on the target architecture.
// The x86 PC uses: a, b (floppy 1 and 2), c (first hard disk), d (first CD-ROM),
// n-p (Etherboot from network adapter 1-4), hard disk boot is the default.
//
// This should not be used together with the WithBootIndex property of devices,
// since the firmware implementations normally do not support both at the same
// time.
func (bo *BootOptions) SetBootOrder(drives ...string) *BootOptions {
	order := strings.Join(drives, "")
	bo.properties = append(bo.properties, queso.NewProperty("order", order))

	return bo
}

func Boot(properties ...*BootProperty) *queso.Option {
	if len(properties) == 0 {
		panic("at least one property is specified for Boot")
	}

	props := make([]*queso.Property, 0)

	for _, property := range properties {
		props = append(props, property.Property)
	}

	return queso.NewOption("boot", "", props...)
}

// BootProperty represents a property that can be used with Boot.
type BootProperty struct {
	*queso.Property
}

// NewBootProperty returns a new instance of BootProperty.
func NewBootProperty(key string, value interface{}) *BootProperty {
	return &BootProperty{
		Property: queso.NewProperty(key, value),
	}
}

// WithBootOrder specifies boot order drives as a string of drive letters for Boot.
// Valid drive letters depend on the target architecture.
// The x86 PC uses: a, b (floppy 1 and 2), c (first hard disk), d (first CD-ROM),
// n-p (Etherboot from network adapter 1-4), hard disk boot is the default.
//
// This should not be used together with the WithBootIndex property of devices,
// since the firmware implementations normally do not support both at the same
// time.
func WithBootOrder(drives ...string) *BootProperty {
	order := strings.Join(drives, "")

	return NewBootProperty("order", order)
}

// WithBootOnce specifies Boot drives as a string of drive letters. It applies
// the boot order only on the first startup. See WithBootOrder for more information.
//
// This should not be used together with the WithBootIndex property of devices,
// since the firmware implementations normally do not support both at the same
// time.
func WithBootOnce(drives ...string) *BootProperty {
	order := strings.Join(drives, "")

	return NewBootProperty("once", order)
}

// IsInteractive specifies whether interactive Boot menus/prompts should be
// enabled as far as firmware/BIOS supports them with Boot. The default is
// non-interactive boot.
func IsInteractive(interactive bool) *BootProperty {
	return NewBootProperty("menu", interactive)
}

// WithSplashImage specifies a splash picture that could be passed to BIOS on Boot,
// enabling user to show it as logo, when used with IsInteractive (if
// firmware/BIOS supports them). Currently, Seabios for X86 system supports it.
//
// Limitation: The splash file could be a JPEG file or a BMP file in 24 BPP
// format (true color). The resolution should be supported by the SVGA mode, so
// the recommended is 320x240, 640x480, or 800x640.
func WithSplashImage(file string) *BootProperty {
	return NewBootProperty("splash", file)
}

// WithSplashTime specifies the amount of time to show the image specified with
// the WithSplashImage property (in milliseconds) on Boot.
func WithSplashTime(milliseconds int) *BootProperty {
	return NewBootProperty("splash-time", milliseconds)
}

// WithRebootTimeout causes the guest to pause for the specified timeout (in
// milliseconds) when Boot failed, then reboot. If the timeout is -1, guest will
// not reboot and QEMU passes -1 to BIOS by default. Currently, Seabios for X86
// system supports it.
func WithRebootTimeout(milliseconds int) *BootProperty {
	return NewBootProperty("reboot-timeout", milliseconds)
}

// IsStrictBoot specifies that strict Boot should be used.
func IsStrictBoot(strict bool) *BootProperty {
	return NewBootProperty("strict", strict)
}
