package queso

// Entity represents a base entity with the ability to manage and set
// properties. This is used as the base for every qemu.Usable.
type Entity struct {
	Properties []*Property
	OptionFlag string
	OptionName string
}

// NewEntity returns a new [Entity] instance.
func NewEntity(flag string, name string) *Entity {
	return &Entity{
		OptionFlag: flag,
		OptionName: name,
		Properties: make([]*Property, 0),
	}
}

// SetOptionFlag sets the flag, which is the first argument after the `qemu-system-*`
// command (i.e. starts with single dash). This method is used mainly to override
// the default flag when combining two entities together using a convenience
// wrapper (e.g. NIC).
//
//	qemu-system-* -<flag> <name>,...
func (e *Entity) SetOptionFlag(flag string) *Entity {
	e.OptionFlag = flag
	return e
}

// SetOptionName sets the name, which immediately follows the flag (i.e. dashed
// argument). This method is used mainly to override the default name when combining
// two entities together using a convenience wrapper (e.g. NIC).
//
//	qemu-system-* -<flag> <name>,...
func (e *Entity) SetOptionName(name string) *Entity {
	e.OptionName = name
	return e
}

// Option returns the [Option] representation of the entity.
func (e *Entity) Option() *Option {
	return NewOption(e.OptionFlag, e.OptionName, e.Properties...)
}

// Property returns the [Property] with the specified key (if found).
func (e *Entity) Property(key string) *Property {
	var match *Property

	for _, property := range e.Properties {
		if property.Key == key {
			match = property
		}
	}

	return match
}

// RemoveProperty removes the [Property] with the specified key from the
// entity's properties.
func (e *Entity) RemoveProperty(key string) *Entity {
	properties := make([]*Property, 0)

	for _, property := range e.Properties {
		if property.Key != key {
			properties = append(properties, property)
		}
	}

	e.Properties = properties
	return e
}

// SetProperty is used to add arbitrary properties to the entity.
func (e *Entity) SetProperty(key string, value interface{}) *Entity {
	e.Properties = append(e.Properties, NewProperty(key, value))
	return e
}

// UpsertProperty either adds a property with the specified key to the properties
// slice if it doesn't exist, otherwise overwrites the existing value.
func (e *Entity) UpsertProperty(key string, value interface{}) *Entity {
	found := false

	properties := e.Properties
	for _, property := range properties {
		if property.Key == key {
			found = true
			property.Value = value
		}
	}

	if !found {
		properties = append(properties, NewProperty(key, value))
	}
	e.Properties = properties
	return e
}
