package audiodev

// SPICEBackend represents a backend that sends audio through SPICE. This backend
// requires the [qemu.SPICE] and automatically selected in that case, so usually
// you can ignore this option. This backend has no backend specific properties.
type SPICEBackend struct {
	*Backend
}

// NewSPICEBackend returns a new instance of [SPICEBackend].
//
//	qemu-system-* -audiodev spice,id=id
func NewSPICEBackend(id string) *SPICEBackend {
	return &SPICEBackend{
		NewBackend("spice", id),
	}
}
