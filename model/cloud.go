package model

// Cloud represents a single collection of machines that provide
// a particular flavor of infrastructure on which to build systems,
// explore configurations, conduct training, and install distributed
// applications (like Kubernetes from kubeadm).
//
type Cloud struct {
	Name        string     `json:"name,omitempty"`        // ascii + .
	Description string     `json:"description,omitempty"` // unlimited
	Instances   []Instance `json:"instances,omitempty"`   // see Instance
}

// Instance represents a single specific instance of a Machine with
// a set of specific parameters that are unique to each (IP, name, etc.)
//
type Instance struct {
	Name    string `json:"name,omitempty"`    // becomes hostname
	Machine string `json:"machine,omitempty"` // Machine.Name
	Address string `json:"address,omitempty"` // (starting) IP address
	Count   int    `json:"count,omitempty"`   // default 1, 250 max
}
