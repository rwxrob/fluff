package fluff

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/rwxrob/cmdbox/util"
	"gopkg.in/yaml.v2"
)

var current_cloud *cloud
var current_provider provider
var longestNameLength int

func init() {
	err := yaml.Unmarshal([]byte(DefaultYAML), Manifest)
	if err != nil {
		panic(err)
	}
	// TODO overlay the local fluff.yaml (if found)
	createMachineIndex()
	createCloudIndex()
	current_cloud = getcloud("default")
	current_provider = detectProvider()
	expandInstances()
	setLongestNameLength()
}

const (
	start_default_cloud = `# default cloud`
	end_default_cloud   = `# end default cloud`
)

type provider interface {
	detect() bool
	create(i instance) error
	destroy(i instance) error
	start(i instance) error
	stop(i instance) error
	snapshot(i instance) error // save a snapshot
	status(i instance) string  // up, down, "" (indeterminate)
	list() error               // name and status
}

var endip = regexp.MustCompile(`(\d+.\d+.\d+\.)(\d+)$`)

func setLongestNameLength() {
	for _, i := range current_cloud.Instances {
		length := len(i.Name)
		if length > longestNameLength {
			longestNameLength = length
		}
	}
}

func incrementip(ip string, n int) string {
	m := endip.FindStringSubmatch(ip)
	if len(m) > 2 {
		i, err := strconv.Atoi(m[2])
		if err != nil {
			return ip
		}
		return fmt.Sprintf("%v%v", m[1], i+n)
	}
	return ip
}

func expandInstances() {
	instances := []instance{}
	for _, i := range current_cloud.Instances {
		if i.Count > 1 {
			for n := 0; n < i.Count; n++ {
				ni := i
				ni.Count = 1
				ni.Address = incrementip(ni.Address, n)
				ni.Name = fmt.Sprintf("%v%v", ni.Name, n+1)
				instances = append(instances, ni)
			}
			continue
		}
		instances = append(instances, i)
	}
	current_cloud.Instances = instances
}

func detectProvider() provider {
	switch {
	case vmware.detect():
		return vmware
	case vbox.detect():
		return vbox
	}
	return nil
}

func setcurrent(name string) error {
	c := getcloud(name)
	if c == nil {
		return fmt.Errorf("cloud not found: %v", name)
	}
	current_cloud = c
	return nil
}

type manifest struct {
	Machines []machine `json:"machines"`
	Clouds   []cloud   `json:"clouds"`
}

func (m manifest) String() string { return util.ToYAML(m) }

var machines = map[string]*machine{}

func createMachineIndex() {
	for n, m := range Manifest.Machines {
		if m.Name == "" {
			continue
		}
		machines[m.Name] = &Manifest.Machines[n]
	}
}

func getmachine(name string) *machine {
	if v, has := machines[name]; has {
		return v
	}
	return nil
}

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

var clouds = map[string]*cloud{}

func createCloudIndex() {
	for n, c := range Manifest.Clouds {
		if c.Name == "" {
			continue
		}
		clouds[c.Name] = &Manifest.Clouds[n]
	}
}

func getcloud(name string) *cloud {
	if v, has := clouds[name]; has {
		return v
	}
	return nil
}

type cloud struct {
	Name      string     `json:"name"`      // ascii + .
	Summary   string     `json:"summary"`   // description
	Instances []instance `json:"instances"` // see Instance
}

func (m cloud) String() string { return util.ToYAML(m) }

type instance struct {
	Name    string `json:"name"`    // becomes hostname
	Summary string `json:"summary"` // description
	Machine string `json:"machine"` // Machine.Name
	Address string `json:"address"` // (starting) IP address
	Count   int    `json:"count"`   // default 1, 250 max
}

func (m instance) String() string { return util.ToYAML(m) }
