// Package netdev is used to create and manage network device backends for use
// with QEMU. See https://www.qemu.org/docs/master/system/invocation.html#hxtool-5
// for more details.
package netdev

import "github.com/mikerourke/queso"

// Backend represents a generic network backend.
type Backend struct {
	*queso.Entity
}

// New returns a new instance of [Backend]. backendType is the type of backend
// to create.
//
//	qemu-system-* -netdev <backendType>
func New(backendType string) *Backend {
	return &Backend{queso.NewEntity("netdev", backendType)}
}
