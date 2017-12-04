package main

import (
	"fmt"
	"os/exec"
)

func main() {
	command := "date"
	exec_cmd(command)
}

func exec_cmd(c string) {
	cmd := exec.Command(c)
	out, outErr := cmd.Output()
	chkerr(outErr)
	fmt.Printf("%s", out)
}

func chkerr(err error) {
	if err != nil {
		panic(err)
	}
}
