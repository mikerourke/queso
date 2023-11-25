package chardev

import "github.com/mikerourke/queso"

// SpicePortBackend connects to a Spice port, allowing a Spice client to handle
// the traffic identified by a name (preferably a fqdn).
type SpicePortBackend struct {
	*Backend
}

// NewSpicePortBackend returns a new instance of [SpicePortBackend].
// id is the unique ID, which can be any string up to 127 characters long.
// debugLevel is the debug level. name is the name of the Spice port to
// connect to.
//
//	qemu-system-* -chardev spiceport,id=id,debug=debug,name=name
func NewSpicePortBackend(id string, debugLevel string, name string) *SpicePortBackend {
	backend := New("spiceport", id)
	backend.properties = append(backend.properties,
		queso.NewProperty("debug", debugLevel),
		queso.NewProperty("name", name))

	return &SpicePortBackend{backend}
}

// SpiceVMCBackend connects to a Spice virtual machine channel, such as vdiport.
type SpiceVMCBackend struct {
	*Backend
}

// NewSpiceVMCBackend returns a new instance of [SpiceVMCBackend].
// id is the unique ID, which can be any string up to 127 characters long.
// debugLevel is the debug level. name is the name of the Spice channel to
// connect to.
//
//	qemu-system-* -chardev spiceport,id=id,debug=debug,name=name
func NewSpiceVMCBackend(id string, debugLevel string, name string) *SpiceVMCBackend {
	backend := New("spicevmc", id)
	backend.properties = append(backend.properties,
		queso.NewProperty("debug", debugLevel),
		queso.NewProperty("name", name))

	return &SpiceVMCBackend{backend}
}
