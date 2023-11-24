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
// get disabled and deactivated. To enable and activate the TPM again afterward,
// the host has to be rebooted and the user is required to enter the firmware's
// menu to enable and activate the TPM. If the TPM is left disabled and/or deactivated
// most TPM commands will fail.
type PassThroughBackend struct {
	*Backend
}

// NewPassThroughBackend returns a new instance of [PassThroughBackend]. id
// is the unique identifier for the TPM device backend.
//
//	qemu-system-* -tpmdev passthrough,id=id
func NewPassThroughBackend(id string) *PassThroughBackend {
	backend := New("passthrough")

	backend.properties = append(backend.properties,
		queso.NewProperty("id", id))

	return &PassThroughBackend{backend}
}

// SetPath specifies the path to the host's TPM device, i.e., on a Linux host
// this would be `/dev/tpm0`. The default value is `/dev/tpm0`.
//
//	qemu-system-* -tpmdev passthrough,path=path
func (b *PassThroughBackend) SetPath(path string) *PassThroughBackend {
	b.properties = append(b.properties, queso.NewProperty("path", path))
	return b
}

// SetCancelPath specifies the path to the host TPM device's sysfs entry allowing
// for cancellation of an ongoing TPM command. By default, QEMU will search for the
// sysfs entry to use.
//
//	qemu-system-* -tpmdev passthrough,cancel-path=path
func (b *PassThroughBackend) SetCancelPath(path string) *PassThroughBackend {
	b.properties = append(b.properties, queso.NewProperty("cancel-path", path))
	return b
}
