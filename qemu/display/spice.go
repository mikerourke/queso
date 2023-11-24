package display

import (
	"strings"

	"github.com/mikerourke/queso/qemu/cli"
)

// TODO: Read more about Spice so you can add comments to SpiceChannel.

// SpiceChannel is a channel.
type SpiceChannel string

const (
	SpiceChannelCursor   SpiceChannel = "cursor"
	SpiceChannelDisplay  SpiceChannel = "display"
	SpiceChannelDefault  SpiceChannel = "default"
	SpiceChannelInputs   SpiceChannel = "inputs"
	SpiceChannelMain     SpiceChannel = "main"
	SpiceChannelPlayback SpiceChannel = "playback"
	SpiceChannelRecord   SpiceChannel = "record"
)

// SpiceDisplay is a display that uses the Spice remote desktop protocol.
type SpiceDisplay struct {
	*Display
}

// NewSpiceDisplay returns a new instance of [SpiceDisplay].
//
//	qemu-system-* -spice
func NewSpiceDisplay() *SpiceDisplay {
	return &SpiceDisplay{New("spice")}
}

// SetAddress sets the IP address spice is listening on. Default is any address.
//
//	qemu-system-* -spice,addr=addr
func (d *SpiceDisplay) SetAddress(addr string) *SpiceDisplay {
	d.properties = append(d.properties, cli.NewProperty("addr", addr))
	return d
}

// SetChannelForPlainText specifies which channel should not be encrypted (i.e.
// plain text). This can be called multiple times to configure multiple channels.
// The special name [SpiceChannelDefault] can be used to set the default mode.
//
// For channels which are not explicitly forced into one mode, the Spice client is
// allowed to pick TLS/plain text as it pleases.
//
//	qemu-system-* -spice,plaintext-channel=main|display|cursor|inputs|record|playback
func (d *SpiceDisplay) SetChannelForPlainText(channel SpiceChannel) *SpiceDisplay {
	d.properties = append(d.properties, cli.NewProperty("plaintext-channel", channel))
	return d
}

// SetChannelForTLS specifies which channel should be encrypted. This can be called
// multiple times to configure multiple channels. The special name [SpiceChannelDefault]
// can be used to set the default mode.
//
// For channels which are not explicitly forced into one mode, the Spice client is
// allowed to pick TLS/plain text as it pleases.
//
//	qemu-system-* -spice,tls-channel=main|display|cursor|inputs|record|playback
func (d *SpiceDisplay) SetChannelForTLS(channel SpiceChannel) *SpiceDisplay {
	d.properties = append(d.properties, cli.NewProperty("tls-channel", channel))
	return d
}

// ImageCompression represents the lossless image compression type used with
// the [SpiceDisplay.SetImageCompression] method.
type ImageCompression string

const (
	ImageCompressionAutoGLZ ImageCompression = "auto_glz"
	ImageCompressionAutoLZ  ImageCompression = "auto_lz"
	ImageCompressionGLZ     ImageCompression = "glz"
	ImageCompressionLZ      ImageCompression = "lz"
	ImageCompressionOff     ImageCompression = "off"
	ImageCompressionQUIC    ImageCompression = "quic"
)

// SetImageCompression configures image compression (lossless). The default
// is [ImageCompressionAutoGLZ].
//
//	qemu-system-* -spice,image-compression=[auto_glz|auto_lz|quic|glz|lz|off]
func (d *SpiceDisplay) SetImageCompression(compression ImageCompression) *SpiceDisplay {
	d.properties = append(d.properties,
		cli.NewProperty("image-compression", string(compression)))
	return d
}

// SetPasswordSecret sets the ID of the Secret object containing the
// password you need to authenticate.
//
//	qemu-system-* -spice,password-secret=secret
func (d *SpiceDisplay) SetPasswordSecret(secret string) *SpiceDisplay {
	d.properties = append(d.properties, cli.NewProperty("password-secret", secret))
	return d
}

// SetPort sets the TCP port Spice is listening on for plaintext channels.
//
//	qemu-system-* -spice,port=port
func (d *SpiceDisplay) SetPort(port int) *SpiceDisplay {
	d.properties = append(d.properties, cli.NewProperty("port", port))
	return d
}

// SetRenderNode sets the DRM render node for OpenGL rendering. If not specified,
// it will pick the first available. (Since 2.9)
//
//	qemu-system-* -spice,rendernode=file
func (d *SpiceDisplay) SetRenderNode(file string) *SpiceDisplay {
	d.properties = append(d.properties, cli.NewProperty("rendernode", file))
	return d
}

