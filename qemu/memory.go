package qemu

import "github.com/mikerourke/queso"

// MemoryPath allocates guest RAM from a temporarily created file in path.
func MemoryPath(path string) *queso.Option {
	return queso.NewOption("mem-path", path)
}

// PreallocateMemory pre-allocates memory when using MemoryPath.
func PreallocateMemory() *queso.Option {
	return queso.NewOption("mem-prealloc", "")
}

// Memory sets guest startup RAM size to specified size in megabytes. Default is 128 MiB.
// Optionally, a suffix of "M" or "G" can be used to signify a value in megabytes
// or gigabytes respectively.
func Memory(size string, properties ...*MemoryProperty) *queso.Option {
	props := make([]*queso.Property, 0)

	for _, property := range properties {
		props = append(props, property.Property)
	}

	return queso.NewOption("m", size, props...)
}

// MemoryProperty represents a property that can be used with Memory.
type MemoryProperty struct {
	*queso.Property
}

// NewMemoryProperty returns a new instance of MemoryProperty.
func NewMemoryProperty(key string, value interface{}) *MemoryProperty {
	return &MemoryProperty{
		Property: queso.NewProperty(key, value),
	}
}

// WithMemorySlots specifies amount of hot-pluggable memory slots.
func WithMemorySlots(count int) *MemoryProperty {
	return NewMemoryProperty("slots", count)
}

// WithMemoryMaximum specifies maximum amount of memory. Note that the size
// must be aligned to the page size.
func WithMemoryMaximum(size string) *MemoryProperty {
	return NewMemoryProperty("maxmem", size)
}
