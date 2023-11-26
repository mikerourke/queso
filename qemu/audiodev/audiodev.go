// Package audiodev is used to manage audio backend drivers.
package audiodev

import (
	"fmt"

	"github.com/mikerourke/queso"
)

// Direction is used to qualify an audio device property for a backend.
type Direction string

const (
	// Input indicates that the property should be applied to the audio device
	// backend's input.
	Input Direction = "in"

	// Output indicates that the property should be applied to the audio device
	// backend's output.
	Output Direction = "out"
)

// Backend represents a generic audio device backend. If possible, try to
// use specified backends (e.g. [ALSABackend]).
type Backend struct {
	*queso.Entity
}

// New returns a new instance of a generic audio [Backend]. driverName
// is the name of the driver associated with the backend. id is a unique
// identifier for the backend.
//
//	qemu-system-* -audiodev <driverName>,id=id
func New(driverName string, id string) *Backend {
	return &Backend{
		queso.NewEntity("audiodev", driverName).SetProperty("id", id),
	}
}

// SetDirectionProperty sets a specific direction property (i.e. [Input] or [Output]).
func (b *Backend) SetDirectionProperty(direction Direction, key string, value interface{}) *Backend {
	b.SetProperty(fmt.Sprintf("%s.%s", direction, key), value)
	return b
}

// SetBufferLength sets the size of the buffer in microseconds.
//
//	qemu-system-* -audiodev <name>,in|out.buffer-length=usecs
func (b *Backend) SetBufferLength(direction Direction, length int) *Backend {
	b.SetDirectionProperty(direction, "buffer-length", length)
	return b
}

// SetChannels specifies the number of channels to use when [Backend.ToggleFixedSettings]
// was called with true. The default is 2 (stereo).
//
//	qemu-system-* -audiodev <name>,in|out.channels=channels
func (b *Backend) SetChannels(direction Direction, count int) *Backend {
	b.SetDirectionProperty(direction, "channels", count)
	return b
}

// SetFrequency specifies the frequency to use when [Backend.ToggleFixedSettings]
// was called with true. The default is 44100Hz.
//
//	qemu-system-* -audiodev <name>,in|out.frequency=frequency
func (b *Backend) SetFrequency(direction Direction, frequency string) *Backend {
	b.SetDirectionProperty(direction, "frequency", frequency)
	return b
}

// SampleFormat represents the sample format (number of bits per sample) to use
// in the [Backend.SetSampleFormat] method.
// See https://www.metadata2go.com/file-info/sample-fmt for more details.
type SampleFormat string

const (
	// S8 represents signed 8 bits sample format.
	S8 SampleFormat = "s8"

	// S16 represents signed 16 bits sample format.
	S16 SampleFormat = "s16"

	// S32 represents signed 32 bits sample format.
	S32 SampleFormat = "s32"

	// U8 represents unsigned 8 bits sample format.
	U8 SampleFormat = "u8"

	// U16 represents unsigned 16 bits sample format.
	U16 SampleFormat = "u16"

	// U32 represents unsigned 32 bits sample format.
	U32 SampleFormat = "u32"

	// F32 represents floating 32 bits sample format.
	F32 SampleFormat = "f32"
)

// SetSampleFormat specifies the sample format to use when [Backend.ToggleFixedSettings]
// was called with true.
//
//	qemu-system-* -audiodev <name>,in|out.format=format
func (b *Backend) SetSampleFormat(direction Direction, format SampleFormat) *Backend {
	b.SetDirectionProperty(direction, "format", format)
	return b
}

// SetTimerPeriod sets the timer period used by the audio subsystem in
// microseconds. The default is 10,000 (10 ms).
//
//	qemu-system-* -audiodev <name>,timer-period=period
func (b *Backend) SetTimerPeriod(period int) *Backend {
	b.SetProperty("timer-period", period)
	return b
}

// SetVoices specifies the number of voices. The default is 1.
//
//	qemu-system-* -audiodev <name>,in|out.voices=voices
func (b *Backend) SetVoices(direction Direction, count int) *Backend {
	b.SetDirectionProperty(direction, "voices", count)
	return b
}

// ToggleFixedSettings enables/disables fixed settings for host audio any backend.
// When disabled, it will change based on how the guest opens the sound card. In
// this case you must not specify frequency, channels or format. This is enabled
// by default.
//
//	qemu-system-* -audiodev <name>,in|out.fixed-settings=on|off
func (b *Backend) ToggleFixedSettings(direction Direction, enabled bool) *Backend {
	b.SetDirectionProperty(direction, "fixed-settings", enabled)
	return b
}

// ToggleMixingEngine enables/disables QEMU's mixing engine to mix all streams inside
// QEMU and convert audio formats when not supported by the backend.
// When disabled, [Backend.ToggleFixedSettings] must be disabled too.
//
// Note that disabling this option means that the selected backend must support
// multiple streams and the audio formats used by the virtual cards, otherwise
// you'll get no sound. It's not recommended disabling this option unless you
// want to use 5.1 or 7.1 audio, as mixing engine only supports mono and stereo
// audio. This property is enabled by default.
//
//	qemu-system-* -audiodev <name>,in|out.mixing-engine=on|off
func (b *Backend) ToggleMixingEngine(direction Direction, enabled bool) {
	b.SetDirectionProperty(direction, "mixing-engine", enabled)
}
