package display

import "github.com/mikerourke/queso"

// GTKDisplay displays video output in a GTK window. This interface provides
// drop-down menus and other UI elements to configure and control the VM during
// runtime.
type GTKDisplay struct {
	*Display
}

// NewGTKDisplay returns a new instance of [GTKDisplay].
//
//	qemu-system-* -display gtk
func NewGTKDisplay() *GTKDisplay {
	return &GTKDisplay{New("gtk")}
}

// ToggleFullScreen toggles whether to start in full-screen mode.
//
//	qemu-system-* -display gtk,full-screen=on|off
func (d *GTKDisplay) ToggleFullScreen(enabled bool) *GTKDisplay {
	d.properties = append(d.properties, queso.NewProperty("full-screen", enabled))
	return d
}

// ToggleOpenGL enables or disables OpenGL for displaying.
//
//	qemu-system-* -display gtk,gl=on|off
func (d *GTKDisplay) ToggleOpenGL(enabled bool) *GTKDisplay {
	d.properties = append(d.properties, queso.NewProperty("gl", enabled))
	return d
}

// ToggleShowCursor toggles whether to force showing the mouse cursor.
//
//	qemu-system-* -display gtk,show-cursor=on|off
func (d *GTKDisplay) ToggleShowCursor(enabled bool) *GTKDisplay {
	d.properties = append(d.properties, queso.NewProperty("show-cursor", enabled))
	return d
}

// ToggleShowMenuBar toggles whether to display the main window menubar, defaults
// to true.
//
//	qemu-system-* -display gtk,show-menubar=on|off
func (d *GTKDisplay) ToggleShowMenuBar(enabled bool) *GTKDisplay {
	d.properties = append(d.properties, queso.NewProperty("show-menubar", enabled))
	return d
}

// ToggleShowTabBar toggles whether to display the tab bar for switching between the
// various graphical interfaces (e.g. VGA and virtual console character devices)
// by default.
//
//	qemu-system-* -display gtk,show-tabs=on|off
func (d *GTKDisplay) ToggleShowTabBar(enabled bool) *GTKDisplay {
	d.properties = append(d.properties, queso.NewProperty("show-tabs", enabled))
	return d
}

// ToggleWindowClose enables or disables the ability to quit QEMU with the window
// close button.
//
//	qemu-system-* -display gtk,window-close=on|off
func (d *GTKDisplay) ToggleWindowClose(enabled bool) *GTKDisplay {
	d.properties = append(d.properties, queso.NewProperty("window-close", enabled))
	return d
}

// ToggleZoomToFit toggles whether to expand video output to the window size.
// The default value is false.
//
//	qemu-system-* -display gtk,zoom-to-fit=on|off
func (d *GTKDisplay) ToggleZoomToFit(enabled bool) *GTKDisplay {
	d.properties = append(d.properties, queso.NewProperty("zoom-to-fit", enabled))
	return d
}
