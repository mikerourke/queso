// Package audiodev is used to manage audio backend drivers.
package audiodev

import (
	"fmt"

	"github.com/mikerourke/queso"
)

const (
	BackendTypeNone        = "none"
	BackendTypeALSA        = "alsa"
	BackendTypeCoreAudio   = "coreaudio"
	BackendTypeDirectSound = "dsound"
	BackendTypeOSS         = "oss"
	BackendTypePulseAudio  = "pa"
	BackendTypeSDL         = "sdl"
	BackendTypeSpice       = "spice"
	BackendTypeWAV         = "wav"
)

// Direction is used to qualify an audio device property for a Backend.
type Direction string

const (
	// Input indicates that the property should be applied to the audio device
	// Backend's input.
	Input Direction = "in"

	// Output indicates that the property should be applied to the audio device
	// Backend's output.
	Output Direction = "out"
)

// Backend adds a new audio backend driver identified by backendType and id.
// There are global and driver specific properties. Some values can be set
// differently for input and output.
func Backend(backendType string, id string, properties ...*Property) *queso.Option {
	props := []*queso.Property{{"id", id}}

	for _, property := range properties {
		props = append(props, property.Property)
	}

	return queso.NewOption("audiodev", backendType, props...)
}

// DummyBackend creates a dummy backend that discards all outputs. This backend
// has no backend specific properties.
func DummyBackend(id string, properties ...*Property) *queso.Option {
	return Backend(BackendTypeNone, id, properties...)
}

// ALSABackend creates backend using the ALSABackend. This backend is only
// available on Linux.
func ALSABackend(id string, properties ...*Property) *queso.Option {
	return Backend(BackendTypeALSA, id, properties...)
}

// CoreAudioBackend creates a backend using Apple's Core Audio. This backend is
// only available on macOS and only supports playback.
func CoreAudioBackend(id string, properties ...*Property) *queso.Option {
	return Backend(BackendTypeCoreAudio, id, properties...)
}

// DirectSoundBackend creates a backend using Microsoft's DirectSoundBackend.
// This backend is only available on Windows and only supports playback.
func DirectSoundBackend(id string, properties ...*Property) *queso.Option {
	return Backend(BackendTypeDirectSound, id, properties...)
}

// OSSBackend creates a backend using OSSBackend. This backend is available on
// most Unix-like systems.
func OSSBackend(id string, properties ...*Property) *queso.Option {
	return Backend(BackendTypeOSS, id, properties...)
}

// PulseAudioBackend creates a backend using PulseAudioBackend. This backend is
// available on most systems.
func PulseAudioBackend(id string, properties ...*Property) *queso.Option {
	return Backend(BackendTypePulseAudio, id, properties...)
}

// SDLBackend creates a backend using SDLBackend. This backend is available on
// most systems, but you should use your platform's native backend if possible.
func SDLBackend(id string, properties ...*Property) *queso.Option {
	return Backend(BackendTypeSDL, id, properties...)
}

// SpiceBackend creates a backend that sends audio through SpiceBackend. This
// backend requires SPICE and automatically selected in that case, so usually
// you can ignore this option. This backend has no backend specific properties.
func SpiceBackend(id string) *queso.Option {
	return Backend(BackendTypeSpice, id)
}

// WAVBackend creates a backend that writes audio to a WAVBackend file specified
// by the path parameter. If path is an empty string, defaults to `qemu.wav`.
func WAVBackend(id string, path string) *queso.Option {
	return Backend(BackendTypeWAV, id, NewProperty("path", path))
}

// Property represents a property that can be used with an audio device Backend
// option.
type Property struct {
	*queso.Property
}

// NewProperty returns a new instance of Property.
func NewProperty(key string, value interface{}) *Property {
	return &Property{
		Property: queso.NewProperty(key, value),
	}
}

func newDirectionProperty(direction Direction, key string, value interface{}) *Property {
	return NewProperty(fmt.Sprintf("%s.%s", direction, key), value)
}

// WithTimerPeriod sets the timer period used by the audio subsystem in
// microseconds for use with any Backend. The default is 10000 (10 ms).
func WithTimerPeriod(microseconds int) *Property {
	return NewProperty("timer-period", microseconds)
}

// IsMixingEngine enables/disables QEMU's mixing engine to mix all streams inside QEMU
// and convert audio formats when not supported by the backend. When disabled,
// IsFixedSettings must be disabled too. Note that disabling this option means
// that the selected backend must support multiple streams and the audio formats
// used by the virtual cards, otherwise you'll get no sound. It's not
// recommended disabling this option unless you want to use 5.1 or 7.1 audio, as
// mixing engine only supports mono and stereo audio. This property is enabled by default.
func IsMixingEngine(direction Direction, enabled bool) *Property {
	return newDirectionProperty(direction, "mixing-engine", enabled)
}

