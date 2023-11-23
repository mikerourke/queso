// Package debug contains debug/expert options.
// See https://qemu.readthedocs.io/en/latest/system/invocation.html#hxtool-9 for more details.
package debug

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/mikerourke/queso/internal/cli"
)

// AcceptGDBConnectionOnDevice accepts a GDB connection on the specified device.
// See https://www.qemu.org/docs/master/system/gdb.html#gdb-usage for more details
// on GDB usage. Note that this option does not pause QEMU execution – if you want
// QEMU to not start the guest until you connect with GDB and issue a continue
// command, you will need to also pass the [SkipCPUStartAtStartup] option to QEMU.
//
// The most usual configuration is to listen on a local TCP socket.
//
//	qemu-system-* -gdb tcp::3771
func AcceptGDBConnectionOnDevice(device string) *cli.Option {
	return cli.NewOption("gdb", device)
}

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
func HostRedirect(source RedirectSource, device string) *cli.Option {
	return cli.NewOption(string(source), device)
}

// PIDFile stores the QEMU process PID in file. It is useful if you launch QEMU
// from a script.
func PIDFile(file string) *cli.Option {
	return cli.NewOption("pidfile", file)
}

// SingleStepMode runs the emulation in single step mode.
func SingleStepMode() *cli.Option {
	return cli.NewOption("singlestep", "")
}

// SkipCPUStartAtStartup does not start CPU at startup (you must type "c" in the
// monitor).
func SkipCPUStartAtStartup() *cli.Option {
	return cli.NewOption("S", "")
}

// IsOverCommitHintForMemory specifies whether to run QEMU with hints about
// host resource overcommit for guest memory. This works when host memory is not
// over-committed and reduces the worst-case latency for guest.
func IsOverCommitHintForMemory(enabled bool) *cli.Option {
	return cli.NewOption("overcommit", "", cli.NewProperty("mem-lock", enabled))
}

// IsOverCommitHintForCPU specifies whether to run QEMU with hints about
// host resource overcommit for CPUs. When used, host estimates of CPU cycle and
// power utilization will be incorrect, not taking into account guest idle time.
func IsOverCommitHintForCPU(enabled bool) *cli.Option {
	return cli.NewOption("overcommit", "", cli.NewProperty("cpu-pm", enabled))
}

// OpenGDBOnTCPPort opens a gdbserver on TCP port 1234.
func OpenGDBOnTCPPort() *cli.Option {
	return cli.NewOption("s", "")
}

// EnableLoggingForItems enables logging of specified items.
func EnableLoggingForItems(items ...string) *cli.Option {
	value := strings.Join(items, ",")

	return cli.NewOption("d", value)
}

// OutputToLogFile outputs log to the specified file instead of to stderr.
func OutputToLogFile(file string) *cli.Option {
	return cli.NewOption("D", file)
}

// FilterDebugOutput filters debug output to that relevant to a range of target
// addresses. The filter spec can be either start+size, start-size or start..end
// where start end and size are the addresses and sizes required.
func FilterDebugOutput(addresses ...string) *cli.Option {
	value := strings.Join(addresses, ",")

	return cli.NewOption("dfilter", value)
}

// SeedWith forces the guest to use a deterministic pseudo-random number generator,
// seeded with the specified seed. This does not affect crypto routines within
// the host.
func SeedWith(seed int) *cli.Option {
	return cli.NewOption("seed", strconv.Itoa(seed))
}

// EnableKVM enables KVM full virtualization support. This option is only available
// if KVM support is enabled when compiling.
func EnableKVM() *cli.Option {
	return cli.NewOption("enable-kvm", "")
}

// XenGuestDomainID specifies Xen guest domain id (Xen only).
func XenGuestDomainID(id string) *cli.Option {
	return cli.NewOption("xen-domid", id)
}

// XenAttach attaches to existing Xen domain. libxl will use this when starting
// QEMU (Xen only). Restrict set of available Xen operations to specified domain
// id (Xen only).
func XenAttach() *cli.Option {
	return cli.NewOption("xen-attach", "")
}

// NoReboot exits instead of rebooting.
func NoReboot() *cli.Option {
	return cli.NewOption("no-reboot", "")
}

// NoShutdown doesn't exit QEMU on guest shutdown, but instead only stops the
// emulation. This allows for instance switching to monitor to commit changes to
// the disk image.
func NoShutdown() *cli.Option {
	return cli.NewOption("no-shutdown", "")
}

// Action modifies QEMU's default behavior for certain events. It provides a generic
// method for specifying the same behaviors that are modified by the NoReboot and
// NoShutdown options.
type Action struct {
	Event  string
	Action string
}

// NewAction returns a new instance of Action. These are passed into the qemu.Use
// method.
//
// # Examples
//
//	qemu.Use(
//		qemu.NewAction("panic", "none"),
//		qemu.NewAction("reboot", "shutdown"),
//		qemu.NewAction("shutdown", "pause"),
//		qemu.NewAction("watchdog", "pause"))
func NewAction(event string, action string) *Action {
	return &Action{
		Event:  event,
		Action: action,
	}
}

func (a *Action) option() *cli.Option {
	return cli.NewOption("action", "", cli.NewProperty(a.Event, a.Action))
}

// LoadVM starts right away with a saved state (`loadvm` in monitor).
func LoadVM(file string) *cli.Option {
	return cli.NewOption("loadvm", file)
}

// Daemonize daemonizes the QEMU process after initialization. QEMU will not
// detach from standard IO until it is ready to receive connections on any of
// its devices. This option is a useful way for external programs to launch QEMU
// without having to cope with initialization race conditions.
func Daemonize() *cli.Option {
	return cli.NewOption("daemonize", "")
}

