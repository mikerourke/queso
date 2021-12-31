package object

import "github.com/mikerourke/queso"

// CryptoBuiltIn creates a cryptodev backend which executes crypto operation from
// the QEMU cipher APIS. The id parameter is a unique ID that will be used to
// reference this cryptodev backend from the virtio-crypto device.
func CryptoBuiltIn(id string, properties ...*CryptoProperty) *queso.Option {
	props := []*queso.Property{queso.NewProperty("id", id)}

	for _, property := range properties {
		props = append(props, property.Property)
	}

	return queso.NewOption("object", "cryptodev-backend-builtin", props...)
}

// CryptoVHostUser creates a vhost-user cryptodev backend, backed by a character
// device with ID chardev. The id parameter is a unique ID that will be used to
// reference this cryptodev backend from the virtio-crypto device. The chardev
// should be a Unix domain socket backed one. The vhost-user uses a specifically
// defined protocol to pass vhost ioctl replacement messages to an application
// on the other end of the socket.
func CryptoVHostUser(id string, chardev string, properties ...*CryptoProperty) *queso.Option {
	props := []*queso.Property{
		queso.NewProperty("id", id),
		queso.NewProperty("chardev", chardev),
	}

	for _, property := range properties {
		props = append(props, property.Property)
	}

	return queso.NewOption("object", "cryptodev-vhost-user", props...)
}

// CryptoProperty represents a property that can be passed to a cryptodev backend.
type CryptoProperty struct {
	*queso.Property
}

// NewCryptoProperty returns a new instance of CryptoProperty.
func NewCryptoProperty(key string, value interface{}) *CryptoProperty {
	return &CryptoProperty{
		Property: queso.NewProperty(key, value),
	}
}

// WithQueues specifies the queue number of cryptodev backend for CryptoBuiltIn or
// CryptoVHostUser. If omitted, the default value is 1.
func WithQueues(queues int) *CryptoProperty {
	return NewCryptoProperty("queues", queues)
}
