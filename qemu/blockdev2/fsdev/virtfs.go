package fsdev

import "github.com/mikerourke/queso/internal/cli"

type VirtualFileSystemDevice struct {
	Type       Type
	properties []*cli.Property
}
