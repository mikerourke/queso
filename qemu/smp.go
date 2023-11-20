package qemu

import "github.com/mikerourke/queso"

// SMP simulates a SMP system with the count of CPUs initially present on the
// machine type board.
type SMP struct {
	properties []*queso.Property
}

// NewSMP returns a new SMP instance that can be used with QEMU.
func NewSMP() *SMP {
	return &SMP{}
}

func (s *SMP) option() *queso.Option {
	if len(s.properties) == 0 {
		panic("either SetCPUCount or at least one of the topology parameters must be specified for UseSMP")
	}

	return queso.NewOption("smp", "", s.properties...)
}

// SetCPUCount specifies the initial CPU count to use. If omitted, the maximum
// number of CPUs will be used.
func (s *SMP) SetCPUCount(count int) *SMP {
	s.properties = append(s.properties, queso.NewProperty("cpus", count))
	return s
}

// SetMaxCPUs enables further CPUs to be added at runtime. If omitted, the maximum
// number of CPUs will be calculated from the provided topology members and the
// initial CPU count will match the maximum number.
func (s *SMP) SetMaxCPUs(count int) *SMP {
	s.properties = append(s.properties, queso.NewProperty("maxcpus", count))
	return s
}

// SetSocketCount specifies the count of sockets to use.
func (s *SMP) SetSocketCount(count int) *SMP {
	s.properties = append(s.properties, queso.NewProperty("sockets", count))
	return s
}

// SetDieCount specifies the count of dies per socket to use.
func (s *SMP) SetDieCount(count int) *SMP {
	s.properties = append(s.properties, queso.NewProperty("dies", count))
	return s
}

// SetCoreCount specifies the count of cores per die to use.
func (s *SMP) SetCoreCount(count int) *SMP {
	s.properties = append(s.properties, queso.NewProperty("cores", count))
	return s
}

// SetThreadCount specifies the count of threads per core to use.
func (s *SMP) SetThreadCount(count int) *SMP {
	s.properties = append(s.properties, queso.NewProperty("threads", count))
	return s
}
