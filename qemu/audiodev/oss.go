package audiodev

import "github.com/mikerourke/queso"

// OSSBackend represents an audio backend using OSS. This backend is available
// on most Unix-like systems.
type OSSBackend struct {
	*Backend
}

// NewOSSBackend returns a new instance of [OSSBackend].
//
//	qemu-system-* -audiodev oss,id=id
func NewOSSBackend(id string) *OSSBackend {
	return &OSSBackend{
		NewBackend("oss", id),
	}
}

// SetBufferCount sets the count of buffers.
//
//	qemu-system-* -audiodev os,buffer-count=count
func (b *OSSBackend) SetBufferCount(direction Direction, count string) *OSSBackend {
	b.properties = append(b.properties, newDirectionProperty("buffer-count", direction, count))
	return b
}

// SetDevice specifies the OSS device to use for input and/or output.
// The default is "/dev/dsp".
//
//	qemu-system-* -audiodev oss,in|out.dev=device
func (b *OSSBackend) SetDevice(direction Direction, device string) *OSSBackend {
	b.properties = append(b.properties, newDirectionProperty("dev", direction, device))
	return b
}

// SetDSPPolicy sets the timing policy (between 0 and 10, where smaller number means
// smaller latency but higher CPU usage). Use -1 to use buffer sizes specified by
// buffer and [OSSBackend.SetBufferCount]. This option is ignored if you do not have OSS 4.
// The default value is 5.
//
//	qemu-system-* -audiodev oss,dsp-policy=policy
func (b *OSSBackend) SetDSPPolicy(policy int) *OSSBackend {
	b.properties = append(b.properties, queso.NewProperty("dsp-policy", policy))
	return b
}

// ToggleExclusive specifies whether to open the device in exclusive mode (VMIX
// won't work in this case). This property is disabled by default.
//
//	qemu-system-* -audiodev oss,exclusive=on|off
func (b *OSSBackend) ToggleExclusive(enabled bool) *OSSBackend {
	b.properties = append(b.properties, queso.NewProperty("exclusive", enabled))
	return b
}

// ToggleTryMMAP will try using memory mapped device access. The default value
// is false.
//
//	qemu-system-* -audiodev oss,try-mmap=on|off
func (b *OSSBackend) ToggleTryMMAP(enabled bool) *OSSBackend {
	b.properties = append(b.properties, queso.NewProperty("try-mmap", enabled))
	return b
}

// ToggleTryPoll will attempt to use poll mode with the device if enabled.
// This property is enabled by default.
//
//	qemu-system-* -audiodev oss,in|out.try-poll=on|off
func (b *OSSBackend) ToggleTryPoll(direction Direction, enabled bool) *OSSBackend {
	b.properties = append(b.properties, newDirectionProperty("try-poll", direction, enabled))
	return b
}
