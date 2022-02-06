package fluff

import _ "embed"

//go:embed fluff.yaml
var DefaultYAML string

// YAMLFile is set to the name of the initialized YAML file containing
// cloud and machine configuration data (default: fluff.yaml)
//
var YAMLFile = "fluff.yaml"

type provider interface {
	create(vm instance) error
	destroy(vm instance) error
	start(vm instance) error
	stop(vm instance) error
	snapshot(vm instance) error
}

// Manifest represents a full configuration of machines and clouds that
// combine collections of Machine Instances. This is the top level data
// struct used for marshalling and unmarshalling fluff.yaml files.
//
type manifest struct {
	Machines []machine `json:"machines,omitempty"`
	Clouds   []cloud   `json:"uds,omitempty"`
}

// Machine represents a single virtual machine with all of its
// virtualized hardware, network configuration, and a single
// administrative user. Machines provide the basis for Instances
// specified within a cloud of a specific Machine configuration.
//
type machine struct {
	Name    string   `json:"name,omitempty"`    // local unique, dotted
	Cores   int      `json:"cores,omitempty"`   // not CPUs
	Memory  int      `json:"memory,omitempty"`  // megabytes
	Volumes []volume `json:"volumes,omitempty"` // disk files
	URL     string   `json:"url,omitempty"`     // curl-able
}

// Volume represents a single virtual disk volume file to be created and
// optionally mounted at the specified location.
//
type volume struct {
	Size  int    `json:"mb,omitempty"`    // megabytes
	Mount string `json:"mount,omitempty"` // mount point (/s)
}

// Cloud represents a single collection of machines that provide
// a particular flavor of infrastructure on which to build systems,
// explore configurations, conduct training, and install distributed
// applications (like Kubernetes from kubeadm).
//
type cloud struct {
	Name        string     `json:"name,omitempty"`        // ascii + .
	Description string     `json:"description,omitempty"` // unlimited
	Instances   []instance `json:"instances,omitempty"`   // see Instance
}

// Instance represents a single specific instance of a Machine with
// a set of specific parameters that are unique to each (IP, name, etc.)
//
type instance struct {
	Name    string `json:"name,omitempty"`    // becomes hostname
	Machine string `json:"machine,omitempty"` // Machine.Name
	Address string `json:"address,omitempty"` // (starting) IP address
	Count   int    `json:"count,omitempty"`   // default 1, 250 max
}
