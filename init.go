package fluff

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/rwxrob/cmdbox/util"
)

func Init(args ...string) error {

	if util.Found(YAMLFile) {
		return fmt.Errorf(
			"cowardly refusing to overwrite existing %s", YAMLFile)
	}

	r := strings.NewReader(DefaultYAML)
	detail := "simple"
	if len(args) > 0 {
		detail = args[0]
	}

	switch detail {
	case "full":
		err := os.WriteFile(YAMLFile,
			[]byte(DefaultYAML), 0600)
		if err != nil {
			return err
		}
		log.Printf("created full %v file", YAMLFile)
		return nil
	case "simple":
		buf := util.GetSection(r, start_default_cloud, end_default_cloud)
		err := os.WriteFile(YAMLFile, []byte(buf), 0600)
		if err != nil {
			return err
		}
		log.Printf("created simple %v file", YAMLFile)
		return nil
	}

	return fmt.Errorf("unexpected arguments: %v", args)
}
