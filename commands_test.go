package main

import (
	"os"
	"os/exec"
	"testing"
)

func TestCmdExit(t *testing.T) {
	c := Config{}
	if os.Getenv("BE_CRASHER") == "1" {
		CommandExit(&c)
		return
	}
	cmd := exec.Command(os.Args[0], "-test.run=TestCmdExit")
	cmd.Env = append(os.Environ(), "BE_CRASHER=1")
	if err := cmd.Run(); err == nil {
		return
	} else {
		t.Fatalf("Process ran with err %v, want exit status 0", err)
	}
}
