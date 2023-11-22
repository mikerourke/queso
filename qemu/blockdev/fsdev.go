package blockdev

import (
	"fmt"

	"github.com/mikerourke/queso"
)

// FileSystemDevice defines a new filesystem device option.
func FileSystemDevice(
	deviceType string,
	id string,
	properties ...*FileSystemDeviceProperty,
) *queso.Option {
	props := []*queso.Property{queso.NewProperty("id", id)}

	for _, property := range properties {
		props = append(props, property.Property)
	}

	return queso.NewOption("fsdev", deviceType, props...)
}

// SecurityModel specifies the security model to be used for the export path.
type SecurityModel string

const (
	// SecurityModelNone is same as SecurityModelPassThrough except the sever won't
	// report failures if it fails to set file attributes like ownership.
	SecurityModelNone SecurityModel = "none"

	// SecurityModelPassThrough indicates that files are stored using the same credentials
	// as they are created on the guest. This requires QEMU to run as root.
	SecurityModelPassThrough SecurityModel = "passthrough"

	// SecurityModelMappedXAttr indicates that some file attributes like uid,
	// gid, mode bits and link target are stored as file attributes.
	SecurityModelMappedXAttr SecurityModel = "mapped-xattr"

	// SecurityModelMappedFile indicates that file attributes are stored in the hidden
	// `.virtfs_metadata` directory. Directories exported by this security model
	// cannot interact with other Unix tools.
	SecurityModelMappedFile SecurityModel = "mapped-file"
)

// LocalFileSystemDevice defines a filesystem device for which accesses to the
// filesystem are done by QEMU. The exportPath parameter represents the export path
// for the filesystem device. Files under this path will be available to the 9p
// client on the guest. The securityModel parameter specifies the security model
// to be used for the export path.
func LocalFileSystemDevice(
	id string,
	exportPath string,
	securityModel SecurityModel,
	properties ...*FileSystemDeviceProperty,
) *queso.Option {
	props := []*FileSystemDeviceProperty{
		NewFileSystemDeviceProperty("path", exportPath),
		NewFileSystemDeviceProperty("security_model", securityModel),
	}

	if properties != nil {
		props = append(props, properties...)
	}

	return FileSystemDevice("local", id, props...)
}

// SocketType is used to specify the socket type for a proxy device.
type SocketType string

const (
	// SocketTypeFile uses a socket file for communicating with virtfs-proxy-helper.
	SocketTypeFile SocketType = "socket"

	// SocketTypeDescriptor uses a socket file descriptor for communicating with
	// virtfs-proxy-helper.
	SocketTypeDescriptor SocketType = "sock_fd"
)

// ProxyFileSystemDevice defines a filesystem device for which accesses to the
// filesystem are done by virtfs-proxy-helper. The socketType parameter specifies
// whether a file path or file descriptor should be used. The socketFileOrDescriptor
// parameter is the file path or fd to use (based on the socketType).
func ProxyFileSystemDevice(
	id string,
	socketType SocketType,
	socketFileOrDescriptor string,
	properties ...*FileSystemDeviceProperty,
) *queso.Option {
	props := []*FileSystemDeviceProperty{
		NewFileSystemDeviceProperty(string(socketType), socketFileOrDescriptor),
	}

	if properties != nil {
		props = append(props, properties...)
	}

	return FileSystemDevice("proxy", id, props...)
}

// SyntheticFileSystemDevice defines a synthetic filesystem, which is only used by
// QTests. The readOnly parameter defaults to true (see IsMountReadOnly), but
// can be overridden.
func SyntheticFileSystemDevice(id string, readOnly bool) *queso.Option {
	return FileSystemDevice("synth", id, IsMountReadOnly(readOnly))
}

// FileSystemDeviceProperty represents a property associated with a filesystem
// device option.
type FileSystemDeviceProperty struct {
	*queso.Property
}

