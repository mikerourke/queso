package queso

import (
	"fmt"
	"reflect"
)

// Property represents a property associated with an Option that gets passed
// into the command line.
type Property struct {
	Key   string
	Value interface{}
}

// NewProperty returns a new instance of a Property.
func NewProperty(key string, value interface{}) *Property {
	return &Property{
		Key:   key,
		Value: value,
	}
}

// Arg converts the property to a value that gets passed into the command line.
func (p *Property) Arg() string {
	stringVal := fmt.Sprintf("%v", p.Value)

	if reflect.TypeOf(p.Value).Kind() == reflect.Bool {
		if reflect.ValueOf(p.Value).Bool() {
			stringVal = "on"
		} else {
			stringVal = "off"
		}
	}

	return fmt.Sprintf("%s=%s", p.Key, stringVal)
}
