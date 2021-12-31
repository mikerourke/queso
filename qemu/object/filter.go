package object

import (
	"fmt"

	"github.com/mikerourke/queso"
)

// FilterBuffer batches the packet delivery. All packets arriving in a given
// interval on the netdev parameter (ID of a network device) are delayed until
// the end of the interval. The interval parameter is in microseconds.
func FilterBuffer(id string, netdev string, interval int, properties ...*FilterProperty) *queso.Option {
	props := []*queso.Property{
		queso.NewProperty("id", id),
		queso.NewProperty("netdev", netdev),
		queso.NewProperty("interval", interval),
	}

	for _, property := range properties {
		props = append(props, property.Property)
	}

	return queso.NewOption("object", "filter-buffer", props...)
}

// FilterMirror mirrors packet from the network device with ID netdev to the
// character device with ID outdev.
func FilterMirror(
	id string,
	netdev string,
	outdev string,
	queue FilterQueueType,
	properties ...*FilterProperty,
) *queso.Option {
	props := []*queso.Property{
		queso.NewProperty("id", id),
		queso.NewProperty("netdev", netdev),
		queso.NewProperty("outdev", outdev),
		queso.NewProperty("queue", queue),
	}

	for _, property := range properties {
		props = append(props, property.Property)
	}

	return queso.NewOption("object", "filter-mirror", props...)
}

// FilterRedirector redirects filter's network packet from the network device
// with ID netdev to the character device with ID indev or outdev. If both the
// indev and outdev parameters are specified, they cannot match. Either one
// can be an empty string, but not both.
func FilterRedirector(
	id string,
	netdev string,
	indev string,
	outdev string,
	queue FilterQueueType,
	properties ...*FilterProperty,
) *queso.Option {
	switch {
	case indev == "" && outdev == "":
		panic("indev and outdev cannot both be empty for a FilterRedirector")

	case indev == outdev:
		panic("indev and outdev cannot be the same for a FilterRedirector")
	}

	props := []*queso.Property{
		queso.NewProperty("id", id),
		queso.NewProperty("netdev", netdev),
		queso.NewProperty("queue", queue),
	}

	if indev != "" {
		props = append(props, queso.NewProperty("indev", indev))
	}

	if outdev != "" {
		props = append(props, queso.NewProperty("outdev", outdev))
	}

	for _, property := range properties {
		props = append(props, property.Property)
	}

	return queso.NewOption("object", "filter-redirector", props...)
}

// FilterRewriter is a part of COLO project. It will rewrite TCP packet to secondary
// from primary to keep secondary TCP connection, and rewrite TCP packet to primary
// from secondary so TCP packet can be handled by client.
func FilterRewriter(
	id string,
	netdev string,
	queue FilterQueueType,
	properties ...*FilterProperty,
) *queso.Option {
	props := []*queso.Property{
		queso.NewProperty("id", id),
		queso.NewProperty("netdev", netdev),
		queso.NewProperty("queue", queue),
	}

	for _, property := range properties {
		props = append(props, property.Property)
	}

	return queso.NewOption("object", "filter-rewriter", props...)
}

// FilterDump dumps the network traffic on network device with ID netdev to the
// specified file. The file format is libpcap, so it can be analyzed with tools
// such as tcpdump or Wireshark.
func FilterDump(
	id string,
	netdev string,
	file string,
	properties ...*FilterProperty,
) *queso.Option {
	props := []*queso.Property{
		queso.NewProperty("id", id),
		queso.NewProperty("netdev", netdev),
		queso.NewProperty("file", file),
	}

	for _, property := range properties {
		props = append(props, property.Property)
	}

	return queso.NewOption("object", "filter-dump", props...)
}

// ColoCompare gets packet from character devices with ID primaryIn and secondaryIn,
// then compares whether the payload of the primary packet and secondary packet are
// the same. If same, it will output primary packet to device with ID outdev, else
// it will notify COLO-framework to do checkpoint and send primary packet to outdev.
// In order to improve efficiency, we need to put the task of comparison in another
// thread specified via the ioThread parameter.
//
// This object must be used with the help of FilterMirror, FilterRedirector, and
// FilterRewriter.
func ColoCompare(
	id string,
	primaryIn string,
	secondaryIn string,
	outdev string,
	ioThread string,
	properties ...*FilterProperty,
) *queso.Option {
	props := []*queso.Property{
		queso.NewProperty("id", id),
		queso.NewProperty("primary_in", primaryIn),
		queso.NewProperty("secondary_in", secondaryIn),
		queso.NewProperty("outdev", outdev),
		queso.NewProperty("iothread", ioThread),
	}

	for _, property := range properties {
		props = append(props, property.Property)
	}

	return queso.NewOption("object", "colo-compare", props...)
}