// SetTLSCiphers specifies which ciphers to use.
//
//	qemu-system-* -spice,tls-ciphers=ciphers
func (d *SpiceDisplay) SetTLSCiphers(ciphers ...string) *SpiceDisplay {
	d.properties = append(d.properties, cli.NewProperty("tls-ciphers", strings.Join(ciphers, ",")))
	return d
}

// SetTLSPort sets the TCP port Spice is listening on for encrypted channels.
//
//	qemu-system-* -spice,tls-port=port
func (d *SpiceDisplay) SetTLSPort(port int) *SpiceDisplay {
	d.properties = append(d.properties, cli.NewProperty("tls-port", port))
	return d
}

// VideoStreamDetection represents the allowed values for the
// [SpiceDisplay.SetVideoStreamDetection] property.
type VideoStreamDetection string

const (
	VideoStreamDetectionOff    VideoStreamDetection = "off"
	VideoStreamDetectionAll    VideoStreamDetection = "all"
	VideoStreamDetectionFilter VideoStreamDetection = "filter"
)

// SetVideoStreamDetection configures video stream detection. The default is
// [VideoStreamDetectionOff].
//
//	qemu-system-* -spice,streaming-video=off|all|filter
func (d *SpiceDisplay) SetVideoStreamDetection(detection VideoStreamDetection) *SpiceDisplay {
	d.properties = append(d.properties, cli.NewProperty("streaming-video", string(detection)))
	return d
}

// TODO: Find out if the way I'm specifying the WAN compression modes will work.

// WANCompressionMode represents the mode to use for WAN image compression set
// using the [SpiceDisplay.SetWANCompressionForJPEG] and
// [SpiceDisplay.SetWANCompressionForZlib] methods.
type WANCompressionMode string

const (
	WANCompressionAuto   WANCompressionMode = "auto"
	WANCompressionAlways WANCompressionMode = "always"
	WANCompressionNever  WANCompressionMode = "never"
)

// SetWANCompressionForJPEG configures WAN image compression for JPEG (lossy
// for slow links). The default is [WANCompressionAuto].
//
//	qemu-system-* -spice,jpeg-wan-compression=auto|never|always
func (d *SpiceDisplay) SetWANCompressionForJPEG(mode WANCompressionMode) *SpiceDisplay {
	d.properties = append(d.properties, cli.NewProperty("jpeg-wan-compression", string(mode)))
	return d
}

// SetWANCompressionForZlib configures WAN image compression for Zlib (lossy
// for slow links). The default is [WANCompressionAuto].
//
//	qemu-system-* -spice,zlib-glz-wan-compression=auto|never|always
func (d *SpiceDisplay) SetWANCompressionForZlib(mode WANCompressionMode) *SpiceDisplay {
	d.properties = append(d.properties, cli.NewProperty("zlib-glz-wan-compression", string(mode)))
	return d
}

// SetX509CACertFile sets the x509 CA certificate file path.
//
//	qemu-system-* -spice,x509-cacert-file=file
func (d *SpiceDisplay) SetX509CACertFile(file string) *SpiceDisplay {
	d.properties = append(d.properties, cli.NewProperty("x509-cacert-file", file))
	return d
}

// SetX509CertFile sets the x509 certificate file path.
//
//	qemu-system-* -spice,x509-cert-file=file
func (d *SpiceDisplay) SetX509CertFile(file string) *SpiceDisplay {
	d.properties = append(d.properties, cli.NewProperty("x509-cert-file", file))
	return d
}

// SetX509DHKeyFile sets the x509 DH key file path.
//
//	qemu-system-* -spice,x509-dh-key-file=file
func (d *SpiceDisplay) SetX509DHKeyFile(file string) *SpiceDisplay {
	d.properties = append(d.properties, cli.NewProperty("x509-dh-key-file", file))
	return d
}

// SetX509KeyFile sets the x509 key file path.
//
//	qemu-system-* -spice,x509-key-file=file
func (d *SpiceDisplay) SetX509KeyFile(file string) *SpiceDisplay {
	d.properties = append(d.properties, cli.NewProperty("x509-key-file", file))
	return d
}

// SetX509KeyPasswordFile sets the x509 key password file path.
//
//	qemu-system-* -spice,x509-key-password=file
func (d *SpiceDisplay) SetX509KeyPasswordFile(file string) *SpiceDisplay {
	d.properties = append(d.properties, cli.NewProperty("x509-key-password", file))
	return d
}

