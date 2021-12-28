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

// AudioDevProperty represents a property that can be used with the audio device
// options.
type AudioDevProperty struct {
	*queso.Property
}

func newAudioDevProperty(key string, value interface{}) *AudioDevProperty {
	return &AudioDevProperty{
		Property: &queso.Property{key, value},
	}
}

func newAudioDevOption(name string, id string, properties ...*AudioDevProperty) *queso.Option {
	props := make([]*queso.Property, 0)

	for _, property := range properties {
		props = append(props, property.Property)
	}

	opt := queso.NewOption("audiodev", name, props...)
	opt.SetID(id)

	return opt
}

// AudioDriver adds a new audio backend driver identified by name and id. There are
// global and driver specific properties. Some values can be set differently for
// input and output.
func AudioDriver(name string, id string, properties ...*AudioDevProperty) *queso.Option {
	return newAudioDevOption(name, id, properties...)
}

// AudioNone creates a dummy backend that discards all outputs. This backend has no
// backend specific properties.
func AudioNone(id string, properties ...*AudioDevProperty) *queso.Option {
	return newAudioDevOption("none", id, properties...)
}

// AudioALSA creates backend using the AudioALSA. This backend is only available on Linux.
func AudioALSA(id string, properties ...*AudioDevProperty) *queso.Option {
	return newAudioDevOption("alsa", id, properties...)
}

// AudioCoreAudio creates a backend using Apple's Core Audio. This backend is only
// available on macOS and only supports playback.
func AudioCoreAudio(id string, properties ...*AudioDevProperty) *queso.Option {
	return newAudioDevOption("coreaudio", id, properties...)
}

// AudioDirectSound creates a backend using Microsoft's DirectSound. This
// backend is only available on Windows and only supports playback.
func AudioDirectSound(id string, properties ...*AudioDevProperty) *queso.Option {
	return newAudioDevOption("dsound", id, properties...)
}

// AudioOSS creates a backend using OSS. This backend is available on most
// Unix-like systems.
func AudioOSS(id string, properties ...*AudioDevProperty) *queso.Option {
	return newAudioDevOption("oss", id, properties...)
}

// AudioPulseAudio creates a backend using PulseAudio. This backend is available on
// most systems.
func AudioPulseAudio(id string, properties ...*AudioDevProperty) *queso.Option {
	return newAudioDevOption("pa", id, properties...)
}

// AudioSDL creates a backend using SDL. This backend is available on most systems, but
// you should use your platform's native backend if possible.
func AudioSDL(id string, properties ...*AudioDevProperty) *queso.Option {
	return newAudioDevOption("sdl", id, properties...)
}

// AudioSPICE creates a backend that sends audio through SPICE. This backend requires
// SPICE and automatically selected in that case, so usually you can ignore
// this option. This backend has no backend specific properties.
func AudioSPICE(id string) *queso.Option {
	return newAudioDevOption("spice", id)
}

// AudioWAV creates a backend that writes audio to a WAV file specified by the path
// parameter. If path is an empty string, defaults to `qemu.wav`.
func AudioWAV(id string, path string) *queso.Option {
	return newAudioDevOption("wav", id, newAudioDevProperty("path", path))
}

func newDirectionProperty(direction Direction, property string, value interface{}) *AudioDevProperty {
	return newAudioDevProperty(fmt.Sprintf("%s.%s", direction, property), value)
}

// WithTimerPeriod sets the timer period used by the audio subsystem in microseconds.
// Default is 10000 (10 ms).
func WithTimerPeriod(microseconds int) *AudioDevProperty {
	return newAudioDevProperty("timer-period", microseconds)
}

// IsMixingEngine enables/disables QEMU's mixing engine to mix all streams inside QEMU
// and convert audio formats when not supported by the backend. When disabled, IsFixedSettings
// must be disabled too. Note that disabling this option means that the selected
// backend must support multiple streams and the audio formats used by the virtual
// cards, otherwise you'll get no sound. It's not recommended disabling this option
// unless you want to use 5.1 or 7.1 audio, as mixing engine only supports mono
// and stereo audio. Default is enabled.
func IsMixingEngine(direction Direction, enabled bool) *AudioDevProperty {
	return newDirectionProperty(direction, "mixing-engine", enabled)
}

