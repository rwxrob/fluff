# ☁🏠☁ Fluff, Happy Little Clouds at Home

*"It's docker-compose for cloud-init VMs."*

* Support VMware Workstation Pro and VirtualBox
* Focus *only* on virtual hardware configuration
* Compliment Ansible for system configuration
* Cater to absolute beginners as well as pros
* Simplest possible CLI with optional web UI
* Highly opinionated defaults
* Only allow cloud-init images
* Static IP network support only
* Secure shell into everything
* Batteries included

## Installation

Fluff is designed to be run from any Windows or Mac workstation
without any prerequisites other than the following:

1. Install a terminal (Windows Terminal, iTerm2, etc.)
1. Install VMware Workstation Pro (recommended) or VirtualBox
1. Install `qemu-img`
1. Install `fluff` binary

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

Otherwise, grab the binary for your computer type:

```
curl ...
```

## Usage

```
fluff init - create a starter fluff.yaml file
fluff lint - check the fluff.yaml file for syntax and more
fluff up   - start a local cloud, apply fluff.yaml if found
fluff down - stop local cloud VMs and save their state
fluff away - destroy and delete all local cloud VMs
fluff shot - take a "snapshot" of all local cloud VMs
fluff list - list all the local cloud VMs with name and IP
fluff copy - copy all local cloud VMs and config to target 
fluff help - display help
```

There are a number of hidden commands that are contained within the
other main commands but can be called individually:

```
fluff validate [PATH]         - validate a fluff.yaml file
fluff fetch [URL]             - retrieve and cache distro image
fluff convert [PATH] [FORMAT] - convert target image to specified format
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

