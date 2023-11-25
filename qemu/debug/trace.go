package debug

import "github.com/mikerourke/queso"

// Trace traces events matching a pattern or from a file and optionally logs the
// output to a specified file.
type Trace struct {
	properties []*queso.Property
}

// NewTrace returns a new instance of [Trace].
//
//	qemu-system-* -trace
func NewTrace() *Trace {
	return &Trace{
		properties: make([]*queso.Property, 0),
	}
}

func (t *Trace) option() *queso.Option {
	return queso.NewOption("trace", "", t.properties...)
}

// SetProperty is used to add arbitrary properties to the [Trace].
func (t *Trace) SetProperty(key string, value interface{}) *Trace {
	t.properties = append(t.properties, queso.NewProperty(key, value))
	return t
}

// MatchPattern immediately enables events matching pattern (either event name or
// a globbing pattern). This property is only available if QEMU has been compiled
// with the "simple", "log", or "ftrace" tracing backend.
//
//	qemu-system-* -trace enable=pattern
func (t *Trace) MatchPattern(pattern string) *Trace {
	t.properties = append(t.properties, queso.NewProperty("enable", pattern))
	return t
}

// EnableEventsInFile immediately enable events listed in a file. The file
// must contain one event name (as listed in the trace-events-all file) per
// line; globbing patterns are accepted too. This property is only available if
// QEMU has been compiled with the "simple", "log", or "ftrace" tracing backend.
//
//	qemu-system-* -trace events=file
func (t *Trace) EnableEventsInFile(file string) *Trace {
	t.properties = append(t.properties, queso.NewProperty("events", file))
	return t
}

// SetOutputFile logs output traces to a file. This property is only
// available if QEMU has been compiled with the "simple" tracing backend.
//
//	qemu-system-* -trace file=file
func (t *Trace) SetOutputFile(file string) *Trace {
	t.properties = append(t.properties, queso.NewProperty("file", file))
	return t
}
