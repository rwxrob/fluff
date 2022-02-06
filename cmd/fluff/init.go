package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/rwxrob/cmdbox"
	"github.com/rwxrob/cmdbox/util"
	"github.com/rwxrob/fluff/model"
)

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

func init() {
	x := cmdbox.Add("init")
	x.Params = []string{"simple", "full"}
	x.Summary = `initialize a default YAML file (` + model.YAMLFile + `)`
	x.Usage = `[simple|full]`
	x.Description = `
	  Creates a ` + model.YAMLFile + ` file with documented defaults
	  suitable for configuring a fluff cloud. Pass the "full" argument if
	  you wish every single default configuration with documentation
	  (including the list of available machine types and sources).`
	x.Method = func(args []string) error {
		if util.Found(model.YAMLFile) {
			return fmt.Errorf(
				"cowardly refusing to overwrite existing %s", model.YAMLFile)
		}
		r := strings.NewReader(model.DefaultYAML)
		detail := "simple"
		if len(args) > 0 {
			detail = args[0]
		}
		switch detail {
		case "full":
			err := os.WriteFile(model.YAMLFile,
				[]byte(model.DefaultYAML), 0600)
			if err != nil {
				return err
			}
			log.Printf("created full %v file", model.YAMLFile)
			return nil
		case "simple":
			buf := util.GetSection(r, StartDefaultCloud, EndDefaultCloud)
			err := os.WriteFile(model.YAMLFile, []byte(buf), 0600)
			if err != nil {
				return err
			}
			log.Printf("created simple %v file", model.YAMLFile)
			return nil
		}
		return x.UsageError()
	}
}
