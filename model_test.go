package fluff

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

func Example_machines() {
	fmt.Println(getmachine("alma8.server").Cores)
	fmt.Println(getmachine("alma8.node").Cores)
	// Output:
	// 2
	// 1
}

func Example_clouds() {
	for _, i := range getcloud("basic").Instances {
		fmt.Println(i.Name)
	}
	// Output:
	// control
	// node
}

func Example_volume() {
	vol := new(volume)

	// yaml
	buf := "size: 1024\n"
	err := yaml.Unmarshal([]byte(buf), vol)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(vol.Size)

	// json
	buf = `{"size": 2048}`
	err = yaml.Unmarshal([]byte(buf), vol)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(vol.Size)

	// Output:
	// 1024
	// 2048

}

func Example_manifest() {
	m := new(manifest)
	buf, err := os.ReadFile("fluff.yaml")
	if err != nil {
		fmt.Println(err)
	}
	err = yaml.Unmarshal(buf, m)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(m)

	// Output:
	// machines:
	// - name: alma8.server
	//   cores: 2
	//   memory: 2048
	//   volumes:
	//   - size: 100
	//   url: |
	//     https://repo.almalinux.org/almalinux/8/cloud/x86_64/images/AlmaLinux-8-GenericCloud-8.5-20211119.x86_64.qcow2
	//   base: ""
	// - name: alma8.node
	//   cores: 1
	//   memory: 0
	//   volumes: []
	//   url: ""
	//   base: alma8.server
	// clouds:
	// - name: basic
	//   summary: |
	//     Simple control and three node mini-cloud suitable for testing
	//     basic endpoint architecture and applications such as Kubernetes
	//     installed with kubeadm. Default IPs: 192.168.132.10-13.
	//   instances:
	//   - name: control
	//     summary: ""
	//     machine: alma8.server
	//     address: 192.168.132.10
	//     count: 1
	//   - name: node
	//     summary: ""
	//     machine: alma8.node
	//     address: 192.168.132.20
	//     count: 3

}

func Example_machine() {
	m := new(machine)
	buf := `
name: alma8.server
cores: 2
memory: 2048
volumes:
- size: 10
- size: 20
url: https://repo.almalinux.org/almalinux/8/cloud/x86_64/images/AlmaLinux-8-GenericCloud-8.5-20211119.x86_64.qcow2
`
	err := yaml.Unmarshal([]byte(buf), m)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(m)

	// Output:
	// name: alma8.server
	// cores: 2
	// memory: 2048
	// volumes:
	// - size: 10
	// - size: 20
	// url: https://repo.almalinux.org/almalinux/8/cloud/x86_64/images/AlmaLinux-8-GenericCloud-8.5-20211119.x86_64.qcow2
	// base: ""

}

func Example_instance() {
	m := new(instance)
	buf := `
name: control
summary: control-plane for high availability
machine: alma8.server
address: 192.168.132.2
count: 3 # quorum
`
	err := yaml.Unmarshal([]byte(buf), m)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(m)

	// Output:
	// name: control
	// summary: control-plane for high availability
	// machine: alma8.server
	// address: 192.168.132.2
	// count: 3

}

func Example_cloud() {
	c := new(cloud)
	buf := `
name: basic
summary: |
  Simple control and three node mini-cloud suitable for testing
  basic endpoint architecture and applications such as Kubernetes
  installed with kubeadm. Default IPs: 192.168.132.10-13.
instances:
- name: control              # control
  summary: control plane
  machine: alma8.server 
  address: 192.168.132.10
  count: 1
- name: node                 # node-1, node-2, node-3
  summary: worker nodes
  machine: alma8.node
  address: 192.168.132.20    # .20, .21, .22
  count: 3
`

	if err := yaml.Unmarshal([]byte(buf), c); err != nil {
		fmt.Println(err)
	}
	fmt.Println(c)

	// Output:
	// name: basic
	// summary: |
	//   Simple control and three node mini-cloud suitable for testing
	//   basic endpoint architecture and applications such as Kubernetes
	//   installed with kubeadm. Default IPs: 192.168.132.10-13.
	// instances:
	// - name: control
	//   summary: control plane
	//   machine: alma8.server
	//   address: 192.168.132.10
	//   count: 1
	// - name: node
	//   summary: worker nodes
	//   machine: alma8.node
	//   address: 192.168.132.20
	//   count: 3

}
