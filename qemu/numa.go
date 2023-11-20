package qemu

import (
	"fmt"
	"strconv"

	"github.com/mikerourke/queso"
)

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

type NUMANode struct {
	Usable
	ID         int
	properties []*queso.Property
}

func NewNUMANode(id int) *NUMANode {
	return &NUMANode{
		ID:         id,
		properties: make([]*queso.Property, 0),
	}
}

func (nn *NUMANode) option() *queso.Option {
	return queso.NewOption("numa", "node", nn.properties...)
}

// SetMemorySize assigns a given RAM amount to a NUMA node.
func (nn *NUMANode) SetMemorySize(size string) *NUMANode {
	nn.properties = append(nn.properties, queso.NewProperty("mem", size))
	return nn
}

// SetMemoryDevice assigns RAM from a given MemoryBackend device ID.
func (nn *NUMANode) SetMemoryDevice(id string) *NUMANode {
	nn.properties = append(nn.properties, queso.NewProperty("memdev", id))
	return nn
}

func (nn *NUMANode) SetCPUs(cpus ...int) *NUMANode {
	value := "0"
	switch len(cpus) {
	case 1:
		value = strconv.Itoa(cpus[0])

	case 2:
		value = fmt.Sprintf("%d-%d", cpus[0], cpus[1])

	default:
		panic("only 1 or 2 CPUs are allowed for SetCPUs")
	}

	nn.properties = append(nn.properties, queso.NewProperty("cpus", value))
	return nn
}

// SetInitiator points to an initiator NUMA node that has the best performance
// (the lowest latency or largest bandwidth) to the NUMA node. Note that this
// option can be set only when the machine property WithHMAT is enabled.
func (nn *NUMANode) SetInitiator(initiator int) *NUMANode {
	nn.properties = append(nn.properties, queso.NewProperty("initiator", initiator))
	return nn
}

type NUMACPU struct {
	Usable
	properties []*queso.Property
}

func NewNUMACPU(node *NUMANode) *NUMACPU {
	return &NUMACPU{
		properties: []*queso.Property{
			queso.NewProperty("node-id", node.ID),
		},
	}
}

func (cpu *NUMACPU) option() *queso.Option {
	return queso.NewOption("numa", "cpu", cpu.properties...)
}

// SetSocketID specifies the CPU socket ID for a CPU.
func (cpu *NUMACPU) SetSocketID(id int) *NUMACPU {
	cpu.properties = append(cpu.properties, queso.NewProperty("socket-id", id))
	return cpu
}

// SetCoreID specifies the CPU core ID for a CPU.
func (cpu *NUMACPU) SetCoreID(id int) *NUMACPU {
	cpu.properties = append(cpu.properties, queso.NewProperty("core-id", id))
	return cpu
}

// SetThreadID specifies the CPU thread ID for a CPU.
func (cpu *NUMACPU) SetThreadID(id int) *NUMACPU {
	cpu.properties = append(cpu.properties, queso.NewProperty("thread-id", id))
	return cpu
}

type NUMADistance struct {
	Usable
	properties []*queso.Property
}

func NewNUMADistance(source *NUMANode, destination *NUMANode, distance string) *NUMADistance {
	properties := []*queso.Property{
		queso.NewProperty("src", source.ID),
		queso.NewProperty("dst", destination.ID),
		queso.NewProperty("val", distance),
	}

	return &NUMADistance{
		properties: properties,
	}
}

func (nd *NUMADistance) option() *queso.Option {
	return queso.NewOption("numa", "dist", nd.properties...)
}

type HMATLB struct {
	Usable
	properties []*queso.Property
}

type HMATLBDataType string

const (
	HMATLBBandwidthAccess HMATLBDataType = "access-bandwidth"
	HMATLBBandwidthRead   HMATLBDataType = "read-bandwidth"
	HMATLBBandwidthWrite  HMATLBDataType = "write-bandwidth"
	HMATLBLatencyAccess   HMATLBDataType = "access-latency"
	HMATLBLatencyRead     HMATLBDataType = "read-latency"
	HMATLBLatencyWrite    HMATLBDataType = "write-latency"
)

func NewHMATLB(initiator *NUMANode, target *NUMANode, hierarchy string, dataType HMATLBDataType) *HMATLB {
	properties := []*queso.Property{
		queso.NewProperty("initiator", initiator.ID),
		queso.NewProperty("target", target.ID),
		queso.NewProperty("hierarchy", hierarchy),
		queso.NewProperty("data-type", dataType),
	}

	return &HMATLB{
		properties: properties,
	}
}

func (h *HMATLB) option() *queso.Option {
	return queso.NewOption("numa", "hmat-lb", h.properties...)
}

func (h *HMATLB) SetLatency(value int) *HMATLB {
	for _, property := range h.properties {
		if property.Key == "bandwidth" {
			panic("cannot set latency on HMATLB if bandwidth already set")
		}
	}

	h.properties = append(h.properties, queso.NewProperty("latency", value))
	return h
}

func (h *HMATLB) SetBandwidth(value string) *HMATLB {
	for _, property := range h.properties {
		if property.Key == "latency" {
			panic("cannot set bandwidth on HMATLB if latency already set")
		}
	}

	h.properties = append(h.properties, queso.NewProperty("bandwidth", value))
	return h
}

type HMATCache struct {
	Usable
	properties []*queso.Property
}

func NewHMATCache(node *NUMANode, size string, level int) *HMATCache {
	properties := []*queso.Property{
		queso.NewProperty("node-id", node.ID),
		queso.NewProperty("size", size),
		queso.NewProperty("level", level),
	}

	return &HMATCache{
		properties: properties,
	}
}

func (h *HMATCache) option() *queso.Option {
	return queso.NewOption("numa", "hmat-cache", h.properties...)
}

// SetAssociativity sets the NUMA cache associativity.
func (h *HMATCache) SetAssociativity(associativity CacheAssociativity) *HMATCache {
	h.properties = append(h.properties, queso.NewProperty("associativity", associativity))
	return h
}

// SetPolicy sets the write policy.
func (h *HMATCache) SetPolicy(policy string) *HMATCache {
	h.properties = append(h.properties, queso.NewProperty("policy", policy))
	return h
}

// SetLineSize sets the write cache line size in bytes.
func (h *HMATCache) SetLineSize(bytes int) *HMATCache {
	h.properties = append(h.properties, queso.NewProperty("line", bytes))
	return h
}

type NUMASystem struct {
	nodes []Usable
}

func NewNUMASystem() *NUMASystem {
	return &NUMASystem{
		nodes: make([]Usable, 0),
	}
}

func (ns *NUMASystem) Add(nodes ...Usable) *NUMASystem {
	for _, node := range nodes {
		ns.nodes = append(ns.nodes, node)
	}
	return ns
}

func (ns *NUMASystem) Nodes() []Usable {
	return ns.nodes
}
