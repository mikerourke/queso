package fsdev

// Type indicates the type of file system device or virtual file system device.
type Type string

const (
	// Local represents a file system device in which accesses to the filesystem
	// are done by QEMU.
	Local Type = "local"

	// Proxy represents a file system device in which accesses to the filesystem are
	// done by virtfs-proxy-helper(1).
	//
	// Deprecated: Use [Local] instead.
	Proxy Type = "proxy"

	// Synth represents a synthetic filesystem, only used by QTests.
	Synth Type = "synth"
)

type SecurityModel string

const (
	SecurityModelNone        SecurityModel = "none"
	SecurityModelPassthrough SecurityModel = "passthrough"
	SecurityModelMappedXAttr SecurityModel = "mapped-xattr"
	SecurityModelMappedFile  SecurityModel = "mapped-file"
)
