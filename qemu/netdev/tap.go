package netdev

// TAPBackend represents a host TAP network backend.
type TAPBackend struct {
	*Backend
}

// NewTAPBackend returns a new instance of [TAPBackend]. id is a unique
// identifier for the backend.
//
//	qemu-system-* -netdev tap,id=id
func NewTAPBackend(id string) *TAPBackend {
	backend := New("tap")

	backend.SetProperty("id", id)

	return &TAPBackend{backend}
}

// SetBridge specifies the bridge device.
//
//	qemu-system-* -netdev tap,br=bridge
func (b *TAPBackend) SetBridge(bridge string) *TAPBackend {
	b.SetProperty("br", bridge)
	return b
}

// SetDownScript sets the network script file to de-configure the backend. The
// default network configure script is "/etc/qemu-ifdown".
//
//	qemu-system-* -netdev tap,downscript=file
func (b *TAPBackend) SetDownScript(file string) *TAPBackend {
	b.SetProperty("downscript", file)
	return b
}

// SetFileDescriptor specifies the handle of an already opened host TAP interface.
//
//	qemu-system-* -netdev tap,fd=fd
func (b *TAPBackend) SetFileDescriptor(fd int) *TAPBackend {
	b.SetProperty("fd", fd)
	return b
}

// SetHelper specifies an executable path to configure the TAP interface and
// attach it to the bridge.
//
//	qemu-system-* -netdev tap,helper=helper
func (b *TAPBackend) SetHelper(helper string) *TAPBackend {
	b.SetProperty("helper", helper)
	return b
}

// SetInterfaceName defines the interface name.
//
//	qemu-system-* -netdev tap,ifname=name
func (b *TAPBackend) SetInterfaceName(name string) *TAPBackend {
	b.SetProperty("ifname", name)
	return b
}

// SetUpScript sets the network script file to configure the backend. The default
// network configure script is "/etc/qemu-ifup".
//
//	qemu-system-* -netdev tap,script=file
func (b *TAPBackend) SetUpScript(file string) *TAPBackend {
	b.SetProperty("script", file)
	return b
}
