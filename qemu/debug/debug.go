// Package debug contains debug/expert options.
// See https://qemu.readthedocs.io/en/latest/system/invocation.html#hxtool-9 for more details.
package debug

// TODO: Add -rtc handling.

// TODO: Add -icount handling.

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/mikerourke/queso"
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
func AcceptGDBConnectionOnDevice(device string) *queso.Option {
	return queso.NewOption("gdb", device)
}

// Chroot chroots to the specified directory immediately before starting guest
// execution. Especially useful in combination with [RunAs].
//
//	qemu-system-* -chroot dir
func Chroot(dir string) *queso.Option {
	return queso.NewOption("chroot", dir)
}

// Daemonize daemonizes the QEMU process after initialization. QEMU will not
// detach from standard IO until it is ready to receive connections on any of
// its devices. This option is a useful way for external programs to launch QEMU
// without having to cope with initialization race conditions.
//
//	qemu-system-* -daemonize
func Daemonize() *queso.Option {
	return queso.NewOption("daemonize", "")
}

// DumpVMStateToFile dumps JSON-encoded VM state information for current machine
// type the specified file.
//
//	qemu-system-* -dump-vmstate file
func DumpVMStateToFile(file string) *queso.Option {
	return queso.NewOption("dump-vmstate", file)
}

