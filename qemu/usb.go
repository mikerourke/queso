package qemu

import "github.com/mikerourke/queso"

// EnableUSB enables USB emulation on machine types with an on-board USB host controller
// (if not enabled by default). Note that on-board USB host controllers may not support
// USB 3.0.
//
//	qemu-system-* -usb
func EnableUSB() *queso.Option {
	return queso.NewOption("usb", "")
}

// USBDeviceName represents the device name that can be passed to the
// [WithUSBDevice] option.
type USBDeviceName string

const (
	// USBDeviceBraille represents a Braille device. This will use BrlAPI to display
	// the braille output on a real or fake device (i.e. it also creates a corresponding
	// braille character device automatically beside the USB device).
	USBDeviceBraille USBDeviceName = "braille"

	// USBDeviceKeyboard represents a standard USB keyboard. It will override the
	// PS/2 keyboard (if present).
	USBDeviceKeyboard USBDeviceName = "keyboard"

	// USBDeviceMouse represents a Virtual Mouse. This will override the PS/2 mouse
	// emulation when activated.
	USBDeviceMouse USBDeviceName = "mouse"

	// USBDeviceTablet represents a pointer device that uses absolute coordinates (like
	// a touchscreen). This means QEMU is able to report the mouse position without
	// having to grab the mouse. Also overrides the PS/2 mouse emulation when activated.
	USBDeviceTablet USBDeviceName = "tablet"

	// USBDeviceWacomTablet represents a Wacom PenPartner USB tablet.
	USBDeviceWacomTablet USBDeviceName = "wacom-tablet"
)

// WithUSBDevice add the USB device name, and enables an on-board USB controller
// if possible and necessary. You can achieve this using NewMachine as well:
//
//	qemu.Use(NewMachine().SetProperty("usb", true))
//
// Note that this option is mainly intended for the user's convenience only.
// More fine-grained control can be achieved by selecting a USB host controller
// (if necessary) and the desired USB device via the device.New option instead.
//
//	qemu-system-* -usbdevice name
func WithUSBDevice(name USBDeviceName) *queso.Option {
	return queso.NewOption("usbdevice", string(name))
}
