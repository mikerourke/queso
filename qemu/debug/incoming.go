package debug

import (
	"fmt"
)

// IncomingTCPOptions represent options passed to IncomingTCPPort.
type IncomingTCPOptions struct {
	Host string
	Port int
	To   int
	IPv4 bool
	IPv6 bool
}

// IncomingTCPPort prepares for incoming migration, listen on a given TCP port.
func IncomingTCPPort(options IncomingTCPOptions) *cli.Option {
	var flag string
	if options.Host != "" {
		flag = fmt.Sprintf("tcp:%s:%d", options.Host, options.Port)
	} else {
		flag = fmt.Sprintf("tcp:%d", options.Port)
	}

	properties := []*cli.Property{
		cli.NewProperty("to", options.To),
		cli.NewProperty("ipv4", options.IPv4),
		cli.NewProperty("ipv6", options.IPv6),
	}

	return cli.NewOption("incoming", flag, properties...)
}

// IncomingSocketPath prepares for incoming migration, listens on a given unix socket.
func IncomingSocketPath(path string) *cli.Option {
	return cli.NewOption("incoming", fmt.Sprintf("unix:%s", path))
}

// IncomingFileDescriptor accepts incoming migration from a given file descriptor.
func IncomingFileDescriptor(fd int) *cli.Option {
	return cli.NewOption("incoming", fmt.Sprintf("fd:%d", fd))
}

// IncomingFile accepts incoming migration from a given file starting at offset.
// offset allows the common size suffixes, or a 0x prefix, but not both.
func IncomingFile(file string, offset int) *cli.Option {
	return cli.NewOption("incoming",
		fmt.Sprintf("file:%s", file),
		cli.NewProperty("offset", offset))
}

// IncomingCommand accepts incoming migration as an output from specified external
// command.
func IncomingCommand(command string) *cli.Option {
	return cli.NewOption("incoming", fmt.Sprintf("exec:%s", command))
}

// IncomingDefer waits for the URI to be specified via migrate_incoming. The monitor
// can be used to change settings (such as migration parameters) prior to issuing
// the migrate_incoming to allow the migration to begin.
func IncomingDefer() *cli.Option {
	return cli.NewOption("incoming", "defer")
}
