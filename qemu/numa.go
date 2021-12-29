package qemu

import (
	"fmt"
	"strconv"

	"github.com/mikerourke/queso"
)

// NewNUMAOption returns an Option instance that can be used to specify a NUMA
// node.
func NewNUMAOption(name string, properties ...*NUMAProperty) *queso.Option {
	props := make([]*queso.Property, 0)

	for _, property := range properties {
		props = append(props, property.Property)
	}

	return queso.NewOption("numa", name, props...)
}

// NUMANode defines a generic NUMA node to which properties can be assigned.
func NUMANode(properties ...*NUMAProperty) *queso.Option {
	return NewNUMAOption("node", properties...)
}

// NUMADistance sets the NUMA distance from a source node to a destination node.
func NUMADistance(source string, destination string, distance string) *queso.Option {
	return NewNUMAOption("dist",
		NewNUMAProperty("src", source),
		NewNUMAProperty("dest", destination),
		NewNUMAProperty("val", distance))
}

// NUMACPU assigns CPU objects to a node using topology layout properties of CPU.
// The set of properties is machine specific, and depends on Machine and SMP options.
func NUMACPU(nodeID string, properties ...*NUMAProperty) *queso.Option {
	props := []*NUMAProperty{NewNUMAProperty("node-id", nodeID)}

	if properties != nil {
		props = append(props, properties...)
	}

	return NewNUMAOption("cpu", props...)
}

// NUMAHMATLB sets System Locality Latency and Bandwidth Information between initiator
// and target NUMA nodes in ACPI Heterogeneous Attribute Memory Table (HMAT). Initiator
// NUMA node can create memory requests, usually it has one or more processors. Target
// NUMA node contains addressable memory. See the QEMU documentation for more details.
func NUMAHMATLB(
	initiator int,
	target int,
	hierarchy string,
	dataType string,
	properties ...*NUMAProperty,
) *queso.Option {
	props := []*NUMAProperty{
		NewNUMAProperty("initiator", initiator),
		NewNUMAProperty("target", target),
		NewNUMAProperty("hierarchy", hierarchy),
		NewNUMAProperty("data-type", dataType),
	}

	if properties != nil {
		props = append(props, properties...)
	}

	return NewNUMAOption("hmat-lb", props...)
}

// NUMAHMATCache sets the cache properties for the ACPI Heterogeneous Attribute
// Memory Table (HMAT). See the QEMU documentation for more details.
func NUMAHMATCache(nodeID int, size string, level int, properties ...*NUMAProperty) *queso.Option {
	props := []*NUMAProperty{
		NewNUMAProperty("node-id", nodeID),
		NewNUMAProperty("size", size),
		NewNUMAProperty("level", level),
	}

	if properties != nil {
		props = append(props, properties...)
	}

	return NewNUMAOption("hmat-cache", props...)
}

// NUMAProperty represents a property that can be used with a NUMA option.
type NUMAProperty struct {
	*queso.Property
}

// NewNUMAProperty returns a new instance of an NUMAProperty.
func NewNUMAProperty(key string, value interface{}) *NUMAProperty {
	return &NUMAProperty{
		Property: queso.NewProperty(key, value),
	}
}

// WithMemorySize assigns a given RAM amount to a NUMA node.
func WithMemorySize(size string) *NUMAProperty {
	return NewNUMAProperty("mem", size)
}

// WithMemoryDevice assigns RAM from a given MemoryBackend device ID.
func WithMemoryDevice(id string) *NUMAProperty {
	return NewNUMAProperty("memdev", id)
}

// WithCPUs is used to assign VCPUs to a NUMA node. The value can be a single
// number or two numbers representing a range.
func WithCPUs(cpus ...int) *NUMAProperty {
	switch len(cpus) {
	case 1:
		return NewNUMAProperty("cpus", strconv.Itoa(cpus[0]))

	case 2:
		return NewNUMAProperty("cpus",
			fmt.Sprintf("%d-%d", cpus[0], cpus[1]))

	default:
		panic("only 1 or 2 CPUs are allowed for WithCPUCount")
	}
}

// WithNUMANodeID specifies the associated NUMA node ID.
func WithNUMANodeID(id int) *NUMAProperty {
	return NewNUMAProperty("nodeid", id)
}

// WithInitiator points to an initiator NUMA node that has the best performance
// (the lowest latency or largest bandwidth) to the NUMA node. Note that this option
// can be set only when the machine property WithHMAT is enabled.
func WithInitiator(initiator int) *NUMAProperty {
	return NewNUMAProperty("initiator", initiator)
}

// WithNUMALatency sets the latency for the NUMA node.
func WithNUMALatency(nanoseconds int) *NUMAProperty {
	return NewNUMAProperty("latency", nanoseconds)
}

// WithNUMABandwidth sets the bandwidth for the NUMA node. The possible value and
// units are NUM[M|G|T], mean that the bandwidth value are NUM byte per second
// (or MB/s, GB/s or TB/s depending on used suffix).
func WithNUMABandwidth(bandwidth string) *NUMAProperty {
	return NewNUMAProperty("bandwidth", bandwidth)
}

// CacheAssociativity represents the possible values for the WithCacheAssociativity
// property.
type CacheAssociativity string

const (
	// CacheAssociativityNone indicates that no associativity should be used.
	CacheAssociativityNone CacheAssociativity = "none"

	// CacheAssociativityDirect indicates that direct-mapped associativity
	// should be used.
	CacheAssociativityDirect CacheAssociativity = "direct"

	// CacheAssociativityComplex indicates that complex cache indexing
	// associativity should be used.
	CacheAssociativityComplex CacheAssociativity = "complex"
)

// WithCacheAssociativity is the NUMA cache associativity.
func WithCacheAssociativity(associativity CacheAssociativity) *NUMAProperty {
	return NewNUMAProperty("associativity", associativity)
}

// WithWritePolicy is the write policy.
func WithWritePolicy(policy string) *NUMAProperty {
	return NewNUMAProperty("policy", policy)
}

// WithWriteCacheLineSize is the write cache line size in bytes.
func WithWriteCacheLineSize(bytes int) *NUMAProperty {
	return NewNUMAProperty("line", bytes)
}
