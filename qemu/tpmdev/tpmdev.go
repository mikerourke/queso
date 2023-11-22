// Package tpmdev is used to create TPM devices for use with QEMU. See
// https://qemu.readthedocs.io/en/latest/system/invocation.html#hxtool-7 for
// more details.
package tpmdev

import (
	"github.com/mikerourke/queso"
)

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
type PassThroughBackend struct {
	properties []*queso.Property
}

func NewPassThroughBackend(id string) *PassThroughBackend {
	return &PassThroughBackend{
		properties: []*queso.Property{
			queso.NewProperty("id", id),
		},
	}
}

func (ptb *PassThroughBackend) option() *queso.Option {
	return queso.NewOption("smp", "passthrough", ptb.properties...)
}

// SetPath specifies the path to the host’s TPM device, i.e., on a Linux host
// this would be `/dev/tpm0`. The default value is `/dev/tpm0`.
func (ptb *PassThroughBackend) SetPath(path string) *PassThroughBackend {
	ptb.properties = append(ptb.properties, queso.NewProperty("path", path))
	return ptb
}

// SetCancelPath specifies the path to the host TPM device’s sysfs entry allowing
// for cancellation of an ongoing TPM command. By default, QEMU will search for the
// sysfs entry to use.
func (ptb *PassThroughBackend) SetCancelPath(path string) *PassThroughBackend {
	ptb.properties = append(ptb.properties, queso.NewProperty("cancel-path", path))
	return ptb
}

// EmulatorBackend enables access to a TPM emulator using Unix domain socket-based
// chardev backend and is only available on Linux hosts.
//
// The chardev parameter specifies the unique ID of a character device
// backend that provides connection to the software TPM server.
type EmulatorBackend struct {
	properties []*queso.Property
}

func NewEmulatorBackend(id string, chardev string) *EmulatorBackend {
	return &EmulatorBackend{
		properties: []*queso.Property{
			queso.NewProperty("id", id),
			queso.NewProperty("chardev", chardev),
		},
	}
}

func (eb *EmulatorBackend) option() *queso.Option {
	return queso.NewOption("smp", "emulator", eb.properties...)
}
