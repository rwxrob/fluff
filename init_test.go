package main

import (
	"os"
	"testing"

	"github.com/rwxrob/cmdbox"
	"github.com/rwxrob/fluff/model"
)

func Test_init_full(t *testing.T) {
	model.YAMLFile = `testdata/fluff.yaml`
	os.Remove(model.YAMLFile)
	defer os.Remove(model.YAMLFile)
	cmdbox.Call(nil, "init", "full")
	got, _ := os.ReadFile(model.YAMLFile)
	want, _ := os.ReadFile("testdata/init_full.yaml")
	if string(got) != string(want) {
		t.Log(string(got))
		t.Fail()
	}
}

func Test_init_simple(t *testing.T) {
	model.YAMLFile = `testdata/fluff.yaml`
	os.Remove(model.YAMLFile)
	defer os.Remove(model.YAMLFile)
	cmdbox.Call(nil, "init", "simple")
	got, _ := os.ReadFile(model.YAMLFile)
	want, _ := os.ReadFile("testdata/init_simple.yaml")
	if string(got) != string(want) {
		t.Log(string(got))
		t.Fail()
	}
}

func Test_init_bad(t *testing.T) {
	model.YAMLFile = `testdata/fluff.yaml`
	os.Remove(model.YAMLFile)
	defer os.Remove(model.YAMLFile)
	err := cmdbox.Call(nil, "init", "bork")
	if err == nil {
		t.Error("woah, bork is not an init command, but no error")
	}
}
