package qemu

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/mikerourke/queso"
)

// This file contains debug/expert options.
// See https://qemu.readthedocs.io/en/latest/system/invocation.html#hxtool-9 for
// more details.

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

// PIDFile stores the QEMU process PID in file. It is useful if you launch QEMU
// from a script.
func PIDFile(file string) *queso.Option {
	return queso.NewOption("pidfile", file)
}

// SingleStepMode runs the emulation in single step mode.
func SingleStepMode() *queso.Option {
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

// SeedWith forces the guest to use a deterministic pseudo-random number generator,
// seeded with the specified seed. This does not affect crypto routines within
// the host.
func SeedWith(seed int) *queso.Option {
	return queso.NewOption("seed", strconv.Itoa(seed))
}

// BIOSFile sets the filename for the BIOS.
func BIOSFile(file string) *queso.Option {
	return queso.NewOption("bios", file)
}

// EnableKVM enables KVM full virtualization support. This option is only available
// if KVM support is enabled when compiling.
func EnableKVM() *queso.Option {
	return queso.NewOption("enable-kvm", "")
}

// XenGuestDomainID specifies Xen guest domain id (Xen only).
func XenGuestDomainID(id string) *queso.Option {
	return queso.NewOption("xen-domid", id)
}

// XenAttach attaches to existing Xen domain. libxl will use this when starting
// QEMU (Xen only). Restrict set of available Xen operations to specified domain
// id (Xen only).
func XenAttach() *queso.Option {
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
// Examples
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

func (a *Action) option() *queso.Option {
	return queso.NewOption("action", "", queso.NewProperty(a.Event, a.Action))
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

// OptionROMFile loads the contents of file as an option ROM. This option is useful
// to load things like EtherBoot.
func OptionROMFile(file string) *queso.Option {
	return queso.NewOption("option-rom", file)
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
func EscapeCharacter(asciiValue int) *queso.Option {
	return queso.NewOption("echr", strconv.Itoa(asciiValue))
}

// IncomingTCPOptions represent options passed to IncomingTCPPort.
type IncomingTCPOptions struct {
	Host string
	Port int
	To   int
	IPv4 bool
	IPv6 bool
}

// IncomingTCPPort prepares for incoming migration, listen on a given TCP port.
func IncomingTCPPort(options IncomingTCPOptions) *queso.Option {
	var flag string
	if options.Host != "" {
		flag = fmt.Sprintf("tcp:%s:%d", options.Host, options.Port)
	} else {
		flag = fmt.Sprintf("tcp:%d", options.Port)
	}

	properties := []*queso.Property{
		queso.NewProperty("to", options.To),
		queso.NewProperty("ipv4", options.IPv4),
		queso.NewProperty("ipv6", options.IPv6),
	}

	return queso.NewOption("incoming", flag, properties...)
}

// IncomingSocketPath prepares for incoming migration, listens on a given unix socket.
func IncomingSocketPath(path string) *queso.Option {
	return queso.NewOption("incoming", fmt.Sprintf("unix:%s", path))
}

// IncomingFileDescriptor accepts incoming migration from a given file descriptor.
func IncomingFileDescriptor(fd int) *queso.Option {
	return queso.NewOption("incoming", fmt.Sprintf("fd:%d", fd))
}

// IncomingFile accepts incoming migration from a given file starting at offset.
// offset allows the common size suffixes, or a 0x prefix, but not both.
func IncomingFile(file string, offset int) *queso.Option {
	return queso.NewOption("incoming", fmt.Sprintf("file:%s", file), queso.NewProperty("offset", offset))
}

// IncomingCommand accepts incoming migration as an output from specified external
// command.
func IncomingCommand(command string) *queso.Option {
	return queso.NewOption("incoming", fmt.Sprintf("exec:%s", command))
}

// IncomingDefer waits for the URI to be specified via migrate_incoming. The monitor
// can be used to change settings (such as migration parameters) prior to issuing
// the migrate_incoming to allow the migration to begin.
func IncomingDefer() *queso.Option {
	return queso.NewOption("incoming", "defer")
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

// NVRAMVariable sets the OpenBIOS nvram variable with the specified name to the
// specified value (PPC, SPARC only).
func NVRAMVariable(name string, value string) *queso.Option {
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
	properties := []*queso.Property{
		queso.NewProperty("enabled", enabled),
		queso.NewProperty("target", target),
	}

	if chardev != "" {
		properties = append(properties, queso.NewProperty("chardev", chardev))
	}

	if arguments != nil {
		for _, argument := range arguments {
			properties = append(properties,
				queso.NewProperty(argument.Name, argument.Value))
		}
	}

	return queso.NewOption("semihosting-config", "", properties...)
}

// OldParamMode uses old param mode (ARM only).
func OldParamMode() *queso.Option {
	return queso.NewOption("old-param", "")
}

// Sandbox enables Seccomp mode 2 system call filter.
type Sandbox struct {
	Filter     bool
	properties []*queso.Property
}

// NewSandbox returns a new instance of Sandbox. Setting "filter" to true will
// enable syscall filtering, and false will disable it.
func NewSandbox(filter bool) *Sandbox {
	return &Sandbox{
		Filter:     filter,
		properties: make([]*queso.Property, 0),
	}
}

func (s *Sandbox) option() *queso.Option {
	return queso.NewOption("sandbox", queso.BoolPropertyToStatus(s.Filter), s.properties...)
}

// ToggleObsoleteSystemCalls enables or disables obsolete system calls.
func (s *Sandbox) ToggleObsoleteSystemCalls(enabled bool) *Sandbox {
	s.properties = append(s.properties, queso.NewProperty("obsolete", enabled))
	return s
}

// ToggleElevatedPrivileges enables or disables set*uid|gid system calls.
func (s *Sandbox) ToggleElevatedPrivileges(enabled bool) *Sandbox {
	// Since a value of "on" _disables_ elevated privileges, we want to negate
	// whatever value was passed in to indicate that elevated privileges are
	// enabled.
	s.properties = append(s.properties, queso.NewProperty("elevateprivileges", !enabled))
	return s
}

// ToggleSpawning enables or disables *fork and execve.
func (s *Sandbox) ToggleSpawning(enabled bool) *Sandbox {
	// Since a value of "on" _disables_ spawning, we want to negate
	// whatever value was passed in to indicate that spawning is enabled.
	s.properties = append(s.properties, queso.NewProperty("spawn", !enabled))
	return s
}

// ToggleResourceControl enables or disables process affinity and schedular priority.
func (s *Sandbox) ToggleResourceControl(enabled bool) *Sandbox {
	// Since a value of "on" _disables_ resource control, we want to negate
	// whatever value was passed in to indicate that resource control is enabled.
	s.properties = append(s.properties, queso.NewProperty("spawn", !enabled))
	return s
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
// output to a specified file.
type Trace struct {
	properties []*queso.Property
}

func NewTrace() *Trace {
	return &Trace{
		properties: make([]*queso.Property, 0),
	}
}

func (t *Trace) option() *queso.Option {
	return queso.NewOption("trace", "", t.properties...)
}

// MatchPattern immediately enables events matching pattern (either event name or
// a globbing pattern) for a Trace. This property is only available if QEMU
// has been compiled with the "simple", "log", or "ftrace" tracing backend.
func (t *Trace) MatchPattern(pattern string) *Trace {
	t.properties = append(t.properties, queso.NewProperty("enable", pattern))
	return t
}

// EnableEventsInFile immediately enable events listed in file for a Trace. The file
// must contain one event name (as listed in the trace-events-all file) per
// line; globbing patterns are accepted too. This property is only available if
// QEMU has been compiled with the "simple", "log", or "ftrace" tracing backend.
func (t *Trace) EnableEventsInFile(file string) *Trace {
	t.properties = append(t.properties, queso.NewProperty("events", file))
	return t
}

// SetOutputFile logs output traces to file for a Trace. This property is only
// available if QEMU has been compiled with the "simple" tracing backend.
func (t *Trace) SetOutputFile(file string) *Trace {
	t.properties = append(t.properties, queso.NewProperty("file", file))
	return t
}

// Plugin loads a plugin from the specified shared library file. Optional arguments
// can be passed to the plugin.
func Plugin(file string, arguments ...*Argument) *queso.Option {
	properties := []*queso.Property{
		queso.NewProperty("file", file),
	}

	for _, argument := range arguments {
		properties = append(properties, queso.NewProperty(argument.Name, argument.Value))
	}

	return queso.NewOption("plugin", "", properties...)
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
	//with RunAs.
	ChrootDir string
}

// RunWith sets QEMU process lifecycle options.
func RunWith(options RunWithOptions) *queso.Option {
	properties := []*queso.Property{
		queso.NewProperty("async-teardown", options.AsyncTeardown),
		queso.NewProperty("chroot", options.ChrootDir),
	}

	return queso.NewOption("run-with", "", properties...)
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
func ErrorMessageFormat(options ErrorMessageFormatOptions) *queso.Option {
	properties := []*queso.Property{
		queso.NewProperty("timestamp", options.TimeStamp),
		queso.NewProperty("guest-name", options.GuestName),
	}

	return queso.NewOption("msg", "", properties...)
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

// PerfMap generates a map file for Linux perf tools that will allow basic profiling
// information to be broken down into basic blocks.
func PerfMap() *queso.Option {
	return queso.NewOption("perfmap", "")
}

// JITDump generate a dump file for Linux perf tools that maps basic blocks to
// symbol names, line numbers and JITted code.
func JITDump() *queso.Option {
	return queso.NewOption("jitdump", "")
}
