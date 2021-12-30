package display

import (
	"net"
	"strings"

	"github.com/mikerourke/queso"
)

// Spice enables the spice remote desktop protocol.
func Spice(properties ...*SpiceProperty) *queso.Option {
	props := make([]*queso.Property, 0)

	for _, property := range properties {
		props = append(props, property.Property)
	}

	return queso.NewOption("spice", "", props...)
}

// SpiceProperty represents a property for use with the Spice option.
type SpiceProperty struct {
	*queso.Property
}

// NewSpiceProperty returns a new instance of SpiceProperty.
func NewSpiceProperty(key string, value interface{}) *SpiceProperty {
	return &SpiceProperty{
		Property: queso.NewProperty(key, value),
	}
}

// WithTCPPort sets the TCP port spice is listening on for plaintext channels.
func WithTCPPort(port int) *SpiceProperty {
	return NewSpiceProperty("port", port)
}

// WithIPAddress set the IP address spice is listening on.
// The default is any address.
func WithIPAddress(addr net.IP) *SpiceProperty {
	return NewSpiceProperty("addr", addr.String())
}

// IPVersion represents the IP version to use with the IsIPVersionUsed property.
type IPVersion string

const (
	IPVersion4    IPVersion = "ipv4"
	IPVersion6    IPVersion = "ipv6"
	IPVersionUnix IPVersion = "unix"
)

// IsIPVersionUsed forces using the specified IP version.
func IsIPVersionUsed(ipVersion IPVersion, enabled bool) *SpiceProperty {
	return NewSpiceProperty(string(ipVersion), enabled)
}

// WithPassword sets the password you need to authenticate.
//
// This property is deprecated and insecure because it leaves the password visible
// in the process listing. Use WithSpicePasswordSecret instead.
func WithPassword(password string) *SpiceProperty {
	return NewSpiceProperty("password", password)
}

// WithSpicePasswordSecret sets the ID of the Secret object containing the
// password you need to authenticate.
func WithSpicePasswordSecret(secret string) *SpiceProperty {
	return NewSpiceProperty("password-secret", secret)
}

// IsSpiceSASL enables/disables the requirement that the client use SASL to authenticate
// with spice. The exact choice of authentication method used is controlled
// from the system/user's SASL configuration file for the `qemu` service. This is
// typically found in `/etc/sasl2/qemu.conf`.
//
// If running QEMU as an unprivileged user, an environment variable `SASL_CONF_PATH`
// can be used to make it search alternate locations for the service config.
// While some SASL auth methods can also provide data encryption (e.g. GSSAPI), it is
// recommended that SASL always be combined with the `tls` and `x509` settings to
// enable use of SSL and server certificates. This ensures a data encryption preventing
// compromise of authentication credentials.
func IsSpiceSASL(enabled bool) *SpiceProperty {
	return NewSpiceProperty("sasl", enabled)
}

// IsTicketingDisabled specifies whether to allow client connects without
// authentication.
func IsTicketingDisabled(disabled bool) *SpiceProperty {
	return NewSpiceProperty("disable-ticketing", disabled)
}

// IsCopyPasteDisabled specifies whether to disable copy/paste between the client
// and the guest.
func IsCopyPasteDisabled(disabled bool) *SpiceProperty {
	return NewSpiceProperty("disable-copy-paste", disabled)
}

// IsAgentFileTransferDisabled specifies whether to disable spice-vdagent based
// file transfer between the client and the guest.
func IsAgentFileTransferDisabled(disabled bool) *SpiceProperty {
	return NewSpiceProperty("disable-agent-file-xfer", disabled)
}

// WithTLSPort sets the TCP port spice is listening on for encrypted channels.
func WithTLSPort(port int) *SpiceProperty {
	return NewSpiceProperty("tls-port", port)
}

// WithX509Directory sets the x509 file directory.
func WithX509Directory(path string) *SpiceProperty {
	return NewSpiceProperty("x509-dir", path)
}

// WithX509KeyFile sets the x509 key file.
func WithX509KeyFile(file string) *SpiceProperty {
	return NewSpiceProperty("x509-key-file", file)
}

// WithX509KeyPasswordFile sets the x509 key password file.
func WithX509KeyPasswordFile(file string) *SpiceProperty {
	return NewSpiceProperty("x509-key-password", file)
}

// WithX509CertificateFile sets the x509 certificate file.
func WithX509CertificateFile(file string) *SpiceProperty {
	return NewSpiceProperty("x509-cert-file", file)
}

// WithX509CACertificateFile sets the x509 CA certificate file.
func WithX509CACertificateFile(file string) *SpiceProperty {
	return NewSpiceProperty("x509-cacert-file", file)
}

// WithX509DHKeyFile sets the x509 DH key file.
func WithX509DHKeyFile(file string) *SpiceProperty {
	return NewSpiceProperty("x509-dh-key-file", file)
}

// WithTLSCiphers specifies which ciphers to use.
func WithTLSCiphers(ciphers ...string) *SpiceProperty {
	value := strings.Join(ciphers, ",")

	return NewSpiceProperty("tls-ciphers", value)
}

