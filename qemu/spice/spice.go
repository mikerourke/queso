package spice

import (
	"strings"

	"github.com/mikerourke/queso"
)

// Use enables the spice remote desktop protocol.
func Use(properties ...*Property) *queso.Option {
	props := make([]*queso.Property, 0)

	for _, property := range properties {
		props = append(props, property.Property)
	}

	return queso.NewOption("spice", "", props...)
}

// Property represents a property for use with the Spice option.
type Property struct {
	*queso.Property
}

// NewProperty returns a new instance of Property.
func NewProperty(key string, value interface{}) *Property {
	return &Property{
		Property: queso.NewProperty(key, value),
	}
}

// WithTCPPort sets the TCP port Spice is listening on for plaintext channels.
func WithTCPPort(port int) *Property {
	return NewProperty("port", port)
}

// WithIPAddress set the IP address Spice is listening on. The default is any
// address.
func WithIPAddress(addr string) *Property {
	return NewProperty("addr", addr)
}

// IPVersion represents the IP version to use with the IsIPVersionUsed property
// for Spice.
type IPVersion string

const (
	IPVersion4    IPVersion = "ipv4"
	IPVersion6    IPVersion = "ipv6"
	IPVersionUnix IPVersion = "unix"
)

// IsIPVersionUsed forces using the specified IP version for Spice.
func IsIPVersionUsed(ipVersion IPVersion, enabled bool) *Property {
	return NewProperty(string(ipVersion), enabled)
}

// WithPassword sets the password you need to authenticate for Spice.
//
// This property is deprecated and insecure because it leaves the password visible
// in the process listing. Use WithPasswordSecret instead.
func WithPassword(password string) *Property {
	return NewProperty("password", password)
}

// WithPasswordSecret sets the ID of the Secret object containing the
// password you need to authenticate for Spice.
func WithPasswordSecret(secret string) *Property {
	return NewProperty("password-secret", secret)
}

// IsSASL enables/disables the requirement that the Spice client use SASL
// to authenticate. The exact choice of authentication method used is controlled
// from the system/user's SASL configuration file for the `qemu` service. This is
// typically found in `/etc/sasl2/qemu.conf`.
//
// If running QEMU as an unprivileged user, an environment variable `SASL_CONF_PATH`
// can be used to make it search alternate locations for the service config.
// While some SASL auth methods can also provide data encryption (e.g. GSSAPI), it is
// recommended that SASL always be combined with the `tls` and `x509` settings to
// enable use of SSL and server certificates. This ensures a data encryption preventing
// compromise of authentication credentials.
func IsSASL(enabled bool) *Property {
	return NewProperty("sasl", enabled)
}

// IsTicketingDisabled specifies whether to allow Spice client connections
// without authentication.
func IsTicketingDisabled(disabled bool) *Property {
	return NewProperty("disable-ticketing", disabled)
}

// IsCopyPasteDisabled specifies whether to disable copy/paste between the
// Spice client and the guest.
func IsCopyPasteDisabled(disabled bool) *Property {
	return NewProperty("disable-copy-paste", disabled)
}

// IsAgentFileTransferDisabled specifies whether to disable spice-vdagent based
// file transfer between the Spice client and the guest.
func IsAgentFileTransferDisabled(disabled bool) *Property {
	return NewProperty("disable-agent-file-xfer", disabled)
}

// WithTLSPort sets the TCP port Spice is listening on for encrypted channels.
func WithTLSPort(port int) *Property {
	return NewProperty("tls-port", port)
}

// WithX509Directory sets the x509 file directory for Spice.
func WithX509Directory(path string) *Property {
	return NewProperty("x509-dir", path)
}

// WithX509KeyFile sets the x509 key file for Spice.
func WithX509KeyFile(file string) *Property {
	return NewProperty("x509-key-file", file)
}

// WithX509KeyPasswordFile sets the x509 key password file for Spice.
func WithX509KeyPasswordFile(file string) *Property {
	return NewProperty("x509-key-password", file)
}

// WithX509CertificateFile sets the x509 certificate file for Spice.
func WithX509CertificateFile(file string) *Property {
	return NewProperty("x509-cert-file", file)
}

// WithX509CACertificateFile sets the x509 CA certificate file for Spice.
func WithX509CACertificateFile(file string) *Property {
	return NewProperty("x509-cacert-file", file)
}

// WithX509DHKeyFile sets the x509 DH key file for Spice.
func WithX509DHKeyFile(file string) *Property {
	return NewProperty("x509-dh-key-file", file)
}

// WithTLSCiphers specifies which ciphers to use for Spice.
func WithTLSCiphers(ciphers ...string) *Property {
	value := strings.Join(ciphers, ",")

	return NewProperty("tls-ciphers", value)
}

