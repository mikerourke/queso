package qemu

import "github.com/mikerourke/queso"

// AddFileDescriptor adds a file descriptor to a fd set. The fd parameter defines
// the file descriptor of which a duplicate is added to fd set. The file
// descriptor cannot be stdin, stdout, or stderr. The set parameter defines the
// ID of the fd set to add the file descriptor to. The opaque parameter defines
// a free-form string that can be used to describe fd, and can be set to an
// empty string to omit.
func AddFileDescriptor(fd int, set int, opaque string) *queso.Option {
	props := []*queso.Property{
		{"fd", fd},
		{"set", set},
	}

	if opaque != "" {
		props = append(props, queso.NewProperty("opaque", opaque))
	}

	return queso.NewOption("add-fd", "", props...)
}
