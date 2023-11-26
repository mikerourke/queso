package numa

import "github.com/mikerourke/queso"

type CPU struct {
	properties []*queso.Property
}

func NewCPU(node *Node) *CPU {
	return &CPU{
		properties: []*queso.Property{
			queso.NewProperty("node-id", node.ID),
		},
	}
}

// Option returns the invoked option that gets converted to an argument when
// passed to QEMU.
func (cpu *CPU) Option() *queso.Option {
	return queso.NewOption("numa", "cpu", cpu.properties...)
}

// SetSocketID specifies the CPU socket ID for a CPU.
func (cpu *CPU) SetSocketID(id int) *CPU {
	cpu.properties = append(cpu.properties, queso.NewProperty("socket-id", id))
	return cpu
}

// SetCoreID specifies the CPU core ID for a CPU.
func (cpu *CPU) SetCoreID(id int) *CPU {
	cpu.properties = append(cpu.properties, queso.NewProperty("core-id", id))
	return cpu
}

// SetThreadID specifies the CPU thread ID for a CPU.
func (cpu *CPU) SetThreadID(id int) *CPU {
	cpu.properties = append(cpu.properties, queso.NewProperty("thread-id", id))
	return cpu
}
