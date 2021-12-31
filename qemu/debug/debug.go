// Package debug contains debug/expert options that can be passed to QEMU.
// See https://qemu.readthedocs.io/en/latest/system/invocation.html#hxtool-9 for
// more details.
package debug

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/mikerourke/queso"
)

// RedirectSource represents the source that gets redirected to a host device.
type RedirectSource string

const (
	// RedirectSourceDebugConsole is used to redirect the debug console to a host device
	// (same devices as the serial port). The debug console is an I/O port which is
	// typically port 0xe9; writing to that I/O port sends output to this device.
	// The default device is `vc` in graphical mode and `stdio` in non-graphical mode.
	RedirectSourceDebugConsole RedirectSource = "debugcon"

	// RedirectSourceMonitor redirects the monitor to a host device (same devices as
	// the serial port). The default device is `vc` in graphical mode and `stdio` in
	// non-graphical mode.
	RedirectSourceMonitor RedirectSource = "monitor"

	// RedirectSourceParallel redirects the virtual parallel port to a host device
	// (same devices as the serial port). On Linux hosts, `/dev/parportN` can be
	// used to use hardware devices connected on the corresponding host parallel port.
	// This option can be used several times to simulate up to 3 parallel ports.
	RedirectSourceParallel RedirectSource = "parallel"

	// RedirectSourceQMP is like RedirectSourceMonitor, but opens in "control" mode.
	RedirectSourceQMP RedirectSource = "qmp"

	// RedirectSourceQMPPretty is like RedirectSourceQMP, but uses pretty JSON
	// formatting.
	RedirectSourceQMPPretty RedirectSource = "qmp-pretty"

	// RedirectSourceSerial redirects the virtual serial port to a host character device.
	// The default device is `vc` in graphical mode and `stdio` in non-graphical mode.
	// This option can be used several times to simulate up to 4 serial ports.
	RedirectSourceSerial RedirectSource = "serial"
)

// HostRedirect redirects the specified source to the specified host device.
// Use "none" for the device parameter to disable all ports if the source
// parameter is RedirectSourceParallel or RedirectSourceSerial.
func HostRedirect(source RedirectSource, device string) *queso.Option {
	return queso.NewOption(string(source), device)
}

// UsePIDFile stores the QEMU process PID in file. It is useful if you launch QEMU
// from a script.
func UsePIDFile(file string) *queso.Option {
	return queso.NewOption("pidfile", file)
}

// UseSingleStepMode runs the emulation in single step mode.
func UseSingleStepMode() *queso.Option {
	return queso.NewOption("singlestep", "")
}

// SkipCPUStartAtStartup does not start CPU at startup (you must type "c" in the
// monitor).
func SkipCPUStartAtStartup() *queso.Option {
	return queso.NewOption("S", "")
}

// IsOverCommitHintForMemory specifies whether to run QEMU with hints about
// host resource overcommit for guest memory. This works when host memory is not
// over-committed and reduces the worst-case latency for guest.
func IsOverCommitHintForMemory(enabled bool) *queso.Option {
	return queso.NewOption("overcommit", "", queso.NewProperty("mem-lock", enabled))
}

// IsOverCommitHintForCPU specifies whether to run QEMU with hints about
// host resource overcommit for CPUs. When used, host estimates of CPU cycle and
// power utilization will be incorrect, not taking into account guest idle time.
func IsOverCommitHintForCPU(enabled bool) *queso.Option {
	return queso.NewOption("overcommit", "", queso.NewProperty("cpu-pm", enabled))
}

// AcceptGDBConnectionOnDevice accepts a GDB connection on the specified device.
// See https://www.qemu.org/docs/master/system/gdb.html#gdb-usage for more details
// on GDB usage. Note that this option does not pause QEMU execution – if you want
// QEMU to not start the guest until you connect with GDB and issue a continue
// command, you will need to also pass the SkipCPUStartAtStartup option to QEMU.
func AcceptGDBConnectionOnDevice(device string) *queso.Option {
	return queso.NewOption("gdb", device)
}

// OpenGDBOnTCPPort opens a gdbserver on TCP port 1234.
func OpenGDBOnTCPPort() *queso.Option {
	return queso.NewOption("s", "")
}

// EnableLoggingForItems enables logging of specified items.
func EnableLoggingForItems(items ...string) *queso.Option {
	value := strings.Join(items, ",")

	return queso.NewOption("d", value)
}

// OutputToLogFile outputs log to the specified file instead of to stderr.
func OutputToLogFile(file string) *queso.Option {
	return queso.NewOption("D", file)
}

// FilterDebugOutput filters debug output to that relevant to a range of target
// addresses. The filter spec can be either start+size, start-size or start..end
// where start end and size are the addresses and sizes required.
func FilterDebugOutput(addresses ...string) *queso.Option {
	value := strings.Join(addresses, ",")

	return queso.NewOption("dfilter", value)
}

// UseSeed forces the guest to use a deterministic pseudo-random number generator,
// seeded with the specified seed. This does not affect crypto routines within
// the host.
func UseSeed(seed int) *queso.Option {
	return queso.NewOption("seed", strconv.Itoa(seed))
}

