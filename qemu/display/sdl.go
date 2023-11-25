package display

import "github.com/mikerourke/queso"

// OpenGLOption is used to specify which OpenGL option to use for [SDLDisplay].
type OpenGLOption string

const (
	// OpenGLOn indicates OpenGL is enabled.
	OpenGLOn OpenGLOption = "on"

	// OpenGLOff indicates OpenGL is disabled.
	OpenGLOff OpenGLOption = "off"

	// OpenGLCore indicates that OpenGL Core (i.e. the core subsset of OpenGL),
	// should be used.
	OpenGLCore OpenGLOption = "core"

	// OpenGLES indicates that OpenGL for Embedded Systems should be used.
	// See https://www.khronos.org/opengles/ for more details.
	OpenGLES OpenGLOption = "es"
)

// MouseGrabbingKeys represents which keys to use to toggle mouse grabbing in
// an [SDLDisplay].
type MouseGrabbingKeys string

const (
	// MouseGrabbingKeysLeft uses left Shift, left Ctrl, left Alt in conjuction
	// with the "g" key to toggle mouse grabbing.
	MouseGrabbingKeysLeft MouseGrabbingKeys = "left"

	// MouseGrabbingKeysRight uses right Ctrl in conjuction with the "g" key
	// to toggle mouse grabbing.
	MouseGrabbingKeysRight MouseGrabbingKeys = "right"
)

// SDLDisplay displays the video output via SDL (usually in a separate graphics
// window; see the SDL documentation for other possibilities).
type SDLDisplay struct {
	*Display
}

// NewSDLDisplay returns a new instance of [SDLDisplay].
//
//	qemu-system-* -display sdl
func NewSDLDisplay() *SDLDisplay {
	return &SDLDisplay{New("sdl")}
}

// SetMouseGrabbingKeys sets the keys used to select the modifier keys for toggling
// the mouse grabbing in conjunction with the “g” key. See [MouseGrabbingKeys]
// for more details.
//
//	qemu-system-* -display sdl,grab-mod=mods
func (d *SDLDisplay) SetMouseGrabbingKeys(mods MouseGrabbingKeys) *SDLDisplay {
	d.properties = append(d.properties, queso.NewProperty("grab-mods", string(mods)))
	return d
}

// SetOpenGL defines the option for using OpenGL for rendering (the D-Bus interface
// will share framebuffers with DMABUF file descriptors).
//
//	qemu-system-* -display sdl,gl=on|off|core|es
func (d *SDLDisplay) SetOpenGL(gl OpenGLOption) *SDLDisplay {
	d.properties = append(d.properties, queso.NewProperty("gl", gl))
	return d
}

// ToggleShowCursor toggles whether to force showing the mouse cursor.
//
//	qemu-system-* -display sdl,show-cursor=on|off
func (d *SDLDisplay) ToggleShowCursor(enabled bool) *SDLDisplay {
	d.properties = append(d.properties, queso.NewProperty("show-cursor", enabled))
	return d
}

// ToggleWindowClose enables or disables the ability to quit QEMU with the window
// close button.
//
//	qemu-system-* -display sdl,window-close=on|off
func (d *SDLDisplay) ToggleWindowClose(enabled bool) *SDLDisplay {
	d.properties = append(d.properties, queso.NewProperty("window-close", enabled))
	return d
}
