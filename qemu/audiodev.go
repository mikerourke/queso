package qemu

import (
	"fmt"

	"github.com/mikerourke/queso"
)

// Direction is used to qualify an audio device property.
type Direction string

const (
	// Input indicates that the property should be applied to the audio device's
	// input.
	Input Direction = "in"

	// Output indicates that the property should be applied to the audio device's
	// output.
	Output Direction = "out"
)

// NewAudioDevOption returns an Option instance that can be used to specify
// audio devices.
func NewAudioDevOption(name string, id string, properties ...*AudioDevProperty) *queso.Option {
	props := []*queso.Property{{"id", id}}

	for _, property := range properties {
		props = append(props, property.Property)
	}

	return queso.NewOption("audiodev", name, props...)
}

// AudioDriver adds a new audio backend driver identified by name and id. There are
// global and driver specific properties. Some values can be set differently for
// input and output.
func AudioDriver(name string, id string, properties ...*AudioDevProperty) *queso.Option {
	return NewAudioDevOption(name, id, properties...)
}

// AudioNone creates a dummy backend that discards all outputs. This backend has no
// backend specific properties.
func AudioNone(id string, properties ...*AudioDevProperty) *queso.Option {
	return NewAudioDevOption("none", id, properties...)
}

// AudioALSA creates backend using the AudioALSA. This backend is only available on Linux.
func AudioALSA(id string, properties ...*AudioDevProperty) *queso.Option {
	return NewAudioDevOption("alsa", id, properties...)
}

// AudioCoreAudio creates a backend using Apple's Core Audio. This backend is only
// available on macOS and only supports playback.
func AudioCoreAudio(id string, properties ...*AudioDevProperty) *queso.Option {
	return NewAudioDevOption("coreaudio", id, properties...)
}

// AudioDirectSound creates a backend using Microsoft's DirectSound. This
// backend is only available on Windows and only supports playback.
func AudioDirectSound(id string, properties ...*AudioDevProperty) *queso.Option {
	return NewAudioDevOption("dsound", id, properties...)
}

// AudioOSS creates a backend using OSS. This backend is available on most
// Unix-like systems.
func AudioOSS(id string, properties ...*AudioDevProperty) *queso.Option {
	return NewAudioDevOption("oss", id, properties...)
}

// AudioPulseAudio creates a backend using PulseAudio. This backend is available on
// most systems.
func AudioPulseAudio(id string, properties ...*AudioDevProperty) *queso.Option {
	return NewAudioDevOption("pa", id, properties...)
}

// AudioSDL creates a backend using SDL. This backend is available on most systems, but
// you should use your platform's native backend if possible.
func AudioSDL(id string, properties ...*AudioDevProperty) *queso.Option {
	return NewAudioDevOption("sdl", id, properties...)
}

// AudioSPICE creates a backend that sends audio through SPICE. This backend requires
// SPICE and automatically selected in that case, so usually you can ignore
// this option. This backend has no backend specific properties.
func AudioSPICE(id string) *queso.Option {
	return NewAudioDevOption("spice", id)
}

// AudioWAV creates a backend that writes audio to a WAV file specified by the path
// parameter. If path is an empty string, defaults to `qemu.wav`.
func AudioWAV(id string, path string) *queso.Option {
	return NewAudioDevOption("wav", id, NewAudioDevProperty("path", path))
}

// AudioDevProperty represents a property that can be used with an audio device
// option.
type AudioDevProperty struct {
	*queso.Property
}

// NewAudioDevProperty returns a new instance of an AudioDevProperty.
func NewAudioDevProperty(key string, value interface{}) *AudioDevProperty {
	return &AudioDevProperty{
		Property: &queso.Property{key, value},
	}
}

func newAudioDevDirectionProperty(direction Direction, key string, value interface{}) *AudioDevProperty {
	return NewAudioDevProperty(fmt.Sprintf("%s.%s", direction, key), value)
}

// WithTimerPeriod sets the timer period used by the audio subsystem in microseconds.
// Default is 10000 (10 ms).
func WithTimerPeriod(microseconds int) *AudioDevProperty {
	return NewAudioDevProperty("timer-period", microseconds)
}

// IsMixingEngine enables/disables QEMU's mixing engine to mix all streams inside QEMU
// and convert audio formats when not supported by the backend. When disabled, IsFixedSettings
// must be disabled too. Note that disabling this option means that the selected
// backend must support multiple streams and the audio formats used by the virtual
// cards, otherwise you'll get no sound. It's not recommended disabling this option
// unless you want to use 5.1 or 7.1 audio, as mixing engine only supports mono
// and stereo audio. Default is enabled.
func IsMixingEngine(direction Direction, enabled bool) *AudioDevProperty {
	return newAudioDevDirectionProperty(direction, "mixing-engine", enabled)
}

