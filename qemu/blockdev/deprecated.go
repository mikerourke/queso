package blockdev

import "github.com/mikerourke/queso"

// WithFDA specifies the file to use for 1st floppy disk drive.
//
//	qemu-system-* -fda file
func WithFDA(file string) *queso.Option {
	return queso.NewOption("fda", file)
}

// WithFDB specifies the file to use for 2nd floppy disk drive.
//
//	qemu-system-* -fdb file
func WithFDB(file string) *queso.Option {
	return queso.NewOption("fdb", file)
}

// WithHDA specifies the file to use for hard disk 0 on the default bus of the
// emulated machine.
//
//	qemu-system-* -hda file
func WithHDA(file string) *queso.Option {
	return queso.NewOption("hda", file)
}

// WithHDB specifies the file to use for hard disk 1 on the default bus of the
// emulated machine.
//
//	qemu-system-* -hdb file
func WithHDB(file string) *queso.Option {
	return queso.NewOption("hdb", file)
}

// WithHDC specifies the file to use for hard disk 2 on the default bus of the
// emulated machine.
//
//	qemu-system-* -hdc file
func WithHDC(file string) *queso.Option {
	return queso.NewOption("hdc", file)
}

// WithHDD specifies the file to use for hard disk 3 on the default bus of the
// emulated machine.
//
//	qemu-system-* -hdd file
func WithHDD(file string) *queso.Option {
	return queso.NewOption("hdd", file)
}

// WithCDROM specifies the file to use as CD-ROM image on the default bus of the
// emulated machine (which is IDE1 master on x86, so you cannot use [WithHDA] and
// WithCDROM at the same time there).
//
// On systems that support it, you can use the host CD-ROM by using
// "/dev/cdrom" as filename.
//
//	qemu-system-* -cdrom file
func WithCDROM(file string) *queso.Option {
	return queso.NewOption("cdrom", file)
}