// IsFixedSettings enables/disables fixed settings for host audio any Backend.
// When disabled, it will change based on how the guest opens the sound card. In
// this case you must not specify frequency, channels or format. This is enabled
// by default.
func IsFixedSettings(direction Direction, enabled bool) *Property {
	return newDirectionProperty(direction, "fixed-settings", enabled)
}

// WithFrequency specifies the frequency to use when using IsFixedSettings any
// Backend. The default is 44100Hz.
func WithFrequency(direction Direction, frequency string) *Property {
	return newDirectionProperty(direction, "frequency", frequency)
}

// WithChannels specifies the number of channels to use when using IsFixedSettings
// for any Backend. The default is 2 (stereo).
func WithChannels(direction Direction, count int) *Property {
	return newDirectionProperty(direction, "channels", count)
}

// SampleFormat represents the sample format (number of bits per sample) to use
// in the WithSampleFormat property for any Backend.
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

// WithSampleFormat specifies the sample format to use when using IsFixedSettings
// for any Backend.
func WithSampleFormat(direction Direction, format SampleFormat) *Property {
	return newDirectionProperty(direction, "format", format)
}

// WithVoices specifies the number of voices to use for any Backend. The default
// is 1.
func WithVoices(direction Direction, count int) *Property {
	return newDirectionProperty(direction, "voices", count)
}

// WithBufferLength sets the size of the buffer in microseconds for any Backend.
func WithBufferLength(direction Direction, microseconds int) *Property {
	return newDirectionProperty(direction, "buffer-length", microseconds)
}

// WithALSADevice specifies the ALSABackend device to use for input and/or
// output. The default is "default".
func WithALSADevice(direction Direction, device string) *Property {
	return newDirectionProperty(direction, "dev", device)
}

// WithOSSDevice specifies the OSSBackend device to use for input and/or output.
// The default is `/dev/dsp`.
func WithOSSDevice(direction Direction, device string) *Property {
	return newDirectionProperty(direction, "dev", device)
}

// WithPeriodLength sets the period length in microseconds for an ALSABackend.
func WithPeriodLength(direction Direction, microseconds int) *Property {
	return newDirectionProperty(direction, "period-length", microseconds)
}

// IsTryPoll will attempt to use poll mode with the device if enabled for an
// ALSABackend or OSSBackend. This property is enabled by default.
func IsTryPoll(direction Direction, enabled bool) *Property {
	return newDirectionProperty(direction, "try-poll", enabled)
}

// WithThreshold specifies the threshold (in microseconds) when playback starts
// for an ALSABackend. The default is 0.
func WithThreshold(direction Direction, microseconds int) *Property {
	return newDirectionProperty(direction, "threshold", microseconds)
}

// WithBufferCount sets the count of the buffers for an OSSBackend or SDLBackend.
func WithBufferCount(direction Direction, count int) *Property {
	return newDirectionProperty(direction, "buffer-count", count)
}

// WithPlaybackLatency adds extra microseconds of latency to playback for
// a DirectSoundBackend.
func WithPlaybackLatency(microseconds int) *Property {
	return NewProperty("latency", microseconds)
}

// IsTryMMAP specifies whether to try using memory mapped device access in an
// OSSBackend. This property is disabled by default.
func IsTryMMAP(enabled bool) *Property {
	return NewProperty("try-mmap", enabled)
}

// IsExclusive specifies whether to open the device in exclusive mode (VMIX won't
// work in this case) for an OSSBackend. This property is disabled by default.
func IsExclusive(enabled bool) *Property {
	return NewProperty("exclusive", enabled)
}

// WithDSPPolicy sets the timing policy (between 0 and 10, where smaller number
// means smaller latency but higher CPU usage) for an OSSBackend. Use -1 to use
// buffer sizes specified by WithBufferCount. This option is ignored if you do
// not have OSSBackend 4. The default is 5.
func WithDSPPolicy(policy int) *Property {
	return NewProperty("dsp-policy", policy)
}

// WithServer sets the PulseAudioBackend server to connect to.
func WithServer(server string) *Property {
	return NewProperty("server", server)
}

// WithServerLatency is the desired latency in microseconds for the
// PulseAudioBackend server.
func WithServerLatency(direction Direction, microseconds int) *Property {
	return newDirectionProperty(direction, "latency", microseconds)
}

// WithSink specifies the source/sink for recording/playback for the
// PulseAudioBackend server.
func WithSink(direction Direction, name string) *Property {
	return newDirectionProperty(direction, "name", name)
}
