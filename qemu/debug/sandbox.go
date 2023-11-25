package debug

import "github.com/mikerourke/queso"

// Sandbox enables Seccomp mode 2 system call filter.
type Sandbox struct {
	Filter     bool
	properties []*queso.Property
}

// NewSandbox returns a new instance of Sandbox. Setting filter to true will
// enable syscall filtering, and false will disable it.
//
//	qemu-system-* -sandbox <filter>
func NewSandbox(filter bool) *Sandbox {
	return &Sandbox{
		Filter:     filter,
		properties: make([]*queso.Property, 0),
	}
}

func (s *Sandbox) option() *queso.Option {
	return queso.NewOption("sandbox",
		queso.BoolPropertyToStatus(s.Filter), s.properties...)
}

// SetProperty can be used to set arbitrary properties on the [Sandbox].
func (s *Sandbox) SetProperty(name string, value interface{}) *Sandbox {
	s.properties = append(s.properties, queso.NewProperty(name, value))
	return s
}

// ToggleObsoleteSystemCalls enables or disables obsolete system calls.
//
//	qemu-system-* -sandbox on,obsolete=on|off
func (s *Sandbox) ToggleObsoleteSystemCalls(enabled bool) *Sandbox {
	s.properties = append(s.properties, queso.NewProperty("obsolete", enabled))
	return s
}

// ToggleElevatedPrivileges enables or disables set*uid|gid system calls.
//
//	qemu-system-* -sandbox on,elevateprivileges=on|off
func (s *Sandbox) ToggleElevatedPrivileges(enabled bool) *Sandbox {
	// Since a value of "on" _disables_ elevated privileges, we want to negate
	// whatever value was passed in to indicate that elevated privileges are
	// enabled.
	s.properties = append(s.properties, queso.NewProperty("elevateprivileges", !enabled))
	return s
}

// ToggleSpawning enables or disables *fork and execve.
//
//	qemu-system-* -sandbox on,spawn=on|off
func (s *Sandbox) ToggleSpawning(enabled bool) *Sandbox {
	// Since a value of "on" _disables_ spawning, we want to negate
	// whatever value was passed in to indicate that spawning is enabled.
	s.properties = append(s.properties, queso.NewProperty("spawn", !enabled))
	return s
}

// ToggleResourceControl enables or disables process affinity and schedular priority.
//
//	qemu-system-* -sandbox on,resourcecontrol=on|off
func (s *Sandbox) ToggleResourceControl(enabled bool) *Sandbox {
	// Since a value of "on" _disables_ resource control, we want to negate
	// whatever value was passed in to indicate that resource control is enabled.
	s.properties = append(s.properties, queso.NewProperty("resourcecontrol", !enabled))
	return s
}
