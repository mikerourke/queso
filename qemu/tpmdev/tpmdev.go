package tpmdev

import "github.com/mikerourke/queso"

// PassThroughBackend enables access to the host's TPM using the passthrough driver
// and is only available on Linux hosts.
//
// Some notes about using the host's TPM with the passthrough driver:
// The TPM device accessed by the passthrough driver must not be used by any
// other application on the host.
//
// Since the host's firmware (BIOS/UEFI) has already initialized the TPM, the VM's
// firmware (BIOS/UEFI) will not be able to initialize the TPM again and may therefore
// not show a TPM-specific menu that would otherwise allow the user to configure
// the TPM, e.g., allow the user to enable/disable or activate/deactivate the TPM.
//
// Further, if TPM ownership is released from within a VM then the host's TPM will
// get disabled and deactivated. To enable and activate the TPM again afterwards,
// the host has to be rebooted and the user is required to enter the firmware's
// menu to enable and activate the TPM. If the TPM is left disabled and/or deactivated
// most TPM commands will fail.
func PassThroughBackend(id string, properties ...*Property) *queso.Option {
	props := []*queso.Property{{"id", id}}

	for _, property := range properties {
		props = append(props, property.Property)
	}

	return queso.NewOption("tpmdev", "passthrough", props...)
}

// EmulatorBackend enables access to a TPM emulator using Unix domain socket-based
// chardev backend and is only available on Linux hosts.
//
// The chardev parameter specifies the unique ID of a character device
// backend that provides connection to the software TPM server.
func EmulatorBackend(id string, chardev string) *queso.Option {
	return queso.NewOption("tpmdev", "emulator",
		queso.NewProperty("id", id),
		queso.NewProperty("chardev", chardev))
}

// Property represents a property to use with a TPM backend option.
type Property struct {
	*queso.Property
}

// NewProperty returns a new instance of Property.
func NewProperty(key string, value interface{}) *Property {
	return &Property{
		Property: queso.NewProperty(key, value),
	}
}

// WithPath specifies the path to the host's TPM device for a PassThroughBackend,
// i.e., on a Linux host this would be `/dev/tpm0`. The default  is `/dev/tpm0`.
func WithPath(path string) *Property {
	return NewProperty("path", path)
}

// WithCancelPath specifies the path to the host TPM device's sysfs entry allowing
// for cancellation of an ongoing TPM command for a PassThroughBackend. If omitted,
// QEMU will search for the sysfs entry to use.
func WithCancelPath(path string) *Property {
	return NewProperty("cancel-path", path)
}