// OptionROMFile loads the contents of file as an option ROM. This option is useful
// to load things like EtherBoot.
func OptionROMFile(file string) *cli.Option {
	return cli.NewOption("option-rom", file)
}

// TODO: Add -rtc handling.

// TODO: Add -icount handling.

// EscapeCharacter changes the escape character used for switching to the monitor
// when using monitor and serial sharing. The default is 0x01 when using the
// -nographic option. 0x01 is equal to pressing Control-a. You can select a
// different character from the ascii control keys where 1 through 26 map to
// Control-a through Control-z. For instance, you could use either of the
// following to change the escape character to Control-t.
//
//	qemu.With(qemu.EscapeCharacter(0x14), qemu.EscapeCharacter(20))
func EscapeCharacter(asciiValue int) *cli.Option {
	return cli.NewOption("echr", strconv.Itoa(asciiValue))
}

// OnlyAllowMigratableDevices only allows migratable devices. Devices will not be
// allowed to enter an un-migratable state.
func OnlyAllowMigratableDevices() *cli.Option {
	return cli.NewOption("only-migratable", "")
}

// NoDefaults doesn't create default devices. Normally, QEMU sets the default devices
// like serial port, parallel port, virtual console, monitor device, VGA adapter,
// floppy and CD-ROM drive and others. This option will disable all those default devices.
func NoDefaults() *cli.Option {
	return cli.NewOption("nodefaults", "")
}

// Chroot chroots to the specified directory immediately before starting guest execution.
// Especially useful in combination with RunAs.
func Chroot(dir string) *cli.Option {
	return cli.NewOption("chroot", dir)
}

// RunAs drops root privileges, switching to the specified user immediately before
// starting guest execution.
func RunAs(user string) *cli.Option {
	return cli.NewOption("runas", user)
}

// NVRAMVariable sets the OpenBIOS nvram variable with the specified name to the
// specified value (PPC, SPARC only).
func NVRAMVariable(name string, value string) *cli.Option {
	return cli.NewOption("prom-env", fmt.Sprintf("%s=%s", name, value))
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

// OldParamMode uses old param mode (ARM only).
func OldParamMode() *cli.Option {
	return cli.NewOption("old-param", "")
}

// ReadConfigurationFile reads device configuration from file. This approach is
// useful when you want to spawn a QEMU process with many command line options,
// but you don’t want to exceed the command line character limit.
func ReadConfigurationFile(file string) *cli.Option {
	return cli.NewOption("readconfig", file)
}

// NoUserConfiguration makes QEMU not load any of the user-provided config files
// on sysconfdir.
func NoUserConfiguration() *cli.Option {
	return cli.NewOption("no-user-config", "")
}

// Plugin loads a plugin from the specified shared library file. Optional arguments
// can be passed to the plugin.
func Plugin(file string, arguments ...*Argument) *cli.Option {
	properties := []*cli.Property{
		cli.NewProperty("file", file),
	}

	for _, argument := range arguments {
		properties = append(properties, cli.NewProperty(argument.Name, argument.Value))
	}

	return cli.NewOption("plugin", "", properties...)
}

// RunWithOptions represent the options struct passed to RunWith.
type RunWithOptions struct {
	// AsyncTeardown enables asynchronous teardown when true. A new process called
	// “cleanup/<QEMU_PID>” will be created at startup sharing the address space
	// with the main QEMU process, using clone. It will wait for the main QEMU
	// process to terminate completely, and then exit. This allows QEMU to
	// terminate very quickly even if the guest was huge, leaving the teardown
	// of the address space to the cleanup process. Since the cleanup process
	// shares the same cgroups as the main QEMU process, accounting is performed
	// correctly.
	//
	// This only works if the cleanup process is not forcefully killed with
	// SIGKILL before the main QEMU process has terminated completely.
	AsyncTeardown bool

	// ChrootDir can be used for doing a chroot to the specified directory immediately
	// before starting the guest execution. This is especially useful in combination
	// with RunAs.
	ChrootDir string
}

// RunWith sets QEMU process lifecycle options.
func RunWith(options RunWithOptions) *cli.Option {
	properties := []*cli.Property{
		cli.NewProperty("async-teardown", options.AsyncTeardown),
		cli.NewProperty("chroot", options.ChrootDir),
	}

	return cli.NewOption("run-with", "", properties...)
}

// ErrorMessageFormatOptions represent the options for error messages.
type ErrorMessageFormatOptions struct {
	// TimeStamp will prefix messages with a timestamp.
	TimeStamp bool

	// GuestName will prefix messages with guest name but only if qemu.Name
	// guest option is set (see standard.go). Otherwise, the option is ignored.
	GuestName bool
}

// ErrorMessageFormat is used to control error message format.
func ErrorMessageFormat(options ErrorMessageFormatOptions) *cli.Option {
	properties := []*cli.Property{
		cli.NewProperty("timestamp", options.TimeStamp),
		cli.NewProperty("guest-name", options.GuestName),
	}

	return cli.NewOption("msg", "", properties...)
}

// DumpVMStateToFile dumps JSON-encoded VM state information for current machine
// type the specified file.
func DumpVMStateToFile(file string) *cli.Option {
	return cli.NewOption("dump-vmstate", file)
}

// EnableSyncProfiling enables synchronization profiling.
func EnableSyncProfiling() *cli.Option {
	return cli.NewOption("enable-sync-profile", "")
}

// PerfMap generates a map file for Linux perf tools that will allow basic profiling
// information to be broken down into basic blocks.
func PerfMap() *cli.Option {
	return cli.NewOption("perfmap", "")
}

// JITDump generate a dump file for Linux perf tools that maps basic blocks to
// symbol names, line numbers and JITted code.
func JITDump() *cli.Option {
	return cli.NewOption("jitdump", "")
}
