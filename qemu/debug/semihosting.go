package debug

import "github.com/mikerourke/queso"

// EnableSemiHosting enables semi-hosting mode (ARM, M68K, Xtensa, MIPS, Nios II,
// RISC-V only).
//
// Note that this allows guest direct access to the host filesystem, so should
// only be used with a trusted guest OS.
//
// See https://www.qemu.org/docs/master/about/emulation.html#semihosting for
// more details.
//
//	qemu-system-* -semihosting
func EnableSemiHosting() *queso.Option {
	return queso.NewOption("semihosting", "")
}

// SemiHostingTarget defines where the semihosting calls will be addressed.
type SemiHostingTarget string

const (
	// SemiHostingTargetAuto means SemiHostingTargetGDB is used during debug
	// sessions and SemiHostingTargetNative otherwise.
	SemiHostingTargetAuto SemiHostingTarget = "auto"

	// SemiHostingTargetGDB represents the GDB target.
	SemiHostingTargetGDB SemiHostingTarget = "gdb"

	// SemiHostingTargetNative represents the QEMU target.
	SemiHostingTargetNative SemiHostingTarget = "native"
)

// SemiHostingConfig enables and configures semihosting (ARM, M68K, Xtensa, MIPS,
// Nios II, RISC-V only).
//
// Note that this allows guest direct access to the host filesystem, so should
// only be used with a trusted guest OS.
//
// On ARM this implements the standard semihosting API, version 2.0. On M68K, this
// implements the "ColdFire GDB" interface used by libgloss. On Xtensa, semihosting
// provides basic file IO calls, such as open/read/write/seek/select.
// Tensilica bare-metal libc for ISS and linux platform "sim" use this interface.
// On RISC-V, this implements the standard semihosting API, version 0.2.
//
// See https://www.qemu.org/docs/master/about/emulation.html#semihosting for
// more details.
type SemiHostingConfig struct {
	properties []*queso.Property
}

// NewSemiHostingConfig returns a new [SemiHostingConfig].
//
//	qemu-system-* -semihosting-config
func NewSemiHostingConfig() *SemiHostingConfig {
	return &SemiHostingConfig{
		properties: make([]*queso.Property, 0),
	}
}

// Option returns the invoked option that gets converted to an argument when
// passed to QEMU.
func (c *SemiHostingConfig) Option() *queso.Option {
	return queso.NewOption("semihosting-config", "", c.properties...)
}

// SetProperty is used to add arbitrary properties to the [SemiHostingConfig].
func (c *SemiHostingConfig) SetProperty(key string, value interface{}) *SemiHostingConfig {
	c.properties = append(c.properties, queso.NewProperty(key, value))
	return c
}

// SetCharacterDeviceOutput send the output to a character device backend output
// for [SemiHostingTargetNative] or [SemiHostingTargetAuto] output when not
// in [SemiHostingTargetGDB].
//
//	qemu-system-* -semihosting-config chardev=chardev
func (c *SemiHostingConfig) SetCharacterDeviceOutput(chardev string) *SemiHostingConfig {
	c.properties = append(c.properties, queso.NewProperty("chardev", chardev))
	return c
}

// SetTarget defines where the semihosting calls will be addressed, to QEMU
// ([SemiHostingTargetNative]) or to GDB ([SemiHostingTargetGDB]). The default
// is [SemiHostingTargetAuto], which means [SemiHostingTargetGDB] during debug
// sessions and [SemiHostingTargetNative] otherwise.
//
//	qemu-system-* -semihosting-config target=native|gdb|auto
func (c *SemiHostingConfig) SetTarget(target SemiHostingTarget) *SemiHostingConfig {
	c.properties = append(c.properties, queso.NewProperty("target", string(target)))
	return c
}

// ToggleEnabled toggles whether the sandbox is enabled or disabled.
//
//	qemu-system-* -semihosting-config enable=on|off
func (c *SemiHostingConfig) ToggleEnabled(enabled bool) *SemiHostingConfig {
	c.properties = append(c.properties, queso.NewProperty("enabled", enabled))
	return c
}

// ToggleUserspace enables or disables code running in guest userspace to access
// the semihosting interface. The default is that only privileged guest code can
// make semihosting calls. Note that setting this to true should only be used if
// all guest code is trusted (for example, in bare-metal test case code).
//
//	qemu-system-* -semihosting-config userspace=on|off
func (c *SemiHostingConfig) ToggleUserspace(enabled bool) *SemiHostingConfig {
	c.properties = append(c.properties, queso.NewProperty("userspace", enabled))
	return c
}
