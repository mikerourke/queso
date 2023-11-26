package main

import (
	"fmt"
	"log"

	"github.com/mikerourke/queso/qemu"
	"github.com/mikerourke/queso/qemu/accel"
	"github.com/mikerourke/queso/qemu/blockdev/fsdev"
	"github.com/mikerourke/queso/qemu/device"
	"github.com/mikerourke/queso/qemu/netdev"
	"github.com/mikerourke/queso/qemu/nic"
)

/*
qemu-system-x86_64 -m 3G -smp 2 \
  -netdev user,id=n,hostfwd=tcp:127.0.0.1:9000-:445 \
  -device e1000,netdev=n \
  -usb -device usb-tablet \
  -k en-us \
  -drive file=some-file.qcow2,media=disk,format=qcow2 \
  -cdrom some-iso.iso

*/

func main() {
	// goPath := os.Getenv("GOPATH")
	// machineFile := filepath.Join(goPath, "src", "github.com", "mikerourke", "queso", "machines", "bunsen.qcow2")

	// if err := diskimage.Create(diskimage.CreateOption
	// 	Format:    diskimage.FileFormatQCOW2,
	// 	File:      machineFile,
	// 	Size:      "4G",
	// 	Overwrite: false,
	// }); err != nil {
	// 	log.Println(err)
	//
	// 	return
	// }

	q := qemu.New("qemu-system-x86_64")

	ip := device.NewIPMIBMCExternal("a", "b")
	ip.SetProperty("a", "b")
	q.Use(ip)

	be := nic.NewAFXDPNIC()
	be.SetProgramAttachMode(netdev.ProgramAttachNative)
	q.Use(be)

	ac := accel.New("test")
	ac.SetProperty("a", "b")
	q.Use(ac)

	synth := fsdev.NewSyntheticFileSystemDevice("test").SetProperty("a", "b")
	q.Use(synth)

	net := nic.NewAFXDPNIC().
		SetModelName("e1000").
		SetMACAddress("52:54:98:76:54:32")
	q.Use(net)

	fmt.Println(q.Args())

	return

	q.With(
		qemu.WithKeyboardLayout(qemu.LanguageEnglishUS),
		qemu.AddFileDescriptor(1, 2, "s"),
		qemu.AddFileDescriptor(1, 2, "s"),
		qemu.WithMonitorRedirect("none"),
		qemu.MemorySize("3G"),

		// Network Settings
		// network.UserBackend("n",
		//	network.WithForwardRule(
		//		network.NewHostForwardRule(network.PortTypeTCP,
		//			9000,
		//			445,
		//		).WithHostIP("127.0.0.1"))),
		// device.Use("e1000",
		//	device.NewProperty("netdev", "n")),

		// USB Settings
		qemu.EnableUSB(),
		qemu.WithUSBDevice(qemu.USBDeviceTablet),

		// Drive Settings
		// blockdev.Drive(
		//	blockdev.WithDiskImageFile("some-file.qcow2"),
		//	blockdev.WithDiskImageFormat(diskimage.FileFormatQCOW2),
		//	blockdev.WithDriveMedia(blockdev.DriveMediaDisk)),
		// blockdev.DiskDrive(blockdev.CDROM, "some-iso.iso")
	)

	log.Println(q.Args())

	if err := q.Cmd().Run(); err != nil {
		log.Println(err)
	}
}
