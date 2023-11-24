package display

import "github.com/mikerourke/queso"

// NoneDisplay does not display video output. The guest will still see an emulated
// graphics card, but its output will not be displayed to the QEMU user.
//
// This option differs from the [WithNoGraphic] option in that it only affects what
// is done with video output; [WithNoGraphic] also changes the destination of the
// serial and parallel port data.
type NoneDisplay struct {
	*Display
}

// NewNoneDisplay returns a new instance of [NoneDisplay].
//
//	qemu-system-* -display none
func NewNoneDisplay() *NoneDisplay {
	return &NoneDisplay{New("none")}
}

// WithNoGraphic totally disables graphical output so that QEMU is a simple command
// line application. The emulated serial port is redirected on the console and
// muxed with the monitor (unless redirected elsewhere explicitly). Therefore, you
// can still use QEMU to debug a Linux kernel with a serial console.
//
//	qemu-system-* -nographic
func WithNoGraphic() *queso.Option {
	return queso.NewOption("nographic", "")
}