// FilterProperty represents a property that can be passed to a filter option.
type FilterProperty struct {
	*queso.Property
}

// NewFilterProperty returns a new instance of FilterProperty.
func NewFilterProperty(key string, value interface{}) *FilterProperty {
	return &FilterProperty{
		Property: queso.NewProperty(key, value),
	}
}

// IsFilterEnabled indicates whether the filter is on or off, the default status
// for the filter is on.
func IsFilterEnabled(enabled bool) *FilterProperty {
	return NewFilterProperty("status", enabled)
}

// FilterQueueType represents a filter type that can be applied to any filter
// via the WithFilterQueue property.
type FilterQueueType string

const (
	// FilterQueueAll represents a filter that is attached both to the receive
	// and the transmit queue of the netdev (default).
	FilterQueueAll FilterQueueType = "all"

	// FilterQueueReceive represents a filter that is attached to the receive
	// queue of the netdev, where it will receive packets sent to the netdev.
	FilterQueueReceive FilterQueueType = "rx"

	// FilterQueueTransmit represents a filter that is attached to the transmit
	// queue of the netdev, where it will receive packets sent by the netdev.
	FilterQueueTransmit FilterQueueType = "tx"
)

// WithFilterQueue specifies to which queue the filter is attached. See the
// FilterQueueType constants for more details.
func WithFilterQueue(filter FilterQueueType) *FilterProperty {
	return NewFilterProperty("queue", filter)
}

// FilterPosition represents where the filter should be inserted in the filter
// list and is specified using the WithFilterPosition property.
type FilterPosition string

const (
	// FilterPositionHead indicates that the filter is inserted at the head of the
	// filter list, before any existing filters.
	FilterPositionHead FilterPosition = "head"

	// FilterPositionTail indicates that the filter is inserted at the tail of the
	// filter list, behind any existing filters (default).
	FilterPositionTail FilterPosition = "tail"

	// FilterPositionInsert indicates that the filter is inserted before or behind
	// a filter with a specified ID.
	FilterPositionInsert FilterPosition = "id"
)

// WithFilterPosition specifies where the filter should be inserted in the filter
// list. It can be applied to any filter. If FilterPositionInsert is specified,
// you must specify a value for the id parameter, otherwise use an empty string.
func WithFilterPosition(position FilterPosition, id string) *FilterProperty {
	value := string(position)

	if position == FilterPositionInsert {
		if id == "" {
			panic("an ID is required if WithFilterPosition is specified with FilterPositionInsert")
		} else {
			value = fmt.Sprintf("%s=%s", position, id)
		}
	}

	return NewFilterProperty("position", value)
}

// FilterInsert represents the location of a filter in the list and is specified
// via the WithFilterInsert property.
type FilterInsert string

const (
	// FilterInsertBefore inserts before the filter specified in WithFilterPosition.
	FilterInsertBefore FilterInsert = "before"

	// FilterInsertBehind inserts behind the filter specified in
	// WithFilterPosition (default).
	FilterInsertBehind FilterInsert = "behind"
)

// WithFilterInsert specifies where to insert the filter with the ID specified
// for the WithFilterPosition property.
func WithFilterInsert(insert FilterInsert) *FilterProperty {
	return NewFilterProperty("insert", insert)
}

// IsVNETHDRSupport specifies whether perform a packet action with vnet_hdr_len
// for a FilterMirror, FilterRedirector, or FilterRewriter.
func IsVNETHDRSupport(enabled bool) *FilterProperty {
	return NewFilterProperty("vnet_hdr_support", enabled)
}

// WithDumpFileMaxLength specifies the maximum length of the dump file for a
// FilterDump. The default is 64k.
func WithDumpFileMaxLength(length string) *FilterProperty {
	return NewFilterProperty("maxlen", length)
}

// WithNotifyDevice notifies a Xen colo-frame to do checkpoint for a ColoCompare
// object using Xen COLO.
func WithNotifyDevice(id string) *FilterProperty {
	return NewFilterProperty("notify_dev", id)
}
