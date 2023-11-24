package cli

import (
	"fmt"
	"reflect"
)

// Property represents a property associated with an [Option] that gets passed
// into the command line.
type Property struct {
	Key   string
	Value interface{}
}

// NewProperty returns a new instance of Property.
func NewProperty(key string, value interface{}) *Property {
	return &Property{
		Key:   key,
		Value: value,
	}
}

// Arg converts the property to a value that gets passed into the command line.
func (p *Property) Arg() string {
	stringVal := p.stringValue()

	return fmt.Sprintf("%s=%s", p.Key, stringVal)
}

// PropertiesTable returns a map of the properties with key of property name
// and value of property value.
func PropertiesTable(properties []*Property) map[string]string {
	table := make(map[string]string)

	for _, property := range properties {
		stringVal := property.stringValue()

		table[property.Key] = stringVal
	}

	return table
}

func (p *Property) stringValue() string {
	stringVal := fmt.Sprintf("%v", p.Value)

	if reflect.TypeOf(p.Value).Kind() == reflect.Bool {
		if reflect.ValueOf(p.Value).Bool() {
			stringVal = "on"
		} else {
			stringVal = "off"
		}
	}

	return stringVal
}

// StatusFromBool returns the ifTrue if the specified property value is true
// and elseFalse if it is false.
func StatusFromBool(value bool, ifTrue string, elseFalse string) string {
	if value {
		return ifTrue
	} else {
		return elseFalse
	}
}
