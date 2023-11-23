package qemu

import (
	"github.com/mikerourke/queso/qemu/cli"
)

// AppendCommandLine uses the specified cmdLine as kernel command line.
//
//	qemu-system-* -append cmdline
func AppendCommandLine(cmdline string) *cli.Option {
	return cli.NewOption("append", cmdline)
}

// BIOSFile sets the filename for the BIOS.
//
//	qemu-system-* -bios file
func BIOSFile(file string) *cli.Option {
	return cli.NewOption("bios", file)
}

// DeviceTreeBinary uses the specified file as a device tree binary (dtb) image
// and passes it to the kernel on boot.
//
//	qemu-system-* -dtb file
func DeviceTreeBinary(file string) *cli.Option {
	return cli.NewOption("dtb", file)
}

// ParallelFlash uses the specified file as a parallel flash image.
//
//	qemu-system-* -pflash file
func ParallelFlash(file string) *cli.Option {
	return cli.NewOption("pflash", file)
}

// Kernel uses the specified bzImage as kernel image. The kernel can be either a
// Linux kernel or in multiboot format.
//
//	qemu-system-* -kernel bzImage
func Kernel(bzImage string) *cli.Option {
	return cli.NewOption("kernel", bzImage)
}

// InitRAMDisk uses specified fileOrArgs as initial ram disk. In multiboot mode,
// you can use the syntax InitRAMDisk("file1 arg=foo,file2").
//
//	qemu-system-* -initrd file
//	qemu-system-* -initrd "file1 arg=foo,file2"
//	qemu-system-* -initrd "bzImage earlyprintk=xen,,keep root=/dev/xvda1,initrd.img"
func InitRAMDisk(fileOrArgs string) *cli.Option {
	return cli.NewOption("initrd", fileOrArgs)
}
