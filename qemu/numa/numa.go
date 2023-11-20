// Package numa is used to define NUMA nodes for use with QEMU.
package numa

import (
	"fmt"
	"strconv"

	"github.com/mikerourke/queso"
)

// NUMA defines a NUMA object.
func NUMA(name string, properties ...*Property) *queso.Option {
	props := make([]*queso.Property, 0)

	for _, property := range properties {
		props = append(props, property.Property)
	}

	return queso.NewOption("numa", name, props...)
}

// Node defines a generic NUMA node to which properties can be assigned.
//
// Example
//
//	qemu.New("qemu-system-x86_64").SetOptions(
//		qemu.Machine("pc"),
//		qemu.SMP(
//			qemu.WithCPUCount(1),
//			qemu.WithSockets(2),
//			qemu.WithMaxCPUs(2)),
//		numa.Node(numa.WithNodeID(0)),
//		numa.Node(numa.WithNodeID(1)),
//		numa.CPU(0, numa.WithCPUSocketID(0)),
//		numa.CPU(1, numa.WithCPUSocketID(1)))
//
// Invocation
//
//	qemu-system-x86_64
//		-machine pc \
//		-smp 1,sockets=2,maxcpus=2 \
//		-numa node,nodeid=0 -numa node,nodeid=1 \
//		-numa cpu,node-id=0,socket-id=0 -numa cpu,node-id=1,socket-id=1
func Node(properties ...*Property) *queso.Option {
	return NUMA("node", properties...)
}

// Distance sets the NUMA distance from a source node to a destination node.
func Distance(source string, destination string, distance string) *queso.Option {
	return NUMA("dist",
		NewProperty("src", source),
		NewProperty("dest", destination),
		NewProperty("val", distance))
}

// CPU assigns CPU objects to a node using topology layout properties of CPU.
// The set of properties is machine specific, and depends on Machine and SMP
// options.
func CPU(nodeID int, properties ...*Property) *queso.Option {
	props := []*Property{NewProperty("node-id", nodeID)}

	if properties != nil {
		props = append(props, properties...)
	}

	return NUMA("cpu", props...)
}

// HMATLB sets System Locality Latency and Bandwidth Information between
// initiator and target NUMA nodes in ACPI Heterogeneous Attribute Memory Table
// (HMAT). Initiator NUMA node can create memory requests, usually it has one or
// more processors. Target NUMA node contains addressable memory. See the QEMU
// documentation for more details.
func HMATLB(
	initiator int,
	target int,
	hierarchy string,
	dataType string,
	properties ...*Property,
) *queso.Option {
	props := []*Property{
		NewProperty("initiator", initiator),
		NewProperty("target", target),
		NewProperty("hierarchy", hierarchy),
		NewProperty("data-type", dataType),
	}

	if properties != nil {
		props = append(props, properties...)
	}

	return NUMA("hmat-lb", props...)
}

// HMATCache sets the cache properties for the ACPI Heterogeneous Attribute
// Memory Table (HMAT). See the QEMU documentation for more details.
func HMATCache(nodeID int, size string, level int, properties ...*Property) *queso.Option {
	props := []*Property{
		NewProperty("node-id", nodeID),
		NewProperty("size", size),
		NewProperty("level", level),
	}

	if properties != nil {
		props = append(props, properties...)
	}

	return NUMA("hmat-cache", props...)
}

// Property represents a property that can be used with a NUMA option.
type Property struct {
	*queso.Property
}

// NewProperty returns a new instance of Property.
func NewProperty(key string, value interface{}) *Property {
	return &Property{
		Property: queso.NewProperty(key, value),
	}
}

// WithMemorySize assigns a given RAM amount to a NUMA node.
func WithMemorySize(size string) *Property {
	return NewProperty("mem", size)
}

// WithMemoryDevice assigns RAM from a given MemoryBackend device ID.
func WithMemoryDevice(id string) *Property {
	return NewProperty("memdev", id)
}

// WithCPUs is used to assign VCPUs to a NUMA node. The value can be a single
// number or two numbers representing a range.
func WithCPUs(cpus ...int) *Property {
	switch len(cpus) {
	case 1:
		return NewProperty("cpus", strconv.Itoa(cpus[0]))

	case 2:
		return NewProperty("cpus",
			fmt.Sprintf("%d-%d", cpus[0], cpus[1]))

	default:
		panic("only 1 or 2 CPUs are allowed for WithCPUCount")
	}
}

// WithNodeID specifies the associated NUMA node ID.
func WithNodeID(id int) *Property {
	return NewProperty("nodeid", id)
}

// WithCPUSocketID specifies the CPU socket ID for a CPU.
func WithCPUSocketID(id int) *Property {
	return NewProperty("socket-id", id)
}

// WithCPUCoreID specifies the CPU core ID for a CPU.
func WithCPUCoreID(id int) *Property {
	return NewProperty("core-id", id)
}

// WithCPUThreadID specifies the CPU thread ID for a CPU.
func WithCPUThreadID(id int) *Property {
	return NewProperty("thread-id", id)
}

// WithInitiator points to an initiator NUMA node that has the best performance
// (the lowest latency or largest bandwidth) to the NUMA node. Note that this
// option can be set only when the machine property WithHMAT is enabled.
func WithInitiator(initiator int) *Property {
	return NewProperty("initiator", initiator)
}

// WithLatency sets the latency for the NUMA node.
func WithLatency(nanoseconds int) *Property {
	return NewProperty("latency", nanoseconds)
}

// WithBandwidth sets the bandwidth for the NUMA node. The possible value and
// units are NUM[M|G|T], mean that the bandwidth value are NUM byte per second
// (or MB/s, GB/s or TB/s depending on used suffix).
func WithBandwidth(bandwidth string) *Property {
	return NewProperty("bandwidth", bandwidth)
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
func WithCacheAssociativity(associativity CacheAssociativity) *Property {
	return NewProperty("associativity", associativity)
}

// WithWritePolicy is the write policy.
func WithWritePolicy(policy string) *Property {
	return NewProperty("policy", policy)
}

// WithWriteCacheLineSize is the write cache line size in bytes.
func WithWriteCacheLineSize(bytes int) *Property {
	return NewProperty("line", bytes)
}
