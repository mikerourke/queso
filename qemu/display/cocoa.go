package display

import "github.com/mikerourke/queso"

// CocoaDisplay displays video output in a Cocoa window. Mac only. This interface
// provides drop-down menus and other UI elements to configure and control the
// VM during runtime.
type CocoaDisplay struct {
	*Display
}

// NewCocoaDisplay returns a new instance of [CocoaDisplay].
//
//	qemu-system-* -display cocoa
func NewCocoaDisplay() *CocoaDisplay {
	return &CocoaDisplay{New("cocoa")}
}

// ToggleLeftCommandKey enables or disables forwarding left command key to host.
//
//	qemu-system-* -display cocoa,left-command-key=on|off
func (d *CocoaDisplay) ToggleLeftCommandKey(enabled bool) *CocoaDisplay {
	d.properties = append(d.properties, queso.NewProperty("left-command-key", enabled))
	return d
}

// ToggleShowCursor toggles whether to force showing the mouse cursor.
//
//	qemu-system-* -display cocoa,show-cursor=on|off
func (d *CocoaDisplay) ToggleShowCursor(enabled bool) *CocoaDisplay {
	d.properties = append(d.properties, queso.NewProperty("show-cursor", enabled))
	return d
}
