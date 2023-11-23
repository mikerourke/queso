package qemu

// Memory sets guest startup RAM size to specified size in megabytes. Default is 128 MiB.
// Optionally, a suffix of "M" or "G" can be used to signify a value in megabytes
// or gigabytes respectively.
type Memory struct {
	Size       string
	properties []*cli.Property
}

// NewMemory returns a new Memory instance (for setting memory properties in VM).
func NewMemory(size string) *Memory {
	return &Memory{
		Size:       size,
		properties: make([]*cli.Property, 0),
	}
}

// Option returns the Memory settings as an Option for use in command line.
func (m *Memory) option() *cli.Option {
	return cli.NewOption("m", "", m.properties...)
}

// SetMemorySlots specifies amount of hot-pluggable memory slots.
func (m *Memory) SetMemorySlots(count int) *Memory {
	m.properties = append(m.properties, cli.NewProperty("slots", count))
	return m
}

// SetMemoryMaximum specifies maximum amount of memory. Note that the size
// must be aligned to the page size.
func (m *Memory) SetMemoryMaximum(size string) *Memory {
	m.properties = append(m.properties, cli.NewProperty("maxmem", size))
	return m
}

// MemoryPath allocates guest RAM from a temporarily created file in path.
func MemoryPath(path string) *cli.Option {
	return cli.NewOption("mem-path", path)
}

// PreallocateMemory pre-allocates memory when using MemoryPath.
func PreallocateMemory() *cli.Option {
	return cli.NewOption("mem-prealloc", "")
}

// MemorySize returns an option used to set the memory size with a string and
// no other options.
func MemorySize(size string) *cli.Option {
	return cli.NewOption("m", size)
}
