package main

import (
	"make/mymake"
	"os"
	"testing"
)

var m *mymake.Make

func TestMain(m *testing.M) {
	Setup()
	os.Exit(m.Run())
}

func Setup() {
	m = mymake.NewMake("./data_seq")
	pm = mymake.NewMake("./data_para")
}
