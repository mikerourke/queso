package debug

import "github.com/mikerourke/queso"

// Trace traces events matching a pattern or from a file and optionally logs the
// output to a specified file.
type Trace struct {
	properties []*queso.Property
}

func NewTrace() *Trace {
	return &Trace{
		properties: make([]*queso.Property, 0),
	}
}

func (t *Trace) option() *queso.Option {
	return queso.NewOption("trace", "", t.properties...)
}

// MatchPattern immediately enables events matching pattern (either event name or
// a globbing pattern) for a Trace. This property is only available if QEMU
// has been compiled with the "simple", "log", or "ftrace" tracing backend.
func (t *Trace) MatchPattern(pattern string) *Trace {
	t.properties = append(t.properties, queso.NewProperty("enable", pattern))
	return t
}

// EnableEventsInFile immediately enable events listed in file for a Trace. The file
// must contain one event name (as listed in the trace-events-all file) per
// line; globbing patterns are accepted too. This property is only available if
// QEMU has been compiled with the "simple", "log", or "ftrace" tracing backend.
func (t *Trace) EnableEventsInFile(file string) *Trace {
	t.properties = append(t.properties, queso.NewProperty("events", file))
	return t
}

// SetOutputFile logs output traces to file for a Trace. This property is only
// available if QEMU has been compiled with the "simple" tracing backend.
func (t *Trace) SetOutputFile(file string) *Trace {
	t.properties = append(t.properties, queso.NewProperty("file", file))
	return t
}
