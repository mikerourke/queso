package qemu

import (
	"strings"

	"github.com/mikerourke/queso"
)

type BootOptions struct {
	properties []*queso.Property
}

func NewBootOptions() *BootOptions {
	return &BootOptions{
		properties: make([]*queso.Property, 0),
	}
}

func (bo *BootOptions) option() *queso.Option {
	if len(bo.properties) == 0 {
		panic("at least one property must be specified for BootOptions")
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

// SetBootOnce specifies Boot drives as a string of drive letters. It applies
// the boot order only on the first startup. See WithBootOrder for more information.
//
// This should not be used together with the WithBootIndex property of devices,
// since the firmware implementations normally do not support both at the same
// time.
func (bo *BootOptions) SetBootOnce(drives ...string) *BootOptions {
	order := strings.Join(drives, "")
	bo.properties = append(bo.properties, queso.NewProperty("once", order))
	return bo
}

// ToggleInteractive specifies whether interactive Boot menus/prompts should be
// enabled as far as firmware/BIOS supports them with Boot. The default is
// non-interactive boot.
func (bo *BootOptions) ToggleInteractive(interactive bool) *BootOptions {
	bo.properties = append(bo.properties, queso.NewProperty("menu", interactive))
	return bo
}

// SetSplashImage specifies a splash picture that could be passed to BIOS on Boot,
// enabling user to show it as logo, when used with IsInteractive (if
// firmware/BIOS supports them). Currently, Seabios for X86 system supports it.
//
// Limitation: The splash file could be a JPEG file or a BMP file in 24 BPP
// format (true color). The resolution should be supported by the SVGA mode, so
// the recommended is 320x240, 640x480, or 800x640.
func (bo *BootOptions) SetSplashImage(file string) *BootOptions {
	bo.properties = append(bo.properties, queso.NewProperty("splash", file))
	return bo
}

// SetSplashTime specifies the amount of time to show the image specified with
// the WithSplashImage property (in milliseconds) on Boot.
func (bo *BootOptions) SetSplashTime(ms int) *BootOptions {
	bo.properties = append(bo.properties, queso.NewProperty("splash-time", ms))
	return bo
}

// SetRebootTimeout causes the guest to pause for the specified timeout (in
// milliseconds) when Boot failed, then reboot. If the timeout is -1, guest will
// not reboot and QEMU passes -1 to BIOS by default. Currently, Seabios for X86
// system supports it.
func (bo *BootOptions) SetRebootTimeout(ms int) *BootOptions {
	bo.properties = append(bo.properties, queso.NewProperty("reboot-timeout", ms))
	return bo
}

// ToggleStrictBoot specifies that strict Boot should be used.
func (bo *BootOptions) ToggleStrictBoot(strict bool) *BootOptions {
	bo.properties = append(bo.properties, queso.NewProperty("strict", strict))
	return bo
}
