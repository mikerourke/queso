package netdev

// VHostVDPABackend represents a device that uses a datapath which complies with
// the virtio specifications with a vendor specific control path. vDPA devices
// can be both physically located on the hardware or emulated by software.
type VHostVDPABackend struct {
	*Backend
}

// NewVHostVDPABackend returns a new instance of [VHostVDPABackend].
//
//	qemu-system-* -netdev vhost-vdpa
func NewVHostVDPABackend() *VHostVDPABackend {
	return &VHostVDPABackend{New("vhost-vdpa")}
}

// SetVHostDevicePath specifies the path to the VHost device.
//
//	qemu-system-* -netdev vhost-vdpa,vhostdev=path
func (b *VHostVDPABackend) SetVHostDevicePath(path string) *VHostVDPABackend {
	b.SetProperty("vhostdev", path)
	return b
}

// SetVHostFD sets the VHost file descriptor.
//
//	qemu-system-* -netdev vhost-vdpa,vhostfd=fd
func (b *VHostVDPABackend) SetVHostFD(fd int) *VHostVDPABackend {
	b.SetProperty("vhostfd", fd)
	return b
}
