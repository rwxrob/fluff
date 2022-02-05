package main

import (
	"os"
	"testing"

	"github.com/rwxrob/cmdbox"
)

func Test_init_full(t *testing.T) {
	YAMLFile = `testdata/fluff.yaml`
	os.Remove(YAMLFile)
	defer os.Remove(YAMLFile)
	cmdbox.Call(nil, "init", "full")
	got, _ := os.ReadFile(YAMLFile)
	want, _ := os.ReadFile("testdata/init_full.yaml")
	if string(got) != string(want) {
		t.Log(string(got))
		t.Fail()
	}
}

func Test_init_simple(t *testing.T) {
	YAMLFile = `testdata/fluff.yaml`
	os.Remove(YAMLFile)
	defer os.Remove(YAMLFile)
	cmdbox.Call(nil, "init", "simple")
	got, _ := os.ReadFile(YAMLFile)
	want, _ := os.ReadFile("testdata/init_simple.yaml")
	if string(got) != string(want) {
		t.Log(string(got))
		t.Fail()
	}
}

func Test_init_bad(t *testing.T) {
	YAMLFile = `testdata/fluff.yaml`
	os.Remove(YAMLFile)
	defer os.Remove(YAMLFile)
	err := cmdbox.Call(nil, "init", "bork")
	if err == nil {
		t.Error("woah, bork is not an init command, but no error")
	}
}
