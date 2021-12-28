package objects

import "github.com/mikerourke/queso"

// NewTLSOption returns an Option that represents a TLS object to pass to QEMU.
func NewTLSOption(name string, id string, properties ...*TLSProperty) *queso.Option {
	props := []*queso.Property{{"id", id}}

	for _, property := range properties {
		props = append(props, property.Property)
	}

	return queso.NewOption("object", name, props...)
}

func newTLSCredentialsOption(
	name string,
	id string,
	endpoint string,
	dir string,
	properties ...*TLSProperty,
) *queso.Option {
	props := []*TLSProperty{
		NewTLSProperty("endpoint", endpoint),
		NewTLSProperty("dir", dir),
	}

	if properties != nil {
		props = append(props, properties...)
	}

	return NewTLSOption(name, id, props...)
}

// TLSCredentialsAnon creates a TLS anonymous credentials object, which can be used to provide
// TLS support on network backends. The id parameter is a unique ID which network backends
// will use to access the credentials. The endpoint is either server or client depending
// on whether the QEMU network backend that uses the credentials will be acting as a
// client or as a server. If verify-peer is enabled (the default) then once the handshake
// is completed, the peer credentials will be verified, though this is a no-op for
// anonymous credentials.
//
// The dir parameter tells QEMU where to find the credential files. For server
// endpoints, this directory may contain a file dh-params.pem providing diffie-hellman
// parameters to use for the TLS server. If the file is missing, QEMU will generate
// a set of DH parameters at startup. This is a computationally expensive operation
// that consumes random pool entropy, so it is recommended that a persistent set of
// parameters be generated upfront and saved.
func TLSCredentialsAnon(id string, endpoint string, dir string, properties ...*TLSProperty) *queso.Option {
	return newTLSCredentialsOption("tls-creds-anon", id, endpoint, dir, properties...)
}

// TLSCredentialsPSK creates a TLS Pre-Shared Keys (PSK) credentials object, which can be
// used to provide TLS support on network backends. The id parameter is a unique ID
// which network backends will use to access the credentials. The endpoint is either
// server or client depending on whether the QEMU network backend that uses the
// credentials will be acting as a client or as a server. For clients only, username
// is the username which will be sent to the server. If omitted it defaults to “qemu”.
//
// The dir parameter tells QEMU where to find the keys file. It is called `dir/keys.psk`
// and contains "username:key" pairs. This file can most easily be created using the
// GnuTLS `psktool` program.
//
// For server endpoints, dir may also contain a file dh-params.pem providing diffie-hellman
// parameters to use for the TLS server. If the file is missing, QEMU will generate a
// set of DH parameters at startup. This is a computationally expensive operation that
// consumes random pool entropy, so it is recommended that a persistent set of parameters
// be generated up front and saved.
func TLSCredentialsPSK(id string, endpoint string, dir string, properties ...*TLSProperty) *queso.Option {
	return newTLSCredentialsOption("tls-creds-psk", id, endpoint, dir, properties...)
}

// TLSCredentialsX509 creates a TLS anonymous credentials object, which can be used to provide
// TLS support on network backends. The id parameter is a unique ID which network backends
// will use to access the credentials. The endpoint is either server or client depending
// on whether the QEMU network backend that uses the credentials will be acting as a
// client or as a server. If verify-peer is enabled (the default) then once the handshake
// is completed, the peer credentials will be verified. With x509 certificates, this
// implies that the clients must be provided with valid client certificates too.
//
// The dir parameter tells QEMU where to find the credential files. For server
// endpoints, this directory may contain a file `dh-params.pem` providing diffie-hellman
// parameters to use for the TLS server. If the file is missing, QEMU will generate
// a set of DH parameters at startup. This is a computationally expensive operation
// that consumes random pool entropy, so it is recommended that a persistent set of
// parameters be generated upfront and saved.
//
// For x509 certificate credentials the directory will contain further files providing
// the x509 certificates. The certificates must be stored in PEM format, in filenames
//  - `ca-cert.pem`
//  - `ca-crl.pem` (optional)
//  - `server-cert.pem` (only servers)
//  - `server-key.pem` (only servers)
//  - `client-cert.pem` (only clients)
//  - `client-key.pem` (only clients)
//
// For the `server-key.pem` and `client-key.pem` files which contain sensitive private
// keys, it is possible to use an encrypted version by providing the WithPasswordID option.
// This provides the ID of a previously created secret object containing the
// password for decryption.
//
// The priority parameter allows to override the global default priority used by
// gnutls. This can be useful if the system administrator needs to use a weaker
// set of crypto priorities for QEMU without potentially forcing the weakness onto
// all applications. Or conversely if one wants a stronger default for QEMU than
// for all other applications, they can do this through this parameter. Its format
// is a gnutls priority string as described at https://gnutls.org/manual/html_node/Priority-Strings.html.
func TLSCredentialsX509(id string, endpoint string, dir string, properties ...*TLSProperty) *queso.Option {
	return newTLSCredentialsOption("tls-creds-x509", id, endpoint, dir, properties...)
}

// TLSCipherSuites creates a TLS cipher suites object, which can be used to control
// the TLS cipher/protocol algorithms that applications are permitted to use.
//
// The id parameter is a unique ID which frontends will use to access the ordered
// list of permitted TLS cipher suites from the host.
//
// The priority parameter allows to override the global default priority used by
// gnutls. This can be useful if the system administrator needs to use a weaker set
// of crypto priorities for QEMU without potentially forcing the weakness onto all
// applications. Or conversely if one wants a stronger default for QEMU than for
// all other applications, they can do this through this parameter. Its format is a
// gnutls priority string as described at https://gnutls.org/manual/html_node/Priority-Strings.html.
//
// An example of use of this object is to control UEFI HTTPS Boot. The tls-cipher-suites
// object exposes the ordered list of permitted TLS cipher suites from the host
// side to the guest firmware, via fw_cfg. The list is represented as an array of
// IANA_TLS_CIPHER objects. The firmware uses the IANA_TLS_CIPHER array for configuring
// guest-side TLS.
func TLSCipherSuites(id string, priority string) *queso.Option {
	return NewTLSOption("tls-cipher-suites", id, NewTLSProperty("priority", priority))
}

// TLSProperty represents a property that can be used with an audio device
// option.
type TLSProperty struct {
	*queso.Property
}

// NewTLSProperty returns a new instance of an TLSProperty.
func NewTLSProperty(key string, value interface{}) *TLSProperty {
	return &TLSProperty{
		Property: &queso.Property{key, value},
	}
}

// IsVerifyPeer indicates that once the handshake is completed, the peer
// credentials will be verified, though this is a no-op for anonymous credentials.
//
// // This property is only valid for TLSCredentialsAnon and TLSCredentialsX509.
func IsVerifyPeer(value bool) *TLSProperty {
	return NewTLSProperty("verify-peer", value)
}

// WithPasswordID provides the ID of a previously created secret object containing
// the password for decryption.
//
// This property is only valid for TLSCredentialsX509.
func WithPasswordID(value string) *TLSProperty {
	return NewTLSProperty("passwordid", value)
}

// WithUsername represents the username which will be sent to the server. If
// omitted it defaults to "qemu".
//
// This property is only valid for TLSCredentialsPSK.
func WithUsername(value string) *TLSProperty {
	return NewTLSProperty("username", value)
}
