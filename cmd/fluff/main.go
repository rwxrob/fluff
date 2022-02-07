package main

import (
	"github.com/rwxrob/cmdbox"
	"github.com/rwxrob/fluff"
)

func init() {
	var x *cmdbox.Command

	// ----------------- fluff command -------------------

	x = cmdbox.Add("fluff", "i|init")
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

	// ----------------- init command --------------------

	x = cmdbox.Add("init")
	x.Params = []string{"simple", "full"}
	x.Summary = `initialize a default YAML file (` + fluff.YAMLFile + `)`
	x.Usage = `[simple|full]`
	x.Description = `
	  Creates a ` + fluff.YAMLFile + ` file with documented defaults
	  suitable for configuring a fluff cloud. Pass the "full" argument if
	  you wish every single default configuration with documentation
	  (including the list of available machine types and sources).`
	x.Method = fluff.Init

}

func main() {
	cmdbox.Execute()
}
