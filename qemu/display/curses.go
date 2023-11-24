package display

import "github.com/mikerourke/queso/qemu/cli"

// CursesDisplay displays video output via curses. For graphics device models which
// support a text mode, QEMU can display this output using a curses/ncurses interface.
//
// Nothing is displayed when the graphics device is in graphical mode or if the graphics
// device does not support a text mode. Generally only the VGA device models support text mode.
type CursesDisplay struct {
	*Display
}

// NewCursesDisplay returns a new instance of [CursesDisplay].
//
//	qemu-system-* -display curses
func NewCursesDisplay() *CursesDisplay {
	return &CursesDisplay{New("curses")}
}

// SetCharset sets the font charset used by the guest. For example, specify
// "CP850" for IBM CP850 encoding. The default is "CP437".
//
//	qemu-system-* -display curses,charset=encoding
func (d *CursesDisplay) SetCharset(encoding string) *CursesDisplay {
	d.properties = append(d.properties, cli.NewProperty("charset", encoding))
	return d
}

// WithCurses displays the VGA output when in text mode using a curses/ncurses
// interface. Nothing is displayed in graphical mode.
//
//	qemu-system-* -curses
func WithCurses() *cli.Option {
	return cli.NewOption("curses", "")
}
