package qemu

import (
	"fmt"
	"strings"

	"github.com/mikerourke/queso/internal/cli"
)

// Name sets the name of the guest. This name will be displayed in the SDL window
// caption. The name will also be used for the VNC server. Also, optionally set
// the top visible process name in Linux. Naming of individual threads can also be
// enabled on Linux to aid debugging.
func Name(name string) *cli.Option {
	return cli.NewOption("name", name)
}

// UUID sets the system UUID.
func UUID(uuid string) *cli.Option {
	return cli.NewOption("uuid", uuid)
}

// CPU specifies which CPU model to use.
func CPU(model string) *cli.Option {
	return cli.NewOption("cpu", model)
}

// SGXEPC defines a SGX EPC section.
// TODO: Verify this is correct.
func SGXEPC(memID string, numaID string) *cli.Option {
	flag := fmt.Sprintf("sgx-epc.0.memdev=@var{%s},sgx-epc.0.node=@var{%s}",
		memID, numaID)

	return cli.NewOption(flag, "")
}

// SetParameter is used to set the specified arg value for the specified id of
// type group.
func SetParameter(group string, id string, arg string, value string) *cli.Option {
	name := fmt.Sprintf("%s.%s.%s=%s", group, id, arg, value)

	return cli.NewOption("set", name)
}

// DriverProperty sets the default value of the specified driver's property
// to the specified value.
//
// In particular, you can use this to set driver properties for devices which are
// created automatically by the machine model. To create a device which is not
// created automatically and set properties on it, use DeviceDriver.
func DriverProperty(driver string, name string, value string) *cli.Option {
	props := []*cli.Property{
		cli.NewProperty("driver", driver),
		cli.NewProperty("property", name),
		cli.NewProperty("value", value),
	}

	return cli.NewOption("global", "", props...)
}

// SoundHardware enables audio and selected sound hardware.
func SoundHardware(card ...string) *cli.Option {
	name := ""

	switch len(card) {
	case 0:
		panic("at least one card is required for SoundHardware")

	case 1:
		name = card[0]

	default:
		name = strings.Join(card, ",")
	}

	return cli.NewOption("soundhw", name)
}

// DataDirectoryPath sets the directory for the BIOS, VGA BIOS and keymaps to
// the specified path. This corresponds to the -L flag passed to the QEMU
// executable.
func DataDirectoryPath(path string) *cli.Option {
	return cli.NewOption("L", path)
}
