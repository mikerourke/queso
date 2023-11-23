package blockdev

import "github.com/mikerourke/queso/qemu/cli"

// FDA specifies the file to use for 1st floppy disk drive.
//
//	qemu-system-* -fda file
func FDA(file string) *cli.Option {
	return cli.NewOption("fda", file)
}

// FDB specifies the file to use for 2nd floppy disk drive.
//
//	qemu-system-* -fdb file
func FDB(file string) *cli.Option {
	return cli.NewOption("fdb", file)
}

// HDA specifies the file to use for hard disk 0 on the default bus of the
// emulated machine.
//
//	qemu-system-* -hda file
func HDA(file string) *cli.Option {
	return cli.NewOption("hda", file)
}

// HDB specifies the file to use for hard disk 1 on the default bus of the
// emulated machine.
//
//	qemu-system-* -hdb file
func HDB(file string) *cli.Option {
	return cli.NewOption("hdb", file)
}

// HDC specifies the file to use for hard disk 2 on the default bus of the
// emulated machine.
//
//	qemu-system-* -hdc file
func HDC(file string) *cli.Option {
	return cli.NewOption("hdc", file)
}

// HDD specifies the file to use for hard disk 3 on the default bus of the
// emulated machine.
//
//	qemu-system-* -hdd file
func HDD(file string) *cli.Option {
	return cli.NewOption("hdd", file)
}

// CDROM specifies the file to use as CD-ROM image on the default bus of the
// emulated machine (which is IDE1 master on x86, so you cannot use [HDA] and
// CDROM at the same time there).
//
// On systems that support it, you can use the host CD-ROM by using "/dev/cdrom" as filename.
func CDROM(file string) *cli.Option {
	return cli.NewOption("cdrom", file)
}
