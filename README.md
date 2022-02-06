# ☁🏠☁ Fluff, Happy Little Clouds at Home

*"It's docker-compose for cloud-init VMs."*

* Focus *only* on virtual hardware configuration
* Compliment Ansible for system configuration
* Only allow cloud-init images
* Primary support for VMware Workstation Pro
* Secondary support for VirtualBox
* Fore-knowledge of static IPs required
* Cater to absolute beginners
* Highly opinionated defaults
* Secure shell into everything
* Batteries included
* Simplest CLI possible

## Installation

```
go install github.com/rwxrob/fluff@latest
```

## Usage

```
fluff init - create a starter fluff.yaml file
fluff up   - start a cloud from the fluff.yaml file, init if not found
fluff down - stop the cloud VMs and save their state
fluff rm   - destroy and delete all of a cloud
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

## Domain Model

**machine** - single instance of a machine with a given *spec*
**cloud** - a collection of *machines*

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

## Networking

The `base.ip` value is incremented by one for each machine in a `group`.

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

