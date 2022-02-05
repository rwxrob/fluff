package model

// Manifest represents a full configuration of Machines and Clouds that
// combine collections of Machine Instances. This is the top level data
// struct used for marshalling and unmarshalling fluff.yaml files.
//
type Manifest struct {
	Machines []Machine `json:"machines,omitempty"`
	Clouds   []Cloud   `json:"clouds,omitempty"`
}
