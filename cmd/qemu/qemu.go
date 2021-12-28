package main

import "github.com/mikerourke/queso/qemu"

func main() {
	q := qemu.New("qemu-system-x86_64")

	q.SetOptions(qemu.Accel(qemu.AccelHAXM, qemu.WithThread(qemu.ThreadMulti)))

	q.Start()
}
