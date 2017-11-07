package main

import (
	//"bytes"
	"fmt"
	//"log"
	"os/exec"
	//"sync"
	//"strings"
)

func main() {
	//wg := new(sync.WaitGroup)
	//wg.Add(1)
	iadd := "hp@192.168.0.15"
	exec_ssh(iadd)
	// defer wg.Done()
}

func exec_ssh(s string) {
	cmd := exec.Command("ssh", s)
	err := cmd.Start()
	chkerr(err)
	err = cmd.Wait()
	chkerr(err)
	fmt.Println("***")
}

func chkerr(e error) {
	if e != nil {
		panic(e)
	}
}
