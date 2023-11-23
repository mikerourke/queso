package display

import (
	"fmt"

	"github.com/mikerourke/queso/qemu/cli"
)

type VNCDisplay struct {
	displayValue string
	properties   []*cli.Property
}

func NewVNCToDisplay(to int) *VNCDisplay {
	return &VNCDisplay{
		displayValue: fmt.Sprintf("to=%d", to),
		properties:   make([]*cli.Property, 0),
	}
}

func NewVNCHostDisplay(display string) *VNCDisplay {
	return &VNCDisplay{
		displayValue: fmt.Sprintf("host=%s", display),
		properties:   make([]*cli.Property, 0),
	}
}

func NewVNCUnixDisplay(path string) *VNCDisplay {
	return &VNCDisplay{
		displayValue: fmt.Sprintf("unix=%s", path),
		properties:   make([]*cli.Property, 0),
	}
}

func NewVNCNoneDisplay() *VNCDisplay {
	return &VNCDisplay{
		displayValue: "none",
		properties:   make([]*cli.Property, 0),
	}
}

// ToggleReverse specifies whether to connect to a listening VNC client via a "reverse"
// connection. The client is specified by the display. For reverse network connections
// NewHostDisplay(<value>), the <value> argument is a TCP port number, not a display
// number.
func (v *VNCDisplay) ToggleReverse(enabled bool) *VNCDisplay {
	v.properties = append(v.properties, cli.NewProperty("reverse", enabled))
	return v
}

// ToggleWebSocket specifies whether to open an additional TCP listening port dedicated
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
func (v *VNCDisplay) ToggleWebSocket(enabled bool) *VNCDisplay {
	v.properties = append(v.properties, cli.NewProperty("websocket", enabled))
	return v
}

// TogglePasswordRequired specifies whether to require that password-based
// authentication is used for client connections.
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
func (v *VNCDisplay) TogglePasswordRequired(required bool) *VNCDisplay {
	v.properties = append(v.properties, cli.NewProperty("password", required))
	return v
}

// SetPasswordSecret sets the ID of the Secret object containing the
// password you need to authenticate.
func (v *VNCDisplay) SetPasswordSecret(secret string) *VNCDisplay {
	v.properties = append(v.properties, cli.NewProperty("password-secret", secret))
	return v
}

// SetTLSCredentials provides the ID of a set of TLS credentials to use to secure
// the VNC server. They will apply to both the normal VNC server socket and the
// WebSocket socket (if enabled).
//
// Setting TLS credentials will cause the VNC server socket to enable the VeNCrypt
// auth mechanism. The credentials should have been previously created using
// object.TLSCredentials* (see object/tls.go).
func (v *VNCDisplay) SetTLSCredentials(id string) *VNCDisplay {
	v.properties = append(v.properties, cli.NewProperty("tls-creds", id))
	return v
}

// SetTLSAuthZ provides the ID of the QAuthZ authorization object against which the
// client's x509 distinguished name will be validated. This object is only resolved
// at time of use, so can be deleted and recreated on the fly while the VNC server
// is active. If missing, it will default to denying access.
func (v *VNCDisplay) SetTLSAuthZ(id string) *VNCDisplay {
	v.properties = append(v.properties, cli.NewProperty("tls-authz", id))
	return v
}

// ToggleSASL enables/disables the requirement that the client use SASL to authenticate
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
func (v *VNCDisplay) ToggleSASL(required bool) *VNCDisplay {
	v.properties = append(v.properties, cli.NewProperty("sasl", required))
	return v
}

// SetSASLAuthZ provides the ID of the QAuthZ authorization object against which
// the client's SASL username will be validated. This object is only resolved at time
// of use, so can be deleted and recreated on the fly while the VNC server is active.
// If missing, it will default to denying access.
func (v *VNCDisplay) SetSASLAuthZ(id string) *VNCDisplay {
	v.properties = append(v.properties, cli.NewProperty("sasl-authz", id))
	return v
}

// ToggleACL is the legacy method for enabling authorization of clients against the x509
// distinguished name and SASL username. It results in the creation of two authz-list
// objects with IDs of vnc.username and vnc.x509dname. The rules for these objects
// must be configured with the HMP ACL commands.
//
// Deprecated: Use SetSASLAuthZ or SetTLSAuthZ instead.
func (v *VNCDisplay) ToggleACL(enabled bool) *VNCDisplay {
	v.properties = append(v.properties, cli.NewProperty("acl", enabled))
	return v
}

// ToggleLossyCompression enables/disables lossy compression methods (gradient,
// JPEG, ...). If this option is set, VNC client may receive lossy framebuffer
// updates depending on its encoding settings. Enabling this option can save a
// lot of bandwidth at the expense of quality.
func (v *VNCDisplay) ToggleLossyCompression(enabled bool) *VNCDisplay {
	v.properties = append(v.properties, cli.NewProperty("lossy", enabled))
	return v
}

// ToggleAdaptiveEncoding enables/disables adaptive encodings. Adaptive encodings are
// enabled by default. An adaptive encoding will try to detect frequently updated
// screen regions, and send updates in these regions using a lossy encoding (like JPEG).
// This can be really helpful to save bandwidth when playing videos. Disabling adaptive
// encodings restores the original static behavior of encodings like Tight.
func (v *VNCDisplay) ToggleAdaptiveEncoding(enabled bool) *VNCDisplay {
	v.properties = append(v.properties, cli.NewProperty("non-adaptive", !enabled))
	return v
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

// SetSharingPolicy sets the display sharing policy. See the comments for each
// SharingPolicy constant for more details.
func (v *VNCDisplay) SetSharingPolicy(policy SharingPolicy) *VNCDisplay {
	v.properties = append(v.properties, cli.NewProperty("share", policy))
	return v
}

// SetKeyboardDelay sets keyboard delay, for key down and key up events, in
// milliseconds. The default is 10.
//
// Keyboards are low-bandwidth devices, so this  slowdown can help the device and
// guest to keep up and not lose events in case events are arriving in bulk.
// Possible causes for the latter are flaky network connections, or scripts for
// automated testing.
func (v *VNCDisplay) SetKeyboardDelay(ms int) *VNCDisplay {
	v.properties = append(v.properties, cli.NewProperty("key-delay-ms", ms))
	return v
}

// SetAudioDevice uses the audio device associated with the specified id when the
// VNC client requests audio transmission. When not using an AudioDevice option
// (see audiodev.go), this  property must be omitted, otherwise is must be present
// and specify a valid audio device.
func (v *VNCDisplay) SetAudioDevice(id string) *VNCDisplay {
	v.properties = append(v.properties, cli.NewProperty("audiodev", id))
	return v
}

// TogglePowerControl permits/prevents the remote client to issue shutdown,
// reboot, or reset power control requests.
func (v *VNCDisplay) TogglePowerControl(enabled bool) *VNCDisplay {
	v.properties = append(v.properties, cli.NewProperty("power-control", enabled))
	return v
}
