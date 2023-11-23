package debug

// EnableSemiHosting enables semi-hosting mode (ARM, M68K, Xtensa, MIPS, Nios II,
// RISC-V only).
//
// Note that this allows guest direct access to the host filesystem, so should
// only be used with a trusted guest OS.
func EnableSemiHosting() *cli.Option {
	return cli.NewOption("semihosting", "")
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
// The target parameter defines where the semihosting calls will be addressed. See
// [SemiHostingTarget] for additional details. The optional chardev sends the output
// to a chardev backend output for native or auto output when not in
// [SemiHostingTargetGDB]. The arguments parameter represent input arguments.
func SemiHostingConfig(
	enabled bool,
	target SemiHostingTarget,
	chardev string,
	arguments ...*Argument,
) *cli.Option {
	properties := []*cli.Property{
		cli.NewProperty("enabled", enabled),
		cli.NewProperty("target", target),
	}

	if chardev != "" {
		properties = append(properties, cli.NewProperty("chardev", chardev))
	}

	if arguments != nil {
		for _, argument := range arguments {
			properties = append(properties,
				cli.NewProperty(argument.Name, argument.Value))
		}
	}

	return cli.NewOption("semihosting-config", "", properties...)
}
