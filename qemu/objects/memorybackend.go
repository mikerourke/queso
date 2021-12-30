package objects

import (
	"strings"

	"github.com/mikerourke/queso"
)

// MemoryBackend returns an Option that represents a Memory Backend object
// to pass to QEMU.
func MemoryBackend(name string, id string, properties ...*MemoryBackendProperty) *queso.Option {
	props := []*queso.Property{{"id", id}}

	for _, property := range properties {
		props = append(props, property.Property)
	}

	return queso.NewOption("object", name, props...)
}

// MemoryBackendFile represents a memory file backend object, which can be used to
// back the guest RAM with huge pages.
func MemoryBackendFile(id string, properties ...*MemoryBackendProperty) *queso.Option {
	return MemoryBackend("memory-backend-file", id, properties...)
}

// MemoryBackendRAM represents a memory backend object, which can be used to back
// the guest RAM. Memory backend objects offer more control than the WithMemory option
// that is traditionally used to define guest RAM.
func MemoryBackendRAM(id string, properties ...*MemoryBackendProperty) *queso.Option {
	return MemoryBackend("memory-backend-ram", id, properties...)
}

// MemoryBackendMemFD represents an anonymous memory file backend object, which
// allows QEMU to share the memory with an external process (e.g. when using
// vhost-user). The memory is allocated with memfd and optional sealing.
// This option can only be used on Linux.
func MemoryBackendMemFD(id string, properties ...*MemoryBackendProperty) *queso.Option {
	return MemoryBackend("memory-backend-memfd", id, properties...)
}

// MemoryBackendProperty represents a property that can be used with an audio device
// option.
type MemoryBackendProperty struct {
	*queso.Property
}

// NewMemoryBackendProperty returns a new instance of an MemoryBackendProperty.
func NewMemoryBackendProperty(key string, value interface{}) *MemoryBackendProperty {
	return &MemoryBackendProperty{
		Property: queso.NewProperty(key, value),
	}
}

// WithMemorySize provides the size of the memory region, and accepts common suffixes, e.g. 500M.
func WithMemorySize(size string) *MemoryBackendProperty {
	return NewMemoryBackendProperty("size", size)
}

// WithMemoryPath provides the path to either a shared memory or huge page
// filesystem mount.
func WithMemoryPath(path string) *MemoryBackendProperty {
	return NewMemoryBackendProperty("mem-path", path)
}

// IsShare determines whether the memory region is marked as private to QEMU,
// or shared. The latter allows a co-operating external process to access the QEMU
// memory region.
func IsShare(enabled bool) *MemoryBackendProperty {
	return NewMemoryBackendProperty("share", enabled)
}

// IsDiscardData indicates whether file contents can be destroyed when QEMU exits, to
// avoid unnecessarily flushing data to the backing file. Note that IsDiscardData is
// only an optimization, and QEMU might not discard file contents if it aborts
// unexpectedly or is terminated using SIGKILL.
func IsDiscardData(enabled bool) *MemoryBackendProperty {
	return NewMemoryBackendProperty("discard-data", enabled)
}

// IsMemoryMerge enables memory merge, also known as MADV_MERGEABLE, so that Kernel
// Samepage Merging will consider the pages for memory deduplication.
func IsMemoryMerge(enabled bool) *MemoryBackendProperty {
	return NewMemoryBackendProperty("merge", enabled)
}

// IsDump indicates whether to include the memory from core dumps. This feature
// is also known as MADV_DONTDUMP.
func IsDump(enabled bool) *MemoryBackendProperty {
	return NewMemoryBackendProperty("dump", enabled)
}

// IsPrealloc indicates whether to enable memory pre-allocation.
func IsPrealloc(enabled bool) *MemoryBackendProperty {
	return NewMemoryBackendProperty("prealloc", enabled)
}

// WithHostNodes binds the memory range to the specified list of NUMA host nodes.
func WithHostNodes(ids []string) *MemoryBackendProperty {
	return NewMemoryBackendProperty("host-nodes", strings.Join(ids, ","))
}

// NUMAPolicy represents the NUMA policy to use for the WithNUMAPolicy property.
type NUMAPolicy string

const (
	// NUMAPolicyDefault is the default host policy.
	NUMAPolicyDefault NUMAPolicy = "default"

	// NUMAPolicyPreferred prefers the given host node list for allocation.
	NUMAPolicyPreferred NUMAPolicy = "preferred"

	// NUMAPolicyBind restricts memory allocation to the given host node list.
	NUMAPolicyBind NUMAPolicy = "bind"

	// NUMAPolicyInterleave interleaves memory allocations across the given host
	// node list.
	NUMAPolicyInterleave = "interleave"
)

// WithNUMAPolicy set the NUMA policy to the specified value. See NUMAPolicy for
// details.
func WithNUMAPolicy(policy NUMAPolicy) *MemoryBackendProperty {
	return NewMemoryBackendProperty("policy", policy)
}

// WithAlign specifies the base address alignment when QEMU mmap(2) mem-path, and
// accepts common suffixes, e.g. 2M. Some backend store specified by WithMemoryPath
// requires an alignment different from the default one used by QEMU, e.g. the
// device DAX /dev/dax0.0 requires 2M alignment rather than 4K. In such cases,
// users can specify the required alignment via this option.
func WithAlign(alignment string) *MemoryBackendProperty {
	return NewMemoryBackendProperty("align", alignment)
}

// IsPersistentMemory specifies whether the backing file specified by WithMemoryPath is
// in host persistent memory that can be accessed using the SNIA NVM programming
// model (e.g. Intel NVDIMM). If PersistentMemory is set to true, QEMU will take
// necessary operations to guarantee the persistence of its own writes to mem-path
// (e.g. in vNVDIMM label emulation and live migration). Also, we will map the
// BackendFile with MAP_SYNC flag, which ensures the file metadata is in sync for
// mem-path in case of host crash or a power failure. MAP_SYNC requires support
// from both the host kernel (since Linux kernel 4.15) and the filesystem of
// mem-path mounted with DAX option.
func IsPersistentMemory(enabled bool) *MemoryBackendProperty {
	return NewMemoryBackendProperty("pmem", enabled)
}

// IsReadOnly specifies whether the backing file is opened read-only or
// read-write (default).
func IsReadOnly(enabled bool) *MemoryBackendProperty {
	return NewMemoryBackendProperty("readonly", enabled)
}

// IsSeal indicates whether to create a sealed-file, that will block further
// resizing the memory. The default is true.
//
// This property is only valid for the MemFD backend.
func IsSeal(enabled bool) *MemoryBackendProperty {
	return NewMemoryBackendProperty("seal", enabled)
}

// IsHugeTLB indicates if the file to be created resides in the hugetlbfs
// filesystem (since Linux 4.14). In some versions of Linux, this option is
// incompatible with the Seal option (requires at least Linux 4.16).
//
// This property is only valid for the MemFD backend.
func IsHugeTLB(enabled bool) *MemoryBackendProperty {
	return NewMemoryBackendProperty("hugetlb", enabled)
}

// WithHugeTLBSize is used to specify the hugetlb page size on systems that support
// multiple hugetlb page sizes (it must be a power of 2 value supported by the system).
//
// This property is only valid for the MemFD backend.
func WithHugeTLBSize(bytes int) *MemoryBackendProperty {
	return NewMemoryBackendProperty("hugetlbsize", bytes)
}
