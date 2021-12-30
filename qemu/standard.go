package qemu

import (
	"fmt"
	"strings"

	"github.com/mikerourke/queso"
)

// Name sets the name of the guest. This name will be displayed in the SDL window
// caption. The name will also be used for the VNC server. Also, optionally set
// the top visible process name in Linux. Naming of individual threads can also be
// enabled on Linux to aid debugging.
func Name(name string) *queso.Option {
	return queso.NewOption("name", name)
}

// UUID sets the system UUID.
func UUID(uuid string) *queso.Option {
	return queso.NewOption("uuid", uuid)
}

// CPU specifies which CPU model to use.
func CPU(model string) *queso.Option {
	return queso.NewOption("cpu", model)
}

// SGXEPC defines a SGX EPC section.
// TODO: Verify this is correct.
func SGXEPC(memoryDeviceID string, numaID string) *queso.Option {
	flag := fmt.Sprintf("sgx-epc.0.memdev=@var{%s},sgx-epc.0.node=@var{%s}", memoryDeviceID, numaID)

	return queso.NewOption(flag, "")
}

// SetParameter is used to set the specified arg value for the specified id of
// type group.
func SetParameter(group string, id string, arg string, value string) *queso.Option {
	name := fmt.Sprintf("%s.%s.%s=%s", group, id, arg, value)

	return queso.NewOption("set", name)
}

// SetDriverProperty sets the default value of the specified driver's property
// to the specified value.
//
// In particular, you can use this to set driver properties for devices which are
// created automatically by the machine model. To create a device which is not
// created automatically and set properties on it, use DeviceDriver.
func SetDriverProperty(driver string, name string, value string) *queso.Option {
	props := []*queso.Property{
		{"driver", driver},
		{"property", name},
		{"value", value},
	}

	return queso.NewOption("global", "", props...)
}

// SoundHardware enables audio and selected sound hardware.
func SoundHardware(card ...string) *queso.Option {
	name := ""

	switch len(card) {
	case 0:
		panic("at least one card is required for SoundHardware")

	case 1:
		name = card[0]

	default:
		name = strings.Join(card, ",")
	}

	return queso.NewOption("soundhw", name)
}
