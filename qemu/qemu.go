package qemu

import (
	"os"
	"os/exec"

	"github.com/mikerourke/queso"
)

// QEMU represents an instance of the QEMU process.
type QEMU struct {
	exePath string
	args    []string
	cmd     *exec.Cmd
}

// New returns a new instance of QEMU. The path parameter represents the path
// to the QEMU executable. This can be a string if QEMU is in the PATH, or an
// absolute/relative path.
func New(path string) *QEMU {
	return &QEMU{
		exePath: path,
		args:    []string{},
	}
}

// With sets the options to use for invoking QEMU.
func (q *QEMU) With(options ...*queso.Option) *QEMU {
	args := make([]string, 0)

	for _, option := range options {
		args = append(args, option.Args()...)
	}

	q.args = args

	return q
}

// Usable represents an item that can be passed to the Use method.
type Usable interface {
	option() *queso.Option
}

// Use adds items as args to the QEMU command. It differs from the With method
// in that it accepts items that are defined with additional properties (as
// opposed to just using options).
func (q *QEMU) Use(usables ...Usable) *QEMU {
	for _, usable := range usables {
		usableArgs := usable.option().Args()
		q.args = append(q.args, usableArgs...)
	}

	return q
}

// Args returns a slice of the args that will be passed to QEMU. This is
// useful for debugging purposes.
func (q *QEMU) Args() []string {
	return q.args
}

// Cmd returns the exec.Cmd instance for QEMU.
func (q *QEMU) Cmd() *exec.Cmd {
	q.cmd = exec.Command(q.exePath, q.args...)

	return q.cmd
}

// Run starts the QEMU executable.
func (q *QEMU) Run() error {
	q.cmd.Stdout = os.Stdout
	q.cmd.Stderr = os.Stderr

	return q.cmd.Run()
}
