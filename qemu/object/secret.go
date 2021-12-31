package object

import (
	"fmt"

	"github.com/mikerourke/queso"
)

// SecretFormat represents the format that the secret is stored in.
type SecretFormat string

const (
	SecretFormatBase64 SecretFormat = "base64"
	SecretFormatRaw    SecretFormat = "raw"
)

// SecretData defines a secret to store a password, encryption key, or some other
// sensitive data by passing the data in directly via the data parameter.
//
// Example
//
//	qemu.New("qemu-system-x86_64").SetOptions(
//		object.SecretData("sec0", "letmein", object.SecretFormatRaw))
//
// Invocation
//
//	qemu-system-x86_64 -object secret,id=sec0,data=letmein,format=raw
func SecretData(
	id string,
	data string,
	format SecretFormat,
	properties ...*SecretProperty,
) *queso.Option {
	props := []*queso.Property{
		queso.NewProperty("id", id),
		queso.NewProperty("data", data),
		queso.NewProperty("format", format),
	}

	for _, property := range properties {
		props = append(props, property.Property)
	}

	return queso.NewOption("object", "secret", props...)
}

// SecretFile defines a secret to store a password, encryption key, or some other
// sensitive data that is read from the specified file parameter.
//
// Example (with AES Encryption)
//
// First, a master key needs to be created in base64 encoding:
//	openssl rand -base64 32 > key.b64
//	KEY=$(base64 -d key.b64 | hexdump  -v -e '/1 "%02X"')
//
// Each secret to be encrypted needs to have a random initialization vector generated.
// These do not need to be kept secret:
//	openssl rand -base64 16 > iv.b64
//	IV=$(base64 -d iv.b64 | hexdump  -v -e '/1 "%02X"')
//
// The secret to be defined can now be encrypted, in this case we’re telling
// openssl to base64 encode the result, but it could be left as raw bytes if desired.
//	SECRET=$(printf "letmein" | openssl enc -aes-256-cbc -a -K $KEY -iv $IV)
//
// To utilize this via the qemu library:
//	qemu.New("qemu-system-x86_64").SetOptions(
//		object.SecretFile("secmaster0", "key.b64", object.SecretFormatBase64),
//		object.SecretData("sec0", "$SECRET", object.SecretFormatBase64,
//			object.WithAESEncryption("secmaster0", "$(<iv.b64)")))
//
// Invocation
//	qemu-system-x86_64 \
//		-object secret,id=secmaster0,format=base64,file=key.b64 \
//		-object secret,id=sec0,keyid=secmaster0,format=base64,data=$SECRET,iv=$(<iv.b64)
func SecretFile(
	id string,
	file string,
	format SecretFormat,
	properties ...*SecretProperty,
) *queso.Option {
	props := []*queso.Property{
		queso.NewProperty("id", id),
		queso.NewProperty("file", file),
		queso.NewProperty("format", format),
	}

	for _, property := range properties {
		props = append(props, property.Property)
	}

	return queso.NewOption("object", "secret", props...)
}

// SecretProperty represents a property that can be passed to Secret.
type SecretProperty struct {
	*queso.Property
}

// NewSecretProperty returns a new instance of SecretProperty.
func NewSecretProperty(key string, value interface{}) *SecretProperty {
	return &SecretProperty{
		Property: queso.NewProperty(key, value),
	}
}

// WithAESEncryption encrypts the data associated with a secret using the
// AES-256-CBC cipher.
func WithAESEncryption(keyID string, iv string) *SecretProperty {
	key := fmt.Sprintf("keyid=%s,iv", keyID)

	return NewSecretProperty(key, iv)
}
