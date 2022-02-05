package model

// Machine represents a single virtual machine with all of its
// virtualized hardware, network configuration, and a single
// administrative user. Machines provide the basis for Instances
// specified within a cloud of a specific Machine configuration.
//
type Machine struct {
	Name    string   `json:"name,omitempty"`    // local unique, dotted
	Cores   int      `json:"cores,omitempty"`   // not CPUs
	Memory  int      `json:"memory,omitempty"`  // megabytes
	Volumes []Volume `json:"volumes,omitempty"` // disk files
	URL     string   `json:"url,omitempty"`     // curl-able
}

// Volume represents a single virtual disk volume file to be created and
// optionally mounted at the specified location.
//
type Volume struct {
	Size  int    `json:"mb,omitempty"`    // megabytes
	Mount string `json:"mount,omitempty"` // mount point (/s)
}
