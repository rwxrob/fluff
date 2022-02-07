package fluff

import (
	"fmt"
	"os"
)

// Note that these tests are rather difficult to set for any testing
// environment since we are testing for the sorts of things that we
// would be using to do the host emulation for such testing.

func ExampleVMware() {

	if os.Getenv("TEST_WITHOUT_VMWARE") != "" {
		fmt.Println(true)
		return
	}

	// make sure you have all the required executables
	// before running this test

	fmt.Println(vmware.detect())

	// Output:
	// true
}

/*
func Example_vmware_create() {
	i := new(instance)
	i.Name = "myinstance"
	i.Machine = "alma8.server"
	err := vmware.create(i)
	if err != nil {
		fmt.Println(err)
	}
	// TODO call list and get output, check for myinstance in output
	// Output:
	// 2
}
*/

func Example_vmware_names() {
	for _, name := range vmware.getnames() {
		fmt.Println(name)
	}
	// Output:
	// control
	// node-1
	// node-2
	// node-3
}

/*
// bring up control and node-1 before running these tests

func Example_vmware_getstatuses() {
	statuses, err := vmware.getstatuses()
	if err != nil {
		fmt.Println(err)
	}
	for name, status := range statuses {
		fmt.Printf("%v %v\n", name, status)
	}
	// Output:
	// control up
	// node-1 up
	// node-2 down
	// node-3 down
}
*/

func Example_vmware_list() {
	err := vmware.list()
	if err != nil {
		fmt.Println(err)
	}
	// Output:
	// control up
	// node-1 up
	// node-2 down
	// node-3 down
}