// FilterDebugOutput filters debug output to that relevant to a range of target
// addresses. The filter spec can be either start+size, start-size or start..end
// where start end and size are the addresses and sizes required.
//
//	qemu-system-* -dfilter addresses
func FilterDebugOutput(addresses ...string) *queso.Option {
	value := strings.Join(addresses, ",")

	return queso.NewOption("dfilter", value)
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
// parameter is [RedirectSourceParallel] or [RedirectSourceSerial].
//
//	qemu-system-* -<source> device
func HostRedirect(source RedirectSource, device string) *queso.Option {
	return queso.NewOption(string(source), device)
}

// IsOverCommitHintForCPU specifies whether to run QEMU with hints about
// host resource overcommit for CPUs. When used, host estimates of CPU cycle and
// power utilization will be incorrect, not taking into account guest idle time.
//
//	qemu-system-* -overcommit cpu-pm=on|off
func IsOverCommitHintForCPU(enabled bool) *queso.Option {
	return queso.NewOption("overcommit", "", queso.NewProperty("cpu-pm", enabled))
}

// IsOverCommitHintForMemory specifies whether to run QEMU with hints about
// host resource overcommit for guest memory. This works when host memory is not
// over-committed and reduces the worst-case latency for guest.
//
//	qemu-system-* -overcommit mem-lock=on|off
func IsOverCommitHintForMemory(enabled bool) *queso.Option {
	return queso.NewOption("overcommit", "", queso.NewProperty("mem-lock", enabled))
}

// JITDump generate a dump file for Linux perf tools that maps basic blocks
// to symbol names, line numbers and JITted code.
//
//	qemu-system-* -jitdump
func JITDump() *queso.Option {
	return queso.NewOption("jitdump", "")
}

// LoadVM starts right away with a saved state (`loadvm` in monitor).
//
//	qemu-system-* -loadvm
func LoadVM(file string) *queso.Option {
	return queso.NewOption("loadvm", file)
}

// NoDefaults doesn't create default devices. Normally, QEMU sets the default devices
// like serial port, parallel port, virtual console, monitor device, VGA adapter,
// floppy and CD-ROM drive and others. This option will disable all those default devices.
//
//	qemu-system-* -nodefaults
func NoDefaults() *queso.Option {
	return queso.NewOption("nodefaults", "")
}

// NoReboot exits instead of rebooting.
//
//	qemu-system-* -no-reboot
func NoReboot() *queso.Option {
	return queso.NewOption("no-reboot", "")
}

// NoShutdown doesn't exit QEMU on guest shutdown, but instead only stops the
// emulation. This allows for instance switching to monitor to commit changes to
// the disk image.
//
//	qemu-system-* -no-shutdown
func NoShutdown() *queso.Option {
	return queso.NewOption("no-shutdown", "")
}

// NoUserConfiguration makes QEMU not load any of the user-provided config
// files on sysconfdir.
//
//	qemu-system-* -no-user-config
func NoUserConfiguration() *queso.Option {
	return queso.NewOption("no-user-config", "")
}

// OldParamMode uses old param mode (ARM only).
//
//	qemu-system-* -old-param
func OldParamMode() *queso.Option {
	return queso.NewOption("old-param", "")
}

// OnlyAllowMigratableDevices only allows migratable devices. Devices will not be
// allowed to enter an un-migratable state.
//
//	qemu-system-* -only-migratable
func OnlyAllowMigratableDevices() *queso.Option {
	return queso.NewOption("only-migratable", "")
}

// OpenGDBOnTCPPort opens a gdbserver on TCP port 1234.
//
//	qemu-system-* -s
func OpenGDBOnTCPPort() *queso.Option {
	return queso.NewOption("s", "")
}

// OptionROMFile loads the contents of file as an option ROM. This option
// is useful to load things like EtherBoot.
//
//	qemu-system-* -option-rom file
func OptionROMFile(file string) *queso.Option {
	return queso.NewOption("option-rom", file)
}

// OutputToLogFile outputs log to the specified file instead of to stderr.
//
//	qemu-system-* -D file
func OutputToLogFile(file string) *queso.Option {
	return queso.NewOption("D", file)
}

// PerfMap generates a map file for Linux perf tools that will allow basic
// profiling information to be broken down into basic blocks.
//
//	qemu-system-* -perfmap
func PerfMap() *queso.Option {
	return queso.NewOption("perfmap", "")
}

// PIDFile stores the QEMU process PID in file. It is useful if you launch
// QEMU from a script.
//
//	qemu-system-* -pidfile file
func PIDFile(file string) *queso.Option {
	return queso.NewOption("pidfile", file)
}

// ReadConfigurationFile reads device configuration from file. This approach is
// useful when you want to spawn a QEMU process with many command line options,
// but you don't want to exceed the command line character limit.
//
//	qemu-system-* -readconfig file
func ReadConfigurationFile(file string) *queso.Option {
	return queso.NewOption("readconfig", file)
}

// RunAs drops root privileges, switching to the specified user immediately before
// starting guest execution.
//
//	qemu-system-* -runas user
func RunAs(user string) *queso.Option {
	return queso.NewOption("runas", user)
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
//
//	qemu-system-* -run-with,async-teardown=on|off,chroot=dir
func RunWith(options RunWithOptions) *queso.Option {
	properties := []*queso.Property{
		queso.NewProperty("async-teardown", options.AsyncTeardown),
		queso.NewProperty("chroot", options.ChrootDir),
	}

	return queso.NewOption("run-with", "", properties...)
}

// SeedWith forces the guest to use a deterministic pseudo-random number generator,
// seeded with the specified seed. This does not affect crypto routines within
// the host.
//
//	qemu-system-* -seed seed
func SeedWith(seed int) *queso.Option {
	return queso.NewOption("seed", strconv.Itoa(seed))
}

// SingleStepMode runs the emulation in single step mode.
//
//	qemu-system-* -singlestep
func SingleStepMode() *queso.Option {
	return queso.NewOption("singlestep", "")
}

// SkipCPUStartAtStartup does not start CPU at startup (you must type "c" in the
// monitor).
//
//	qemu-system-* -S
func SkipCPUStartAtStartup() *queso.Option {
	return queso.NewOption("S", "")
}

// ErrorMessageFormat represent the options for error messages.
type ErrorMessageFormat struct {
	// TimeStamp will prefix messages with a timestamp.
	TimeStamp bool

	// GuestName will prefix messages with guest name but only if qemu.Name
	// guest option is set (see standard.go). Otherwise, the option is ignored.
	GuestName bool
}

// WithErrorMessageFormat is used to control error message format.
//
//	qemu-system-* -msg timestamp=timestamp,guestname=guestname
func WithErrorMessageFormat(options ErrorMessageFormat) *queso.Option {
	properties := []*queso.Property{
		queso.NewProperty("timestamp", options.TimeStamp),
		queso.NewProperty("guest-name", options.GuestName),
	}

	return queso.NewOption("msg", "", properties...)
}

// WithEscapeCharacter changes the escape character used for switching to the
// monitor when using monitor and serial sharing. The default is 0x01 when using
// the display.WithNoGraphic option. 0x01 is equal to pressing Control-a. You
// can select a different character from the ascii control keys where 1 through
// 26 map to Control-a through Control-z. For instance, you could use either of
// the following to change the escape character to Control-t.
//
//	qemu.With(qemu.WithEscapeCharacter(0x14), qemu.WithEscapeCharacter(20))
//
//	qemu-system-* -echr asciivalue
func WithEscapeCharacter(asciiValue int) *queso.Option {
	return queso.NewOption("echr", strconv.Itoa(asciiValue))
}

// WithKVMSupport enables KVM full virtualization support. This option is only available
// if KVM support is enabled when compiling.
//
//	qemu-system-* -enable-kvm
func WithKVMSupport() *queso.Option {
	return queso.NewOption("enable-kvm", "")
}

// WithLoggingForItems enables logging of specified items.
//
//	qemu-system-* -d items
func WithLoggingForItems(items ...string) *queso.Option {
	return queso.NewOption("d", strings.Join(items, ","))
}

// WithNVRAMVariable sets the OpenBIOS nvram variable with the specified name
// to the specified value (PPC, SPARC only).
//
//	qemu-system-* -prom-env <name>=<value>
func WithNVRAMVariable(name string, value string) *queso.Option {
	return queso.NewOption("prom-env", fmt.Sprintf("%s=%s", name, value))
}

// WithSyncProfiling enables synchronization profiling.
//
//	qemu-system-* -enable-sync-profile
func WithSyncProfiling() *queso.Option {
	return queso.NewOption("enable-sync-profile", "")
}

// XenAttach attaches to existing Xen domain. libxl will use this when starting
// QEMU (Xen only). Restrict set of available Xen operations to specified domain
// id (Xen only).
//
//	qemu-system-* -xen-attach
func XenAttach() *queso.Option {
	return queso.NewOption("xen-attach", "")
}

// XenGuestDomainID specifies Xen guest domain id (Xen only).
//
//	qemu-system-* -xen-domid id
func XenGuestDomainID(id string) *queso.Option {
	return queso.NewOption("xen-domid", id)
}
