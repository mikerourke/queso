package numa

import (
	"github.com/mikerourke/queso"
	"github.com/mikerourke/queso/qemu"
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

type HMATCache struct {
	qemu.Usable
	properties []*queso.Property
}

func NewHMATCache(node *Node, size string, level int) *HMATCache {
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
