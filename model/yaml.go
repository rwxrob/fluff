package model

import _ "embed"

//go:embed fluff.yaml
var DefaultYAML string

// YAMLFile is set to the name of the initialized YAML file containing
// cloud and machine configuration data (default: fluff.yaml)
//
var YAMLFile = "fluff.yaml"
