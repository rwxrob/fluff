package main

import (
	_ "embed"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/rwxrob/cmdbox"
	"github.com/rwxrob/cmdbox/util"
)

// YAMLFile is set to the name of the initialized YAML file containing
// cloud and machine configuration data (default: fluff.yaml)
//
var YAMLFile = "fluff.yaml"

// StartDefaultCloud is the regular expression that will be checked for
// the beginning of a line to indicate the beginning of the default
// cloud section in the YAMLFile.
//
var StartDefaultCloud = `# default cloud`

// EndDefaultCloud is the regular expression that will be checked for
// at beginning of a line to indicate the ending of the default
// cloud section in the YAMLFile.
//
var EndDefaultCloud = `# end default cloud`

//go:embed fluff.yaml
var _YAML string

func init() {
	// TODO parse the _YAML file into a struct
}

func init() {
	x := cmdbox.Add("init", "simple", "full")
	x.Params = []string{"simple", "full"}
	x.Summary = `initialize a default YAML file (` + YAMLFile + `)`
	x.Usage = `[simple|full]`
	x.Description = `
		Creates a ` + YAMLFile + ` file with documented defaults
		("commented") suitable for configuring a fluff cloud. Pass the
		"simple" argument if you prefer no comments or "full" if you wish
		every single default configuration with documentation (including the
		list of available machine types and sources).`

	x.Method = func(args []string) error {
		if util.Found(YAMLFile) {
			return fmt.Errorf(
				"cowardly refusing to overwrite existing %s", YAMLFile)
		}
		r := strings.NewReader(_YAML)
		detail := "simple"
		if len(args) > 0 {
			detail = args[0]
		}
		switch detail {
		case "full":
			err := os.WriteFile(YAMLFile, []byte(_YAML), 0600)
			if err != nil {
				return err
			}
			log.Printf("created full %v file", YAMLFile)
			return nil
		case "simple":
			buf := util.GetSection(r, StartDefaultCloud, EndDefaultCloud)
			err := os.WriteFile(YAMLFile, []byte(buf), 0600)
			if err != nil {
				return err
			}
			log.Printf("created simple %v file", YAMLFile)
			return nil
		}
		return x.UsageError()
	}
}
