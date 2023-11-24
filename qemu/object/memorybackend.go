package object

import (
	"strings"

	"github.com/mikerourke/queso"
	"github.com/mikerourke/queso/qemu/numa"
)

// MemoryBackend is the generic memory backend upon which other backends are
// built.
type MemoryBackend struct {
	Type       string
	properties []*queso.Property
}

// NewMemoryBackend returns a new instance of [MemoryBackend]. The id is a unique ID
// that will be used to reference this memory region in other parameters,
// e.g. -numa, -device nvdimm, etc.
func NewMemoryBackend(backendType string, id string) *MemoryBackend {
	return &MemoryBackend{
		Type: backendType,
		properties: []*queso.Property{
			queso.NewProperty("id", id),
		},
	}
}

func (b *MemoryBackend) option() *queso.Option {
	return queso.NewOption("object", b.Type, b.properties...)
}

// SetHostNodes binds the memory range to the specified list of NUMA host nodes.
//
//	qemu-system-* -object memory-backend-*,host-nodes=host-nodes
func (b *MemoryBackend) SetHostNodes(ids []string) *MemoryBackend {
	b.properties = append(b.properties,
		queso.NewProperty("host-nodes", strings.Join(ids, ",")))
	return b
}

// SetNUMAPolicy defines the NUMA policy. See [numa.Policy] for additional details.
//
//	qemu-system-* -object memory-backend-*,policy=default|preferred|bind|interleave
func (b *MemoryBackend) SetNUMAPolicy(policy numa.Policy) *MemoryBackend {
	b.properties = append(b.properties, queso.NewProperty("policy", string(policy)))
	return b
}

// SetSize provides the size of the memory region, and accepts common
// suffixes, e.g. "500M".
//
//	qemu-system-* -object memory-backend-*,size=size
func (b *MemoryBackend) SetSize(size string) *MemoryBackend {
	b.properties = append(b.properties, queso.NewProperty("size", size))
	return b
}

// ToggleDump excludes the memory from core dumps when set to false. This feature
// is also known as MADV_DONTDUMP.
//
//	qemu-system-* -object memory-backend-*,dump=on|off
func (b *MemoryBackend) ToggleDump(enabled bool) *MemoryBackend {
	b.properties = append(b.properties, queso.NewProperty("dump", enabled))
	return b
}

// ToggleMerging enables or disables memory merge, also known as MADV_MERGEABLE, so
// that Kernel Samepage Merging will consider the pages for memory deduplication.
//
//	qemu-system-* -object memory-backend-*,merge=on|off
func (b *MemoryBackend) ToggleMerging(enabled bool) *MemoryBackend {
	b.properties = append(b.properties, queso.NewProperty("merge", enabled))
	return b
}

// TogglePrealloc enables or disables memory preallocation.
//
//	qemu-system-* -object memory-backend-*,prealloc=on|off
func (b *MemoryBackend) TogglePrealloc(enabled bool) *MemoryBackend {
	b.properties = append(b.properties, queso.NewProperty("prealloc", enabled))
	return b
}

// ToggleSharing determines whether the memory region is marked as private to QEMU, or
// shared. The latter allows a co-operating external process to access the
// QEMU memory region.
//
// This option is also required for pvrdma devices due to limitations in the
// RDMA API provided by Linux.
//
// Setting enabled to true might affect the ability to configure NUMA bindings for the memory
// backend under some circumstances, see Documentation/vm/numa_memory_policy.txt on the
// Linux kernel source tree for additional details
//
//	qemu-system-* -object memory-backend-*,share=on|off
func (b *MemoryBackend) ToggleSharing(enabled bool) *MemoryBackend {
	b.properties = append(b.properties, queso.NewProperty("share", enabled))
	return b
}

// MemoryBackendFile represents a memory file backend object, which can be used to
// back the guest RAM with huge pages.
type MemoryBackendFile struct {
	*MemoryBackend
}

// NewMemoryBackendFile represents a memory file backend object, creates a new
// instance of MemoryBackendFile.
//
//	qemu-system-* -object memory-backend-file
func NewMemoryBackendFile(id string) *MemoryBackendFile {
	return &MemoryBackendFile{
		MemoryBackend: &MemoryBackend{
			Type: "memory-backend-file",
			properties: []*queso.Property{
				queso.NewProperty("id", id),
			},
		},
	}
}

