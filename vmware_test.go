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
