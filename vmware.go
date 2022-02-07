package fluff

import (
	"os/exec"
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

func (p *vmwareProvider) create(i instance) error   { return nil }
func (p *vmwareProvider) destroy(i instance) error  { return nil }
func (p *vmwareProvider) start(i instance) error    { return nil }
func (p *vmwareProvider) stop(i instance) error     { return nil }
func (p *vmwareProvider) snapshot(i instance) error { return nil }
