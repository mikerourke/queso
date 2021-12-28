package objects

import "github.com/mikerourke/queso"

// RNGBuiltIn creates a random number generator backend which obtains entropy from
// QEMU builtin functions. The id parameter is a unique ID that will be used to
// reference this entropy backend from the virtio-rng device. By default, the
// virtio-rng device uses this RNG backend.
func RNGBuiltIn(id string) *queso.Option {
	return queso.NewOption("object", "rng-built-in", queso.NewProperty("id", id))
}

// RNGRandom creates a random number generator backend which obtains entropy from a
// device on the host. The id parameter is a unique ID that will be used to
// reference this entropy backend from the virtio-rng device.
//
// The filename parameter represents the file to get random values from. If empty,
// uses `/dev/urandom`.
func RNGRandom(id string, filename string) *queso.Option {
	props := []*queso.Property{{"id", id}}

	if filename != "" {
		props = append(props, queso.NewProperty("filename", filename))
	}

	return queso.NewOption("object", "rng-random", props...)
}

// RNGExternalDaemon creates a random number generator backend which obtains entropy
// from an external daemon running on the host. The id parameter is a unique ID
// that will be used to reference this entropy backend from the virtio-rng device.
// The chardev parameter is the unique ID of a character device backend that
// provides the connection to the RNG daemon.
func RNGExternalDaemon(id string, chardev string) *queso.Option {
	return queso.NewOption("object", "rng-egd", queso.NewProperty("chardev", chardev))
}
