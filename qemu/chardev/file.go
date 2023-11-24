package chardev

import "github.com/mikerourke/queso/qemu/cli"

// FileBackend logs all traffic received from the guest to a file.
type FileBackend struct {
	*Backend
	// Path specifies the path of the file to be opened.
	Path string
}

// NewFileBackend returns a new instance of [FileBackend]. id is the unique ID,
// which can be any string up to 127 characters long. path specifies the path of
// the file to be opened. This file will be created if it does not already
// exist, and overwritten if it does.
//
//	qemu-system-* -chardev file,id=id,path=path
func NewFileBackend(id string, path string) *FileBackend {
	return &FileBackend{
		Backend: NewBackend("file", id),
		Path:    path,
	}
}

// SetInputPath specifies the path of a second file which will be used for input.
// If this path is not specified, no input will be available from the character device.
//
// Note that this is not supported on Windows hosts.
//
//	qemu-system-* -chardev file,input-path=path
func (b *FileBackend) SetInputPath(path string) *FileBackend {
	b.properties = append(b.properties, cli.NewProperty("input-path", path))
	return b
}
