package display

import (
	"fmt"

	"github.com/mikerourke/queso"
)

// VNCDisplay represents one of the available VNC display options. Normally, if QEMU
// is compiled with graphical window support, it displays output such as guest graphics,
// guest console, and the QEMU monitor in a window.
//
// With this Display, you can have QEMU listen on a VNC display and redirect the VGA
// display over the VNC session. It is very useful to enable the USB tablet device
// when using this option (option qemu.USBDevice). When using the VNC display, you
// must use the qemu.KeyboardLayout parameter to set the keyboard layout if you
// are not using qemu.LanguageEnglishUS.
type VNCDisplay struct {
	// Type is the type of the display.
	Type         string
	properties   []*queso.Property
	displayValue string
}

// NewVNCHostDisplay returns a [VNCDisplay] for which TCP connections will only be
// allowed from a host on the specified port. By convention the TCP port is
// 5900 + <port>. host can be set to "", in which case the server will accept
// connections from any host.
//
//	qemu-system-* -vnc host:port
func NewVNCHostDisplay(port int, host string) *VNCDisplay {
	return &VNCDisplay{
		Type:         "vnc",
		properties:   make([]*queso.Property, 0),
		displayValue: fmt.Sprintf("%s:%d", host, port),
	}
}

// NewVNCNoneDisplay returns a [VNCDisplay] in which VNC is initialized but not
// started. The monitor change command can be used to later start the VNC server.
//
//	qemu-system-* -vnc none
func NewVNCNoneDisplay() *VNCDisplay {
	return &VNCDisplay{
		Type:         "vnc",
		properties:   make([]*queso.Property, 0),
		displayValue: "none",
	}
}

// NewVNCToDisplay returns a [VNCDisplay] for which QEMU will try next available
// VNC displays, until the number to, if the origianlly defined display is not
// available, e.g. port 5900 + <port> is already used by another application.
// By default, to is 0.
//
//	qemu-system-* -vnc to=to
func NewVNCToDisplay(to int) *VNCDisplay {
	return &VNCDisplay{
		Type:         "vnc",
		properties:   make([]*queso.Property, 0),
		displayValue: fmt.Sprintf("to=%d", to),
	}
}

// NewVNCUnixDisplay returns a [VNCDisplay] in which connections will be allowed
// over UNIX domain sockets where path is the location of a Unix socket to listen
// for connections on.
//
//	qemu-system-* -vnc unix:path
func NewVNCUnixDisplay(path string) *VNCDisplay {
	return &VNCDisplay{
		Type:         "vnc",
		properties:   make([]*queso.Property, 0),
		displayValue: fmt.Sprintf("unix:%s", path),
	}
}

func (d *VNCDisplay) option() *queso.Option {
	return queso.NewOption("vnc", d.displayValue, d.properties...)
}

// SetProperty sets arbitrary properties on the [VNCDisplay].
func (d *VNCDisplay) SetProperty(key string, value interface{}) *VNCDisplay {
	d.properties = append(d.properties, queso.NewProperty(key, value))
	return d
}

// SetAudioDevice uses the audio device associated with the specified id when the
// VNC client requests audio transmission. When not using an AudioDevice option
// (see audiodev.go), this  property must be omitted, otherwise is must be present
// and specify a valid audio device.
//
//	qemu-system-* -vnc <display>,audiodev=id
func (d *VNCDisplay) SetAudioDevice(id string) *VNCDisplay {
	d.properties = append(d.properties, queso.NewProperty("audiodev", id))
	return d
}

// SetKeyboardDelay sets keyboard delay, for key down and key up events, in
// milliseconds. The default is 10.
//
// Keyboards are low-bandwidth devices, so this  slowdown can help the device and
// guest to keep up and not lose events in case events are arriving in bulk.
// Possible causes for the latter are flaky network connections, or scripts for
// automated testing.
//
//	qemu-system-* -vnc <display>,key-delay-ms=ms
func (d *VNCDisplay) SetKeyboardDelay(ms int) *VNCDisplay {
	d.properties = append(d.properties, queso.NewProperty("key-delay-ms", ms))
	return d
}

// SetPasswordSecret sets the ID of the Secret object containing the
// password you need to authenticate.
//
//	qemu-system-* -vnc <display>,password-secret=secret
func (d *VNCDisplay) SetPasswordSecret(secret string) *VNCDisplay {
	d.properties = append(d.properties, queso.NewProperty("password-secret", secret))
	return d
}

