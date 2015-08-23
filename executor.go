package main

import (
	"fmt"
	osexec "os/exec"
)

type Executor interface {
	exec(cmd string, args ...string)
}

type runner struct {
	dir string
}

func NewExecutor(d string) Executor {
	return &runner{
		dir: d,
	}

}

func (*runner) exec(cmd string, args ...string) {
	out, err := osexec.Command(cmd, args...).CombinedOutput()
	if err != nil {
		fmt.Printf("%s", err)
	} else {
		fmt.Printf("%s", out)
	}
}
