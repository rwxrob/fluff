package main

import (
	"github.com/rwxrob/cmdbox"
	"github.com/rwxrob/fluff"
)

func init() {
	var x *cmdbox.Command

	// ----------------- fluff command -------------------

	x = cmdbox.Add("fluff", "i|init", "start|u|up", "list")
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

	// ------------------ up command --------------------

	x = cmdbox.Add("up")
	x.Summary = `startup a local cloud of virtual machines`
	x.Usage = `[CLOUD]`
	x.Description = `
		Starts up a local cloud of virtual machines using the detected
		provider and optionally creates them as needed. By default, starts
		up the "basic" cloud consisting of one "control" and three "node"
		machines. To specify another cloud configuration pass the name as an
		optional argument. See the init command for more about how to create
		custom cloud configurations in the ` + fluff.YAMLFile + `.`
	x.Method = fluff.Up

	// ----------------- list command -------------------

	x = cmdbox.Add("list")
	x.Summary = `list names of virtual machines in current local cloud`
	x.Description = `
		List and the names of all the virtual machines and whether they are
		up or down.`
	x.Method = fluff.List

}

func main() {
	cmdbox.Execute()
}
