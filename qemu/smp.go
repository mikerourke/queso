package qemu

import "github.com/mikerourke/queso"

// SMP simulates a SMP system with the count of CPUs initially present on the
// machine type board.
func SMP(properties ...*SMPProperty) *queso.Option {
	if len(properties) == 0 {
		panic("either WithCPUCount or at least one of the topology parameters must be specified for SMP")
	}

	props := make([]*queso.Property, 0)

	for _, property := range properties {
		props = append(props, property.Property)
	}

	return queso.NewOption("accel", "", props...)
}

// SMPProperty represents a property that can be used with the SMP option.
type SMPProperty struct {
	*queso.Property
}

// NewSMPProperty returns a new instance of SMPProperty.
func NewSMPProperty(key string, value interface{}) *SMPProperty {
	return &SMPProperty{
		Property: queso.NewProperty(key, value),
	}
}

// WithCPUCount specifies the initial CPU count to use. If omitted, the maximum
// number of CPUs will be used.
func WithCPUCount(count int) *SMPProperty {
	return NewSMPProperty("cpus", count)
}

// WithMaxCPUs enables further CPUs to be added at runtime. If omitted, the maximum
// number of CPUs will be calculated from the provided topology members and the
// initial CPU count will match the maximum number.
func WithMaxCPUs(count int) *SMPProperty {
	return NewSMPProperty("maxcpus", count)
}

// WithSockets specifies the count of sockets to use.
func WithSockets(count int) *SMPProperty {
	return NewSMPProperty("sockets", count)
}

// WithDies specifies the count of dies per socket to use.
func WithDies(count int) *SMPProperty {
	return NewSMPProperty("dies", count)
}

// WithCores specifies the count of cores per die to use.
func WithCores(count int) *SMPProperty {
	return NewSMPProperty("cores", count)
}

// WithThreads specifies the count of threads per core to use.
func WithThreads(count int) *SMPProperty {
	return NewSMPProperty("threads", count)
}