// Channel is used to define channels with or without TLS encryption for
// the WithTLSChannel and WithPlainTextChannel properties.
type Channel string

const (
	ChannelDefault  Channel = "default"
	ChannelCursor   Channel = "cursor"
	ChannelDisplay  Channel = "display"
	ChannelInputs   Channel = "inputs"
	ChannelMain     Channel = "main"
	ChannelPlayback Channel = "playback"
	ChannelRecord   Channel = "record"
)

// WithTLSChannel forces specific channel to be used with TLS encryption for Spice.
// The property can be specified multiple times to configure multiple channels.
// The special name ChannelDefault can be used to set the default mode.
//
// For channels which are not explicitly forced into one mode the spice client
// is allowed to pick tls/plaintext as desired.
func WithTLSChannel(channel Channel) *Property {
	return NewProperty("tls-channel", channel)
}

// WithPlainTextChannel forces specific channel to be used without TLS encryption
// for Spice. The property can be specified multiple times to configure multiple
// channels. The special name ChannelDefault can be used to set the default
// mode.
//
// For channels which are not explicitly forced into one mode the Spice client
// is allowed to pick tls/plaintext as desired.
func WithPlainTextChannel(channel Channel) *Property {
	return NewProperty("plaintext-channel", channel)
}

// ImageCompressionType represents the lossless image compression type used with
// the WithImageCompressionType property for Spice.
type ImageCompressionType string

const (
	ImageCompressionAutoGLZ ImageCompressionType = "auto_glz"
	ImageCompressionAutoLZ  ImageCompressionType = "auto_lz"
	ImageCompressionQUIC    ImageCompressionType = "quic"
	ImageCompressionGLZ     ImageCompressionType = "glz"
	ImageCompressionLZ      ImageCompressionType = "lz"
	ImageCompressionOff     ImageCompressionType = "off"
)

// WithImageCompressionType configures image compression (lossless) for Spice.
// The default is ImageCompressionAutoGLZ.
func WithImageCompressionType(compression ImageCompressionType) *Property {
	return NewProperty("image-compression", compression)
}

// WANCompressionMode represents the mode to use for wan image compression for
// the WithJPEGWANCompressionMode and WithZLIBGLZWANCompressionMode properties
// for Spice.
type WANCompressionMode string

const (
	WANCompressionModeAuto   WANCompressionMode = "auto"
	WANCompressionModeAlways WANCompressionMode = "always"
	WANCompressionModeNever  WANCompressionMode = "never"
)

// WithJPEGWANCompressionMode configures wan image compression for JPEG files
// (lossy for slow links) for Spice. The default is WANCompressionModeAuto.
func WithJPEGWANCompressionMode(mode WANCompressionMode) *Property {
	return NewProperty("jpeg-wan-compression", mode)
}

// WithZLIBGLZWANCompressionMode configures wan image compression for ZLIB GLZ
// files (lossy for slow links) for Spice. The default is WANCompressionModeAuto.
func WithZLIBGLZWANCompressionMode(mode WANCompressionMode) *Property {
	return NewProperty("zlib-glz-wan-compression", mode)
}

// VideoStreamDetection represents the options for the WithVideoStreamDetection
// property for Spice.
type VideoStreamDetection string

const (
	VideoStreamDetectionOff    VideoStreamDetection = "off"
	VideoStreamDetectionAll    VideoStreamDetection = "all"
	VideoStreamDetectionFilter VideoStreamDetection = "filter"
)

// WithVideoStreamDetection configures video stream detection for Spice.
// The default is VideoStreamDetectionOff.
func WithVideoStreamDetection(detection VideoStreamDetection) *Property {
	return NewProperty("streaming-video", detection)
}

// IsPassMouseEventsViaAgent enables/disables passing mouse events via vdagent
// for Spice. This property is enabled by default.
func IsPassMouseEventsViaAgent(enabled bool) *Property {
	return NewProperty("agent-mouse", enabled)
}

// IsAudioStreamCompression enables/disables audio stream compression
// (using celt 0.5.1) for Spice. This property is enabled by default.
func IsAudioStreamCompression(enabled bool) *Property {
	return NewProperty("playback-compression", enabled)
}

// IsSeamlessMigration enables/disables spice seamless migration for Spice. This
// property is disabled by default.
func IsSeamlessMigration(enabled bool) *Property {
	return NewProperty("seamless-migration", enabled)
}

// IsOpenGL enables/disables OpenGL context for Spice. This property is disabled
// by default.
func IsOpenGL(enabled bool) *Property {
	return NewProperty("gl", enabled)
}

// WithDRMRenderNode specifies the DRM render node for OpenGL rendering for Spice.
// If not specified, it will pick the first available.
func WithDRMRenderNode(file string) *Property {
	return NewProperty("rendernode", file)
}
