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

// NewProperty returns a new instance of Property.
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

type PropertiesTable = map[string]string

// ToPropertiesTable returns a map of the properties with key of property name
// and value of property value.
func ToPropertiesTable(properties []*Property) PropertiesTable {
	table := make(map[string]string)

	for _, property := range properties {
		stringVal := fmt.Sprintf("%v", property.Value)

		if reflect.TypeOf(property.Value).Kind() == reflect.Bool {
			if reflect.ValueOf(property.Value).Bool() {
				stringVal = "on"
			} else {
				stringVal = "off"
			}
		}

		table[property.Key] = stringVal
	}

	return table
}
