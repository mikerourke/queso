package netdev

import (
	"strconv"
	"strings"
)

// AFXDPBackend represents an AF_XDP network backend. AF_XDP is an address family
// that is optimized for high performance packet processing.
//
// Using the XDP_REDIRECT action from an XDP program, the program can redirect ingress
// frames to other XDP enabled network devices, using the bpf_redirect_map() function.
//
// AF_XDP sockets enable the possibility for XDP programs to redirect frames to a
// memory buffer in a user-space application.
//
// See https://www.kernel.org/doc/html/next/networking/af_xdp.html for more details.
type AFXDPBackend struct {
	*Backend
}

// NewAFXDPBackend returns a new instance of [AFXDPBackend]. id is the unique
// identifier for the backend. interfaceName is the name of the network interface
// to connect to.
//
//	qemu-system-* -netdev af-xdp,id=id,ifname=interfaceName
func NewAFXDPBackend(id string, interfaceName string) *AFXDPBackend {
	backend := New("af-xdp")

	backend.SetProperty("id", id).SetProperty("ifname", interfaceName)

	return &AFXDPBackend{backend}
}

// ProgramAttachMode defines a program attach mode for the XDP program. It is
// set from the [AFXDPBackend.SetProgramAttachMode] method.
type ProgramAttachMode string

const (
	// ProgramAttachNative is used if the driver has support for XDP, it will be used
	// by the AF_XDP code to provide better performance, but there is still a copy
	// of the data into user space.
	ProgramAttachNative ProgramAttachMode = "native"

	// ProgramAttachSKB uses SKBs together with the generic XDP support and copies
	// out the data to user space. A fallback mode that works for any network device.
	ProgramAttachSKB ProgramAttachMode = "skb"
)

// SetProgramAttachMode forces a specific program attach mode for a default XDP
// program. This value defaults to best-effort, where the likely most performant
// mode will be in use.
//
//	qemu-system-* -netdev af-xdp,mode=native|skb
func (b *AFXDPBackend) SetProgramAttachMode(mode ProgramAttachMode) *AFXDPBackend {
	b.SetProperty("mode", string(mode))
	return b
}

// SetQueueCount sets the number of queues, which should generally match the number
// of queues in the interface. Traffic arriving at non-configured device queues
// will not be delivered to the network backend.
//
//	qemu-system-* -netdev af-xdp,queues=count
func (b *AFXDPBackend) SetQueueCount(count int) *AFXDPBackend {
	b.SetProperty("queues", count)
	return b
}

// SetQueueStart can be specified if a particular range of queues
// [<start>, <start> + <count>]  should be in use. For example, this is may be
// necessary in order to use certain  NICs in [ProgramAttachNative] mode.
//
// The kernel allows the driver to create a separate set of XDP queues on top of
// regular ones, and only these queues can be used for AF_XDP sockets.
// NICs that work this way may also require an additional traffic redirection
// with ethtool to these special queues.
//
//	qemu-system-* -netdev af-xdp,start-queue=start
func (b *AFXDPBackend) SetQueueStart(start int) *AFXDPBackend {
	b.SetProperty("start-queue", start)
	return b
}

// SetSocketFDs sets the file descriptors for already open but not bound XDP
// sockets already added to a socket map for corresponding queues. You must
// call [AFXDPBackend.ToggleInhibit] with true to use this.
//
//	qemu-system-* -netdev af-xdp,sock-fds=x:y:...:z
func (b *AFXDPBackend) SetSocketFDs(fds ...int) *AFXDPBackend {
	fdStrings := make([]string, len(fds))
	for _, fd := range fds {
		fdStrings = append(fdStrings, strconv.Itoa(fd))
	}

	b.SetProperty("sock-fds", strings.Join(fdStrings, ","))
	return b
}

// ToggleForceCopy enables or disables force copy.
//
//	qemu-system-* -netdev af-xdp,force-copy=on|off
func (b *AFXDPBackend) ToggleForceCopy(enabled bool) *AFXDPBackend {
	b.SetProperty("force-copy", enabled)
	return b
}

// ToggleInhibit enables or disables loading the XDP program externally. If set
// to true, [AFXDPBackend.SetSocketFDs] should be provided with file descriptors
// for already open but not bound XDP sockets already added to a socket map for
// corresponding queues. One socket per queue.
//
//	qemu-system-* -netdev af-xdp,inhibit=on|off
func (b *AFXDPBackend) ToggleInhibit(enabled bool) *AFXDPBackend {
	b.SetProperty("inhibit", enabled)
	return b
}
