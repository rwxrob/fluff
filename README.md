# ☁🏠☁ Fluff, Happy Little Clouds at Home

*Update: After reading more about the libvirt project, kcli (python),
and virt-lightning (also python) so eventually want to support libvirt,
but everything is direct for now.

*"It's docker-compose for cloud-init VMs."*

* Support VMware Workstation Pro and VirtualBox
* Focus *only* on virtual hardware configuration
* Compliment Ansible for system configuration
* Cater to absolute beginners as well as pros
* Simplest possible CLI with optional web UI
* Local cloud VMs organized in `$HOME/Fluff`
* Bridged static IP network support only
* Highly opinionated defaults
* Only allow cloud-init images
* Secure shell into everything
* Batteries included

## Prerequisites

Fluff is primarily designed to be run from any modern Windows or Mac
computer. It will also run from most Linux desktop distros as well:


## Installation


1. Install a terminal (Windows Terminal, iTerm2, etc.)
1. Install VMware Workstation Pro (recommended) or VirtualBox
1. Install `qemu-img`
1. Install `fluff` binary

Fluff will also run from most Linux 

> ⚠️
> Note Windows users do *not* need to setup WSL2 but it won't hurt.

First, you'll need to install a good terminal to make best use of your
home local cloud virtual machines so that you can connect to them from
`ssh` (which is installed by default on all major desktop operating
systems these days).

Second, install one of the following industry standard desktop
virtualization applications:

* VMware Workstation Pro (recommended)
* Oracle VirtualBox

Third, install `qemu-img` depending on your operating system. This
allows `fluff` to convert between virtual machine images designed for
the cloud (`qcow2`) into images that work with VMware (`vmdk`) or
VirtualBox (`vdi`). (This application also has different licensing
requiring it to be installed separately.)

And finally, install the `fluff` binary using any of the following
methods:

If you have Go already installed:

```
go install github.com/rwxrob/fluff@latest
```

Otherwise, grab the binary for your computer and add put it someplace
your computer will know to run it.

```
curl ...
```

## Usage

Here are the main user-facing commands:

```
fluff help [COMMAND] - display help
fluff init           - create a starter fluff.yaml file
fluff lint           - check the fluff.yaml file for syntax and more
fluff up [FILE|URL]  - start a local cloud, display description
fluff down - stop local cloud VMs and save their state
fluff away - destroy and delete all local cloud VMs
fluff shot - take a "snapshot" of all local cloud VMs
fluff list - list all the local cloud VMs with name and IP
fluff copy - copy all local cloud VMs and config to target 
```

There are a number of other (hidden) commands that allow users to do
specific things that are a part of the other main user commands:

```
fluff cached                  - print path to cached image directory
fluff cache URL              - retrieve and cache distro image
fluff convert (vbox|vmware) TARGET   - convert target image to specified format

fluff iso [PATH]              - create a cloudinit.iso
fluff volume [MB] [FORMAT]    - create volume file of size and format
fluff ssh-config [INSTANCE]   - output ~/.ssh/config Host for instance
```

And a few debugging hidden commands:

```
fluff err  - prepare large detailed report of last error
```

And a few expected aliases:

```
fluff start   - up
fluff stop    - down
fluff off     - down
fluff vet     - lint
fluff destroy - away
fluff rm      - away
fluff snap    - shot
fluff ps      - list
fluff cp      - copy
```

## Domain Model and Terminology

**machine** - specification of a virtual machine
**instance** - a specific instance of a configured virtual machine  
**cloud** - a collection of *instances*

## Machine (Box) Specification Types

A *machine* consists of all the hardware specification information.
Certain default specs are included within the binary while users may
create their own in the `specs.yaml` file. Here are the current
defaults:

```yaml
machines:
  - name: alma8.server 
    cores: 2
    memory: 2048
    volumes:
      - size: 100 
    url: |
      https://repo.almalinux.org/almalinux/8/cloud/x86_64/images/AlmaLinux-8-GenericCloud-8.5-20211119.x86_64.qcow2
  - name: alma8.node
    base: alma8.server
    cores: 1

clouds:
  - name: basic
    description: |
      Simple control and three node mini-cloud suitable for testing
      basic endpoint architecture and applications such as Kubernetes
      installed with kubeadm. Default IPs: 192.168.132.10-13.
    instances:
      - name: control
        base: alma8.server 
        ip: 192.168.132.10
    group:
      - name: node           # node-1, node-2, node-3
        base: alma8.node
        count: 3
        startip: 192.168.132.11
```

> Note that this exact YAML file is embedded in the Go binary.

## Primary Use Cases

* Administrators, operators, cloud-native engineers, systems engineers
  to experiment for testing and learning purposes 
* Experimenting and testing any cloud-init enabled virtual machine image
* Simulate specific, real-world networks and traffic within them
  including boding of interfaces, etc.

## YAML Configuration Files

## "Why not use Vagrant?"

❌ It's scope of use-cases is too large  
❌ It's ancient (lots of technical debt)  
❌ It adds layer of unnecessary abstraction  
❌ It wasn't conceived with `cloud-init` in mind  
❌ It doesn't play nice with WSL2  
❌ It has a screwed up network model  
❌ It isn't a single executable    
❌ It's in Ruby 

