package debug

import (
	"fmt"

	"github.com/mikerourke/queso"
)

// IncomingMigration represents an incoming migration.
type IncomingMigration struct {
	Type       string
	properties []*queso.Property
}

// NewIncomingMigration returns a new instance of [IncomingMigration].
// incomingType is the type of migration.
//
//	qemu-system-* -incoming <incomingType>
func NewIncomingMigration(incomingType string) *IncomingMigration {
	return &IncomingMigration{
		Type:       incomingType,
		properties: make([]*queso.Property, 0),
	}
}

// Option returns the invoked option that gets converted to an argument when
// passed to QEMU.
func (i *IncomingMigration) Option() *queso.Option {
	return queso.NewOption("incoming", i.Type, i.properties...)
}

// IncomingMigrationCommand accepts incoming migration as an output from
// specified external command.
type IncomingMigrationCommand struct {
	*IncomingMigration
}

// NewIncomingMigrationCommand returns a new instance of [IncomingMigrationCommand].
//
//	qemu-system-* -incoming exec:cmdline
func NewIncomingMigrationCommand(command string) *IncomingMigrationCommand {
	return &IncomingMigrationCommand{
		NewIncomingMigration(fmt.Sprintf("exec:%s", command)),
	}
}

// IncomingMigrationDefer waits for the URI to be specified via migrate_incoming.
// The monitor can be used to change settings (such as migration parameters) prior
// to issuing the migrate_incoming to allow the migration to begin.
type IncomingMigrationDefer struct {
	*IncomingMigration
}

// NewIncomingMigrationDefer returns a new instance of [IncomingMigrationDefer].
//
//	qemu-system-* -incoming defer
func NewIncomingMigrationDefer() *IncomingMigrationDefer {
	return &IncomingMigrationDefer{
		NewIncomingMigration("defer"),
	}
}

// IncomingMigrationFile accepts incoming migration from a given file starting
// at an optional offset.
type IncomingMigrationFile struct {
	*IncomingMigration
}

// NewIncomingMigrationFile returns a new instance of [IncomingMigrationFile].
// fd is the file descriptor for the file accepting incoming migration.
//
//	qemu-system-* -incoming file:filename
func NewIncomingMigrationFile(file string) *IncomingMigrationFile {
	return &IncomingMigrationFile{
		NewIncomingMigration(fmt.Sprintf("fd:%s", file)),
	}
}

// SetOffset sets the offset from the start of the file to accept incoming
// migrations. The value allows the common size suffixes, or a 0x prefix, but
// not both.
//
//	qemu-system-* -incoming file:filename,offset=offset
func (i *IncomingMigrationFile) SetOffset(offset string) *IncomingMigrationFile {
	i.properties = append(i.properties, queso.NewProperty("offset", offset))
	return i
}

// IncomingMigrationFileDescriptor accepts incoming migration from a file
// descriptor.
type IncomingMigrationFileDescriptor struct {
	*IncomingMigration
}

// NewIncomingMigrationFileDescriptor returns a new instance of [IncomingMigrationFileDescriptor].
// fd is the file descriptor for the file accepting incoming migration.
//
//	qemu-system-* -incoming fd:fd
func NewIncomingMigrationFileDescriptor(fd int) *IncomingMigrationFileDescriptor {
	return &IncomingMigrationFileDescriptor{
		NewIncomingMigration(fmt.Sprintf("fd:%d", fd)),
	}
}

// IncomingMigrationTCPPort prepares for incoming migration, listen on a given
// TCP port.
type IncomingMigrationTCPPort struct {
	*IncomingMigration
}

// NewIncomingTCPPort returns a new instance of [IncomingMigrationTCPPort].
// port is the TCP port number. host is the host address. Set host to "" to
// use the default.
//
//	qemu-system-* -incoming tcp:[host]:port
func NewIncomingTCPPort(port int, host string) *IncomingMigrationTCPPort {
	var incomingType string
	if host != "" {
		incomingType = fmt.Sprintf("tcp:%s:%d", host, port)
	} else {
		incomingType = fmt.Sprintf("tcp:%d", port)
	}

	return &IncomingMigrationTCPPort{
		NewIncomingMigration(incomingType),
	}
}

// SetToPort sets the maximum port to listen on.
//
//	qemu-system-* -incoming tcp:port,to=maxport
func (i *IncomingMigrationTCPPort) SetToPort(port int) *IncomingMigrationTCPPort {
	i.properties = append(i.properties, queso.NewProperty("to", port))
	return i
}

// ToggleIPv4 specifies if IPv4 may be used.
//
//	qemu-system-* -incoming tcp:port,ipv4=on|off
func (i *IncomingMigrationTCPPort) ToggleIPv4(enabled bool) *IncomingMigrationTCPPort {
	i.properties = append(i.properties, queso.NewProperty("ipv4", enabled))
	return i
}

// ToggleIPv6 specifies if IPv6 may be used.
//
//	qemu-system-* -incoming tcp:port,ipv6=on|off
func (i *IncomingMigrationTCPPort) ToggleIPv6(enabled bool) *IncomingMigrationTCPPort {
	i.properties = append(i.properties, queso.NewProperty("ipv6", enabled))
	return i
}

// IncomingMigrationSocket prepares for an incoming migration and listens on a
// Unix socket.
type IncomingMigrationSocket struct {
	*IncomingMigration
}

// NewIncomingMigrationSocket returns a new instance of [IncomingMigrationSocket].
// path is the path of the Unix socket.
//
//	qemu-system-* -incoming unix:socketpath
func NewIncomingMigrationSocket(path string) *IncomingMigrationSocket {
	return &IncomingMigrationSocket{
		NewIncomingMigration(fmt.Sprintf("unix:%s", path)),
	}
}
