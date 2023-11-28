package audiodev

// ALSABackend represents an audio backend using ALSA. This backend is only
// available on Linux.
type ALSABackend struct {
	*Backend
}

// NewALSABackend returns a new instance of [ALSABackend].
//
//	qemu-system-* -audiodev alsa,id=id
func NewALSABackend(id string) *ALSABackend {
	return &ALSABackend{
		New("alsa", id),
	}
}

// SetDevice specifies the ALSA device to use for input and/or output.
// The default is "default".
//
//	qemu-system-* -audiodev alsa,in|out.dev=device
func (b *ALSABackend) SetDevice(direction Direction, device string) *ALSABackend {
	b.SetDirectionProperty(direction, "dev", device)
	return b
}

// SetPeriodLength sets the period length in microseconds.
//
//	qemu-system-* -audiodev alsa,in|out.period-length=length
func (b *ALSABackend) SetPeriodLength(direction Direction, length int) *ALSABackend {
	b.SetDirectionProperty(direction, "period-length", length)
	return b
}

// SetThreshold specifies the threshold in microseconds when playback starts.
//
//	qemu-system-* -audiodev alsa,threshold=threshold
func (b *ALSABackend) SetThreshold(threshold int) *ALSABackend {
	b.SetProperty("threshold", threshold)
	return b
}

// ToggleTryPoll will attempt to use poll mode with the device if enabled.
// This property is enabled by default.
//
//	qemu-system-* -audiodev alsa,in|out.try-poll=on|off
func (b *ALSABackend) ToggleTryPoll(direction Direction, enabled bool) *ALSABackend {
	b.SetDirectionProperty(direction, "try-poll", enabled)
	return b
}
