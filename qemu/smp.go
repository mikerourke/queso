package qemu

import (
	"strconv"
)

// SMP simulates a SMP system with the count of CPUs initially present on the
// machine type board.
type SMP struct {
	CPUCount   int
	properties []*cli.Property
}

// NewSMP returns a new SMP instance that can be used with QEMU. The specified
// cpuCount represents the count of CPUs to use in the SMP.
func NewSMP(cpuCount int) *SMP {
	if cpuCount < 1 {
		panic("CPU count must be at least 1 for the SMP")
	}

	return &SMP{
		CPUCount:   cpuCount,
		properties: make([]*cli.Property, 0),
	}
}

func (s *SMP) option() *cli.Option {
	return cli.NewOption("smp", strconv.Itoa(s.CPUCount), s.properties...)
}

// SetCPUCount specifies the initial CPU count to use. If omitted, the maximum
// number of CPUs will be used.
func (s *SMP) SetCPUCount(count int) *SMP {
	s.properties = append(s.properties, cli.NewProperty("cpus", count))
	return s
}

// SetMaxCPUs enables further CPUs to be added at runtime. If omitted, the maximum
// number of CPUs will be calculated from the provided topology members and the
// initial CPU count will match the maximum number.
func (s *SMP) SetMaxCPUs(count int) *SMP {
	s.properties = append(s.properties, cli.NewProperty("maxcpus", count))
	return s
}

// SetSocketCount specifies the count of sockets to use.
func (s *SMP) SetSocketCount(count int) *SMP {
	s.properties = append(s.properties, cli.NewProperty("sockets", count))
	return s
}

// SetDieCount specifies the count of dies per socket to use.
func (s *SMP) SetDieCount(count int) *SMP {
	s.properties = append(s.properties, cli.NewProperty("dies", count))
	return s
}

// SetCoreCount specifies the count of cores per die to use.
func (s *SMP) SetCoreCount(count int) *SMP {
	s.properties = append(s.properties, cli.NewProperty("cores", count))
	return s
}

// SetThreadCount specifies the count of threads per core to use.
func (s *SMP) SetThreadCount(count int) *SMP {
	s.properties = append(s.properties, cli.NewProperty("threads", count))
	return s
}
