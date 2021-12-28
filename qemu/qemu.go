package qemu

import (
	"fmt"

	"github.com/mikerourke/queso"
)

// QEMU represents an instance of the QEMU process.
type QEMU struct {
	exePath string
	args    []string
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

// Start starts the QEMU executable.
func (q *QEMU) Start() {
	fmt.Println(q.args)
}
