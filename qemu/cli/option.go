package cli

import (
	"fmt"
	"strings"
)

// Option represents an option passed into the command line.
type Option struct {
	Flag          string
	Name          string
	Properties    []*Property
	noLeadingDash bool
}

// NewOption returns a new instance of [Option].
func NewOption(flag string, name string, properties ...*Property) *Option {
	return &Option{
		Flag:       flag,
		Name:       name,
		Properties: properties,
	}
}

// OmitLeadingDash ensures the flag doesn't include a `-` before it when it is
// converted to an arg. This is only required for a very small set of flags,
// which _may_ actually be incorrectly documented.
func (opt *Option) OmitLeadingDash() *Option {
	opt.noLeadingDash = true
	return opt
}

// Args converts the [Option] to a string that can be passed into a QEMU tool via
// the command line.
func (opt *Option) Args() []string {
	// If the leading dash was explicitly excluded, make sure we just use the
	// Flag wihtout the leading dash.
	flagArg := opt.Flag
	if !opt.noLeadingDash {
		flagArg = fmt.Sprintf("-%s", flagArg)
	}

	args := []string{flagArg}

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

// ArgsString returns a string argument that gets passed to QEMU. This is used
// primarily for testing purposes.
func (opt *Option) ArgsString() string {
	args := opt.Args()
	return strings.Join(args, " ")
}

// Table returns a map with key of property name and value of property value
// (as a string).
func (opt *Option) Table() map[string]string {
	table := make(map[string]string)

	if opt.Properties != nil {
		for _, property := range opt.Properties {
			table[property.Key] = fmt.Sprint(property.Value)
		}
	}

	return table
}
