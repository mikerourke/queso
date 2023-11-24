package chardev

// BrailleBackend connects to a local BrlAPI server. This backend does not
// take any options.
type BrailleBackend struct {
	*Backend
}

// NewBrailleBackend returns a new instance of [BrailleBackend].
// id is the unique ID, which can be any string up to 127 characters long.
//
//	qemu-system-* -chardev braille,id=id
func NewBrailleBackend(id string) *BrailleBackend {
	return &BrailleBackend{
		NewBackend("braille", id),
	}
}
