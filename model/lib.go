package model

import _ "embed"

//go:embed fluff.yaml
var DefaultYAML string

// YAMLFile is set to the name of the initialized YAML file containing
// cloud and machine configuration data (default: fluff.yaml)
//
var YAMLFile = "fluff.yaml"

type Manifest struct {
	Machines []Machine
	Clouds   []Cloud
}

type Machine struct{}

type Cloud struct{}

func parseYAML(buf string) {
}
