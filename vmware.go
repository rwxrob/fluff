package fluff

import "os/exec"

var vmware = new(vmwareProvider)

type vmwareProvider struct {
	vmrun        string
	vdiskmanager string
}

// should always find the run binaries
func (p *vmwareProvider) detect() bool {

	// first look for windows in the common path
	if exe, _ := exec.LookPath("vmrun.exe"); len(exe) > 0 {
		p.vmrun = exe
		// TODO look for vdiskmanager
		return true
	}

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

// TODO implement these
func (p *vmwareProvider) create(i instance) error   { return nil }
func (p *vmwareProvider) destroy(i instance) error  { return nil }
func (p *vmwareProvider) start(i instance) error    { return nil }
func (p *vmwareProvider) stop(i instance) error     { return nil }
func (p *vmwareProvider) snapshot(i instance) error { return nil }
