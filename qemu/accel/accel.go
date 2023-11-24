// Package accel is used to manage hardware acceleration in a QEMU guest image.
package accel

import (
	"fmt"

	"github.com/mikerourke/queso/qemu/cli"
)

// Type represents the type of accelerator.
type Type string

const (
	// TypeHVF is the Hypervisor.framework accelerator on macOS.
	// See https://developer.apple.com/documentation/hypervisor for more details.
	TypeHVF Type = "hvf"

	// TypeKVM is the Kernel Virtual Machine, which is a Linux kernel module.
	// See https://wiki.qemu.org/Features/KVM for more details.
	TypeKVM Type = "kvm"

	// TypeNVMM is a Type-2 hypervisor and hypervisor platform for NetBSD.
	// See https://m00nbsd.net/4e0798b7f2620c965d0dd9d6a7a2f296.html for more details.
	TypeNVMM Type = "nvmm"

	// TypeTCG is the Tiny Code Generator, which is the core binary translation
	// engine for QEMU. See https://wiki.qemu.org/Features/TCG for more details.
	TypeTCG Type = "tcg"

	// TypeWHPX is the Windows Hypervisor Platform accelerator (Hyper-V).
	// See https://docs.microsoft.com/en-us/xamarin/android/get-started/installation/android-emulator/hardware-acceleration
	// for more details.
	TypeWHPX Type = "whpx"

	// TypeXen is a Type-1 (bare metal) hypervisor.
	// See https://wiki.xenproject.org/wiki/Xen_Project_Software_Overview for
	// more details.
	TypeXen Type = "xen"

	// TypeHAXM is the Intel Hardware Accelerated Execution Manager.
	// See https://github.com/intel/haxm for more details.
	//
	// Deprecated: No longer supported, but kept for older versions of QEMU.
	TypeHAXM Type = "hax"
)

// WithAccelerator can be used in conjunction with the qemu.With method
// to define an accelerator with no additional properties.
//
//	qemu-system-* -accel type
func WithAccelerator(accelType Type) *cli.Option {
	return cli.NewOption("accel", string(accelType))
}

// Accelerator represents any of the available hardware accelerators.
type Accelerator struct {
	Type       Type
	properties []*cli.Property
}

// New returns a new [Accelerator] instance with the specified accelerator type.
//
//	qemu-system-* -accel type
func New(accelType Type) *Accelerator {
	return &Accelerator{
		Type:       accelType,
		properties: make([]*cli.Property, 0),
	}
}

func (a *Accelerator) option() *cli.Option {
	return cli.NewOption("accel", string(a.Type), a.properties...)
}

// SetProperty is used to add arbitrary properties to the [Accelerator].
func (a *Accelerator) SetProperty(key string, value interface{}) *Accelerator {
	a.properties = append(a.properties, cli.NewProperty(key, value))
	return a
}

// NotifyOnVMExitOption represents the options for notifying when the VM exits.
type NotifyOnVMExitOption string

const (
	// NotifyOnVMExitRun enables VM exit support on x86 host. "window" specifies the
	// corresponding notification window of time to trigger the VM exit if enabled.
	//
	// This feature can mitigate the CPU stuck issue due to event windows that don't
	// open up for a specified amount of time.
	NotifyOnVMExitRun NotifyOnVMExitOption = "run"

	// NotifyOnVMExitInternalError does not notify and continues if the exit happens,
	// but it raises an internal error.
	NotifyOnVMExitInternalError NotifyOnVMExitOption = "internal-error"

	// NotifyOnVMExitDisable does not notify and does nothing when the exit happens.
	NotifyOnVMExitDisable NotifyOnVMExitOption = "disable"
)

// SetNotifyOnVMExit sets the notify approach as well as the notification window
// (if applicable). If you're disabling the exit notification, use `0` for the
// window.
//
//	qemu-system-* -accel name,notify-vmexit=run|internal-error|disable,notify-window=n
func (a *Accelerator) SetNotifyOnVMExit(option NotifyOnVMExitOption, window int) *Accelerator {
	var value string
	if option == NotifyOnVMExitRun {
		value = fmt.Sprintf("run,notify-window=%d", window)
	} else {
		value = string(option)
	}

	a.properties = append(a.properties, cli.NewProperty("notify-vmexit", value))
	return a
}
