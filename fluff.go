/*

Package fluff defines the fluff domain model and fluff.Provider
interface fulfilled by the VirtualBox and VMware providers that are used
by the cross-platform fluff command utility for provisioning "happy,
little clouds at home".

*/
package fluff

import (
	_ "embed"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/rwxrob/cmdbox/util"
)

const (
	EMOJI_INSTANCE_UP       = '\U000026A1' // ⚡
	EMOJI_INSTANCE_DOWN     = '\U0001F4A4' // 💤
	EMOJI_INSTANCE_NOTFOUND = '\U00002753' // ❓
	EMOJI_CLOUD_UP          = '\U0001F329' // 🌩️
	EMOJI_CLOUD_DOWN        = '\U00002601' // ☁️
)

// YAMLFile is set to the name of the initialized YAML file containing
// cloud and machine configuration data (default: fluff.yaml)
//
var YAMLFile = "fluff.yaml"

// DefaultYAML contains the YAML file maintained with the library that
// contains all included machines and cloud configurations. This data is
// written verbatim to fluff.yaml when Init("full") is called.
//go:embed fluff.yaml
var DefaultYAML string

// Manifest starts out as the parsed DefaultYAML data (which is embedded
// in the binary) but is then modified by overlaying the local
// fluff.yaml file.
//
var Manifest = new(manifest)

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

// List prints the current local cloud virtual machine instances by name
// and whether each is up or down.
//
func List(args ...string) error { return current_provider.list() }

// Up optionally accepts the name of a single cloud from the default or
// custom configuration file and concurrently starts up all the VM
// instances specified. Up will automatically create any instances that
// have not yet been created. If no cloud name is provided the name
// "default" will be assumed. Up will stop on the first instance if
// unable to start or create it and will output the error from the
// provider. Up calls List automatically whether successful or not.
//
func Up(args ...string) error {
	defer List()
	if len(args) > 0 {
		err := setcurrent(args[0])
		if err != nil {
			return err
		}
	}
	for _, i := range current_cloud.Instances {
		// TODO make this concurrent
		err := current_provider.create(i)
		if err != nil {
			return err
		}
	}
	return nil
}
