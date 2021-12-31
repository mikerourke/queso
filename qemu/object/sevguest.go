package object

import "github.com/mikerourke/queso"

// SEVGuest creates a Secure Encrypted Virtualization (SEV) guest object, which
// can be used to provide the guest memory encryption support on AMD processors.
//
// When memory encryption is enabled, one of the physical address bit (aka the
// C-bit) is utilized to mark if a memory page is protected. The cBitPosition
// parameter is used to provide the C-bit position. The C-bit position is Host
// family dependent hence user must provide this value. On EPYC, the value
// should be 47.
//
// When memory encryption is enabled, we loose certain bits in physical address
// space. The reducedPhysBits parameter is used to provide the number of bits
// we lose in physical address space. Similar to C-bit, the value is Host family
// dependent. On EPYC, the value should be 5.
func SEVGuest(
	id string,
	cBitPosition int,
	reducedPhysBits int,
	properties ...*SEVProperty,
) *queso.Option {
	props := []*queso.Property{
		queso.NewProperty("id", id),
		queso.NewProperty("cbitpos", cBitPosition),
		queso.NewProperty("reduced-phys-bits", reducedPhysBits),
	}

	for _, property := range properties {
		props = append(props, property.Property)
	}

	return queso.NewOption("object", "sev-guest", props...)
}

// SEVProperty is used to define properties for a SEVGuest.
type SEVProperty struct {
	*queso.Property
}

// NewSEVProperty returns a new instance of SEVProperty.
func NewSEVProperty(key string, value interface{}) *SEVProperty {
	return &SEVProperty{
		Property: queso.NewProperty(key, value),
	}
}

// WithSEVDeviceFile provides the device file to use for communicating with the
// SEV firmware running inside AMD Secure Processor for a SEVGuest. The default
// device is `/dev/sev`. If hardware supports memory encryption then /dev/sev
// devices are created by CCP driver.
func WithSEVDeviceFile(file string) *SEVProperty {
	return NewSEVProperty("sev-device", file)
}

// WithPolicy provides the guest policy to be enforced by the SEV firmware and
// restrict what configuration and operational commands can be performed on this
// guest by the hypervisor for a SEVGuest. The policy should be provided by the
// guest owner and  is bound to the guest and cannot be changed throughout the
// lifetime of the guest. The default is 0.
func WithPolicy(policy int) *SEVProperty {
	return NewSEVProperty("policy", policy)
}

// WithHandle provides the handle of the guest from which to share the key when
// the policy specified with WithPolicy allows sharing the key with another
// SEVGuest.
func WithHandle(handle string) *SEVProperty {
	return NewSEVProperty("handle", handle)
}

// WithDHCertFile provides the guest owner's Public Diffie-Hillman key defined
// in the SEV specification for a SEVGuest. The file must be encoded in base64.
// This property should be used with WithSessionFile.
func WithDHCertFile(file string) *SEVProperty {
	return NewSEVProperty("dh-cert-file", file)
}

// WithSessionFile provides the guest owner's Public Diffie-Hillman key defined
// in the SEV specification for a SEVGuest. The file must be encoded in base64.
// This property should be used with WithDHCertFile.
func WithSessionFile(file string) *SEVProperty {
	return NewSEVProperty("session-file", file)
}

// IsKernelHashesAdded specifies whether to add the hashes of given kernel/initrd/
// cmdline to a designated guest firmware page for measured Linux boot with
// qemu.Kernel. This property is disabled by default.
func IsKernelHashesAdded(added bool) *SEVProperty {
	return NewSEVProperty("kernel-hashes", added)
}
