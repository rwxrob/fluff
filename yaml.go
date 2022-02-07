package fluff

import (
	_ "embed"
)

//go:embed fluff.yaml
var DefaultYAML string

// YAMLFile is set to the name of the initialized YAML file containing
// cloud and machine configuration data (default: fluff.yaml)
//
var YAMLFile = "fluff.yaml"

const (
	start_default_cloud = `# default cloud`
	end_default_cloud   = `# end default cloud`
)
