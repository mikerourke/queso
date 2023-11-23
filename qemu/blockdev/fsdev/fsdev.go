// Package fsdev is used to define file system devices and virtual file system
// devices.
//
// Virtual file system devices are exposed to the guest using a virtio-9p-device
// (a.k.a. 9pfs), which essentially means that a certain directory
// on host is made directly accessible by guest as a pass-through file system by
// using the 9P network protocol for communication between host and guests,
// if desired even accessible, shared by several guests simultaneously.
package fsdev

// Type indicates the type of file system device or virtual file system device.
type Type string

const (
	// Local represents a file system device in which accesses to the filesystem
	// are done by QEMU.
	Local Type = "local"

	// Proxy represents a file system device in which accesses to the filesystem
	// are done by virtfs-proxy-helper(1).
	//
	// Deprecated: Use Local instead.
	Proxy Type = "proxy"

	// Synthetic represents a synthetic filesystem, only used by QTests.
	Synthetic Type = "synth"
)

// SecurityModel is used to specify the security model to be used for the export
// path in file system devices. Specifying this is mandatory only for
// [LocalFileSystemDevice] and [VirtualLocalFileSystemDevice]. Other drivers
// (like [ProxyFileSystemDevice]) don’t take security model as a parameter.
type SecurityModel string

const (
	// SecurityModelNone is the same as SecurityModelPassthrough except the server
	// won’t report failures if it fails to set file attributes like ownership.
	SecurityModelNone SecurityModel = "none"

	// SecurityModelPassthrough represents a model in which files are stored using
	// the same credentials as they are created on the guest.
	// This requires QEMU to run as root.
	SecurityModelPassthrough SecurityModel = "passthrough"

	// SecurityModelMappedXAttr represents a model in which some of the file attributes
	// like uid, gid, mode bits and link target are stored as file attributes.
	SecurityModelMappedXAttr SecurityModel = "mapped-xattr"

	// SecurityModelMappedFile represents a model in which some of the file attributes
	// are stored in the hidden .virtfs_metadata directory. Directories exported by
	// this security model cannot interact with other Unix tools
	SecurityModelMappedFile SecurityModel = "mapped-file"
)

// SocketInterfaceType is the field name used to identify whether the socket is
// using a socket file path or socket file descriptor.
type SocketInterfaceType string

const (
	// SocketInterfacePath enables proxy file system driver to use passed socket
	// file for communicating with virtfs-proxy-helper(1).
	//
	// When using VirtualLocalFileSystemDevice, a helper like libvirt will create
	// socket pair and pass one of the fds as SocketInterfaceDescriptor.
	SocketInterfacePath SocketInterfaceType = "socket"

	// SocketInterfaceDescriptor enables proxy file system driver to use passed
	// socket descriptor for communicating with virtfs-proxy-helper(1).
	//
	// When using LocalFileSystemDevice, a helper like libvirt will create socket
	// pair and pass one of the fds as SocketInterfaceDescriptor.
	SocketInterfaceDescriptor SocketInterfaceType = "sock_fd"
)