// IsFixedSettings enables/disables fixed settings for host audio. When disabled, it
// will change based on how the guest opens the sound card. In this case you must
// not specify frequency, channels or format. Default is enabled.
func IsFixedSettings(direction Direction, enabled bool) *AudioDevProperty {
	return newAudioDevDirectionProperty(direction, "fixed-settings", enabled)
}

// WithFrequency specifies the frequency to use when using IsFixedSettings.
// Default is 44100Hz.
func WithFrequency(direction Direction, frequency string) *AudioDevProperty {
	return newAudioDevDirectionProperty(direction, "frequency", frequency)
}

// WithChannels specifies the number of channels to use when using IsFixedSettings.
// Default is 2 (stereo).
func WithChannels(direction Direction, count int) *AudioDevProperty {
	return newAudioDevDirectionProperty(direction, "channels", count)
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

// WithSampleFormat specifies the sample format to use when using IsFixedSettings.
func WithSampleFormat(direction Direction, format SampleFormat) *AudioDevProperty {
	return newAudioDevDirectionProperty(direction, "format", format)
}

// WithVoices specifies the number of voices to use. Default is 1.
func WithVoices(direction Direction, count int) *AudioDevProperty {
	return newAudioDevDirectionProperty(direction, "voices", count)
}

// WithAudioBufferLength sets the size of the buffer in microseconds.
func WithAudioBufferLength(direction Direction, microseconds int) *AudioDevProperty {
	return newAudioDevDirectionProperty(direction, "buffer-length", microseconds)
}

// WithALSADevice specifies the ALSA device to use for input and/or output.
// Default is "default".
func WithALSADevice(direction Direction, device string) *AudioDevProperty {
	return newAudioDevDirectionProperty(direction, "dev", device)
}

// WithOSSDevice specifies the OSS device to use for input and/or output.
// Default is `/dev/dsp`.
func WithOSSDevice(direction Direction, device string) *AudioDevProperty {
	return newAudioDevDirectionProperty(direction, "dev", device)
}

// WithPeriodLength sets the period length in microseconds.
func WithPeriodLength(direction Direction, microseconds int) *AudioDevProperty {
	return newAudioDevDirectionProperty(direction, "period-length", microseconds)
}

// IsTryPoll will attempt to use poll mode with the device if enabled.
// Default is enabled.
func IsTryPoll(direction Direction, enabled bool) *AudioDevProperty {
	return newAudioDevDirectionProperty(direction, "try-poll", enabled)
}

// WithThreshold specifies the threshold (in microseconds) when playback starts.
// Default is 0.
func WithThreshold(direction Direction, microseconds int) *AudioDevProperty {
	return newAudioDevDirectionProperty(direction, "threshold", microseconds)
}

// WithAudioBufferCount sets the count of the buffers.
func WithAudioBufferCount(direction Direction, count int) *AudioDevProperty {
	return newAudioDevDirectionProperty(direction, "buffer-count", count)
}

// WithAudioLatency adds extra microseconds of latency to playback.
func WithAudioLatency(microseconds int) *AudioDevProperty {
	return NewAudioDevProperty("latency", microseconds)
}

// IsTryMMAP tries using memory mapped device access if enabled. Default is disabled.
func IsTryMMAP(enabled bool) *AudioDevProperty {
	return NewAudioDevProperty("try-mmap", enabled)
}

// IsExclusive specifies whether to open the device in exclusive mode (VMIX won't
// work in this case). Default is disabled.
func IsExclusive(enabled bool) *AudioDevProperty {
	return NewAudioDevProperty("exclusive", enabled)
}

// WithDSPPolicy sets the timing policy (between 0 and 10, where smaller number means
// smaller latency but higher CPU usage). Use -1 to use buffer sizes specified by
// WithAudioBufferCount. This option is ignored if you do not have OSS 4. Default is 5.
func WithDSPPolicy(policy int) *AudioDevProperty {
	return NewAudioDevProperty("dsp-policy", policy)
}

// WithPulseAudioServer sets the PulseAudio server to connect to.
func WithPulseAudioServer(server string) *AudioDevProperty {
	return NewAudioDevProperty("server", server)
}

// WithPulseAudioLatency is the desired latency in microseconds for the PulseAudio
// server.
func WithPulseAudioLatency(direction Direction, microseconds int) *AudioDevProperty {
	return newAudioDevDirectionProperty(direction, "latency", microseconds)
}

// WithPulseAudioSink specifies the source/sink for recording/playback.
func WithPulseAudioSink(direction Direction, name string) *AudioDevProperty {
	return newAudioDevDirectionProperty(direction, "name", name)
}
