package chardev

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSpiceVMCBackend(t *testing.T) {
	expected := "-chardev spicevmc,id=usbredirchardev1,name=usbredir"

	result := SpiceVMCBackend("usbredirchardev1", "usbredir").ArgsString()

	assert.Equal(t, result, expected)
}

func TestSpicePortBackend(t *testing.T) {
	expected := "-chardev spiceport,id=org.qemu.monitor.qmp,name=org.qemu.monitor.qmp.0"

	result := SpicePortBackend(
		"org.qemu.monitor.qmp",
		"org.qemu.monitor.qmp.0",
	).ArgsString()

	assert.Equal(t, result, expected)
}
