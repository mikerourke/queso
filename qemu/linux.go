package qemu

import "github.com/mikerourke/queso"

// Kernel uses the specified bzImage as kernel image. The kernel can be either a
// Linux kernel or in multiboot format.
func Kernel(bzImage string) *queso.Option {
	return queso.NewOption("kernel", bzImage)
}

// Append uses the specified cmdLine as kernel command line.
func Append(cmdLine string) *queso.Option {
	return queso.NewOption("append", cmdLine)
}

// InitRAMDisk uses specified fileOrArgs as initial ram disk. In multiboot mode,
// you can use the syntax InitRAMDisk("file1 arg=foo,file2").
func InitRAMDisk(fileOrArgs string) *queso.Option {
	return queso.NewOption("initrd", fileOrArgs)
}

// DeviceTreeBinary uses the specified file as a device tree binary (dtb) image
// and passes it to the kernel on boot.
func DeviceTreeBinary(file string) *queso.Option {
	return queso.NewOption("dtb", file)
}

// DataDirectoryPath sets the directory for the BIOS, VGA BIOS and keymaps to
// the specified path.
func DataDirectoryPath(path string) *queso.Option {
	return queso.NewOption("L", path)
}
