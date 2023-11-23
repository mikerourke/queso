package qemu

import (
	"fmt"

	"github.com/mikerourke/queso/qemu/cli"
)

// Name sets the name of the guest. This name will be displayed in the SDL window
// caption. The name will also be used for the VNC server. Also, optionally set
// the top visible process name in Linux. Naming of individual threads can also be
// enabled on Linux to aid debugging.
//
//	qemu-system-* -name <name>
func Name(name string) *cli.Option {
	return cli.NewOption("name", name)
}

// UUID sets the system UUID.
//
//	qemu-system-* -uuid <uuid>
func UUID(uuid string) *cli.Option {
	return cli.NewOption("uuid", uuid)
}

// CPU specifies which CPU model to use.
//
//	qemu-system-* -cpu <model>
func CPU(model string) *cli.Option {
	return cli.NewOption("cpu", model)
}

// SGXEPC defines a SGX EPC section.
//
//	qemu-system-* sgx-epc.0.memdev=@var{<memid>},sgx-epc.0.node=@var{<numaid>}
//
// TODO: Verify this is correct.
func SGXEPC(memID string, numaID string) *cli.Option {
	flag := fmt.Sprintf("sgx-epc.0.memdev=@var{%s},sgx-epc.0.node=@var{%s}",
		memID, numaID)

	return cli.NewOption(flag, "").OmitLeadingDash()
}

// SetParameter is used to set the specified arg value for the specified id of
// type group.
//
//	qemu-system-* -set <group>.<id>.<arg>=<value>
func SetParameter(group string, id string, arg string, value string) *cli.Option {
	name := fmt.Sprintf("%s.%s.%s=%s", group, id, arg, value)

	return cli.NewOption("set", name)
}

// DriverProperty sets the default value of the specified driver's property
// to the specified value.
//
// In particular, you can use this to set driver properties for devices which are
// created automatically by the machine model. To create a device which is not
// created automatically and set properties on it, use blockdev.NewDriver.
//
//	qemu-system-* -global driver=<driver>,property=<property>,value=<value>
func DriverProperty(driver string, name string, value string) *cli.Option {
	props := []*cli.Property{
		cli.NewProperty("driver", driver),
		cli.NewProperty("property", name),
		cli.NewProperty("value", value),
	}

	return cli.NewOption("global", "", props...)
}

// Audio is used to configure a default audio backend that will be used whenever the
// audiodev property is not set on a device or machine. In particular, Audio("none")
// ensures that no audio is produced even for machines that have embedded sound hardware.
//
// The following two examples do exactly the same thing, to show how Audio can
// be used to reduce the amount of code required:
//
//	qemu.With(qemu.Audio("pa", cli.NewProperty("model", "sb16)))
//
//	qemu.Use(
//		audiodev.NewPulseAudioBackend("pa"),
//		device.New("sb16").SetProperty("audiodev", "pa"))
//
//	qemu-system-* -audio <driver>
func Audio(driver string, properties ...*cli.Property) *cli.Option {
	allProperties := []*cli.Property{
		cli.NewProperty("driver", driver),
	}
	allProperties = append(allProperties, properties...)

	return cli.NewOption("audio", "", allProperties...)
}

// DataDirectoryPath sets the directory for the BIOS, VGA BIOS and keymaps to
// the specified path. This corresponds to the -L flag passed to the QEMU
// executable.
//
//	qemu-system-* -L <path>
func DataDirectoryPath(path string) *cli.Option {
	return cli.NewOption("L", path)
}
