package qemu

// AddFileDescriptor adds a file descriptor to a fd set. The fd parameter defines
// the file descriptor of which a duplicate is added to fd set. The file
// descriptor cannot be stdin, stdout, or stderr. The set parameter defines the
// ID of the fd set to add the file descriptor to. The opaque parameter defines
// a free-form string that can be used to describe fd, and can be set to an
// empty string to omit.
func AddFileDescriptor(fd int, set int, opaque string) *cli.Option {
	properties := []*cli.Property{
		cli.NewProperty("fd", fd),
		cli.NewProperty("set", set),
	}

	if opaque != "" {
		properties = append(properties, cli.NewProperty("opaque", opaque))
	}

	return cli.NewOption("add-fd", "", properties...)
}