// SetMemoryPath provides the path to either a shared memory or huge page filesystem mount.
//
//	qemu-system-* -object memory-backend-file,mem-path=dir
func (b *MemoryBackendFile) SetMemoryPath(dir string) *MemoryBackendFile {
	b.properties = append(b.properties, queso.NewProperty("mem-path", dir))
	return b
}

type ROMStatus string

const (
	// ROMOn indicates that ROM should be created.
	ROMOn ROMStatus = "on"

	// ROMOff indicates that writable RAM should be created.
	ROMOff ROMStatus = "off"

	// ROMAuto uses the value set by MemoryBackendFile.ToggleReadOnly.
	ROMAuto ROMStatus = "auto"
)

// SetROMStatus specifies whether to create Read Only Memory (ROM) that cannot be modified
// by the VM. Any write attempts to such ROM will be denied. Most use cases want
// proper RAM instead of ROM. However, selected use cases, like R/O NVDIMMs, can
// benefit from ROM.
//
// The [ROMAuto] option is primarily helpful when we want to have writable RAM in
// configurations that would traditionally create ROM before the ROM option was
// introduced: VM templating, where we want to open a file read-only
// (ToggleReadOnly called with true) and mark the memory to be private
// for QEMU (ToggleSharing = false). For this use case, we need
// writable RAM instead of ROM, and want to also set ToggleReadOnly = false.
//
//	qemu-system-* -object memory-backend-file,rom=on|off|auto
func (b *MemoryBackendFile) SetROMStatus(status ROMStatus) *MemoryBackendFile {
	b.properties = append(b.properties, queso.NewProperty("rom", status))
	return b
}

// ToggleDiscardData indicates if the file contents can be destroyed when QEMU exits, to
// avoid unnecessarily flushing data to the backing file.
//
// Note that setting this to true is only an optimization, and QEMU might not
// discard file contents if it aborts unexpectedly or is terminated using SIGKILL.
//
//	qemu-system-* -object memory-backend-file,discard-data=on|off
func (b *MemoryBackendFile) ToggleDiscardData(enabled bool) *MemoryBackendFile {
	b.properties = append(b.properties, queso.NewProperty("discard-data", enabled))
	return b
}

// SetAlignment specifies the base address alignment when QEMU mmap(2)
// [MemoryBackendFile.SetMemoryPath], and accepts common suffixes, e.g. 2M.
// Some backend store specified by [MemoryBackendFile.SetMemoryPath] requires
// an alignment different from the default one used by QEMU, e.g. the
// device DAX /dev/dax0.0 requires 2M alignment rather than 4K.
//
// In such cases, users can specify the required alignment via this option.
//
//	qemu-system-* -object memory-backend-file,align=align
func (b *MemoryBackendFile) SetAlignment(align string) *MemoryBackendFile {
	b.properties = append(b.properties, queso.NewProperty("align", align))
	return b
}

// SetOffset specifies the offset into the target file that the region starts at.
// You can use this parameter to back multiple regions with a single file.
//
//	qemu-system-* -object memory-backend-file,offset=offset
func (b *MemoryBackendFile) SetOffset(offset int) *MemoryBackendFile {
	b.properties = append(b.properties, queso.NewProperty("offset", offset))
	return b
}

// ToggleReadOnly specifies whether the backing file is opened read-only or read-write (default).
//
//	qemu-system-* -object memory-backend-file,readonly=on|off
func (b *MemoryBackendFile) ToggleReadOnly(enabled bool) *MemoryBackendFile {
	b.properties = append(b.properties, queso.NewProperty("readonly", enabled))
	return b
}

// MemoryBackendRAM represents a memory backend object which can be used to back the
// guest RAM. Memory backend objects offer more control than using qemu.NewMemory,
// which is traditionally used to define guest RAM.
type MemoryBackendRAM struct {
	*MemoryBackend
}

// NewMemoryBackendRAM represents a memory file backend object, creates a new
// instance of MemoryBackendFile.
//
//	qemu-system-* -object memory-backend-ram
func NewMemoryBackendRAM(id string) *MemoryBackendFile {
	return &MemoryBackendFile{
		MemoryBackend: &MemoryBackend{
			Type: "memory-backend-ram",
			properties: []*queso.Property{
				queso.NewProperty("id", id),
			},
		},
	}
}