// UseDataDirectory sets the directory for the BIOS, VGA BIOS and keymaps.
func UseDataDirectory(path string) *queso.Option {
	return queso.NewOption("L", path)
}

// UseBIOSFile sets the filename for the BIOS.
func UseBIOSFile(file string) *queso.Option {
	return queso.NewOption("bios", file)
}

// EnableKVM enables KVM full virtualization support. This option is only available
// if KVM support is enabled when compiling.
func EnableKVM() *queso.Option {
	return queso.NewOption("enable-kvm", "")
}

// UseXENGuestDomainID specifies XEN guest domain id (XEN only).
func UseXENGuestDomainID(id string) *queso.Option {
	return queso.NewOption("xen-domid", id)
}

// UseXENAttach attaches to existing xen domain. libxl will use this when starting
// QEMU (XEN only). Restrict set of available xen operations to specified domain
// id (XEN only).
func UseXENAttach() *queso.Option {
	return queso.NewOption("xen-attach", "")
}

// NoReboot exits instead of rebooting.
func NoReboot() *queso.Option {
	return queso.NewOption("no-reboot", "")
}

// NoShutdown doesn't exit QEMU on guest shutdown, but instead only stops the
// emulation. This allows for instance switching to monitor to commit changes to
// the disk image.
func NoShutdown() *queso.Option {
	return queso.NewOption("no-shutdown", "")
}

// Action modifies QEMU's default behavior for certain events. Instances of
// Action are passed to the WithAction option.
type Action struct {
	Event  string
	Action string
}

// NewAction returns a new instance of Action.
func NewAction(event string, action string) *Action {
	return &Action{
		Event:  event,
		Action: action,
	}
}

// WithAction serves to modify QEMU's default behavior when certain guest events occur.
// It provides a generic method for specifying the same behaviors that are modified
// by the NoReboot and NoShutdown options.
//
// Examples
//
//	qemu.WithAction(qemu.NewAction("panic", "none"))
//	qemu.WithAction(qemu.NewAction("reboot", "shutdown"), qemu.NewAction("shutdown", "pause"))
//	qemu.WithAction(qemu.NewAction("watchdog", "pause"))
func WithAction(actions ...*Action) *queso.Option {
	props := make([]*queso.Property, 0)

	for _, action := range actions {
		props = append(props, queso.NewProperty(action.Event, action.Action))
	}

	return queso.NewOption("action", "", props...)
}

// LoadVM starts right away with a saved state (`loadvm` in monitor).
func LoadVM(file string) *queso.Option {
	return queso.NewOption("loadvm", file)
}

// Daemonize daemonizes the QEMU process after initialization. QEMU will not
// detach from standard IO until it is ready to receive connections on any of
// its devices. This option is a useful way for external programs to launch QEMU
// without having to cope with initialization race conditions.
func Daemonize() *queso.Option {
	return queso.NewOption("daemonize", "")
}

// UseOptionROM loads the contents of file as an option ROM. This option is useful
// to load things like EtherBoot.
func UseOptionROM(file string) *queso.Option {
	return queso.NewOption("option-rom", file)
}

// UseForIncomingMigration prepares/accepts incoming migration on the specified
// receiver.
func UseForIncomingMigration(receiver string) *queso.Option {
	return queso.NewOption("incoming", receiver)
}

// OnlyAllowMigratableDevices only allows migratable devices. Devices will not be
// allowed to enter an un-migratable state.
func OnlyAllowMigratableDevices() *queso.Option {
	return queso.NewOption("only-migratable", "")
}

// NoDefaults doesn't create default devices. Normally, QEMU sets the default devices
// like serial port, parallel port, virtual console, monitor device, VGA adapter,
// floppy and CD-ROM drive and others. This option will disable all those default devices.
func NoDefaults() *queso.Option {
	return queso.NewOption("nodefaults", "")
}

// Chroot chroots to the specified directory immediately before starting guest execution.
// Especially useful in combination with RunAs.
func Chroot(dir string) *queso.Option {
	return queso.NewOption("chroot", dir)
}

// RunAs drops root privileges, switching to the specified user immediately before
// starting guest execution.
func RunAs(user string) *queso.Option {
	return queso.NewOption("runas", user)
}

// SetEnvVariable sets the OpenBIOS nvram variable with the specified name to the
// specified value (PPC, SPARC only).
func SetEnvVariable(name string, value string) *queso.Option {
	return queso.NewOption("prom-env", fmt.Sprintf("%s=%s", name, value))
}

// EnableSemiHosting enables semi-hosting mode (ARM, M68K, Xtensa, MIPS, Nios II,
// RISC-V only).
//
// Note that this allows guest direct access to the host filesystem, so should
// only be used with a trusted guest OS.
func EnableSemiHosting() *queso.Option {
	return queso.NewOption("semihosting", "")
}

// Argument represents a name/value pair that is passed into a Plugin or
// SemiHostingConfig.
type Argument struct {
	// Name is the name of the argument.
	Name string

	// Value is the associated value of the argument.
	Value string
}