// IsFixedSettings enables/disables fixed settings for host audio. When disabled, it
// will change based on how the guest opens the sound card. In this case you must
// not specify frequency, channels or format. Default is enabled.
func IsFixedSettings(direction Direction, enabled bool) *AudioDevProperty {
	return newDirectionProperty(direction, "fixed-settings", enabled)
}

// WithFrequency specifies the frequency to use when using IsFixedSettings.
// Default is 44100Hz.
func WithFrequency(direction Direction, frequency string) *AudioDevProperty {
	return newDirectionProperty(direction, "frequency", frequency)
}

// WithChannels specifies the number of channels to use when using IsFixedSettings.
// Default is 2 (stereo).
func WithChannels(direction Direction, count int) *AudioDevProperty {
	return newDirectionProperty(direction, "channels", count)
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
	return newDirectionProperty(direction, "format", format)
}

// WithVoices specifies the number of voices to use. Default is 1.
func WithVoices(direction Direction, count int) *AudioDevProperty {
	return newDirectionProperty(direction, "voices", count)
}

// WithBufferLength sets the size of the buffer in microseconds.
func WithBufferLength(direction Direction, microseconds int) *AudioDevProperty {
	return newDirectionProperty(direction, "buffer-length", microseconds)
}

// WithALSADevice specifies the ALSA device to use for input and/or output.
// Default is "default".
func WithALSADevice(direction Direction, device string) *AudioDevProperty {
	return newDirectionProperty(direction, "dev", device)
}

// WithOSSDevice specifies the OSS device to use for input and/or output.
// Default is `/dev/dsp`.
func WithOSSDevice(direction Direction, device string) *AudioDevProperty {
	return newDirectionProperty(direction, "dev", device)
}

// WithPeriodLength sets the period length in microseconds.
func WithPeriodLength(direction Direction, microseconds int) *AudioDevProperty {
	return newDirectionProperty(direction, "period-length", microseconds)
}

// IsTryPoll will attempt to use poll mode with the device if enabled.
// Default is enabled.
func IsTryPoll(direction Direction, enabled bool) *AudioDevProperty {
	return newDirectionProperty(direction, "try-poll", enabled)
}

// WithThreshold specifies the threshold (in microseconds) when playback starts.
// Default is 0.
func WithThreshold(direction Direction, microseconds int) *AudioDevProperty {
	return newDirectionProperty(direction, "threshold", microseconds)
}

// WithBufferCount sets the count of the buffers.
func WithBufferCount(direction Direction, count int) *AudioDevProperty {
	return newDirectionProperty(direction, "buffer-count", count)
}

// WithLatency adds extra microseconds of latency to playback.
func WithLatency(microseconds int) *AudioDevProperty {
	return newAudioDevProperty("latency", microseconds)
}

// IsTryMMAP tries using memory mapped device access if enabled. Default is disabled.
func IsTryMMAP(enabled bool) *AudioDevProperty {
	return newAudioDevProperty("try-mmap", enabled)
}

// IsExclusive specifies whether to open the device in exclusive mode (VMIX won't
// work in this case). Default is disabled.
func IsExclusive(enabled bool) *AudioDevProperty {
	return newAudioDevProperty("exclusive", enabled)
}

// WithDSPPolicy sets the timing policy (between 0 and 10, where smaller number means
// smaller latency but higher CPU usage). Use -1 to use buffer sizes specified by
// WithBufferCount. This option is ignored if you do not have OSS 4. Default is 5.
func WithDSPPolicy(policy int) *AudioDevProperty {
	return newAudioDevProperty("dsp-policy", policy)
}

// WithServer sets the PulseAudio server to connect to.
func WithServer(server string) *AudioDevProperty {
	return newAudioDevProperty("server", server)
}

// WithSink specifies the source/sink for recording/playback.
func WithSink(direction Direction, name string) *AudioDevProperty {
	return newDirectionProperty(direction, "name", name)
}