// ToggleAgentFileTransfers enables or disables spice-vdagent based file transfer
// between the client and the guest.
//
//	qemu-system-* -spice,disable-agent-file-xfer=on|off
func (d *SpiceDisplay) ToggleAgentFileTransfers(enabled bool) *SpiceDisplay {
	// Negating enabled to ensure if this is called with true, it does _not_ disable file transfer.
	d.properties = append(d.properties, cli.NewProperty("disable-agent-file-xfer", !enabled))
	return d
}

// ToggleConnectWithoutAuth enables or disables client connections without
// authentication.
//
//	qemu-system-* -spice,disable-ticketing=on|off
func (d *SpiceDisplay) ToggleConnectWithoutAuth(enabled bool) *SpiceDisplay {
	d.properties = append(d.properties, cli.NewProperty("disable-ticketing", enabled))
	return d
}

// ToggleIPv4 specifies if IPv4 may be used.
//
//	qemu-system-* -spice,ipv4=on|off
func (d *SpiceDisplay) ToggleIPv4(enabled bool) *SpiceDisplay {
	d.properties = append(d.properties, cli.NewProperty("ipv4", enabled))
	return d
}

// ToggleIPv6 specifies if IPv6 may be used.
//
//	qemu-system-* -spice,ipv6=on|off
func (d *SpiceDisplay) ToggleIPv6(enabled bool) *SpiceDisplay {
	d.properties = append(d.properties, cli.NewProperty("ipv6", enabled))
	return d
}

// ToggleMouseEventPassing enables or disables passing mouse events via vdagent
// for Spice. This property is enabled by default.
//
//	qemu-system-* -spice,agent-mouse=on|off
func (d *SpiceDisplay) ToggleMouseEventPassing(enabled bool) *SpiceDisplay {
	d.properties = append(d.properties, cli.NewProperty("agent-mouse", enabled))
	return d
}

// ToggleOpenGL enables or disables OpenGL for displaying.
//
//	qemu-system-* -spice,gl=on|off
func (d *SpiceDisplay) ToggleOpenGL(enabled bool) *SpiceDisplay {
	d.properties = append(d.properties, cli.NewProperty("gl", enabled))
	return d
}

// TogglePlaybackCompression enables or disables audio stream compression
// (using celt 0.5.1). This property is enabled by default.
//
//	qemu-system-* -spice,playback-compression=on|off
func (d *SpiceDisplay) TogglePlaybackCompression(enabled bool) *SpiceDisplay {
	d.properties = append(d.properties, cli.NewProperty("playback-compression", enabled))
	return d
}

// ToggleSASL enables or disables the requirement that the client use SASL to authenticate
// with the Spice. The exact choice of authentication method used is controlled from
// the system / user’s SASL configuration file for the ‘qemu’ service. This is
// typically found in `/etc/sasl2/qemu.conf`.
//
// If running QEMU as an unprivileged user, an environment variable `SASL_CONF_PATH`
// can be used to make it search alternate locations for the service config.
// While some SASL auth methods can also provide data encryption (e.g. GSSAPI), it
// is recommended that SASL always be combined with the "tls" and "x509" settings
// to enable use of SSL and server certificates. This ensures a data encryption
// preventing compromise of authentication credentials.
//
//	qemu-system-* -spice,sasl=on|off
func (d *SpiceDisplay) ToggleSASL(required bool) *SpiceDisplay {
	d.properties = append(d.properties, cli.NewProperty("sasl", required))
	return d
}

// ToggleSeamlessMigration enables or disables Spice seamless migration. This
// property is disabled by default.
func (d *SpiceDisplay) ToggleSeamlessMigration(enabled bool) *SpiceDisplay {
	d.properties = append(d.properties, cli.NewProperty("seamless-migration", enabled))
	return d
}

// ToggleSharedClipboard enables or disables copy/paste between the client and
// the guest.
//
//	qemu-system-* -spice,disable-copy-paste=on|off
func (d *SpiceDisplay) ToggleSharedClipboard(enabled bool) *SpiceDisplay {
	// Negating enabled to ensure if this is called with true, it does _not_ disable copy/paste.
	d.properties = append(d.properties, cli.NewProperty("disable-copy-paste", !enabled))
	return d
}

// ToggleUnix specifies if Unix socket may be used.
//
//	qemu-system-* -spice,unix=on|off
func (d *SpiceDisplay) ToggleUnix(enabled bool) *SpiceDisplay {
	d.properties = append(d.properties, cli.NewProperty("unix", enabled))
	return d
}
