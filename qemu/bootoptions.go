package qemu

import (
	"strings"

	"github.com/mikerourke/queso"
)

// BootOptions represent the options for booting a VM, such as boot order and
// splash image.
type BootOptions struct {
	properties []*queso.Property
}

// NewBootOptions returns a new instance of BootOptions.
//
//	qemu-system-* -boot
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

// SetBootOnce specifies Boot drives as a string of drive letters. It applies
// the boot order only on the first startup. See [BootOptions.SetBootOrder] for
// more information.
//
//	qemu-system-* -boot once=drives
func (bo *BootOptions) SetBootOnce(drives ...string) *BootOptions {
	order := strings.Join(drives, "")
	bo.properties = append(bo.properties, queso.NewProperty("once", order))
	return bo
}

// SetBootOrder specifies boot order drives as a string of drive letters for BootOptions.
// Valid drive letters depend on the target architecture.
// The x86 PC uses: a, b (floppy 1 and 2), c (first hard disk), d (first CD-ROM),
// n-p (Etherboot from network adapter 1-4), hard disk boot is the default.
//
//	qemu-system-* -boot order=drives
func (bo *BootOptions) SetBootOrder(drives ...string) *BootOptions {
	order := strings.Join(drives, "")
	bo.properties = append(bo.properties, queso.NewProperty("order", order))
	return bo
}

// SetRebootTimeout causes the guest to pause for the specified timeout (in
// milliseconds) when Boot failed, then reboot. If the timeout is -1, guest will
// not reboot and QEMU passes -1 to BIOS by default. Currently, Seabios for X86
// system supports it.
//
//	qemu-system-* -boot reboot-timeout=timeout
func (bo *BootOptions) SetRebootTimeout(timeout int) *BootOptions {
	bo.properties = append(bo.properties, queso.NewProperty("reboot-timeout", timeout))
	return bo
}

// SetSplashImage specifies a splash picture that could be passed to BIOS on Boot,
// enabling user to show it as logo, when used with IsInteractive (if
// firmware/BIOS supports them). Currently, Seabios for X86 system supports it.
//
// Limitation: The splash file could be a JPEG file or a BMP file in 24 BPP
// format (true color). The resolution should be supported by the SVGA mode, so
// the recommended is 320x240, 640x480, or 800x640.
//
//	qemu-system-* -boot splash=file
func (bo *BootOptions) SetSplashImage(file string) *BootOptions {
	bo.properties = append(bo.properties, queso.NewProperty("splash", file))
	return bo
}

// SetSplashTime specifies the amount of time to show the image specified with
// the WithSplashImage property (in milliseconds) on Boot.
//
//	qemu-system-* -boot splash-time=time
func (bo *BootOptions) SetSplashTime(time int) *BootOptions {
	bo.properties = append(bo.properties, queso.NewProperty("splash-time", time))
	return bo
}

// ToggleInteractive specifies whether interactive Boot menus/prompts should be
// enabled as far as firmware/BIOS supports them with Boot. The default is
// non-interactive boot.
//
//	qemu-system-* -boot menu=on|off
func (bo *BootOptions) ToggleInteractive(enabled bool) *BootOptions {
	bo.properties = append(bo.properties, queso.NewProperty("menu", enabled))
	return bo
}

// ToggleStrictBoot specifies that strict boot should be used.
//
//	qemu-system-* -boot strict=on|off
func (bo *BootOptions) ToggleStrictBoot(enabled bool) *BootOptions {
	bo.properties = append(bo.properties, queso.NewProperty("strict", enabled))
	return bo
}