// SpiceChannel is used to define channels with or without TLS encryption for
// the WithTLSChannel and WithPlainTextChannel properties.
type SpiceChannel string

const (
	SpiceChannelDefault  SpiceChannel = "default"
	SpiceChannelMain     SpiceChannel = "main"
	SpiceChannelDisplay  SpiceChannel = "display"
	SpiceChannelCursor   SpiceChannel = "cursor"
	SpiceChannelInputs   SpiceChannel = "inputs"
	SpiceChannelRecord   SpiceChannel = "record"
	SpiceChannelPlayback SpiceChannel = "playback"
)

// WithTLSChannel forces specific channel to be used with TLS encryption.
// The property can be specified multiple times to configure multiple channels.
// The special name SpiceChannelDefault can be used to set the default mode.
//
// For channels which are not explicitly forced into one mode the spice client
// is allowed to pick tls/plaintext as desired.
func WithTLSChannel(channel SpiceChannel) *SpiceProperty {
	return NewSpiceProperty("tls-channel", channel)
}

// WithPlainTextChannel forces specific channel to be used without TLS encryption.
// The property can be specified multiple times to configure multiple channels.
// The special name SpiceChannelDefault can be used to set the default mode.
//
// For channels which are not explicitly forced into one mode the spice client
// is allowed to pick tls/plaintext as desired.
func WithPlainTextChannel(channel SpiceChannel) *SpiceProperty {
	return NewSpiceProperty("plaintext-channel", channel)
}

// ImageCompressionType represents the lossless image compression type used with
// the WithImageCompressionType property.
type ImageCompressionType string

const (
	ImageCompressionAutoGLZ ImageCompressionType = "auto_glz"
	ImageCompressionAutoLZ  ImageCompressionType = "auto_lz"
	ImageCompressionQUIC    ImageCompressionType = "quic"
	ImageCompressionGLZ     ImageCompressionType = "glz"
	ImageCompressionLZ      ImageCompressionType = "lz"
	ImageCompressionOff     ImageCompressionType = "off"
)

// WithImageCompressionType configures image compression (lossless).
// The default is ImageCompressionAutoGLZ.
func WithImageCompressionType(compression ImageCompressionType) *SpiceProperty {
	return NewSpiceProperty("image-compression", compression)
}

// WANCompressionMode represents the mode to use for wan image compression for
// the WithJPEGWANCompressionMode and WithZLIBGLZWANCompressionMode properties.
type WANCompressionMode string

const (
	WANCompressionModeAuto   WANCompressionMode = "auto"
	WANCompressionModeAlways WANCompressionMode = "always"
	WANCompressionModeNever  WANCompressionMode = "never"
)

// WithJPEGWANCompressionMode configures wan image compression for JPEG files (lossy
// for slow links). The default is WANCompressionModeAuto.
func WithJPEGWANCompressionMode(mode WANCompressionMode) *SpiceProperty {
	return NewSpiceProperty("jpeg-wan-compression", mode)
}

// WithZLIBGLZWANCompressionMode configures wan image compression for ZLIB GLZ
// files (lossy for slow links). The default is WANCompressionModeAuto.
func WithZLIBGLZWANCompressionMode(mode WANCompressionMode) *SpiceProperty {
	return NewSpiceProperty("zlib-glz-wan-compression", mode)
}

// VideoStreamDetection represents the options for the WithVideoStreamDetection
// property.
type VideoStreamDetection string

const (
	VideoStreamDetectionOff    VideoStreamDetection = "off"
	VideoStreamDetectionAll    VideoStreamDetection = "all"
	VideoStreamDetectionFilter VideoStreamDetection = "filter"
)

// WithVideoStreamDetection configures video stream detection.
// The default is VideoStreamDetectionOff.
func WithVideoStreamDetection(detection VideoStreamDetection) *SpiceProperty {
	return NewSpiceProperty("streaming-video", detection)
}

// IsPassMouseEventsViaAgent enables/disables passing mouse events via vdagent.
// This is enabled by default.
func IsPassMouseEventsViaAgent(enabled bool) *SpiceProperty {
	return NewSpiceProperty("agent-mouse", enabled)
}

// IsAudioStreamCompression enables/disables audio stream compression (using celt 0.5.1).
// This is enabled by default.
func IsAudioStreamCompression(enabled bool) *SpiceProperty {
	return NewSpiceProperty("playback-compression", enabled)
}

// IsSeamlessMigration enables/disables spice seamless migration.
// This is disabled by default.
func IsSeamlessMigration(enabled bool) *SpiceProperty {
	return NewSpiceProperty("seamless-migration", enabled)
}

// IsOpenGL enables/disables OpenGL context. This is disabled by default.
func IsOpenGL(enabled bool) *SpiceProperty {
	return NewSpiceProperty("gl", enabled)
}

// WithDRMRenderNode specifies the DRM render node for OpenGL rendering. If not
// specified, it will pick the first available.
func WithDRMRenderNode(file string) *SpiceProperty {
	return NewSpiceProperty("rendernode", file)
}
