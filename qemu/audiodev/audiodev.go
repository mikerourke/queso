// Package audiodev is used to manage audio backend drivers.
package audiodev

import (
	"fmt"

	"github.com/mikerourke/queso/internal/cli"
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

type Backend struct {
	DriverName string
	properties []*cli.Property
}

func NewBackend(driverName string, id string) *Backend {
	return &Backend{
		DriverName: driverName,
		properties: []*cli.Property{
			cli.NewProperty("id", id),
		},
	}
}

func (b *Backend) option() *cli.Option {
	return cli.NewOption("audiodev", b.DriverName, b.properties...)
}

// SetBufferLength sets the size of the buffer in microseconds.
//
//	qemu-system-* -audiodev pa,in|out.buffer-length=usecs
func (b *Backend) SetBufferLength(direction Direction, usecs int) *Backend {
	b.properties = append(b.properties, newDirectionProperty("buffer-length", direction, usecs))
	return b
}

// SetChannels specifies the number of channels to use when [Backend.ToggleFixedSettings]
// was called with true. The default is 2 (stereo).
//
//	qemu-system-* -audiodev pa,in|out.channels=channels
func (b *Backend) SetChannels(direction Direction, count int) *Backend {
	b.properties = append(b.properties, newDirectionProperty("channels", direction, count))
	return b
}

// SetFrequency specifies the frequency to use when [Backend.ToggleFixedSettings]
// was called with true. The default is 44100Hz.
//
//	qemu-system-* -audiodev pa,in|out.frequency=frequency
func (b *Backend) SetFrequency(direction Direction, frequency string) *Backend {
	b.properties = append(b.properties, newDirectionProperty("frequency", direction, frequency))
	return b
}

// SampleFormat represents the sample format (number of bits per sample) to use
// in the WithSampleFormat property.
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
//	qemu-system-* -audiodev pa,in|out.format=format
func (b *Backend) SetSampleFormat(direction Direction, format SampleFormat) *Backend {
	b.properties = append(b.properties, newDirectionProperty("format", direction, format))
	return b
}

// SetTimerPeriod sets the timer period used by the audio subsystem in
// microseconds. The default is 10,000 (10 ms).
//
//	qemu-system-* -audiodev pa,timer-period=period
func (b *Backend) SetTimerPeriod(period int) *Backend {
	b.properties = append(b.properties, cli.NewProperty("timer-period", period))
	return b
}

// SetVoices specifies the number of voices. The default is 1.
//
//	qemu-system-* -audiodev pa,in|out.voices=voices
func (b *Backend) SetVoices(direction Direction, count int) *Backend {
	b.properties = append(b.properties, newDirectionProperty("voices", direction, count))
	return b
}

// ToggleFixedSettings enables/disables fixed settings for host audio any backend.
// When disabled, it will change based on how the guest opens the sound card. In
// this case you must not specify frequency, channels or format. This is enabled
// by default.
//
//	qemu-system-* -audiodev pa,in|out.fixed-settings=on|off
func (b *Backend) ToggleFixedSettings(direction Direction, enabled bool) *Backend {
	b.properties = append(b.properties, newDirectionProperty("fixed-settings", direction, enabled))
	return b
}

// ToggleMixingEngine enables/disables QEMU's mixing engine to mix all streams inside QEMU
// and convert audio formats when not supported by the backend. When disabled,
// [Backend.ToggleFixedSettings] must be disabled too. Note that disabling this option means
// that the selected backend must support multiple streams and the audio formats
// used by the virtual cards, otherwise you'll get no sound. It's not
// recommended disabling this option unless you want to use 5.1 or 7.1 audio, as
// mixing engine only supports mono and stereo audio. This property is enabled by default.
//
//	qemu-system-* -audiodev pa,in|out.mixing-engine=on|off
func (b *Backend) ToggleMixingEngine(direction Direction, enabled bool) *Backend {
	b.properties = append(b.properties, newDirectionProperty("mixing-engine", direction, enabled))
	return b
}

func newDirectionProperty(key string, direction Direction, value interface{}) *cli.Property {
	return cli.NewProperty(fmt.Sprintf("%s.%s", direction, key), value)
}