// NewArgument returns a new instance of Argument.
func NewArgument(name string, value string) *Argument {
	return &Argument{
		Name:  name,
		Value: value,
	}
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
// On Arm this implements the standard semihosting API, version 2.0. On M68K, this
// implements the "ColdFire GDB" interface used by libgloss. On Xtensa, semihosting
// provides basic file IO calls, such as open/read/write/seek/select.
// Tensilica bare-metal libc for ISS and linux platform "sim" use this interface.
// On RISC-V, this implements the standard semihosting API, version 0.2.
//
// The target parameter defines where the semihosting calls will be addressed. See
// SemiHostingTarget for additional details. The optional chardev sends the output
// to a chardev backend output for native or auto output when not in
// SemiHostingTargetGDB. The arguments parameter represent input arguments.
func SemiHostingConfig(
	enabled bool,
	target SemiHostingTarget,
	chardev string,
	arguments ...*Argument,
) *queso.Option {
	props := []*queso.Property{
		queso.NewProperty("enabled", enabled),
		queso.NewProperty("target", target),
	}

	if chardev != "" {
		props = append(props, queso.NewProperty("chardev", chardev))
	}

	if arguments != nil {
		for _, argument := range arguments {
			props = append(props, queso.NewProperty(argument.Name, argument.Value))
		}
	}

	return queso.NewOption("semihosting-config", "", props...)
}

// OldParamMode uses old param mode (ARM only).
func OldParamMode() *queso.Option {
	return queso.NewOption("old-param", "")
}

// ReadConfigurationFile reads device configuration from file. This approach is
// useful when you want to spawn a QEMU process with many command line options,
// but you don’t want to exceed the command line character limit.
func ReadConfigurationFile(file string) *queso.Option {
	return queso.NewOption("readconfig", file)
}

// NoUserConfiguration makes QEMU not load any of the user-provided config files
// on sysconfdir.
func NoUserConfiguration() *queso.Option {
	return queso.NewOption("no-user-config", "")
}

// Trace traces events matching a pattern or from a file and optionally logs the
// output to a specified file. See the WithPattern, WithEventsFile, and WithOutputFile
// properties for more details.
func Trace(properties ...TraceProperty) *queso.Option {
	props := make([]*queso.Property, 0)

	for _, property := range properties {
		props = append(props, property.Property)
	}

	return queso.NewOption("trace", "", props...)
}

// TraceProperty represents a property that can be passed to the Trace option.
type TraceProperty struct {
	*queso.Property
}

// NewTraceProperty returns a new instance of TraceProperty.
func NewTraceProperty(key string, value interface{}) *TraceProperty {
	return &TraceProperty{
		Property: queso.NewProperty(key, value),
	}
}

// WithPattern immediately enables events matching pattern (either event name or
// a globbing pattern) for a Trace. This property is only available if QEMU
// has been compiled with the "simple", "log", or "ftrace" tracing backend.
func WithPattern(pattern string) *TraceProperty {
	return NewTraceProperty("enable", pattern)
}

// WithEventsFile immediately enable events listed in file for a Trace. The file
// must contain one event name (as listed in the trace-events-all file) per
// line; globbing patterns are accepted too. This property is only available if
// QEMU has been compiled with the "simple", "log", or "ftrace" tracing backend.
func WithEventsFile(file string) *TraceProperty {
	return NewTraceProperty("events", file)
}

// WithOutputFile logs output traces to file for a Trace. This property is only
// available if QEMU has been compiled with the "simple" tracing backend.
func WithOutputFile(file string) *TraceProperty {
	return NewTraceProperty("file", file)
}

// Plugin loads a plugin from the specified shared library file. Optional arguments
// can be passed to the plugin.
func Plugin(file string, arguments ...*Argument) *queso.Option {
	props := []*queso.Property{queso.NewProperty("file", file)}

	for _, argument := range arguments {
		props = append(props, queso.NewProperty(argument.Name, argument.Value))
	}

	return queso.NewOption("plugin", "", props...)
}

// EnableFIPS enables FIPS 140-2 compliance mode.
func EnableFIPS() *queso.Option {
	return queso.NewOption("enable-fips", "")
}

// ErrorMessageFormatOptions represent the options for error messages.
type ErrorMessageFormatOptions struct {
	// TimeStamp will prefix messages with a timestamp.
	TimeStamp bool

	// GuestName will prefix messages with guest name but only if qemu.Name
	// guest option is set (see standard.go). Otherwise, the option is ignored.
	GuestName bool
}

func UseErrorMessageFormat(options ErrorMessageFormatOptions) *queso.Option {
	props := []*queso.Property{
		queso.NewProperty("timestamp", options.TimeStamp),
		queso.NewProperty("guest-name", options.GuestName),
	}

	return queso.NewOption("msg", "", props...)
}

// DumpVMStateToFile dumps JSON-encoded VM state information for current machine
// type the specified file.
func DumpVMStateToFile(file string) *queso.Option {
	return queso.NewOption("dump-vmstate", file)
}

// EnableSyncProfiling enables synchronization profiling.
func EnableSyncProfiling() *queso.Option {
	return queso.NewOption("enable-sync-profile", "")
}
