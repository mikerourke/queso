package numa

import (
	"fmt"
	"strconv"

	"github.com/mikerourke/queso"
)

type Node struct {
	ID         int
	properties []*queso.Property
}

func NewNode(id int) *Node {
	return &Node{
		ID:         id,
		properties: make([]*queso.Property, 0),
	}
}

func (n *Node) option() *queso.Option {
	return queso.NewOption("numa", "node", n.properties...)
}

// SetMemorySize assigns a given RAM amount to a NUMA node.
func (n *Node) SetMemorySize(size string) *Node {
	n.properties = append(n.properties, queso.NewProperty("mem", size))
	return n
}

// SetMemoryDevice assigns RAM from a given MemoryBackend device ID.
func (n *Node) SetMemoryDevice(id string) *Node {
	n.properties = append(n.properties, queso.NewProperty("memdev", id))
	return n
}

func (n *Node) SetCPUs(cpus ...int) *Node {
	value := "0"
	switch len(cpus) {
	case 1:
		value = strconv.Itoa(cpus[0])

	case 2:
		value = fmt.Sprintf("%d-%d", cpus[0], cpus[1])

	default:
		panic("only 1 or 2 CPUs are allowed for SetCPUs")
	}

	n.properties = append(n.properties, queso.NewProperty("cpus", value))
	return n
}

// SetInitiator points to an initiator NUMA node that has the best performance
// (the lowest latency or largest bandwidth) to the NUMA node. Note that this
// option can be set only when the machine property WithHMAT is enabled.
func (n *Node) SetInitiator(initiator int) *Node {
	n.properties = append(n.properties, queso.NewProperty("initiator", initiator))
	return n
}
