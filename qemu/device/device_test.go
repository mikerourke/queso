package device

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUse(t *testing.T) {
	result := Use("e1000",
		NewProperty("mac", "8E:4E:E1:EA:B7:BB"),
		NewProperty("netdev", "net0")).ArgsString()
	assert.Equal(t, result, "-device e1000,mac=8E:4E:E1:EA:B7:BB,netdev=net0")

	result = Use("virtio-vga").ArgsString()
	assert.Equal(t, result, "-device virtio-vga")

	result = Use("intel-hda").ArgsString()
	assert.Equal(t, result, "-device intel-hda")

	result = Use("hda-duplex", NewProperty("audiodev", "audio0")).ArgsString()
	assert.Equal(t, result, "-device hda-duplex,audiodev=audio0")

	result = Use("usb-tablet", WithBus("usb-bus.0")).ArgsString()
	assert.Equal(t, result, "-device usb-tablet,bus=usb-bus.0")

	result = Use("usb-mouse", WithBus("usb-bus.0")).ArgsString()
	assert.Equal(t, result, "-device usb-mouse,bus=usb-bus.0")

	result = Use("nec-usb-xhci", WithID("usb-controller-0")).ArgsString()
	assert.Equal(t, result, "-device nec-usb-xhci,id=usb-controller-0")
}
