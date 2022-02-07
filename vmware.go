package fluff

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"regexp"
	"sort"
)

var vmware = new(vmwareProvider)

type vmwareProvider struct {
	vmrun        string
	vdiskmanager string
	mkisofs      string
	qemuimg      string
}

func (p *vmwareProvider) findExecutables() bool {

	// look for windows specific exes in the PATH

	if exe, _ := exec.LookPath("vmrun.exe"); len(exe) > 0 {
		p.vmrun = exe
	}
	if exe, _ := exec.LookPath("vmware-vdiskmanager.exe"); len(exe) > 0 {
		p.vdiskmanager = exe
	}
	if exe, _ := exec.LookPath("mkisofs.exe"); len(exe) > 0 {
		p.mkisofs = exe
	}
	if exe, _ := exec.LookPath("qemu-img.exe"); len(exe) > 0 {
		p.qemuimg = exe
	}

	if p.vmrun != "" && p.vdiskmanager != "" && p.mkisofs != "" && p.qemuimg != "" {
		return true
	}

	// look for specific known paths that are not in PATH

	// look for windows wsl2 in the usual locations
	// /mnt/c/Program Files (x86)/VMware/VMware Workstation/vmrun.exe

	// look for windows in the usual locations

	// try mac path and mac expected location
	// /Applications/VMware Fusion.app/Contents/Public/vmrun
	// /Applications/VMware Fusion.app/Contents/Library/vmrun

	// try linux path

	// try linux expected locations for ubuntu

	// try linux expected locations for rhel

	// try linux expected locations for arch

	return false
}

func (p *vmwareProvider) detect() bool { return p.findExecutables() }

func (p *vmwareProvider) create(i instance) error {
	machine := getmachine(i.Machine)
	fmt.Println(machine.Name)

	return nil
}

func (p *vmwareProvider) destroy(i instance) error  { return nil }
func (p *vmwareProvider) start(i instance) error    { return nil }
func (p *vmwareProvider) stop(i instance) error     { return nil }
func (p *vmwareProvider) snapshot(i instance) error { return nil }

func (p *vmwareProvider) status(i instance) string {
	statuses, err := p.getstatuses()
	if err != nil {
		log.Println(err)
		return ""
	}
	return statuses[i.Name]
}

func (p *vmwareProvider) getnames() []string {
	names := []string{}
	for _, i := range current_cloud.Instances {
		name := i.Name
		if i.Count > 1 {
			for n := 0; n < i.Count; n++ {
				name = fmt.Sprintf("%v-%v", i.Name, n+1)
				names = append(names, name)
			}
			continue
		}
		names = append(names, name)
	}
	sort.Strings(names)
	return names
}

func (p *vmwareProvider) getstatuses() (map[string]string, error) {
	names := p.getnames()

	statuses := map[string]string{}
	for _, name := range names {
		statuses[name] = "down"
	}

	out, err := exec.Command(p.vmrun, "list").Output()
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(bytes.NewReader(out))
	for scanner.Scan() {
		line := scanner.Text()
		regx := regexp.MustCompile(`([\w-.]+)\.vmx$`)
		matches := regx.FindStringSubmatch(line)
		if len(matches) > 0 {
			statuses[matches[1]] = "up"
		}
	}

	return statuses, nil
}

func (p *vmwareProvider) list() error {
	names := p.getnames()
	statuses, err := p.getstatuses()
	if err != nil {
		return err
	}
	for _, name := range names {
		fmt.Printf("%v %v\n", name, statuses[name])
	}
	return nil
}
