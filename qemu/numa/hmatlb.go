package numa

import (
	"github.com/mikerourke/queso"
)

type HMATLB struct {
	properties []*queso.Property
}

type HMATLBDataType string

const (
	AccessBandwidth HMATLBDataType = "access-bandwidth"
	ReadBandwidth   HMATLBDataType = "read-bandwidth"
	WriteBandwidth  HMATLBDataType = "write-bandwidth"
	AccessLatency   HMATLBDataType = "access-latency"
	ReadLatency     HMATLBDataType = "read-latency"
	WriteLatency    HMATLBDataType = "write-latency"
)

func NewHMATLB(initiator *Node, target *Node, hierarchy string, dataType HMATLBDataType) *HMATLB {
	properties := []*queso.Property{
		queso.NewProperty("initiator", initiator.ID),
		queso.NewProperty("target", target.ID),
		queso.NewProperty("hierarchy", hierarchy),
		queso.NewProperty("data-type", dataType),
	}

	return &HMATLB{
		properties: properties,
	}
}

func (h *HMATLB) option() *queso.Option {
	return queso.NewOption("numa", "hmat-lb", h.properties...)
}

func (h *HMATLB) SetLatency(value int) *HMATLB {
	for _, property := range h.properties {
		if property.Key == "bandwidth" {
			panic("cannot set latency on HMATLB if bandwidth already set")
		}
	}

	h.properties = append(h.properties, queso.NewProperty("latency", value))
	return h
}

func (h *HMATLB) SetBandwidth(value string) *HMATLB {
	for _, property := range h.properties {
		if property.Key == "latency" {
			panic("cannot set bandwidth on HMATLB if latency already set")
		}
	}

	h.properties = append(h.properties, queso.NewProperty("bandwidth", value))
	return h
}
