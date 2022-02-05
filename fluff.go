package main

import (
	"github.com/rwxrob/cmdbox"
)

func init() {
	x := cmdbox.Add("fluff", "i|init")
	x.Summary = `happy little clouds at home`
	x.Description = `
		Create and explore different cloud virtual machine configurations
		locally at home on your VMware enabled workstation.

    * Focus *only* on virtual hardware configuration
    * Compliment Ansible for system configuration
    * Only allow cloud-init images
    * Primary support for VMware Workstation Pro
    * Secondary support for VirtualBox
    * Fore-knowledge of static IPs required
    * Highly opinionated defaults
    * Batteries included
    * Simplest CLI possible

		`
}
