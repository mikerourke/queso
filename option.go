package queso

import (
	"fmt"
	"strings"
)

// Option represents an option passed into the command line.
type Option struct {
	Flag       string
	Name       string
	Properties []*Property
}

// NewOption returns a new instance of Option.
func NewOption(flag string, name string, properties ...*Property) *Option {
	return &Option{
		Flag:       flag,
		Name:       name,
		Properties: properties,
	}
}

// Args converts the Option to a string that can be passed into a QEMU tool via
// the command line.
func (opt *Option) Args() []string {
	args := []string{fmt.Sprintf("-%s", opt.Flag)}

	props := make([]string, 0)

	if opt.Name != "" {
		props = append(props, opt.Name)
	}

	if opt.Properties != nil {
		for _, property := range opt.Properties {
			props = append(props, property.Arg())
		}
	}

	if len(props) != 0 {
		args = append(args, strings.Join(props, ","))
	}

	return args
}
