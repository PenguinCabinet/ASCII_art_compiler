package main

import (
	"os/exec"
	"testing"
)

func one_cli_run(t *testing.T) {
	err := exec.Command("ls", "-la").Run()
	if err != nil {
		t.Error(err)
	}
}

func TestRun_cli(t *testing.T) {
	one_cli_run("../../acc build")
	one_cli_run("../../acc build")
}
