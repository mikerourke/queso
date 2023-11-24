package numa

import (
	"github.com/mikerourke/queso"
	"github.com/mikerourke/queso/qemu"
)

// System represents all the nodes, CPUs, HMAT-LB, and HMAT caches in a NUMA
// configuration.
type System struct {
	nodes []qemu.Usable
}

// NewSystem returns a new instance of [System].
func NewSystem() *System {
	return &System{
		nodes: make([]qemu.Usable, 0),
	}
}

// Add adds a node, CPU, HMAT-LB, or HMAT cache to the system.
func (s *System) Add(nodes ...qemu.Usable) *System {
	for _, node := range nodes {
		s.nodes = append(s.nodes, node)
	}
	return s
}

// Nodes returns the nodes in the system.
func (s *System) Nodes() []qemu.Usable {
	return s.nodes
}

// UseIn is a convenience function for using the NUMA elements defined in a
// [System].
//
// Example
//
//	// Instead of:
//	q := qemu.New()
//	ns := numa.NewSystem()
//	// ... add stuff to system here ...
//	qemu.Use(ns.Nodes()...)
//
//	You can do this:
//	q := qemu.New()
//	ns := numa.NewSystem()
//	ns.UseIn(q)
func (s *System) UseIn(q *qemu.QEMU) {
	q.Use(s.nodes...)
}

// SetDistanceBetweenNodes sets the NUMA distance from a source node to a
// destination node.
func (s *System) SetDistanceBetweenNodes(
	source *Node,
	destination *Node,
	distance string,
) *System {
	dist := &nodeDistance{
		properties: []*queso.Property{
			queso.NewProperty("src", source.ID),
			queso.NewProperty("dst", destination.ID),
			queso.NewProperty("val", distance),
		},
	}
	s.nodes = append(s.nodes, dist)
	return s
}

type nodeDistance struct {
	qemu.Usable
	properties []*queso.Property
}

func (nd *nodeDistance) option() *queso.Option {
	return queso.NewOption("numa", "dist", nd.properties...)
}
