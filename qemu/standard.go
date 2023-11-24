package qemu

import (
	"fmt"

	"github.com/mikerourke/queso"
)

// Name sets the name of the guest. This name will be displayed in the SDL window
// caption. The name will also be used for the VNC server. Also, optionally set
// the top visible process name in Linux. Naming of individual threads can also be
// enabled on Linux to aid debugging.
//
//	qemu-system-* -name <name>
func Name(name string) *queso.Option {
	return queso.NewOption("name", name)
}

// UUID sets the system UUID.
//
//	qemu-system-* -uuid <uuid>
func UUID(uuid string) *queso.Option {
	return queso.NewOption("uuid", uuid)
}

// CPU specifies which CPU model to use.
//
//	qemu-system-* -cpu <model>
func CPU(model string) *queso.Option {
	return queso.NewOption("cpu", model)
}

// SGXEPC defines a SGX EPC section.
//
//	qemu-system-* sgx-epc.0.memdev=@var{<memid>},sgx-epc.0.node=@var{<numaid>}
//
// TODO: Verify this is correct.
func SGXEPC(memID string, numaID string) *queso.Option {
	flag := fmt.Sprintf("sgx-epc.0.memdev=@var{%s},sgx-epc.0.node=@var{%s}",
		memID, numaID)

	return queso.NewOption(flag, "").OmitLeadingDash()
}

// SetParameter is used to set the specified arg value for the specified id of
// type group.
//
//	qemu-system-* -set <group>.<id>.<arg>=<value>
func SetParameter(group string, id string, arg string, value string) *queso.Option {
	name := fmt.Sprintf("%s.%s.%s=%s", group, id, arg, value)

	return queso.NewOption("set", name)
}

// DriverProperty sets the default value of the specified driver's property
// to the specified value.
//
// In particular, you can use this to set driver properties for devices which are
// created automatically by the machine model. To create a device which is not
// created automatically and set properties on it, use blockdev.NewDriver.
//
//	qemu-system-* -global driver=<driver>,property=<property>,value=<value>
func DriverProperty(driver string, name string, value string) *queso.Option {
	props := []*queso.Property{
		queso.NewProperty("driver", driver),
		queso.NewProperty("property", name),
		queso.NewProperty("value", value),
	}

	return queso.NewOption("global", "", props...)
}

// Audio is used to configure a default audio backend that will be used whenever the
// audiodev property is not set on a device or machine. In particular, Audio("none")
// ensures that no audio is produced even for machines that have embedded sound hardware.
//
// The following two examples do exactly the same thing, to show how Audio can
// be used to reduce the amount of code required:
//
//	qemu.With(qemu.Audio("pa", queso.NewProperty("model", "sb16)))
//
//	qemu.Use(
//		audiodev.NewPulseAudioBackend("pa"),
//		device.New("sb16").SetProperty("audiodev", "pa"))
//
//	qemu-system-* -audio <driver>
func Audio(driver string, properties ...*queso.Property) *queso.Option {
	allProperties := []*queso.Property{
		queso.NewProperty("driver", driver),
	}
	allProperties = append(allProperties, properties...)

	return queso.NewOption("audio", "", allProperties...)
}

// DataDirectoryPath sets the directory for the BIOS, VGA BIOS and keymaps to
// the specified path. This corresponds to the -L flag passed to the QEMU
// executable.
//
//	qemu-system-* -L <path>
func DataDirectoryPath(path string) *queso.Option {
	return queso.NewOption("L", path)
}
