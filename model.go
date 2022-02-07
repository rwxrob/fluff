package fluff

import (
	"github.com/rwxrob/cmdbox/util"
	"gopkg.in/yaml.v2"
)

// Manifest starts out as the parsed DefaultYAML data (which is embedded
// in the binary) but is then modified by overlaying the local
// fluff.yaml file.
//
var Manifest = new(manifest)

func init() {
	err := yaml.Unmarshal([]byte(DefaultYAML), Manifest)
	if err != nil {
		panic(err)
	}
}

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
	Machines []machine `json:"machines"`
	Clouds   []cloud   `json:"clouds"`
}

func (m manifest) String() string { return util.ToYAML(m) }

// Machine represents a single virtual machine with all of its
// virtualized hardware, network configuration, and a single
// administrative user. Machines provide the basis for Instances
// specified within a cloud of a specific Machine configuration.
//
type machine struct {
	Name    string   `json:"name"`    // local unique, dotted
	Cores   int      `json:"cores"`   // not CPUs
	Memory  int      `json:"memory"`  // megabytes
	Volumes []volume `json:"volumes"` // storage volumes
	URL     string   `json:"url"`     // curl-able
	Base    string   `json:"base"`    // overlay and existing machine
}

func (m machine) String() string { return util.ToYAML(m) }

type volume struct {
	Size int `json:"size"` // megabytes
}

func (m volume) String() string { return util.ToYAML(m) }

// Cloud represents a single collection of machines that provide
// a particular flavor of infrastructure on which to build systems,
// explore configurations, conduct training, and install distributed
// applications (like Kubernetes from kubeadm).
//
type cloud struct {
	Name      string     `json:"name"`      // ascii + .
	Summary   string     `json:"summary"`   // description
	Instances []instance `json:"instances"` // see Instance
}

func (m cloud) String() string { return util.ToYAML(m) }

// Instance represents a single specific instance of a Machine with
// a set of specific parameters that are unique to each (IP, name, etc.)
//
type instance struct {
	Name    string `json:"name"`    // becomes hostname
	Summary string `json:"summary"` // description
	Machine string `json:"machine"` // Machine.Name
	Address string `json:"address"` // (starting) IP address
	Count   int    `json:"count"`   // default 1, 250 max
}

func (m instance) String() string { return util.ToYAML(m) }
