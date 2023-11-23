package debug

import (
	"github.com/mikerourke/queso/internal/cli"
	"github.com/mikerourke/queso/qemu/cli"
)

// Sandbox enables Seccomp mode 2 system call filter.
type Sandbox struct {
	Filter     bool
	properties []*cli.Property
}

// NewSandbox returns a new instance of Sandbox. Setting filter to true will
// enable syscall filtering, and false will disable it.
func NewSandbox(filter bool) *Sandbox {
	return &Sandbox{
		Filter:     filter,
		properties: make([]*cli.Property, 0),
	}
}

func (s *Sandbox) option() *cli.Option {
	return cli.NewOption("sandbox", cli.BoolPropertyToState(s.Filter), s.properties...)
}

// ToggleObsoleteSystemCalls enables or disables obsolete system calls.
func (s *Sandbox) ToggleObsoleteSystemCalls(enabled bool) *Sandbox {
	s.properties = append(s.properties, cli.NewProperty("obsolete", enabled))
	return s
}

// ToggleElevatedPrivileges enables or disables set*uid|gid system calls.
func (s *Sandbox) ToggleElevatedPrivileges(enabled bool) *Sandbox {
	// Since a value of "on" _disables_ elevated privileges, we want to negate
	// whatever value was passed in to indicate that elevated privileges are
	// enabled.
	s.properties = append(s.properties, cli.NewProperty("elevateprivileges", !enabled))
	return s
}

// ToggleSpawning enables or disables *fork and execve.
func (s *Sandbox) ToggleSpawning(enabled bool) *Sandbox {
	// Since a value of "on" _disables_ spawning, we want to negate
	// whatever value was passed in to indicate that spawning is enabled.
	s.properties = append(s.properties, cli.NewProperty("spawn", !enabled))
	return s
}

// ToggleResourceControl enables or disables process affinity and schedular priority.
func (s *Sandbox) ToggleResourceControl(enabled bool) *Sandbox {
	// Since a value of "on" _disables_ resource control, we want to negate
	// whatever value was passed in to indicate that resource control is enabled.
	s.properties = append(s.properties, cli.NewProperty("spawn", !enabled))
	return s
}
