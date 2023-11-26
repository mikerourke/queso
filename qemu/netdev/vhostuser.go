package netdev

// VHostUserBackend uses a specifically defined protocol to pass vhost ioctl
// replacement messages to an application on the other end of the socket.
type VHostUserBackend struct {
	*Backend
}

// NewVHostUserBackend returns a new instance of [VHostUserBackend]. chardev
// is the ID fo the backing character device. The device should be backed by a
// Unix domain socket.
//
//	qemu-system-* -netdev vhost-user,chardev=chardev
func NewVHostUserBackend(chardev string) *VHostUserBackend {
	backend := New("vhost-user")

	backend.SetProperty("chardev", chardev)

	return &VHostUserBackend{backend}
}

// SetQueueCount specifies the number of queues to be created for multi-queue
// vhost-user.
//
//	qemu-system-* -netdev vhost-user,queues=count
func (b *VHostUserBackend) SetQueueCount(count int) *VHostUserBackend {
	b.SetProperty("queues", count)
	return b
}

// ToggleForceVHost enables or disables forcing the use of a specifically
// defined protocol to pass vhost ioctl replecmement messages to an application
// on the other end of a socket on non-MSIX guests.
//
//	qemu-system-* vhost-user,vhostforce=on|off
func (b *VHostUserBackend) ToggleForceVHost(enabled bool) *VHostUserBackend {
	b.SetProperty("vhostforce", enabled)
	return b
}
