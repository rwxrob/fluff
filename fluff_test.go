package fluff_test

import (
	"fmt"
	"os"

	"github.com/rwxrob/fluff"
)

func ExampleInit_simple() {
	fluff.YAMLFile = `testdata/fluff.yaml`
	defer os.Remove(fluff.YAMLFile)

	err := fluff.Init("simple")
	if err != nil {
		fmt.Println(err)
	}

	got, _ := os.ReadFile(fluff.YAMLFile)
	want, _ := os.ReadFile(`testdata/init_simple.yaml`)

	fmt.Println(string(got) == string(want))

	// Output:
	// true
}

func ExampleInit_full() {
	fluff.YAMLFile = `testdata/fluff.yaml`
	defer os.Remove(fluff.YAMLFile)

	err := fluff.Init("full")
	if err != nil {
		fmt.Println(err)
	}

	got, _ := os.ReadFile(fluff.YAMLFile)
	want, _ := os.ReadFile(`testdata/init_full.yaml`)

	fmt.Println(string(got) == string(want))

	// Output:
	// true
}

func ExampleInit_bork() {
	fluff.YAMLFile = `testdata/fluff.yaml`
	defer os.Remove(fluff.YAMLFile)
	err := fluff.Init("bork")
	fmt.Println(err)
	// Output:
	// unexpected arguments: [bork]
}

func ExampleUp_default() {
	fluff.Up()

	// Output:
	// control
	// node-1
	// node-2
	// node-3
}
