// Package vals provides common values for properties that adds extra clarity
// to the code (rather than use hardcoded strings).
package vals

// Status is used to indicate the status for a property that can be enabled,
// disabled, or automatically based on some factor.
type Status string

const (
	// On indicates the property is enabled.
	On Status = "on"

	// Off indicates the property is disabled.
	Off Status = "off"

	// Auto indicates the property is enabled/disabled automatically.
	Auto Status = "auto"
)

// String returns the string representation of [Status].
func (s Status) String() string {
	return string(s)
}

// BoolToStatus returns "on" if the specified property value is true
// and "off" if it is false.
func BoolToStatus(value bool) string {
	if value {
		return On.String()
	} else {
		return Off.String()
	}
}
