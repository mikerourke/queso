package display

import "github.com/mikerourke/queso"

// DBusDisplay exports the display over D-Bus interfaces. (Since 7.0)
//
// The connection is registered with the “org.qemu” name (and queued when already
// owned).
type DBusDisplay struct {
	*Display
}

// NewDBusDisplay returns a new instance of [DBusDisplay].
//
//	qemu-system-* -display dbus
func NewDBusDisplay() *DBusDisplay {
	return &DBusDisplay{New("dbus")}
}

// SetAddress sets the D-Bus bus address to connect to.
//
//	qemu-system-* -display dbus,addr=addr
func (d *DBusDisplay) SetAddress(addr string) *DBusDisplay {
	d.properties = append(d.properties, queso.NewProperty("addr", addr))
	return d
}

// SetOpenGL defines the option for using OpenGL for rendering (the D-Bus interface
// will share framebuffers with DMABUF file descriptors).
//
//	qemu-system-* -display dbus,gl=on|off|core|es
func (d *DBusDisplay) SetOpenGL(gl OpenGLOption) *DBusDisplay {
	d.properties = append(d.properties, queso.NewProperty("gl", gl))
	return d
}

// TogglePeerToPeer specifies if peer-to-peer connection should be used, accepted
// via QMP add_client.
//
//	qemu-system-* -display dbus,p2p=yes|no
func (d *DBusDisplay) TogglePeerToPeer(enabled bool) *DBusDisplay {
	value := queso.StatusFromBool(enabled, "yes", "no")
	d.properties = append(d.properties, queso.NewProperty("p2p", value))
	return d
}
