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

func Example_vmware_list() {

	// Output:
	// some
}
*/
