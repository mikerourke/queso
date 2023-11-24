package display

import "github.com/mikerourke/queso"

// EGLHeadlessDisplay offloads all OpenGL operations to a local Direct Rendering
// Infrastructure (DRI) device. For any graphical display, this display needs to
// be paired with either VNC or Spice displays.
type EGLHeadlessDisplay struct {
	*Display
}

// NewEGLHeadlessDisplay returns a new instance of [EGLHeadlessDisplay].
//
//	qemu-system-* -display egl-headless
func NewEGLHeadlessDisplay() *EGLHeadlessDisplay {
	return &EGLHeadlessDisplay{New("egl-headless")}
}

// SetRenderNode sets the sets the DRM render node for OpenGL rendering.
//
//	qemu-system-* -display egl-headless,rendernode=file
func (d *EGLHeadlessDisplay) SetRenderNode(file string) *EGLHeadlessDisplay {
	d.properties = append(d.properties, queso.NewProperty("rendernode", file))
	return d
}
