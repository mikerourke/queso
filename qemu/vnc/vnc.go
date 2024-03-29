package vnc

import (
	"fmt"
	"strconv"

	"github.com/mikerourke/queso"
)

type DisplayType interface {
	Property() string
}

// ToDisplay corresponds to the "to" option for the VNC display.
type ToDisplay struct {
	value int
}

// UseTo returns a new instance of ToDisplay. With this option, QEMU will
// try next available VNC displays, until the value parameter specified, if the originally
// defined VNC display is not available, e.g. port 5900 + display is already used by
// another application. By default, value is 0.
func UseTo(value int) *ToDisplay {
	return &ToDisplay{value}
}

// Property returns the property value of the VNC Display property.
func (td *ToDisplay) Property() string {
	return fmt.Sprintf("to=%d", td.value)
}

// HostDisplay corresponds to the "host" option for the VNC display.
type HostDisplay struct {
	value int
}

// UseHost returns a new instance of HostDisplay. With this option, TCP
// connections will only be allowed from host on the displayOrPortNumber parameter
// specified.
//
// By convention, the TCP port is 5900 + displayOrPortNumber. Optionally, this
// option can be omitted in which case the server will accept connections from
// any host.
func UseHost(displayOrPortNumber int) *HostDisplay {
	return &HostDisplay{displayOrPortNumber}
}

// Property returns the property value of the VNC Display property.
func (hd *HostDisplay) Property() string {
	return fmt.Sprintf("host:%d", hd.value)
}

// UnixDisplay corresponds to the "unix" option for the VNC display.
type UnixDisplay struct {
	value string
}

// UseUnix returns a new instance of UnixDisplay. With this option,
// connections will be allowed over Unix domain sockets where the path parameter
// is the location of a Unix socket to listen for connections on.
func UseUnix(path string) *UnixDisplay {
	return &UnixDisplay{path}
}

// Property returns the property value of the VNC Display property.
func (ud *UnixDisplay) Property() string {
	return fmt.Sprintf("unix:%s", ud.value)
}

// NoneDisplay corresponds to the "none" option for the VNC display.
type NoneDisplay struct {
	value string
}

// UseNone returns a new instance of NoneDisplay. With this option,
// VNC is initialized but not started. The monitor change command can be used to
// later start the VNC server.
func UseNone(value string) *NoneDisplay {
	return &NoneDisplay{value}
}

// Property returns the property value of the VNC Display property.
func (nd *NoneDisplay) Property() string {
	return "none"
}

// VNC listens on the specified DisplayType and redirects the VGA display over
// the VNC session.
func VNC(display DisplayType, properties ...*Property) *queso.Option {
	props := make([]*queso.Property, 0)

	for _, property := range properties {
		props = append(props, property.Property)
	}

	return queso.NewOption("vnc", display.Property(), props...)
}

// Property represents a property for use with the VNC option.
type Property struct {
	*queso.Property
}

// NewProperty returns a new instance of a VNC Property.
func NewProperty(key string, value interface{}) *Property {
	return &Property{
		Property: queso.NewProperty(key, value),
	}
}

// IsReverse specifies whether to connect to a listening VNC client via a "reverse"
// connection. The client is specified by the display. For reverse network connections
// UseHost(<value>), the <value> argument is a TCP port number, not a display
// number.
func IsReverse(enabled bool) *Property {
	return NewProperty("reverse", enabled)
}

// IsWebSocket specifies whether to open an additional TCP listening port dedicated
// to VNC WebSocket connections. If a bare WebSocket option is given, the WebSocket
// port is 5700 + display.
//
// If HostDisplay is specified, connections will only be allowed from this host.
// It is possible to control the WebSocket listen address independently, using the
// syntax WithWebSocket("<host>:<port>").
//
// If no TLS credentials are provided, the WebSocket connection runs in unencrypted
// mode. If TLS credentials are provided, the WebSocket connection requires encrypted
// client connections.
func IsWebSocket(enabled bool) *Property {
	return NewProperty("websocket", enabled)
}

// WithWebSocket specifies the port number or host:port value to use for the WebSocket.
// Set the host parameter to and empty string to only use the port. See IsWebSocket
// for more details.
func WithWebSocket(port int, host string) *Property {
	value := strconv.Itoa(port)
	if host != "" {
		value = fmt.Sprintf("%s:%d", host, port)
	}

	return NewProperty("websocket", value)
}

// IsPasswordRequired specifies whether to require that password-based authentication
// is used for client connections.
//
// The password must be set separately using the `set_password` command in the QEMU Monitor.
// The syntax to change your password is:
//
//	set_password <protocol> <password>
//
// Where <protocol> could be either "vnc" or "spice".
//
// If you would like to change <protocol> password expiration, you should use:
//
//	expire_password <protocol> <expiration-time>
//
// Where expiration time could be one of the following options: "now", "never", +seconds
// or Unix time of expiration, e.g. +60 to make password expire in 60 seconds, or 1335196800
// to make password expire on "Mon Apr 23 12:00:00 EDT 2012" (Unix time for this date and time).
//
// You can also use keywords "now" or "never" for the expiration time to allow <protocol>
// password to expire immediately or never expire.
func IsPasswordRequired(enabled bool) *Property {
	return NewProperty("password", enabled)
}