// SetSASLAuthZ provides the ID of the QAuthZ authorization object against which
// the client's SASL username will be validated. This object is only resolved at time
// of use, so can be deleted and recreated on the fly while the VNC server is active.
// If missing, it will default to denying access.
//
//	qemu-system-* -vnc <display>,sasl-authz=id
func (d *VNCDisplay) SetSASLAuthZ(id string) *VNCDisplay {
	d.properties = append(d.properties, queso.NewProperty("sasl-authz", id))
	return d
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
//
//	qemu-system-* -vnc <display>,share=policy
func (d *VNCDisplay) SetSharingPolicy(policy SharingPolicy) *VNCDisplay {
	d.properties = append(d.properties, queso.NewProperty("share", policy))
	return d
}

// SetTLSAuthZ provides the ID of the QAuthZ authorization object against which the
// client's x509 distinguished name will be validated. This object is only resolved
// at time of use, so can be deleted and recreated on the fly while the VNC server
// is active. If missing, it will default to denying access.
//
//	qemu-system-* -vnc <display>,tls-authz=id
func (d *VNCDisplay) SetTLSAuthZ(id string) *VNCDisplay {
	d.properties = append(d.properties, queso.NewProperty("tls-authz", id))
	return d
}

// SetTLSCredentials provides the ID of a set of TLS credentials to use to secure
// the VNC server. They will apply to both the normal VNC server socket and the
// WebSocket socket (if enabled).
//
// Setting TLS credentials will cause the VNC server socket to enable the VeNCrypt
// auth mechanism. The credentials should have been previously created using
// object.TLSCredentials* (see object/tls.go).
//
//	qemu-system-* -vnc <display>,tls-creds=id
func (d *VNCDisplay) SetTLSCredentials(id string) *VNCDisplay {
	d.properties = append(d.properties, queso.NewProperty("tls-creds", id))
	return d
}

// ToggleACL is the legacy method for enabling authorization of clients against the x509
// distinguished name and SASL username. It results in the creation of two authz-list
// objects with IDs of vnc.username and vnc.x509dname. The rules for these objects
// must be configured with the HMP ACL commands.
//
// Deprecated: Use SetSASLAuthZ or SetTLSAuthZ instead.
//
//	qemu-system-* -vnc <display>,acl=on|off
func (d *VNCDisplay) ToggleACL(enabled bool) *VNCDisplay {
	d.properties = append(d.properties, queso.NewProperty("acl", enabled))
	return d
}

// ToggleAdaptiveEncoding enables/disables adaptive encodings. Adaptive encodings are
// enabled by default. An adaptive encoding will try to detect frequently updated
// screen regions, and send updates in these regions using a lossy encoding (like JPEG).
// This can be really helpful to save bandwidth when playing videos. Disabling adaptive
// encodings restores the original static behavior of encodings like Tight.
//
//	qemu-system-* -vnc <display>,non-adaptive=on|off
func (d *VNCDisplay) ToggleAdaptiveEncoding(enabled bool) *VNCDisplay {
	d.properties = append(d.properties, queso.NewProperty("non-adaptive", !enabled))
	return d
}

// ToggleLossyCompression enables/disables lossy compression methods (gradient,
// JPEG, ...). If this option is set, VNC client may receive lossy framebuffer
// updates depending on its encoding settings. Enabling this option can save a
// lot of bandwidth at the expense of quality.
//
//	qemu-system-* -vnc <display>,lossy=on|off
func (d *VNCDisplay) ToggleLossyCompression(enabled bool) *VNCDisplay {
	d.properties = append(d.properties, queso.NewProperty("lossy", enabled))
	return d
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
//
//	qemu-system-* -vnc <display>,password=on|off
func (d *VNCDisplay) TogglePasswordRequired(required bool) *VNCDisplay {
	d.properties = append(d.properties, queso.NewProperty("password", required))
	return d
}

// TogglePowerControl permits/prevents the remote client to issue shutdown,
// reboot, or reset power control requests.
//
//	qemu-system-* -vnc <display>,power-control=on|off
func (d *VNCDisplay) TogglePowerControl(enabled bool) *VNCDisplay {
	d.properties = append(d.properties, queso.NewProperty("power-control", enabled))
	return d
}

// ToggleReverse specifies whether to connect to a listening VNC client via a "reverse"
// connection. The client is specified by the display. For reverse network connections
// NewHostDisplay(<value>), the <value> argument is a TCP port number, not a display
// number.
//
//	qemu-system-* -vnc <display>,reverse=on|off
func (d *VNCDisplay) ToggleReverse(enabled bool) *VNCDisplay {
	d.properties = append(d.properties, queso.NewProperty("reverse", enabled))
	return d
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
//
//	qemu-system-* -vnc <display>,sasl=on|off
func (d *VNCDisplay) ToggleSASL(required bool) *VNCDisplay {
	d.properties = append(d.properties, queso.NewProperty("sasl", required))
	return d
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
//
//	qemu-system-* -vnc <display>,websocket=on|off
func (d *VNCDisplay) ToggleWebSocket(enabled bool) *VNCDisplay {
	d.properties = append(d.properties, queso.NewProperty("websocket", enabled))
	return d
}
