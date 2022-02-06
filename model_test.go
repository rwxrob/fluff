package fluff

import (
	"fmt"

	"gopkg.in/yaml.v2"
)

func ExampleVolume() {
	vol := new(volume)

	// yaml
	buf := "size: 1024\nmount: /s\n"
	err := yaml.Unmarshal([]byte(buf), vol)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(vol.Size)
	fmt.Println(vol.Mount)

	// json
	buf = `{"size": 2048, "mount": "/m"}`
	err = yaml.Unmarshal([]byte(buf), vol)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(vol.Size)
	fmt.Println(vol.Mount)

	// Output:
	// 1024
	// /s
	// 2048
	// /m
}
