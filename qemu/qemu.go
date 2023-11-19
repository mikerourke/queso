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
	}
}

// SetOptions sets the options to use for invoking QEMU.
func (q *QEMU) SetOptions(options ...*queso.Option) {
	args := make([]string, 0)

	for _, option := range options {
		args = append(args, option.Args()...)
	}

	q.args = args
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