// WithPasswordSecret sets the ID of the Secret object containing the
// password you need to authenticate.
func WithPasswordSecret(secret string) *Property {
	return NewProperty("password-secret", secret)
}

// WithTLSCredentials provides the ID of a set of TLS credentials to use to secure
// the VNC server. They will apply to both the normal VNC server socket and the
// WebSocket socket (if enabled).
//
// Setting TLS credentials will cause the VNC server socket to enable the VeNCrypt
// auth mechanism. The credentials should have been previously created using
// object.TLSCredentials* (see object/tls.go).
func WithTLSCredentials(id string) *Property {
	return NewProperty("tls-creds", id)
}

// WithTLSAuthZ provides the ID of the QAuthZ authorization object against which the
// client's x509 distinguished name will be validated. This object is only resolved
// at time of use, so can be deleted and recreated on the fly while the VNC server
// is active. If missing, it will default to denying access.
func WithTLSAuthZ(id string) *Property {
	return NewProperty("tls-authz", id)
}

// IsSASL enables/disables the requirement that the client use SASL to authenticate
// with the VNC server. The exact choice of authentication method used is controlled
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

// WithSASLAuthZ provides the ID of the QAuthZ authorization object against which
// the client's SASL username will be validated. This object is only resolved at time
// of use, so can be deleted and recreated on the fly while the VNC server is active.
// If missing, it will default to denying access.
func WithSASLAuthZ(id string) *Property {
	return NewProperty("sasl-authz", id)
}

// IsACL is the legacy method for enabling authorization of clients against the x509
// distinguished name and SASL username. It results in the creation of two authz-list
// objects with IDs of vnc.username and vnc.x509dname. The rules for these objects
// must be configured with the HMP ACL commands.
//
// This option is deprecated and should no longer be used. The WithSASLAuthZ and
// WithTLSAuthZ properties are a replacement.
func IsACL(enabled bool) *Property {
	return NewProperty("acl", enabled)
}

// IsLossyCompression enables/disables lossy compression methods (gradient,
// JPEG, ...). If this option is set, VNC client may receive lossy framebuffer
// updates depending on its encoding settings. Enabling this option can save a
// lot of bandwidth at the expense of quality.
func IsLossyCompression(enabled bool) *Property {
	return NewProperty("lossy", enabled)
}

// IsAdaptiveEncoding enables/disables adaptive encodings. Adaptive encodings are
// enabled by default. An adaptive encoding will try to detect frequently updated
// screen regions, and send updates in these regions using a lossy encoding (like JPEG).
// This can be really helpful to save bandwidth when playing videos. Disabling adaptive
// encodings restores the original static behavior of encodings like Tight.
func IsAdaptiveEncoding(enabled bool) *Property {
	return NewProperty("non-adaptive", !enabled)
}

// SharingPolicy represents the options that can be used with the WithSharingPolicy
// property.
type SharingPolicy string

const (
	// SharingPolicyAllowExclusive allows clients to ask for exclusive access.
	// As suggested by the RFB specification, this is implemented by dropping other
	// connections. Connecting multiple clients in parallel requires all clients
	// asking for a shared session (`-shared` switch). This is the default.
	SharingPolicyAllowExclusive SharingPolicy = "allow-exclusive"

	// SharingPolicyForceShared disables exclusive client access. Useful for
	// shared desktop sessions, where you don't want someone forgetting to specify
	// `-shared` and disconnect everybody else.
	SharingPolicyForceShared SharingPolicy = "force-shared"

	// SharingPolicyIgnored completely ignores the shared flag and allows everybody
	// to connect unconditionally. This option doesn't conform to the RFB
	// specification, but is traditional QEMU behavior.
	SharingPolicyIgnored SharingPolicy = "ignored"
)

// WithSharingPolicy sets the display sharing policy. See the comments for each
// SharingPolicy constant for more details.
func WithSharingPolicy(policy SharingPolicy) *Property {
	return NewProperty("share", policy)
}

// WithKeyboardDelay sets keyboard delay, for key down and key up events, in
// milliseconds. The default is 10.
//
// Keyboards are low-bandwidth devices, so this  slowdown can help the device and
// guest to keep up and not lose events in case events are arriving in bulk.
// Possible causes for the latter are flaky network connections, or scripts for
// automated testing.
func WithKeyboardDelay(milliseconds int) *Property {
	return NewProperty("key-delay-ms", milliseconds)
}

// WithAudioDevice uses the audio device associated with the specified id when the
// VNC client requests audio transmission. When not using an AudioDevice option
// (see audiodev.go), this  property must be omitted, otherwise is must be present
// and specify a valid audio device.
func WithAudioDevice(id string) *Property {
	return NewProperty("audiodev", id)
}

// IsPowerControl permits/prevents the remote client to issue shutdown, reboot,
// or reset power control requests.
func IsPowerControl(enabled bool) *Property {
	return NewProperty("power-control", enabled)
}
