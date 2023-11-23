package accel

import "github.com/mikerourke/queso/qemu/cli"

// TCGAccelerator represents an accelerator using tiny code generation (TCG).
type TCGAccelerator struct {
	*Accelerator
}

// NewTCGAccelerator returns a new instace of TCGAccelerator.
//
//	qemu-system-* -accel tcg
func NewTCGAccelerator() *TCGAccelerator {
	return &TCGAccelerator{
		NewAccelerator(TCG),
	}
}

// ThreadingOption represents the type of TCG threads to use.
type ThreadingOption string

const (
	// SingleThreaded indicates that a single thread should be used with TCG.
	SingleThreaded ThreadingOption = "single"

	// MultiThreaded indicates that multiple threads should be used with TCG.
	MultiThreaded ThreadingOption = "multi"
)

// SetThreads controls number of TCG threads. When the TCG is multithreaded,
// there will be one thread per vCPU therefore taking advantage of additional
// host cores. The default is to enable multi-threading where both the back-end
// and front-ends support it and no incompatible TCG features have been enabled
// (e.g. icount/replay).
//
//	qemu-system-* -accel tcg thread=single|multi
func (a *TCGAccelerator) SetThreads(option ThreadingOption) *TCGAccelerator {
	a.properties = append(a.properties, cli.NewProperty("thread", option))
	return a
}

// SetTranslationBlockCacheSize controls the size (in MiB) of the TCG
// translation block cache.
//
//	qemu-system-* -accel tcg tb-size=mb
func (a *TCGAccelerator) SetTranslationBlockCacheSize(mb int) *TCGAccelerator {
	a.properties = append(a.properties, cli.NewProperty("tb-size", mb))
	return a
}

// ToggleOneInstructionPerTranslation makes the TCG accelerator put only one guest
// instruction into each translation block. This slows down emulation a lot, but
// can be useful in some situations, such as when trying to analyse the logs
// produced during debugging.
//
//	qemu-system-* -accel tcg one-insn-per-tb=on|off
func (a *TCGAccelerator) ToggleOneInstructionPerTranslation(enabled bool) *TCGAccelerator {
	a.properties = append(a.properties, cli.NewProperty("one-insn-per-tb", enabled))
	return a
}

// ToggleSplitWX controls the use of split w^x mapping for the TCG code generation
// buffer. Some operating systems require this to be enabled, and in
// such a case this will default to true. On other operating systems, this will
// default to false, but one may enable this for testing or debugging.
//
//	qemu-system-* -accel tcg split-wx=on|off
func (a *TCGAccelerator) ToggleSplitWX(enabled bool) *TCGAccelerator {
	a.properties = append(a.properties, cli.NewProperty("split-wx", enabled))
	return a
}