// NewFileSystemDeviceProperty creates a new instance of FileSystemDeviceProperty.
func NewFileSystemDeviceProperty(key string, value interface{}) *FileSystemDeviceProperty {
	return &FileSystemDeviceProperty{
		Property: queso.NewProperty(key, value),
	}
}

// UseWriteOut indicates that the host page cache will be used to read and write data
// but write notification will be sent to the guest only when the data has been
// reported as written by the storage subsystem.
func UseWriteOut() *FileSystemDeviceProperty {
	return NewFileSystemDeviceProperty("writeout", "immediate")
}

// IsMountReadOnly enables/disables exporting 9p share as a readonly mount for guests.
// By default, read-write access is given.
func IsMountReadOnly(enabled bool) *FileSystemDeviceProperty {
	return NewFileSystemDeviceProperty("readonly", enabled)
}

// WithSocketFile enables proxy filesystem driver to use the specified socket filename
// for communicating with virtfs-proxy-helper
func WithSocketFile(file string) *FileSystemDeviceProperty {
	return NewFileSystemDeviceProperty("socket", file)
}

// WithSocketDescriptor enables proxy filesystem driver to use the specified socket
// descriptor for communicating with virtfs-proxy-helper. Usually a helper like
// libvirt will create `socketpair` and pass one of the fds as sock_fd.
func WithSocketDescriptor(fd string) *FileSystemDeviceProperty {
	return NewFileSystemDeviceProperty("sock_fd", fd)
}

// WithFileMode specifies the default mode for newly created files on the host.
// Works only with security models SecurityModelMappedXAttr and SecurityModelMappedFile.
func WithFileMode(mode string) *FileSystemDeviceProperty {
	return NewFileSystemDeviceProperty("fmode", mode)
}

// WithDirectoryMode specifies the default mode for newly created directories on the host.
// Works only with security models SecurityModelMappedXAttr and SecurityModelMappedFile.
func WithDirectoryMode(mode string) *FileSystemDeviceProperty {
	return NewFileSystemDeviceProperty("dmode", mode)
}

// ThrottlingUnit represents the unit type of the value being throttled.
type ThrottlingUnit string

const (
	ThrottleBytesPerSecond    ThrottlingUnit = "bps"
	ThrottleRequestsPerSecond ThrottlingUnit = "iops"
)

// ThrottlingContext represents the context of the value being throttled.
type ThrottlingContext int

const (
	ThrottlingContextBandwidth ThrottlingContext = iota
	ThrottlingContextBursts
)

// ThrottlingOptions define the parameters for the WithThrottling property.
type ThrottlingOptions struct {
	Context   ThrottlingContext
	Unit      ThrottlingUnit
	Operation IOOperation
	Value     int
}

// WithThrottling can be passed into the File System Device multiple times to
// specify different contexts or units. For example, to throttle the bandwidth
// for reads, pass this into the qemu.With call:
//
//	blockdev.WithThrottling(blockdev.ThrottlingOptions{
//		Context:   ThrottlingContextBandwidth,
//		Unit:      ThrottleBytesPerSecond,
//		Operation: IOOperationRead,
//		Value:     50,
//	})
func WithThrottling(options ThrottlingOptions) *FileSystemDeviceProperty {
	suffix := ""
	if options.Context == ThrottlingContextBursts {
		suffix = "-max"
	}

	op := ""
	switch options.Operation {
	case IOOperationAll:
		op = "total"

	case IOOperationRead:
		op = "read"

	case IOOperationWrite:
		op = "write"
	}

	key := fmt.Sprintf("throttling.%s-%s%s", string(options.Unit), op, suffix)

	return NewFileSystemDeviceProperty(key, options.Value)
}

// WithRequestSizeThrottling sets the bytes of a request count as a new request
// for iops throttling purposes.
func WithRequestSizeThrottling(bytes int) *FileSystemDeviceProperty {
	return NewFileSystemDeviceProperty("throttling.iops-size", bytes)
}
